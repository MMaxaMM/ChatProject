import React from 'react';
import { FC } from 'react';
import { useSelector } from '@store';
import { getIsAuthenticated, getUserId } from '@slices';
import { Navigate, useMatch } from 'react-router-dom';
import { getCookie } from '../../utils/cookie';

type ProtectedRouteProps = {
  children: React.ReactElement;
};

export const ProtectedRoute: FC<ProtectedRouteProps> = ({
  children
}: ProtectedRouteProps) => {
  const isAuthChecked = useSelector(getIsAuthenticated); // isAuthCheckedSelector — селектор получения состояния загрузки пользователя
  const userId = useSelector(getUserId);
  const loginMatch = useMatch('/login');
  const registerMatch = useMatch('/register');
  const token = getCookie('accessToken');

  if (loginMatch || registerMatch) {
    return children;
  }

  if (!token && !isAuthChecked && !userId) {
    // если пользователь на странице авторизации и данных в хранилище нет, то делаем редирект
    return <Navigate replace to='/login' />;
  }

  return children;
};
