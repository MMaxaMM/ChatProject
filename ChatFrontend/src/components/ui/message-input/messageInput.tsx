import { forwardRef } from 'react';
import styles from './messageInput.module.css';
import arrowStart from '../../../images/arrowStart.svg';
import { TMessageInputProps } from './type';

export const MessageInputUI = forwardRef<
  HTMLTextAreaElement,
  TMessageInputProps
>(({ message, handleChange, handleKeyDown }, ref) => (
  <div className={styles.message_input}>
    <textarea
      ref={ref}
      value={message}
      onChange={handleChange}
      onKeyDown={handleKeyDown}
      placeholder='Сообщить ChatGPT'
      rows={1}
      className={styles.message_input__text}
    />
    <button
      className={styles.message_input__button}
      disabled={!message.trim()} // Кнопка отключена, если сообщение пустое
    >
      <img src={arrowStart} className={styles.message_input__button_icon} />
    </button>
  </div>
));
