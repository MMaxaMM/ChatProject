import { FC, useState } from 'react';
import { ChatUI } from '@ui-pages';
import { ChatList } from '@components';

export const Chat: FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleOpen = () => {
    setIsOpen(!isOpen);
  };

  const onCreateChat = () => void 0;

  return (
    <>
      <ChatList
        isOpen={isOpen}
        onClose={toggleOpen}
        onCreateChat={onCreateChat}
      />
      <ChatUI isAsideOpen={isOpen} onOpenTab={toggleOpen} />
    </>
  );
};
