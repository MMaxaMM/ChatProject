import { FC } from 'react';
import { NavLink } from 'react-router-dom';
import styles from './register.module.css';
import openai_logo from '../../../../images/openai_logo.svg';
import error_icon from '../../../../images/error_icon.svg';

export const RegisterUI: FC = () => (
  <div className={styles['page-wrapper']}>
    <div className={styles['main-container']}>
      <section className={styles['content-wrapper']}>
        <img src={openai_logo} alt='Логотип OpenAI' />
        <div className={styles['title-wrapper']}>
          <h1 className={styles.title}>Создать учетную запись</h1>
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
              Придумайте пароль
            </label>
            <div className={styles['invalid-email-error-message']}>
              <img className={styles['error-icon']} src={error_icon} />
              Недопустимый пароль.
            </div>
          </div>
          <button className={styles['continue-btn']} disabled>
            Зарегистрироваться
          </button>
          <p className={styles['other-page']}>
            У вас уже есть учетная запись?{' '}
            <NavLink className={styles['other-page-link']} to={'/login'}>
              Войти
            </NavLink>
          </p>
        </div>
      </section>
    </div>
  </div>
);
