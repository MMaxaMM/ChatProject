import { TChat, TMessage, ChatType } from '@utils-types';
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { getChatsApi } from '@api';

export type ChatState = {
  userId: number | null;
  chats: TChat[];
  chatRequest: boolean;
  currentChatId: number;
  currentChatType: ChatType;
};

const initialState: ChatState = {
  userId: null,
  currentChatId: -1,
  currentChatType: ChatType.typeChat,
  chats: [
    {
      userId: 1,
      chatId: 0,
      chatType: ChatType.typeChat,
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
      chatType: ChatType.typeChat,
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
  return ans;
});

const chatSlice = createSlice({
  name: 'chatSlice',
  initialState,
  reducers: {
    sendMessage: (
      state,
      action: PayloadAction<{ chatId: number; message: TMessage }>
    ) => {
      state.chats[action.payload.chatId].messages.push(action.payload.message);
    },
    createChat: (state, action: PayloadAction<string>) => {
      const newChat: TChat = {
        userId: state.userId ? state.userId : -1,
        chatId: state.chats.length,
        chatType: state.currentChatType,
        messages: [
          {
            role: 'user',
            content: action.payload
          }
        ]
      };
      state.chats.push(newChat);
      state.currentChatId = newChat.chatId;
    },
    setChatId: (state, action: PayloadAction<number>) => {
      state.currentChatId = action.payload;
    },
    setChatType: (state, action: PayloadAction<ChatType>) => {
      state.currentChatType = action.payload;
    }
  },
  selectors: {
    getStoreChats: (state) => state.chats,
    getCurrentChatId: (state) => state.currentChatId,
    getCurrentChatType: (state) => state.currentChatType
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

export const { getStoreChats, getCurrentChatId, getCurrentChatType } =
  chatSlice.selectors;
export const { sendMessage, createChat, setChatId, setChatType } =
  chatSlice.actions;
export const chatReducer = chatSlice.reducer;
