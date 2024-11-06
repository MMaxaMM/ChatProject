import { FC } from 'react';
import styles from './login.module.css';
import openai_logo from '../../../../images/openai_logo.svg';
import error_icon from '../../../../images/error_icon.svg';
import { NavLink } from 'react-router-dom';

export const LoginUI: FC = () => (
  <div className={styles['page-wrapper']}>
    <div className={styles['main-container']}>
      <section className={styles['content-wrapper']}>
        <img src={openai_logo} alt='Логотип OpenAI' />
        <div className={styles['title-wrapper']}>
          <h1 className={styles.title}>С возвращением</h1>
        </div>
        <div className={styles['login-container']}>
          <div className={styles['input-wrapper']}>
            <input
              className={styles['email-input']}
              inputMode='email'
              type='email'
              id='email-input'
              name='email'
              autoComplete='username'
              autoCapitalize='none'
              spellCheck='false'
              required
              placeholder=''
            />
            <label className={styles['email-label']} htmlFor='email-input'>
              Адрес электронной почты
            </label>
            <div className={styles['invalid-email-error-message']}>
              <img className={styles['error-icon']} src={error_icon} />
              Недопустимый адрес электронной почты.
            </div>
          </div>
          <div className={styles['input-wrapper']}>
            <input
              className={styles['email-input']}
              inputMode='text'
              type='password'
              id='password-input'
              name='password'
              autoComplete='username'
              autoCapitalize='none'
              spellCheck='false'
              required
              placeholder=''
            />
            <label className={styles['email-label']} htmlFor='password-input'>
              Пароль
            </label>
            <div className={styles['invalid-email-error-message']}>
              <img className={styles['error-icon']} src={error_icon} />
              Недопустимый пароль.
            </div>
          </div>
          <button className={styles['continue-btn']} disabled>
            Войти
          </button>
          <p className={styles['other-page']}>
            У вас нет учетной записи?{' '}
            <NavLink className={styles['other-page-link']} to={'/register'}>
              Зарегистрироваться
            </NavLink>
          </p>
        </div>
      </section>
    </div>
  </div>
);
