import styles from './app.module.css';
import logo from '../../images/logo.svg';
import arrowStart from '../../images/arrowStart.svg';
import '../../index.css';
import { useCallback, useEffect, useState } from 'react';

function App() {
  const lines = [
    [
      'Ка',
      'к ',
      'сд',
      'ать',
      ' г',
      'ра',
      'фы',
      ' Ми',
      'ха',
      'илу',
      ' Анд',
      'рее',
      'вичу',
      '?'
    ],
    [
      'Ко',
      'гда',
      ' Дми',
      'трий',
      ' Иго',
      'рев',
      'ич',
      ' п',
      'рий',
      'дёт',
      ' на',
      ' па',
      'ру',
      '?'
    ],
    [
      'Что',
      ' та',
      'кое',
      ' си',
      'нта',
      'ксич',
      'еск',
      'ая ',
      'омо',
      'ним',
      'ия',
      '?'
    ],
    [
      'Как ',
      'сде',
      'ла',
      'ть ',
      'НИРС',
      ' за',
      ' два ',
      'ча',
      'са ',
      'до ',
      'сда',
      'чи?'
    ],
    [
      'Как ',
      'под',
      'нять ',
      'лок',
      'аль',
      'ный ',
      'сер',
      'вер ',
      'в ',
      'Cou',
      'nter',
      '-Str',
      'ike?'
    ],
    [
      'Как ',
      'пра',
      'вил',
      'ьно ',
      'ста',
      'вить ',
      'уда',
      'рен',
      'ие: ',
      'обес',
      'пече',
      'ние ',
      'или ',
      'обес',
      'пече',
      'ние?'
    ],
    [
      'Сущ',
      'ест',
      'вует',
      ' ли ',
      'Лен',
      'инск',
      'ая ',
      'ком',
      'ната ',
      'и ',
      'как ',
      'её ',
      'най',
      'ти?'
    ]
  ];

  const [text, setText] = useState('');
  const [lineNumber, setLineNumber] = useState(0);

  const shuffle = useCallback(() => {
    const index = Math.floor(Math.random() * lines.length);
    setLineNumber(index);
  }, []);
  let currentIndex = -1;
  useEffect(() => {
    const interval = setInterval(shuffle, 3000);
    const id = setInterval(() => {
      currentIndex += 1;
      setText((prev) => prev + lines[lineNumber][currentIndex]);
      if (currentIndex === lines[lineNumber].length - 1) {
        clearInterval(id);
      }
    }, 120);
    return () => {
      clearInterval(id);
      clearInterval(interval);
      setText('');
    };
  }, [lineNumber]);

  return (
    <div className={styles.app}>
      <div className={styles.main}>
        <img src={logo} className={styles.logo} />

        <p className={styles.start_text}>Чем я могу помочь?</p>
        <div className={styles.message_input}>
          <p className={styles.message_input__text}>{text}</p>
          <button className={styles.message_input__button}>
            <img
              src={arrowStart}
              className={styles.message_input__button_icon}
            />
          </button>
        </div>
      </div>
    </div>
  );
}

export default App;
