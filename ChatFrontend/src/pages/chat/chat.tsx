import { FC, useState } from 'react';
import { ChatUI, ChatOpenUI } from '@ui-pages';
import { ChatList } from '@components';
import { TChat } from '@utils-types';
import { useSelector } from '@store';
import { getStoreChats } from '@slices';
import { useParams, useNavigate } from 'react-router-dom';

export const Chat: FC = () => {
  const [isOpen, setIsOpen] = useState(true);
  const params = useParams();
  const index = parseInt(params.id ? params.id : '-1');
  const navigate = useNavigate();
  console.log(index);

  const toggleOpen = () => {
    setIsOpen(!isOpen);
  };

  const onCreateChat = () => {
    navigate('/chat');
  };
  const onSendMessage = (message: string) => void 1;
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
