import { FC, memo } from 'react';
import { ChatListItemUI } from '@ui';
import { TChatListItemProps } from './type';
import { useDispatch } from '@store';
import { setChatId, deleteChat } from '@slices';
import { useNavigate } from 'react-router-dom';

export const ChatListItem: FC<TChatListItemProps> = memo(({ chat }) => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const onClick = () => {
    dispatch(setChatId(chat.chat_id));
    // dispatch(setChatType(chat.chat_type));
  };
  const onDelete = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.stopPropagation();
    dispatch(deleteChat({ chatId: chat.chat_id, chatType: chat.chat_type }));
    dispatch(setChatId(-1));
    navigate('/chat');
  };
  return <ChatListItemUI chat={chat} onClick={onClick} onDelete={onDelete} />;
});
