import { FC } from 'react';
import { Link, useMatch } from 'react-router-dom';
import { TChatListItemUIProps } from './type';
import styles from './chatListItem.module.css';
import chatLogo from '../../../images/chatLogo.svg';
import deleteIcon from '../../../images/deleteIcon.svg';
import audioLogo from '../../../images/audioLogo.svg';
import videoLogo from '../../../images/videoLogo.svg';
import docLogo from '../../../images/docLogo.svg';
import { useState } from 'react';
import { ChatType } from '@utils-types';

export const ChatListItemUI: FC<TChatListItemUIProps> = ({
  chat,
  onClick,
  onDelete
}) => {
  const { content, chat_id, chat_type } = chat;
  const [isHovered, setIsHovered] = useState(false);
  let logo: string;
  switch (chat_type) {
    case ChatType.typeAudio:
      logo = audioLogo;
      break;
    case ChatType.typeRAG:
      logo = docLogo;
      break;
    case ChatType.typeVideo:
      logo = videoLogo;
      break;

    default:
      logo = chatLogo;
      break;
  }
  return (
    <li
      className={`${useMatch(`/chat/${chat_id}`) && styles.chat_list_item_active} ${styles.chat_list_item}`}
      onClick={onClick}
    >
      <Link to={`/chat/${chat_id}`} className={styles.chat_list_item__title}>
        <span>{content}</span>
      </Link>

      {useMatch(`/chat/${chat_id}`) ? (
        <button
          className={styles.delete_button}
          onMouseOver={() => setIsHovered(true)}
          onMouseOut={() => setIsHovered(false)}
          onClick={onDelete}
        >
          <img
            src={!isHovered ? logo : deleteIcon}
            className={styles.chat_type_image}
          />
        </button>
      ) : (
        <img src={logo} className={styles.chat_type_image} />
      )}
    </li>
  );
};
