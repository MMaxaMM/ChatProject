import { FC, useState } from 'react';
import { ChatListUI } from '@ui';

export const ChatList: FC = () => {
  const [isOpen, setIsOpen] = useState(true);

  const onClose = () => {
    setIsOpen(!isOpen);
  };

  const onCreateChat = () => void 0;

  return (
    <ChatListUI isOpen={isOpen} onClose={onClose} onCreateChat={onCreateChat} />
  );
};
