import { FC } from 'react';
import { Link, useMatch } from 'react-router-dom';
import { TChatListItemUIProps } from './type';
import styles from './chatListItem.module.css';

export const ChatListItemUI: FC<TChatListItemUIProps> = ({ chat, onClick }) => {
  const { content, chat_id, chat_type } = chat;
  return (
    <li
      className={`${useMatch(`/chat/${chat_id}`) && styles.chat_list_item_active} ${styles.chat_list_item}`}
      onClick={onClick}
    >
      <Link to={`/chat/${chat_id}`} className={styles.chat_list_item__title}>
        <span>{content}</span>
      </Link>
    </li>
  );
};
