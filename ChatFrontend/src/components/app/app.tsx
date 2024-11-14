import styles from './app.module.css';
import { Start, Register, Login, Chat } from '@pages';
import '../../index.css';
import { Route, Routes } from 'react-router-dom';

function App() {
  return (
    <main className={styles.app}>
      <Routes>
        <Route path='/' element={<Start />} />
        <Route path='/register' element={<Register />} />
        <Route path='/login' element={<Login />} />
        {['/chat', '/chat/:id'].map((path) => (
          <Route path={path} element={<Chat />} key='Chat' />
        ))}
      </Routes>
    </main>
  );
}

export default App;
