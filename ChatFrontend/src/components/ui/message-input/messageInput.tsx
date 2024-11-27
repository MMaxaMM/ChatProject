import { forwardRef, useImperativeHandle } from 'react';
import styles from './messageInput.module.css';
import arrowStart from '../../../images/arrowStart.svg';
import { TMessageInputProps, MultiRefHandle } from './type';
import { useRef } from 'react';

export const MessageInputUI = forwardRef<MultiRefHandle, TMessageInputProps>(
  (
    {
      message,
      handleChange,
      handleKeyDown,
      handleSend,
      handleFileChange,
      handleClickFile
    },
    ref
  ) => {
    const fileRef = useRef<HTMLInputElement>(null);
    const textRef = useRef<HTMLTextAreaElement>(null);

    useImperativeHandle(ref, () => ({
      fileRef: fileRef.current,
      textRef: textRef.current
    }));
    return (
      <div className={styles.message_input}>
        <input
          type='file'
          accept='audio/*'
          ref={fileRef}
          className={styles.file_input} // Скрываем стандартный input
          onChange={handleFileChange}
        />
        {/* Стилизованный элемент */}
        <button
          onClick={handleClickFile}
          className={styles.file_input__button}
        />
        <textarea
          ref={textRef}
          value={message}
          onChange={handleChange}
          onKeyDown={handleKeyDown}
          placeholder='Сообщить ChatGPT'
          rows={1}
          className={styles.message_input__text}
        />
        <button
          onClick={handleSend}
          className={styles.message_input__button}
          disabled={!message.trim()} // Кнопка отключена, если сообщение пустое
        >
          <img src={arrowStart} className={styles.message_input__button_icon} />
        </button>
      </div>
    );
  }
);
