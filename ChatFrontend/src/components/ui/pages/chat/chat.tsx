import styles from './chat.module.css';
import logo from '../../../../images/logo.svg';
import { FC } from 'react';
import closeIcon from '../../../../images/closeIcon.svg';
import { TChatUIProps } from './type';
import clsx from 'clsx';
import { MessageInput } from '@components';
import { Typewriter } from 'react-simple-typewriter';
import { useState } from 'react';

export const ChatUI: FC<TChatUIProps> = ({
  isAsideOpen,
  onOpenTab,
  onSendMessage
}) => {
  const [showCursor, setShowCursor] = useState(true);

  const handleType = (count: number) => {
    // Скрываем курсор, когда текущая фраза закончена
    console.log(count);
    setShowCursor(count === 0);
  };
  return (
    <div
      className={clsx(styles.main, {
        [styles.main__open]: isAsideOpen
      })}
    >
      <nav className={styles.header}>
        <div className={styles.tooltip_container}>
          <button
            className={styles.nav_button}
            onClick={onOpenTab}
            style={{ visibility: isAsideOpen ? 'hidden' : 'visible' }}
          >
            <img src={closeIcon} />
          </button>
          <span className={styles.tooltip_open_text}>
            Открыть боковую панель
          </span>
        </div>
        <div className={styles.nav_user_logo}>
          <div className={styles.nav_user_name}>D</div>
        </div>
      </nav>
      <div className={styles.content}>
        <img src={logo} className={styles.logo} />
        <p className={styles.start_text}>
          <Typewriter
            words={['Чем я могу помочь?']}
            loop={1} // Сколько раз повторять (1 — одноразово)
            cursor={showCursor}
            cursorBlinking={false}
            cursorStyle='●'
            typeSpeed={50} // Скорость ввода символов
            deleteSpeed={0} // Скорость удаления текста
            delaySpeed={100} // Задержка между текстами
            // onDelay={() => setShowCursor(false)}
            onType={(count) => {
              // Скрываем курсор после завершения фразы
              if (count === 0) setShowCursor(false);
            }}
            onLoopDone={() => setShowCursor(false)}
          />
        </p>
        <MessageInput onSendMessage={onSendMessage} />
      </div>
    </div>
  );
};
