import { TUser } from '@utils-types';
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { setCookie } from '../../utils/cookie';
import { registerUserApi, loginUserApi } from '@api';

export type UserState = {
  userId: number | null;
  isAuthChecked: boolean;
  isAuthenticated: boolean;
  loginUserRequest: boolean;
};

const initialState: UserState = {
  userId: null,
  isAuthChecked: false,
  isAuthenticated: false,
  loginUserRequest: false
};

export const registerUser = createAsyncThunk(
  'user/register',
  async (data: TUser) => {
    const ans = await registerUserApi(data);
    console.log(ans);
    return ans.user_id;
  }
);

export const loginUser = createAsyncThunk('user/login', async (data: TUser) => {
  const ans = await loginUserApi(data);
  setCookie('accessToken', ans.token);
  return;
});

const userSlice = createSlice({
  name: 'userSlice',
  initialState,
  reducers: {
    logout: (state) => {
      state = initialState;
    }
  },
  selectors: {
    getIsAuthenticated: (state) => state.isAuthenticated,
    getUserId: (state) => state.userId
  },
  extraReducers: (builder) => {
    builder
      .addCase(registerUser.pending, (state) => {
        state.loginUserRequest = true;
      })
      .addCase(registerUser.rejected, (state, action) => {
        state.loginUserRequest = false;
        state.isAuthChecked = false;
      })
      .addCase(registerUser.fulfilled, (state, action) => {
        state.userId = action.payload;
        state.loginUserRequest = false;
        state.isAuthenticated = true;
        state.isAuthChecked = true;
      })

      .addCase(loginUser.pending, (state) => {
        state.loginUserRequest = true;
      })
      .addCase(loginUser.rejected, (state, action) => {
        state.loginUserRequest = false;
        state.isAuthChecked = false;
        console.log(action);
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        state.loginUserRequest = false;
        state.isAuthenticated = true;
        state.isAuthChecked = true;
      });
  }
});

export const { getIsAuthenticated, getUserId } = userSlice.selectors;
export const userReducer = userSlice.reducer;
