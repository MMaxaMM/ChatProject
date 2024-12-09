import { FC } from 'react';
import { NavLink } from 'react-router-dom';
import styles from './register.module.css';
import openai_logo from '../../../../images/openai_logo.svg';
import error_icon from '../../../../images/error_icon.svg';
import { TRegisterUIProps } from './type';

export const RegisterUI: FC<TRegisterUIProps> = ({
  user,
  error,
  password,
  handeleSubmit,
  setUsername,
  setpassword
}) => (
  <div className={styles['page-wrapper']}>
    <div className={styles['main-container']}>
      <section className={styles['content-wrapper']}>
        <img src={openai_logo} alt='Логотип OpenAI' />
        <div className={styles['title-wrapper']}>
          <h1 className={styles.title}>Создать учетную запись</h1>
        </div>
        <form
          className={styles['login-container']}
          name='login'
          onSubmit={handeleSubmit}
        >
          <div className={styles['input-wrapper']}>
            <input
              className={styles['email-input']}
              inputMode='text'
              type='username'
              id='user-input'
              name='username'
              autoComplete='username'
              autoCapitalize='none'
              spellCheck='false'
              required
              placeholder=''
              value={user}
              onChange={setUsername}
            />
            <label className={styles['email-label']} htmlFor='user-input'>
              Имя пользователя
            </label>
            {/* <div className={styles['invalid-email-error-message']}>
              <img className={styles['error-icon']} src={error_icon} />
              Недопустимое имя пользователя.
            </div> */}
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
              value={password}
              onChange={setpassword}
            />
            <label className={styles['email-label']} htmlFor='password-input'>
              Придумайте пароль
            </label>
            {error && (
              <div className={styles['invalid-email-error-message']}>
                <img className={styles['error-icon']} src={error_icon} />
                {error}
              </div>
            )}
          </div>
          <button className={styles['continue-btn']} type='submit'>
            Зарегистрироваться
          </button>
          <p className={styles['other-page']}>
            У вас уже есть учетная запись?{' '}
            <NavLink className={styles['other-page-link']} to={'/login'}>
              Войти
            </NavLink>
          </p>
        </form>
      </section>
    </div>
  </div>
);
