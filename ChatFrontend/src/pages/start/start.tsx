import { useCallback, useEffect, useState, FC } from 'react';
import { StartUI } from '@ui-pages';
import { replace, useNavigate } from 'react-router-dom';

export const Start: FC = () => {
  const lines = [
    'Как сдать графы Михаилу Андреевичу?',
    'Когда Дмитрий Игоревич прийдёт на пару?',
    'Что такое синтаксическая омонимия?',
    'Как сделать НИРС за два часа до сдачи?',
    'Как поднять локальный сервер в Counter-Strike?',
    'Как правильно ставить ударение: обеспечение или обеспечение?',
    'Существует ли Ленинская комната и как её найти?'
  ];

  const navigate = useNavigate();

  const onLogin = () => {
    navigate('/login', { replace: true });
    return;
  };

  const onRegister = () => {
    navigate('/register', { replace: true });
    return;
  };

  return (
    <>
      <StartUI text={lines} onLogin={onLogin} onRegister={onRegister} />
    </>
  );
};
