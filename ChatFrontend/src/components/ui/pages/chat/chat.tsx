import styles from './chat.module.css';
import logo from '../../../../images/logo.svg';
import arrowStart from '../../../../images/arrowStart.svg';
import { FC } from 'react';
import closeIcon from '../../../../images/closeIcon.svg';
import { TChatUIProps } from './type';
import clsx from 'clsx';

export const ChatUI: FC<TChatUIProps> = ({ isAsideOpen, onOpenTab }) => (
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
      <div className={styles.message_input}>
        <p className={styles.message_input__text}>Сообщить ChatGPT</p>
        <button className={styles.message_input__button}>
          <img src={arrowStart} className={styles.message_input__button_icon} />
        </button>
      </div>
    </div>
  </div>
);
