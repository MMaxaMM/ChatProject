import { forwardRef, useImperativeHandle } from 'react';
import styles from './messageInput.module.css';
import arrowStart from '../../../images/arrowStart.svg';
import { TMessageInputProps, MultiRefHandle } from './type';
import { useRef } from 'react';
import { ChatType } from '@utils-types';

export const MessageInputUI = forwardRef<MultiRefHandle, TMessageInputProps>(
  (
    {
      message,
      chatType,
      selectedFile,
      progress,
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
          accept='audio/*,video/*'
          ref={fileRef}
          className={styles.file_input} // Скрываем стандартный input
          onChange={handleFileChange}
        />
        {/* Стилизованный элемент */}
        <button
          onClick={handleClickFile}
          className={styles.file_input__button}
          disabled={
            ChatType.typeAudio !== chatType && ChatType.typeVideo !== chatType
          }
        />
        {chatType === ChatType.typeChat ? (
          <textarea
            ref={textRef}
            value={message}
            onChange={handleChange}
            onKeyDown={handleKeyDown}
            placeholder='Сообщить ChatGPT'
            rows={1}
            className={styles.message_input__text}
          />
        ) : (
          <span className={styles.progress_message}>{selectedFile?.name}</span>
        )}
        <button
          onClick={handleSend}
          className={styles.message_input__button}
          disabled={
            ChatType.typeChat === chatType ? !message.trim() : !selectedFile
          } // Кнопка отключена, если сообщение пустое
        >
          <img src={arrowStart} className={styles.message_input__button_icon} />
        </button>
      </div>
    );
  }
);
