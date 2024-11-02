import styles from './start.module.css';
import logo from '../../../../images/logo.svg';
import arrowStart from '../../../../images/arrowStart.svg';
import { FC } from 'react';
import { TStartUIProps } from './type';

export const StartUI: FC<TStartUIProps> = ({ text, onLogin, onRegister }) => (
  <div className={styles.main}>
    <div className={styles.header}>
      <button className={styles.header__button} onClick={onLogin}>
        Войти
      </button>
      <button className={styles.header__button} onClick={onRegister}>
        Зарегистрироваться
      </button>
    </div>
    <div className={styles.content}>
      <img src={logo} className={styles.logo} />
      <p className={styles.start_text}>Чем я могу помочь?</p>
      <div className={styles.message_input}>
        <p className={styles.message_input__text}>
          {`${text}${text[text.length - 1] !== '?' ? '●' : ''}`}
        </p>
        <button className={styles.message_input__button}>
          <img src={arrowStart} className={styles.message_input__button_icon} />
        </button>
      </div>
    </div>
  </div>
);
