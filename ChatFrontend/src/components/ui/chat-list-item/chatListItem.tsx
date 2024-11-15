import { FC } from 'react';
import { Link, useMatch } from 'react-router-dom';
import { TChatListItemUIProps } from './type';
import styles from './chatListItem.module.css';

export const ChatListItemUI: FC<TChatListItemUIProps> = ({ chat, onClick }) => {
  const { userId, chatId, messages } = chat;
  const title = chatId;
  return (
    <li
      className={`${useMatch(`/chat/${chatId}`) && styles.chat_list_item_active} ${styles.chat_list_item}`}
      onClick={onClick}
    >
      <Link to={`/chat/${chatId}`} className={styles.chat_list_item__title}>
        <span>{title}</span>
      </Link>
    </li>
  );
};
