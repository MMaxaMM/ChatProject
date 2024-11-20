import React, { FC, memo, useState, useEffect } from 'react';
import { TChatTypeModalProps } from './type';
import { ChatTypeModalUI } from '@ui';
import ReactDOM from 'react-dom';
import { useSelector } from '@store';
import { getCurrentChatType } from '@slices';
import { ChatType } from '@utils-types';

const modalRoot = document.getElementById('modals');

export const ChatTypeModal: FC<TChatTypeModalProps> = memo(
  ({ onSelectChats, onClose }) => {
    const currentChat = useSelector(getCurrentChatType);
    const chatType = Object.values(ChatType).includes(currentChat as ChatType)
      ? (currentChat as ChatType)
      : ChatType.typeChat;
    const [selectedChat, setSelectedChat] = useState<string>(chatType);
    useEffect(() => {
      const handleEsc = (e: KeyboardEvent) => {
        e.key === 'Escape' && onClose();
      };

      document.addEventListener('keydown', handleEsc);
      return () => {
        document.removeEventListener('keydown', handleEsc);
      };
    }, [onClose]);
    // Обработчик изменения состояния чекбоксов
    const handleCheckboxChange = (
      event: React.ChangeEvent<HTMLInputElement>
    ) => {
      const { value, checked } = event.target;
      setSelectedChat((prevSelectedChats: string) =>
        checked ? value : prevSelectedChats
      );
    };

    // Обработчик подтверждения выбора
    const handleConfirmSelection = () => {
      onSelectChats(selectedChat);
      onClose();
    };
    return ReactDOM.createPortal(
      <ChatTypeModalUI
        selectedChat={selectedChat}
        onClose={onClose}
        handleCheckboxChange={handleCheckboxChange}
        handleConfirmSelection={handleConfirmSelection}
      />,
      modalRoot as HTMLDivElement
    );
  }
);
