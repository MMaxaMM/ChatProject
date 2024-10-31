import styles from './app.module.css';
import { Start, Register, Login } from '@pages';
import '../../index.css';
import { Route, Routes } from 'react-router-dom';

function App() {
  return (
    <main className={styles.app}>
      <Routes>
        <Route path='/' element={<Start />} />
        <Route path='/register' element={<Register />} />
        <Route path='/login' element={<Login />} />
      </Routes>
    </main>
  );
}

export default App;
