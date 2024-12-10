import styles from './chatOpen.module.css';
import { FC, useEffect, useRef } from 'react';
import closeIcon from '../../../../images/closeIcon.svg';
import { TChatOpenUIProps } from './type';
import clsx from 'clsx';
import { Message, MessageInput } from '@components';
import { TMessage } from '@utils-types';
import { useSelector } from '@store';
import { getUsername } from '@slices';
import { UserModal } from '@components';

export const ChatOpenUI: FC<TChatOpenUIProps> = ({
  isAsideOpen,
  chat,
  isUserModalOpen,
  onOpenUserModal,
  onCloseUserModal,
  onOpenTab,
  onSendMessage,
  onSendFile
}) => {
  const username = useSelector(getUsername);
  const messages: TMessage[] = chat.messages;
  const messagesEndRef = useRef<HTMLDivElement | null>(null);
  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [messages]);
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
        <div className={styles.nav_user_logo} onClick={onOpenUserModal}>
          <div className={styles.nav_user_name}>{username[0]}</div>
        </div>
        {isUserModalOpen && <UserModal onClose={onCloseUserModal} />}
      </nav>
      <div className={styles.content}>
        <ul className={styles.chat_list}>
          {messages.map((message, index) => (
            <Message message={message} key={index} />
          ))}
          <div ref={messagesEndRef} />
        </ul>
        <div className={styles.message_input_wrapper}>
          <MessageInput onSendMessage={onSendMessage} onSendFile={onSendFile} />
        </div>
      </div>
    </div>
  );
};
