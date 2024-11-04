import { FC } from 'react';
import { TMessageUIProps } from './type';
import styles from './message.module.css';

export const MessageUI: FC<TMessageUIProps> = ({ message }) => (
  <li
    className={`${styles.message} ${message.role === 'user' ? styles.message_user : styles.message_ai}`}
  >
    <p className={styles.message_text}>{message.content}</p>
  </li>
);
