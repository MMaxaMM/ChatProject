import { FC, memo } from 'react';
import { TChatTypeModalUIProps } from './type';
import styles from './chatTypeModal.module.css';

// Компонент модального окна с чекбоксами для выбора типов чатов
export const ChatTypeModalUI: FC<TChatTypeModalUIProps> = memo(
  ({ selectedChat, onClose, handleCheckboxChange, handleConfirmSelection }) => (
    <div className={styles['modal-overlay']} onClick={onClose}>
      <div
        className={styles['modal-content']}
        onClick={(e) => e.stopPropagation()}
      >
        <h2>Выберите тип чата</h2>

        <div className={styles['checkbox-container']}>
          <label className={styles['custom-checkbox']}>
            <input
              type='radio'
              value='CHAT'
              checked={selectedChat === 'CHAT'}
              onChange={handleCheckboxChange}
            />
            <span className={styles.checkmark} />
            Чат бот
          </label>

          <label className={styles['custom-checkbox']}>
            <input
              type='radio'
              value='RAG'
              checked={selectedChat === 'RAG'}
              onChange={handleCheckboxChange}
            />
            <span className={styles.checkmark} />
            RAG
          </label>

          <label className={styles['custom-checkbox']}>
            <input
              type='radio'
              value='AUDIO'
              checked={selectedChat === 'AUDIO'}
              onChange={handleCheckboxChange}
            />
            <span className={styles.checkmark} />
            Audio
          </label>

          <label className={styles['custom-checkbox']}>
            <input
              type='radio'
              value='VIDEO'
              checked={selectedChat === 'VIDEO'}
              onChange={handleCheckboxChange}
            />
            <span className={styles.checkmark} />
            Video
          </label>
        </div>

        <div className={styles['modal-buttons']}>
          <button onClick={handleConfirmSelection}>Создать чат</button>
        </div>
      </div>
    </div>
  )
);
