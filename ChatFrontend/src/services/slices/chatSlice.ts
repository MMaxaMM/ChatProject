import { TChat } from '@utils-types';
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { setCookie } from '../../utils/cookie';
import { getChatsApi, loginUserApi } from '@api';

export type ChatState = {
  userId: number | null;
  chats: TChat[];
  chatRequest: boolean;
};

const initialState: ChatState = {
  userId: null,
  chats: [
    {
      userId: 1,
      chatId: 0,
      messages: [
        {
          role: 'user',
          content: 'hello'
        }
      ]
    },
    {
      userId: 2,
      chatId: 1,
      messages: [
        {
          role: 'user',
          content: 'Привет!'
        },
        {
          role: 'assistent',
          content: 'Ответ ассистента...'
        },
        {
          role: 'user',
          content: 'Привет!'
        },
        {
          role: 'assistent',
          content: 'Ответ ассистента...'
        },
        {
          role: 'user',
          content: 'Привет!'
        },
        {
          role: 'assistent',
          content: 'Ответ ассистента...'
        }
      ]
    }
  ],
  chatRequest: false
};

export const getChats = createAsyncThunk('chat/start', async () => {
  const ans = await getChatsApi();
  console.log(ans);
  return ans;
});

const chatSlice = createSlice({
  name: 'chatSlice',
  initialState,
  reducers: {
    
  },
  selectors: {
    getStoreChats: (state) => state.chats
  },
  extraReducers: (builder) => {
    builder
      .addCase(getChats.pending, (state) => {
        state.chatRequest = true;
      })
      .addCase(getChats.rejected, (state, action) => {
        state.chatRequest = false;
      })
      .addCase(getChats.fulfilled, (state, action) => {
        state.userId = action.payload.userId;
        state.chatRequest = false;
        state.chats = action.payload.chats;
      });
  }
});

export const { getStoreChats } = chatSlice.selectors;
export const chatReducer = chatSlice.reducer;
