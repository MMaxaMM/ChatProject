import styles from './chatOpen.module.css';
import arrowStart from '../../../../images/arrowStart.svg';
import { FC } from 'react';
import closeIcon from '../../../../images/closeIcon.svg';
import { TChatOpenUIProps } from './type';
import clsx from 'clsx';
import { Message } from '@components';
import { TMessage } from '@utils-types';

export const ChatOpenUI: FC<TChatOpenUIProps> = ({
  isAsideOpen,
  chat,
  onOpenTab
}) => {
  const messages: TMessage[] = chat.messages;
  return (
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
        <ul className={styles.chat_list}>
          {messages.map((message, index) => (
            <Message message={message} key={index} />
          ))}
        </ul>
        <div className={styles.message_input}>
          <p className={styles.message_input__text}>Сообщить ChatGPT</p>
          <button className={styles.message_input__button}>
            <img
              src={arrowStart}
              className={styles.message_input__button_icon}
            />
          </button>
        </div>
      </div>
    </div>
  );
};
