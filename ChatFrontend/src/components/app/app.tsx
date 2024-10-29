import styles from './app.module.css';
import { Start, Register } from '@pages';
import '../../index.css';
import { Route, Routes } from 'react-router-dom';

function App() {
  return (
    <main className={styles.app}>
      <Routes>
        <Route path='/' element={<Start />} />
        <Route path='/register' element={<Register />} />
      </Routes>
    </main>
  );
}

export default App;
