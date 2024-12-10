import { TUser } from '@utils-types';
import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { setCookie, deleteCookie } from '../../utils/cookie';
import { registerUserApi, loginUserApi } from '@api';

export type UserState = {
  userId: number | null;
  username: string;
  isAuthChecked: boolean;
  isAuthenticated: boolean;
  loginUserRequest: boolean;
};

const initialState: UserState = {
  userId: null,
  username: '',
  isAuthChecked: false,
  isAuthenticated: false,
  loginUserRequest: false
};

export const registerUser = createAsyncThunk(
  'user/register',
  async (data: TUser) => {
    const ans = await registerUserApi(data);
    return ans.user_id;
  }
);

export const loginUser = createAsyncThunk(
  'user/login',
  async (data: TUser, { dispatch }) => {
    const ans = await loginUserApi(data);
    dispatch(setUsername(data.username));
    setCookie('accessToken', ans.token);
    localStorage.setItem('username', ans.username);
    return ans;
  }
);

const userSlice = createSlice({
  name: 'userSlice',
  initialState,
  reducers: {
    logout: (state) => {
      state.isAuthChecked = false;
      state.isAuthenticated = false;
      state.username = '';
      state.userId = null;
      deleteCookie('accessToken');
      localStorage.clear();
    },
    setUsername: (state, action: PayloadAction<string>) => {
      state.username = action.payload;
    },
    refreshUsername: (state) => {
      const username = localStorage.getItem('username');
      state.username = username ? username : '';
    }
  },
  selectors: {
    getIsAuthenticated: (state) => state.isAuthenticated,
    getUserId: (state) => state.userId,
    getUsername: (state) => state.username
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
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        state.username = action.payload.username;
        state.loginUserRequest = false;
        state.isAuthenticated = true;
        state.isAuthChecked = true;
      });
  }
});

export const { getIsAuthenticated, getUserId, getUsername } =
  userSlice.selectors;
export const { logout, setUsername, refreshUsername } = userSlice.actions;
export const userReducer = userSlice.reducer;
