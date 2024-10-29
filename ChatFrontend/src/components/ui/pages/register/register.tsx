import { FC } from 'react';
import { Link } from 'react-router-dom';
import styles from './register.module.css';
import openai_logo from '../../../../images/openai_logo.svg';
import error_icon from '../../../../images/error_icon.svg';

export const RegisterUI: FC = () => (
  <div className={styles['page-wrapper']}>
    <main className={styles['main-container']}>
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
            <a
              className={styles['other-page-link']}
              href='https://auth.openai.com/authorize?client_id=TdJIcbe16WoTHtN95nyywh5E4yOo6ItG&amp;scope=openid+email+profile+offline_access+model.request+model.read+organization.read+organization.write&amp;response_type=code&amp;redirect_uri=https%3A%2F%2Fchatgpt.com%2Fapi%2Fauth%2Fcallback%2Flogin-web&amp;audience=https%3A%2F%2Fapi.openai.com%2Fv1&amp;device_id=53a9f1dd-c42e-45f5-8842-305b20ebc03a&amp;prompt=login&amp;ext-oai-did=53a9f1dd-c42e-45f5-8842-305b20ebc03a&amp;flow=control&amp;ext-login-allow-phone=true&amp;country_code=US&amp;state=kdxaNfTh_LiNzpjjvFNU6wXlXE1r6i0j-9bFimJV27U&amp;code_challenge=z1ErbiB3NwbrNC94usE1qY4PbSCxGLa2kkxbEiQ4QFg&amp;code_challenge_method=S256'
            >
              Войти
            </a>
          </p>
        </div>
      </section>
    </main>
  </div>
);
