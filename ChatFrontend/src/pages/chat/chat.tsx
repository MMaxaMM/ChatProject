import { FC, useEffect, useState } from 'react';
import { ChatUI, ChatOpenUI } from '@ui-pages';
import { ChatList } from '@components';
import { TChat, TMessage, ChatType } from '@utils-types';
import { useSelector, useDispatch } from '@store';
import {
  getStoreChats,
  sendMessage,
  createChat,
  getCurrentChatId,
  setChatId,
  setChatType
} from '@slices';
import { useParams, useNavigate } from 'react-router-dom';

export const Chat: FC = () => {
  const [isOpen, setIsOpen] = useState(true);
  const [isOpenModal, setIsOpenModal] = useState(false);
  const index = useSelector(getCurrentChatId);
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const toggleOpen = () => {
    setIsOpen(!isOpen);
  };

  const onCloseModal = () => {
    setIsOpenModal(false);
  };

  const onCreateChat = () => {
    setIsOpenModal(true);
  };

  const onSelectChats = (selectedChat: string) => {
    const chatType = Object.values(ChatType).includes(selectedChat as ChatType)
      ? (selectedChat as ChatType)
      : ChatType.typeChat;
    dispatch(setChatType(chatType));
    navigate('/chat');
    dispatch(setChatId(-1));
  };

  useEffect(() => {
    if (index !== -1) {
      navigate(`/chat/${index}`);
    }
  }, [index]);

  const onSendMessage = (message: string) => {
    const data: TMessage = {
      role: 'user',
      content: message
    };
    if (index === -1) {
      dispatch(createChat(message));
    } else {
      dispatch(sendMessage({ chatId: index, message: data }));
    }
  };
  const chats: TChat[] = useSelector(getStoreChats);

  return (
    <>
      <ChatList
        chats={chats}
        isOpenModal={isOpenModal}
        onCloseModal={onCloseModal}
        isOpen={isOpen}
        onClose={toggleOpen}
        onCreateChat={onCreateChat}
        onSelectChat={onSelectChats}
      />
      {index >= 0 ? (
        <ChatOpenUI
          isAsideOpen={isOpen}
          chat={chats[index]}
          onOpenTab={toggleOpen}
          onSendMessage={onSendMessage}
        />
      ) : (
        <ChatUI
          isAsideOpen={isOpen}
          onSendMessage={onSendMessage}
          onOpenTab={toggleOpen}
        />
      )}
    </>
  );
};
