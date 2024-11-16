import { FC } from 'react';
import styles from './chatList.module.css';
import closeIcon from '../../../images/closeIcon.svg';
import newChatIcon from '../../../images/newChatIcon.svg';
import { TChatListUIProps } from './type';
import { ChatListItem } from '@components';

export const ChatListUI: FC<TChatListUIProps> = ({
  chats,
  isOpen,
  onClose,
  onCreateChat
}) => (
  <aside className={`${styles.container} ${isOpen && styles.container_open}`}>
    <div className={styles.content}>
      <nav className={styles.nav_buttons}>
        <button className={styles.nav_button} onClick={onClose}>
          <img src={closeIcon} />
        </button>
        <div className={styles.tooltip_container}>
          <button className={styles.nav_button} onClick={onCreateChat}>
            <img src={newChatIcon} />
          </button>
          <span className={styles.tooltip_text}>Новый чат</span>
        </div>
      </nav>
      <div className={styles.chat_list}>
        <h3 className={styles.chat_list_header}>Список чатов</h3>
        <ul className={styles.chat_list_content}>
          {chats.map((chat) => (
            <ChatListItem chat={chat} key={chat.chatId} />
          ))}
        </ul>
      </div>
    </div>
  </aside>
);
