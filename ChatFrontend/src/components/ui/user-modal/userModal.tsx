import { FC, memo } from 'react';
import { TUserModalUIProps } from './type';
import styles from './userModal.module.css';
import exitIcon from '../../../images/exitIcon.svg';
import userIcon from '../../../images/userIcon.svg';

// Компонент модального окна с чекбоксами для выбора типов чатов
export const UserModalUI: FC<TUserModalUIProps> = memo(
  ({ username, onClose, onLogout }) => (
    <div className={styles['modal-overlay']} onClick={onClose}>
      <div
        className={styles['modal-content']}
        onClick={(e) => e.stopPropagation()}
      >
        <h2>
          <img src={userIcon} className={styles.chat_type_image} />
          {username}
        </h2>

        <div className={styles['modal-buttons']}>
          <button onClick={onLogout}>
            <img src={exitIcon} className={styles.chat_type_image} />
            Выйти
          </button>
        </div>
      </div>
    </div>
  )
);
