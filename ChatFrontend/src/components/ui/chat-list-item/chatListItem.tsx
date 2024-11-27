import { FC } from 'react';
import { Link, useMatch } from 'react-router-dom';
import { TChatListItemUIProps } from './type';
import styles from './chatListItem.module.css';
import chatLogo from '../../../images/chatLogo.svg';
import deleteIcon from '../../../images/deleteIcon.svg';
import { useState } from 'react';

export const ChatListItemUI: FC<TChatListItemUIProps> = ({ chat, onClick }) => {
  const { content, chat_id, chat_type } = chat;
  const [isHovered, setIsHovered] = useState(false);
  return (
    <li
      className={`${useMatch(`/chat/${chat_id}`) && styles.chat_list_item_active} ${styles.chat_list_item}`}
      onClick={onClick}
    >
      <Link to={`/chat/${chat_id}`} className={styles.chat_list_item__title}>
        <span>{content}</span>
      </Link>

      {useMatch(`/chat/${chat_id}`) && (
        <button
          className={styles.delete_button}
          onMouseOver={() => setIsHovered(true)}
          onMouseOut={() => setIsHovered(false)}
        >
          <img src={!isHovered ? chatLogo : deleteIcon} />
        </button>
      )}
    </li>
  );
};
