import { FC } from 'react';
import { ChatListUI } from '@ui';
import { TChatListProps } from './types';

export const ChatList: FC<TChatListProps> = ({
  isOpen,
  onClose,
  onCreateChat
}) => (
  <ChatListUI isOpen={isOpen} onClose={onClose} onCreateChat={onCreateChat} />
);
