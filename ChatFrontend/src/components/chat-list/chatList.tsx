import { FC } from 'react';
import { ChatListUI } from '@ui';
import { TChatListProps } from './type';

export const ChatList: FC<TChatListProps> = ({
  chats,
  isOpen,
  isOpenModal,
  onClose,
  onCreateChat,
  onCloseModal,
  onSelectChat
}) => (
  <ChatListUI
    chats={chats}
    isOpen={isOpen}
    onClose={onClose}
    isOpenModal={isOpenModal}
    onCreateChat={onCreateChat}
    onCloseModal={onCloseModal}
    onSelectChat={onSelectChat}
  />
);
