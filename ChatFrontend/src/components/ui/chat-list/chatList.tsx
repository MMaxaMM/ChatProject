import { FC } from 'react';
import styles from './chatList.module.css';
import closeIcon from '../../../images/closeIcon.svg';
import newChatIcon from '../../../images/newChatIcon.svg';
import { TChatListUIProps } from './type';
import { ChatListItem, ChatTypeModal } from '@components';

export const ChatListUI: FC<TChatListUIProps> = ({
  chats,
  isOpen,
  isOpenModal,
  onClose,
  onCreateChat,
  onCloseModal,
  onSelectChat
}) => (
  <aside className={`${styles.container} ${isOpen && styles.container_open}`}>
    <div className={styles.content}>
      <nav className={styles.nav_buttons}>
        <div className={styles.tooltip_container}>
          <button className={styles.nav_button} onClick={onClose}>
            <img src={closeIcon} />
          </button>
          <span className={styles.tooltip_close_text}>
            Закрыть боковую панель
          </span>
        </div>
        <div className={styles.tooltip_container}>
          <button className={styles.nav_button} onClick={onCreateChat}>
            <img src={newChatIcon} />
          </button>
          <span className={styles.tooltip_create_text}>Новый чат</span>
        </div>
        {isOpenModal && (
          <ChatTypeModal onClose={onCloseModal} onSelectChats={onSelectChat} />
        )}
      </nav>
      <div className={styles.chat_list}>
        <h3 className={styles.chat_list_header}>Список чатов</h3>
        <ul className={styles.chat_list_content}>
          {chats.map((chat) => (
            <ChatListItem chat={chat} key={chat.chat_id} />
          ))}
        </ul>
      </div>
    </div>
  </aside>
);
