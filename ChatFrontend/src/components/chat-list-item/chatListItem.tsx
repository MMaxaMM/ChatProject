import { FC, memo } from 'react';
import { ChatListItemUI } from '@ui';
import { TChatListItemProps } from './type';
import { useDispatch } from '@store';
import { setChatId } from '@slices';

export const ChatListItem: FC<TChatListItemProps> = memo(({ chat }) => {
  const dispatch = useDispatch();
  const onClick = () => {
    dispatch(setChatId(chat.chatId));
  };
  return <ChatListItemUI chat={chat} onClick={onClick} />;
});
