import styles from './chat.module.css';
import logo from '../../../../images/logo.svg';
import arrowStart from '../../../../images/arrowStart.svg';
import { FC } from 'react';

export const ChatUI: FC = () => (
  <div className={styles.main}>
    <div className={styles.header}>
      <div className={styles.user_logo} />
    </div>
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
