import { FC } from 'react';
import { ChatUI } from '@ui-pages';
import { ChatList } from '@components';

export const Chat: FC = () => (
  <>
    <ChatList />
    <ChatUI />
  </>
);
