import { getCookie } from './cookie';
import { ChatType, getIndexByChatType, TChat, TMessage, TUser } from './types';

const URL = 'http://127.0.0.1:5050';

const checkResponse = <T>(res: Response): Promise<T> =>
  res.ok ? res.json() : res.json().then((err) => Promise.reject(err));

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
      chat_id: getIndexByChatType(chatType)
    })
  }).then((res) => checkResponse<TChatCreateResponse>(res));

export const deleteChatApi = (chatId: number) =>
  fetch(`${URL}/chat/delete`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
      authorization: getCookie('accessToken')
    } as HeadersInit,
    body: JSON.stringify({
      chat_id: chatId
    })
  })
    .then((res) => checkResponse<TServerResponse<{}>>(res))
    .then((data) => {
      if (data?.ok) return data;
      return Promise.reject(data);
    });

export type TPostMessageQuery = {
  chat_id: number;
  message: TMessage;
};

type TPostMessageResponse = {
  user_id: number;
  chat_id: number;
  message: TMessage;
};

type TChatHistory = Omit<TChat, 'chatType'>;

export const postChatMessageApi = (data: TPostMessageQuery) =>
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
  }).then((res) => checkResponse<TAuthResponse>(res));
