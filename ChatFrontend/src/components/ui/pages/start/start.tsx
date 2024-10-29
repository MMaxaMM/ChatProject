import styles from './start.module.css';
import logo from '../../../../images/logo.svg';
import arrowStart from '../../../../images/arrowStart.svg';
import { FC } from 'react';
import { TStartUIProps } from './type';

export const StartUI: FC<TStartUIProps> = ({ text }) => (
  <div className={styles.main}>
    <img src={logo} className={styles.logo} />

    <p className={styles.start_text}>Чем я могу помочь?</p>
    <div className={styles.message_input}>
      <p className={styles.message_input__text}>{text}</p>
      <button className={styles.message_input__button}>
        <img src={arrowStart} className={styles.message_input__button_icon} />
      </button>
    </div>
  </div>
);
