import { FC } from 'react';
import { TMessageUIProps } from './type';
import styles from './message.module.css';
import { useState, useEffect } from 'react';
import { Typewriter } from 'react-simple-typewriter';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
export const MessageUI: FC<TMessageUIProps> = ({ message }) => {
  const [displayText, setDisplayText] = useState<string>(
    message.isNew ? '' : message.content
  ); // Текущий текст для отображения
  const [currentIndex, setCurrentIndex] = useState<number>(0); // Индекс текущего символа

  useEffect(() => {
    if (message.isNew && currentIndex < message.content.length) {
      const timer = setTimeout(() => {
        setDisplayText((prev) => prev + message.content[currentIndex]);
        setCurrentIndex((prev) => prev + 1);
      }, 50); // Скорость появления символов (в мс)

      return () => clearTimeout(timer); // Чистим таймер
    }
  }, [currentIndex, message.content]);
  return (
    <li
      className={`${styles.message} ${message.role === 'user' ? styles.message_user : styles.message_ai}`}
    >
      {message.content_type === 1 && (
        <ReactMarkdown
          children={displayText}
          remarkPlugins={[remarkGfm]}
          className={styles.message_text}
        />
      )}
      {message.content_type === 2 && (
        <audio controls>
          <source src={message.content} type='audio/mpeg' />
          Ваш браузер не поддерживает элемент <code>audio</code>.
        </audio>
      )}
    </li>
  );
};
