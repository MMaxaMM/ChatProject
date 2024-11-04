import { FC, memo } from 'react';
import { ChatListItemUI } from '@ui';
import { TChatListItemProps } from './type';

export const ChatListItem: FC<TChatListItemProps> = memo(({ chat }) => (
  <ChatListItemUI chat={chat} />
));
