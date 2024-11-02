import { FC, useState } from 'react';
import { ChatUI } from '@ui-pages';
import { ChatList } from '@components';
import { TChat } from '@utils-types';

export const Chat: FC = () => {
  const [isOpen, setIsOpen] = useState(true);

  const toggleOpen = () => {
    setIsOpen(!isOpen);
  };

  const onCreateChat = () => void 0;

  const chats: TChat[] = [
    {
      userId: 1,
      chatId: 1,
      messages: [
        {
          role: 'user',
          content: 'hello'
        }
      ]
    },
    {
      userId: 2,
      chatId: 2,
      messages: [
        {
          role: 'user',
          content: 'hello'
        }
      ]
    }
  ];

  return (
    <>
      <ChatList
        chats={chats}
        isOpen={isOpen}
        onClose={toggleOpen}
        onCreateChat={onCreateChat}
      />
      <ChatUI isAsideOpen={isOpen} onOpenTab={toggleOpen} />
    </>
  );
};
