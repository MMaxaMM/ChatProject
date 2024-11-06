import { setCookie, getCookie } from './cookie';
import { TChat, TMessage, TUser } from './types';

const URL = process.env.CHAT_API_URL;

const checkResponse = <T>(res: Response): Promise<T> =>
  res.ok ? res.json() : res.json().then((err) => Promise.reject(err));

type TServerResponse<T> = {
  success: boolean;
} & T;

type TChatStart = {
  chatId: number;
  content: string;
};

type TChatStartResponse = TServerResponse<{
  userId: number;
  chats: TChatStart[];
}>;

type TChatCreateResponse = TServerResponse<{
  userId: number;
  chatId: number;
}>;

export const getChatsApi = () =>
  fetch(`${URL}/chat/start`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
      authorization: getCookie('accessToken')
    } as HeadersInit
  })
    .then((res) => checkResponse<TChatStartResponse>(res))
    .then((data) => {
      if (data?.success) return data;
      return Promise.reject(data);
    });

export const createChatApi = () =>
  fetch(`${URL}/chat/create`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
      authorization: getCookie('accessToken')
    } as HeadersInit
  })
    .then((res) => checkResponse<TChatCreateResponse>(res))
    .then((data) => {
      if (data?.success) return data;
      return Promise.reject(data);
    });

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
      if (data?.success) return data;
      return Promise.reject(data);
    });

type TPostMessageQuery = {
  chatId: number;
  message: TMessage;
};

type TPostMessageResponse = TServerResponse<{
  userId: number;
  chatId: number;
  message: TMessage;
}>;

type TChatHistory = TServerResponse<TChat>;

export const postChatMessageApi = (data: TPostMessageQuery) =>
  fetch(`${URL}/chat/message`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
      authorization: getCookie('accessToken')
    } as HeadersInit,
    body: JSON.stringify(data)
  })
    .then((res) => checkResponse<TPostMessageResponse>(res))
    .then((data) => {
      if (data?.success) return data;
      return Promise.reject(data);
    });

export const getChatApi = (chatId: number) =>
  fetch(`${URL}/chat/history`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
      authorization: getCookie('accessToken')
    } as HeadersInit,
    body: JSON.stringify({
      chat_id: chatId
    })
  })
    .then((res) => checkResponse<TChatHistory>(res))
    .then((data) => {
      if (data?.success) return data;
      return Promise.reject(data);
    });

type TAuthResponse = TServerResponse<{
  token: string;
}>;

export const registerUserApi = (data: TUser) =>
  fetch(`${URL}/auth/sign-up`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    },
    body: JSON.stringify(data)
  })
    .then((res) => checkResponse<TServerResponse<{ userId: number }>>(res))
    .then((data) => {
      if (data?.success) return data;
      return Promise.reject(data);
    });

export const loginUserApi = (data: TUser) =>
  fetch(`${URL}/auth/sign-up`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8'
    },
    body: JSON.stringify(data)
  })
    .then((res) => checkResponse<TAuthResponse>(res))
    .then((data) => {
      if (data?.success) return data;
      return Promise.reject(data);
    });
