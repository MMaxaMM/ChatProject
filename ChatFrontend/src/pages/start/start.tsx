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
    'Как правильно ставить ударение: обеспEчение или обеспечEние?',
    'Существует ли Ленинская комната и как её найти?'
  ];

  const [text, setText] = useState('');
  const [lineNumber, setLineNumber] = useState(0);
  const navigate = useNavigate();

  const shuffle = useCallback(() => {
    const index = Math.floor(Math.random() * lines.length);
    setLineNumber(index);
  }, []);
  let currentIndex = -1;
  useEffect(() => {
    const interval = setInterval(shuffle, 4000);
    const id = setInterval(() => {
      currentIndex += 1;
      setText((prev) => prev + lines[lineNumber][currentIndex]);
      if (currentIndex === lines[lineNumber].length - 1) {
        clearInterval(id);
      }
    }, 60);
    return () => {
      clearInterval(id);
      clearInterval(interval);
      setText('');
    };
  }, [lineNumber]);

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
      <StartUI
        text={text}
        onClickLogin={onLogin}
        onClickRegister={onRegister}
      />
    </>
  );
};
