import styles from './chat.module.css';
import logo from '../../../../images/logo.svg';
import { FC, forwardRef } from 'react';
import closeIcon from '../../../../images/closeIcon.svg';
import { TChatUIProps } from './type';
import clsx from 'clsx';
import { MessageInput } from '@components';

export const ChatUI: FC<TChatUIProps> = ({
  isAsideOpen,
  onOpenTab,
  onSendMessage
}) => (
  <div
    className={clsx(styles.main, {
      [styles.main__open]: isAsideOpen
    })}
  >
    <nav className={styles.header}>
      <button
        className={styles.nav_button}
        onClick={onOpenTab}
        style={{ visibility: isAsideOpen ? 'hidden' : 'visible' }}
      >
        <img src={closeIcon} />
      </button>
      <div className={styles.nav_user_logo}>
        <div className={styles.nav_user_name}>D</div>
      </div>
    </nav>
    <div className={styles.content}>
      <img src={logo} className={styles.logo} />
      <p className={styles.start_text}>Чем я могу помочь?</p>
      <MessageInput onSendMessage={onSendMessage} />
    </div>
  </div>
);
