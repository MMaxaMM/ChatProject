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
  chats: [],
  chatRequest: false
};

export const getChats = createAsyncThunk('chat/start', async () => {
  const ans = await getChatsApi();
  console.log(ans);
  return ans;
});

const userSlice = createSlice({
  name: 'chatSlice',
  initialState,
  reducers: {},
  selectors: {},
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

export const chatReducer = userSlice.reducer;
