import { FC, memo } from 'react';
import { TChatTypeModalUIProps } from './type';
import styles from './chatTypeModal.module.css';

// Компонент модального окна с чекбоксами для выбора типов чатов
export const ChatTypeModalUI: FC<TChatTypeModalUIProps> = memo(
  ({ onClose, handleCheckboxChange, handleConfirmSelection }) => (
    <div className={styles['modal-overlay']} onClick={onClose}>
      <div
        className={styles['modal-content']}
        onClick={(e) => e.stopPropagation()}
      >
        <h2>Выберите тип чатов</h2>

        <div className={styles['checkbox-container']}>
          <label>
            <input
              type='checkbox'
              value='General Chat'
              onChange={handleCheckboxChange}
            />
            Общий чат
          </label>

          <label>
            <input
              type='checkbox'
              value='Customer Support'
              onChange={handleCheckboxChange}
            />
            Поддержка
          </label>

          <label>
            <input
              type='checkbox'
              value='Technical Support'
              onChange={handleCheckboxChange}
            />
            Техническая поддержка
          </label>

          <label>
            <input
              type='checkbox'
              value='Sales Chat'
              onChange={handleCheckboxChange}
            />
            Продажа
          </label>
        </div>

        <div className={styles['modal-buttons']}>
          <button onClick={handleConfirmSelection}>Подтвердить</button>
          <button className={styles['close-button']} onClick={onClose}>
            Закрыть
          </button>
        </div>
      </div>
    </div>
  )
);
