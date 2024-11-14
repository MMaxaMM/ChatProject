import { getCookie } from './cookie';
import { TChat, TMessage, TUser } from './types';

const URL = 'http://127.0.0.1:5050';

const checkResponse = <T>(res: Response): Promise<T> =>
  res.ok ? res.json() : res.json().then((err) => Promise.reject(err));

type TServerResponse<T> = {
  ok: boolean;
} & T;

type TChatStartResponse = {
  userId: number;
  chats: TChat[];
};

type TChatCreateResponse = {
  userId: number;
  chatId: number;
};

export const getChatsApi = () =>
  fetch(`${URL}/chat/start`, {
    method: 'GET',
    headers: {
      authorization: getCookie('accessToken')
    } as HeadersInit
  })
    .then((res) => checkResponse<TChatStartResponse>(res))
    .then((data) => data);

export const createChatApi = () =>
  fetch(`${URL}/chat/create`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
      authorization: getCookie('accessToken')
    } as HeadersInit
  })
    .then((res) => checkResponse<TChatCreateResponse>(res))
    .then((data) => data);

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
      if (data?.ok) return data;
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
      if (data?.ok) return data;
      return Promise.reject(data);
    });

type TAuthResponse = {
  token: string;
};

export const registerUserApi = (data: TUser) =>
  fetch(`${URL}/auth/sign-up`, {
    method: 'POST',
    body: JSON.stringify(data)
  })
    .then((res) => checkResponse<{ user_id: number }>(res))
    .then((data) => data);

export const loginUserApi = (data: TUser) =>
  fetch(`${URL}/auth/sign-in`, {
    method: 'POST',
    body: JSON.stringify(data)
  })
    .then((res) => checkResponse<TAuthResponse>(res))
    .then((data) => data);
