import { FC } from 'react';
import { ChatListUI } from '@ui';
import { TChatListProps } from './type';

export const ChatList: FC<TChatListProps> = ({
  chats,
  isOpen,
  onClose,
  onCreateChat
}) => (
  <ChatListUI
    chats={chats}
    isOpen={isOpen}
    onClose={onClose}
    onCreateChat={onCreateChat}
  />
);
