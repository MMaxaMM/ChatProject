import React, { FC, memo, useState } from 'react';
import { TChatTypeModalProps } from './type';
import { ChatTypeModalUI } from '@ui';
import ReactDOM from 'react-dom';

const modalRoot = document.getElementById('modals');

export const ChatTypeModal: FC<TChatTypeModalProps> = memo(
  ({ onSelectChats, onClose }) => {
    const [selectedChats, setSelectedChats] = useState<string[]>([]);

    // Обработчик изменения состояния чекбоксов
    const handleCheckboxChange = (
      event: React.ChangeEvent<HTMLInputElement>
    ) => {
      const { value, checked } = event.target;
      setSelectedChats((prevSelectedChats) =>
        checked
          ? [...prevSelectedChats, value]
          : prevSelectedChats.filter((chat) => chat !== value)
      );
    };

    // Обработчик подтверждения выбора
    const handleConfirmSelection = () => {
      onSelectChats(selectedChats);
      onClose();
    };
    return ReactDOM.createPortal(
      <ChatTypeModalUI
        onClose={onClose}
        handleCheckboxChange={handleCheckboxChange}
        handleConfirmSelection={handleConfirmSelection}
      />,
      modalRoot as HTMLDivElement
    );
  }
);
