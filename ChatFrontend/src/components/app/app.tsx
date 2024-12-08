import styles from './app.module.css';
import { Start, Register, Login, Chat } from '@pages';
import '../../index.css';
import { Route, Routes } from 'react-router-dom';
import { ProtectedRoute } from '@components';

function App() {
  return (
    <main className={styles.app}>
      <Routes>
        <Route path='/' element={<Start />} />
        <Route
          path='/register'
          element={
            <ProtectedRoute>
              <Register />
            </ProtectedRoute>
          }
        />
        <Route
          path='/login'
          element={
            <ProtectedRoute>
              <Login />
            </ProtectedRoute>
          }
        />
        {['/chat', '/chat/:id'].map((path) => (
          <Route
            path={path}
            element={
              <ProtectedRoute>
                <Chat />
              </ProtectedRoute>
            }
            key='Chat'
          />
        ))}
      </Routes>
    </main>
  );
}

export default App;
