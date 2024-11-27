import { FC } from 'react';
import { TMessageUIProps } from './type';
import styles from './message.module.css';
import { useState } from 'react';
import { Typewriter } from 'react-simple-typewriter';
export const MessageUI: FC<TMessageUIProps> = ({ message }) => {
  const [showCursor, setShowCursor] = useState(true);
  return (
    <li
      className={`${styles.message} ${message.role === 'user' ? styles.message_user : styles.message_ai}`}
    >
      <p className={styles.message_text}>
        {message.role !== 'user' && message.isNew ? (
          <Typewriter
            words={[message.content]}
            loop={1} // Сколько раз повторять (1 — одноразово)
            cursor={false}
            cursorBlinking={false}
            cursorStyle='●'
            typeSpeed={50} // Скорость ввода символов
            deleteSpeed={0} // Скорость удаления текста
            delaySpeed={10} // Задержка между текстами
            // onDelay={() => setShowCursor(false)}
          />
        ) : (
          message.content
        )}
      </p>
    </li>
  );
};
