import { TChat, TMessage, ChatType, getChatTypeByIndex } from '@utils-types';
import {
  createSlice,
  createAsyncThunk,
  PayloadAction,
  createSelector
} from '@reduxjs/toolkit';
import {
  getChatsApi,
  getChatHistoryApi,
  createChatApi,
  TPostMessageQuery,
  postChatMessageApi
} from '@api';

export type ChatState = {
  userId: number | null;
  chats: TChat[];
  chatRequest: boolean;
  currentChatId: number;
  currentChatType: ChatType;
};

export const getChatHistory = createAsyncThunk(
  'chat/getHistory',
  async (chatId: number) => {
    const ans = await getChatHistoryApi(chatId);
    return ans;
  }
);

export const getChats = createAsyncThunk('chat/start', async () => {
  const ans = await getChatsApi();
  return ans;
});

export const createChat = createAsyncThunk(
  'chat/createChat',
  async (chatType: ChatType) => {
    const ans = await createChatApi(chatType);
    return ans;
  }
);

export const postMessage = createAsyncThunk(
  'chat/postMessage',
  async (data: TPostMessageQuery) => {
    const ans = await postChatMessageApi(data);
    return ans;
  }
);

const initialState: ChatState = {
  userId: null,
  currentChatId: -1,
  currentChatType: ChatType.typeChat,
  chats: [],
  chatRequest: false
};

const chatSlice = createSlice({
  name: 'chatSlice',
  initialState,
  reducers: {
    sendMessage: (
      state,
      action: PayloadAction<{ chat_id: number; message: TMessage }>
    ) => {
      state.chats.map((chat) => {
        if (chat.chat_id === action.payload.chat_id) {
          chat.messages.push(action.payload.message);
        }
      });
    },
    createChatStore: (state, action: PayloadAction<string>) => {
      const newChat: TChat = {
        user_id: state.userId ? state.userId : -1,
        chat_id: state.chats.length,
        chat_type: state.currentChatType,
        content: action.payload,
        date: new Date().toISOString(),
        messages: [
          {
            role: 'user',
            content: action.payload
          }
        ]
      };
      state.chats.push(newChat);
      state.currentChatId = newChat.chat_id;
    },
    setChatId: (state, action: PayloadAction<number>) => {
      state.currentChatId = action.payload;
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
        state.userId = action.payload.user_id;
        state.chatRequest = false;
        state.chats = action.payload.chats;
        state.chats.map((chat) => {
          chat.content = 'Новый чат';
        });
      })

      .addCase(getChatHistory.pending, (state) => {
        state.chatRequest = true;
      })
      .addCase(getChatHistory.rejected, (state, action) => {
        state.chatRequest = false;
      })
      .addCase(getChatHistory.fulfilled, (state, action) => {
        state.chatRequest = false;
        state.chats.map((chat) => {
          if (chat.chat_id === action.payload.chat_id) {
            chat.messages = action.payload.messages;
          }
        });
      })

      .addCase(createChat.pending, (state) => {
        state.chatRequest = true;
      })
      .addCase(createChat.rejected, (state, action) => {
        state.chatRequest = false;
      })
      .addCase(createChat.fulfilled, (state, action) => {
        state.chatRequest = false;
        state.currentChatId = action.payload.chat_id;
        state.currentChatType = getChatTypeByIndex(action.payload.chat_type);
        // const newChat: TChat = {
        //   user_id: action.payload.user_id,
        //   chat_id: action.payload.chat_id,
        //   chat_type: getChatTypeByIndex(action.payload.chat_type),
        //   content: 'Новый чат',
        //   date: new Date().toISOString(),
        //   messages: []
        // };
        // state.chats.push(newChat);
      })

      .addCase(postMessage.pending, (state) => {
        state.chatRequest = true;
      })
      .addCase(postMessage.rejected, (state, action) => {
        state.chatRequest = false;
      })
      .addCase(postMessage.fulfilled, (state, action) => {
        state.chatRequest = false;
        state.chats.map((chat) => {
          if (chat.chat_id === action.payload.chat_id) {
            chat.messages.push(action.payload.message);
          }
        });
      });
  }
});

export const { getStoreChats, getCurrentChatId, getCurrentChatType } =
  chatSlice.selectors;
export const { sendMessage, createChatStore, setChatId } = chatSlice.actions;
export const chatReducer = chatSlice.reducer;
export const selectChatById = createSelector(
  [getStoreChats, (_, chatId) => chatId], // Аргументы
  (chats, chatId) => chats.find((chat) => chat.chat_id === chatId)
);
