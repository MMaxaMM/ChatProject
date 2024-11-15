import { FC, useEffect, useState } from 'react';
import { ChatUI, ChatOpenUI } from '@ui-pages';
import { ChatList } from '@components';
import { TChat, TMessage } from '@utils-types';
import { useSelector, useDispatch } from '@store';
import {
  getStoreChats,
  sendMessage,
  createChat,
  getCurrentChatId,
  setChatId
} from '@slices';
import { useParams, useNavigate } from 'react-router-dom';

export const Chat: FC = () => {
  const [isOpen, setIsOpen] = useState(true);
  const index = useSelector(getCurrentChatId);
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const toggleOpen = () => {
    setIsOpen(!isOpen);
  };

  const onCreateChat = () => {
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
        isOpen={isOpen}
        onClose={toggleOpen}
        onCreateChat={onCreateChat}
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
