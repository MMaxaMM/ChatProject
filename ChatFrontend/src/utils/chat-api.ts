import { error } from 'console';
import { getCookie } from './cookie';
import { ChatType, getIndexByChatType, TChat, TMessage, TUser } from './types';

const URL = 'http://127.0.0.1:5050';

const checkResponse = <T>(res: Response): Promise<T> =>
  res.ok
    ? res.json()
    : res.json().then((err) => {
        console.log(err);
        return Promise.reject({ message: err.error } as Error);
      });

type TServerResponse<T> = {
  ok: boolean;
} & T;

type TChatStartResponse = {
  user_id: number;
  chats: TChat[];
};

type TChatCreateResponse = {
  user_id: number;
  chat_type: number;
  chat_id: number;
};

export const getChatsApi = () =>
  fetch(`${URL}/control/start`, {
    method: 'GET',
    headers: {
      authorization: `Bearer ${getCookie('accessToken')}`
    } as HeadersInit
  })
    .then((res) => checkResponse<TChatStartResponse>(res))
    .then((data) => data);

export const createChatApi = (chatType: ChatType) =>
  fetch(`${URL}/control/create`, {
    method: 'POST',
    headers: {
      authorization: `Bearer ${getCookie('accessToken')}`
    } as HeadersInit,
    body: JSON.stringify({
      chat_type: getIndexByChatType(chatType)
    })
  }).then((res) => checkResponse<TChatCreateResponse>(res));

export type deleteChaTRequest = { chatId: number; chatType: ChatType };

export const deleteChatApi = (data: deleteChaTRequest) =>
  fetch(`${URL}/control/delete`, {
    method: 'DELETE',
    headers: {
      authorization: `Bearer ${getCookie('accessToken')}`
    } as HeadersInit,
    body: JSON.stringify({
      chat_type: getIndexByChatType(data.chatType),
      chat_id: data.chatId
    })
  }).then((res) => checkResponse<{ chat_id: number }>(res));

export type TPostMessageRequest = {
  chat_id: number;
  message: TMessage;
};

type TPostMessageResponse = {
  user_id: number;
  chat_id: number;
  message: TMessage;
};

type TChatHistory = Omit<TChat, 'chatType'>;

export const postChatMessageApi = (data: TPostMessageRequest) =>
  fetch(`${URL}/chat/message`, {
    method: 'POST',
    headers: {
      authorization: `Bearer ${getCookie('accessToken')}`
    } as HeadersInit,
    body: JSON.stringify(data)
  }).then((res) => checkResponse<TPostMessageResponse>(res));

export const getChatHistoryApi = (chatId: number) =>
  fetch(`${URL}/control/history`, {
    method: 'POST',
    headers: {
      authorization: `Bearer ${getCookie('accessToken')}`
    } as HeadersInit,
    body: JSON.stringify({
      chat_id: chatId
    })
  }).then((res) => checkResponse<TChatHistory>(res));

export type TPostAudioRequest = {
  chat_id: number;
  formData: FormData;
};

export const postAudioApi = (
  data: TPostAudioRequest,
  onProgress: (progress: number) => void
) =>
  fetchWithProgress(
    `${URL}/audio/recognize?chat_id=${data.chat_id}`,
    {
      method: 'POST',
      headers: {
        authorization: `Bearer ${getCookie('accessToken')}`
      } as HeadersInit,
      body: data.formData
    },
    onProgress
  ).then((res) => checkResponse<TPostMessageResponse>(res));

export const postVideoApi = (
  data: TPostAudioRequest,
  onProgress: (progress: number) => void
) =>
  fetchWithProgress(
    `${URL}/video/recognize?chat_id=${data.chat_id}`,
    {
      method: 'POST',
      headers: {
        authorization: `Bearer ${getCookie('accessToken')}`
      } as HeadersInit,
      body: data.formData
    },
    onProgress
  ).then((res) => checkResponse<TPostMessageResponse>(res));

type TAuthResponse = {
  token: string;
};

export const registerUserApi = (data: TUser) =>
  fetch(`${URL}/auth/sign-up`, {
    method: 'POST',
    body: JSON.stringify(data)
  }).then((res) => checkResponse<{ user_id: number }>(res));

export const loginUserApi = (data: TUser) =>
  fetch(`${URL}/auth/sign-in`, {
    method: 'POST',
    body: JSON.stringify(data)
  })
    .then((res) => checkResponse<TAuthResponse>(res))
    .then((data) => {
      if (data?.token) return data;
      return Promise.reject(data);
    });

// Функция для fetch с отслеживанием прогресса
async function fetchWithProgress(
  url: string,
  options: RequestInit,
  onProgress: (progress: number) => void
): Promise<Response> {
  const request = new XMLHttpRequest();
  request.open(options.method || 'POST', url);

  if (options.headers) {
    Object.entries(options.headers).forEach(([key, value]) => {
      request.setRequestHeader(key, value as string);
    });
  }
  return new Promise((resolve, reject) => {
    request.upload.onprogress = (event) => {
      if (event.lengthComputable) {
        const percentCompleted = Math.round((event.loaded * 100) / event.total);
        onProgress(percentCompleted);
      }
    };

    request.onload = () => {
      resolve(new Response(request.responseText, { status: request.status }));
    };

    request.onerror = () => reject(new Error('Ошибка при загрузке файла'));

    // Передаем тело запроса
    if (options.body instanceof FormData) {
      request.send(options.body);
    } else {
      reject(new Error('Body должен быть FormData'));
    }
  });
}
