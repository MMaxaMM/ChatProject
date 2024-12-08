import { FC } from 'react';
import { ChatListUI } from '@ui';
import { TChatListProps } from './type';
import { useSelector } from '@store';
import { getStoreChats } from '@slices';
import { TChat } from '@utils-types';

export const ChatList: FC<TChatListProps> = ({
  isOpen,
  isOpenModal,
  onClose,
  onCreateChat,
  onCloseModal,
  onSelectChat
}) => {
  const chats: TChat[] = useSelector(getStoreChats);
  return (
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
};
