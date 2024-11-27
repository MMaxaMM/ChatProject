import { FC, memo } from 'react';
import { ChatListItemUI } from '@ui';
import { TChatListItemProps } from './type';
import { useDispatch } from '@store';
import { setChatId, deleteChat } from '@slices';

export const ChatListItem: FC<TChatListItemProps> = memo(({ chat }) => {
  const dispatch = useDispatch();
  const onClick = () => {
    dispatch(setChatId(chat.chat_id));
  };
  const onDelete = () => {
    dispatch(deleteChat({ chatId: chat.chat_id, chatType: chat.chat_type }));
  };
  return <ChatListItemUI chat={chat} onClick={onClick} onDelete={onDelete} />;
});
