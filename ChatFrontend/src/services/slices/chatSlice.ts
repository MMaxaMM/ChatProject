import {
  TChat,
  TMessage,
  ChatType,
  getChatTypeByIndex,
  getChatTypeFromString
} from '@utils-types';
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
  TPostMessageRequest,
  postChatMessageApi,
  deleteChaTRequest,
  deleteChatApi,
  TPostAudioRequest,
  postAudioApi,
  postVideoApi
} from '@api';

export type ChatState = {
  userId: number | null;
  chats: TChat[];
  chatRequest: boolean;
  currentChatId: number;
  currentChatType: ChatType;
  progress: number | null;
  error: string | null;
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
  async (chatType: ChatType, { dispatch }) => {
    const ans = await createChatApi(chatType);
    dispatch(createChatStore({ chatId: ans.chat_id, chatType: chatType }));
    return ans;
  }
);

export const postMessage = createAsyncThunk(
  'chat/postMessage',
  async (data: TPostMessageRequest) => {
    const ans = await postChatMessageApi(data);
    return ans;
  }
);

export const deleteChat = createAsyncThunk(
  'chat/deleteChat',
  async (data: deleteChaTRequest, { dispatch }) => {
    const ans = await deleteChatApi(data);
    dispatch(getChats());
    return ans;
  }
);

export const postAudio = createAsyncThunk(
  'chat/postAudio',
  async (data: TPostAudioRequest, { dispatch }) => {
    const ans = await postAudioApi(data, (progress) => {
      // Обновляем прогресс через дополнительный экшен
      dispatch(setProgress(progress));
    });
    return ans;
  }
);

export const postVideo = createAsyncThunk(
  'chat/postVideo',
  async (data: TPostAudioRequest, { dispatch }) => {
    const ans = await postVideoApi(data, (progress) => {
      // Обновляем прогресс через дополнительный экшен
      dispatch(setProgress(progress));
    });
    return ans;
  }
);

const initialState: ChatState = {
  userId: null,
  currentChatId: -1,
  currentChatType: ChatType.typeChat,
  chats: [],
  chatRequest: false,
  progress: null,
  error: null
};

const chatSlice = createSlice({
  name: 'chatSlice',
  initialState,
  reducers: {
    sendMessage: (
      state,
      action: PayloadAction<{ chat_id: number; message: TMessage }>
    ) => {
      console.log('message local');
      state.chats.map((chat) => {
        if (chat.chat_id === action.payload.chat_id) {
          chat.messages.push(action.payload.message);
          console.log('message local paste');
        }
      });
    },
    createChatStore: (
      state,
      action: PayloadAction<{ chatId: number; chatType: ChatType }>
    ) => {
      console.log('local chat');
      state.currentChatId = action.payload.chatId;
      const chat: TChat = {
        user_id: state.userId ? state.userId : 0,
        chat_id: action.payload.chatId,
        chat_type: action.payload.chatType,
        date: new Date().toISOString(),
        content: 'Новый чат',
        messages: []
      };
      state.chats.push(chat);
    },
    setProgress(state, action) {
      state.progress = action.payload;
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
    getCurrentChatType: (state) => state.currentChatType,
    getProgress: (state) => state.progress
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
          chat.content = chat.content ? chat.content : 'Новый чат';
          chat.chat_type = getChatTypeByIndex(parseInt(chat.chat_type));
          chat.messages = chat.messages?.length ? chat.messages : [];
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
            chat.messages =
              chat.messages.length > action.payload.messages.length
                ? chat.messages
                : action.payload.messages;
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
        state.currentChatType = getChatTypeByIndex(action.payload.chat_type);
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
            const message: TMessage = {
              role: action.payload.message.role,
              content: action.payload.message.content,
              isNew: true,
              content_type: 1
            };
            chat.messages.push(message);
          }
        });
      })

      .addCase(deleteChat.pending, (state) => {
        state.chatRequest = true;
      })
      .addCase(deleteChat.rejected, (state, action) => {
        state.chatRequest = false;
      })
      .addCase(deleteChat.fulfilled, (state, action) => {
        state.chatRequest = false;
        state.chats.filter((chat) => chat.chat_id !== action.payload.chat_id);
      })

      .addCase(postAudio.pending, (state) => {
        state.progress = 0;
        state.error = null;
      })
      .addCase(postAudio.rejected, (state, action) => {
        state.error = action.payload as string;
      })
      .addCase(postAudio.fulfilled, (state, action) => {
        state.progress = 100;
        state.chats.map((chat) => {
          if (chat.chat_id === action.payload.chat_id) {
            const message: TMessage = {
              role: action.payload.message.role,
              content: action.payload.message.content,
              isNew: true,
              content_type: 1
            };
            chat.messages.push(message);
          }
        });
      })

      .addCase(postVideo.pending, (state) => {
        state.progress = 0;
        state.error = null;
      })
      .addCase(postVideo.rejected, (state, action) => {
        state.error = action.payload as string;
      })
      .addCase(postVideo.fulfilled, (state, action) => {
        state.progress = 100;
        state.chats.map((chat) => {
          if (chat.chat_id === action.payload.chat_id) {
            const message: TMessage = {
              role: action.payload.message.role,
              content: action.payload.message.content,
              isNew: true,
              content_type: 3
            };
            chat.messages.push(message);
          }
        });
      });
  }
});

export const {
  getStoreChats,
  getCurrentChatId,
  getCurrentChatType,
  getProgress
} = chatSlice.selectors;
export const {
  sendMessage,
  setChatId,
  setProgress,
  setChatType,
  createChatStore
} = chatSlice.actions;
export const chatReducer = chatSlice.reducer;
export const selectChatById = createSelector(
  [getStoreChats, (_, chatId) => chatId], // Аргументы
  (chats, chatId) => chats.find((chat) => chat.chat_id === chatId)
);
