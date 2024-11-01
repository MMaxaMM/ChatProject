import { FC } from 'react';
import styles from './chatList.module.css';
import closeIcon from '../../../images/closeIcon.svg';
import newChatIcon from '../../../images/newChatIcon.svg';
import { TChatListUIProps } from './type';

export const ChatListUI: FC<TChatListUIProps> = ({
  isOpen,
  onClose,
  onCreateChat
}) => (
  <aside className={`${styles.container} ${isOpen && styles.container_open}`}>
    <div className={styles.content}>
      <nav>
        <div className={styles.nav_buttons}>
          <button className={styles.nav_button} onClick={onClose}>
            <img src={closeIcon} />
          </button>
          <button className={styles.nav_button} onClick={onCreateChat}>
            <img src={newChatIcon} />
          </button>
        </div>
      </nav>
    </div>
  </aside>
);
