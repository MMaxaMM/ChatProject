import React, { FC, memo, useState, useEffect } from 'react';
import { TUserModalProps } from './type';
import { UserModalUI } from '@ui';
import ReactDOM from 'react-dom';
import { useDispatch, useSelector } from '@store';
import { logout, getUsername } from '@slices';
import { useNavigate } from 'react-router-dom';

const modalRoot = document.getElementById('modals');

export const UserModal: FC<TUserModalProps> = memo(({ onClose }) => {
  const username = useSelector(getUsername);
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const onLogout = () => {
    dispatch(logout());
    navigate('/');
  };

  useEffect(() => {
    const handleEsc = (e: KeyboardEvent) => {
      e.key === 'Escape' && onClose();
    };

    document.addEventListener('keydown', handleEsc);
    return () => {
      document.removeEventListener('keydown', handleEsc);
    };
  }, [onClose]);
  return ReactDOM.createPortal(
    <UserModalUI username={username} onClose={onClose} onLogout={onLogout} />,
    modalRoot as HTMLDivElement
  );
});
