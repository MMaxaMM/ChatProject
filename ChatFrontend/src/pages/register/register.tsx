import { FC, SyntheticEvent, useState } from 'react';
import { RegisterUI } from '@ui-pages';
import { registerUser } from '@slices';
import { useDispatch } from '../../services/store';
import { useNavigate } from 'react-router-dom';

export const Register: FC = () => {
  const navigate = useNavigate();
  const [user, setUser] = useState('');
  const [password, setPassword] = useState('');
  const dispatch = useDispatch();
  const [error, setError] = useState<Error | null>(null);

  const handeleSubmit = (e: SyntheticEvent) => {
    e.preventDefault();
    dispatch(registerUser({ username: user, password: password }))
      .unwrap()
      .then(() => {
        navigate('/login');
      })
      .catch((err) => setError(err));
  };

  const onChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const onChangeUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUser(e.target.value);
  };

  return (
    <RegisterUI
      user={user}
      error={error?.message}
      password={password}
      handeleSubmit={handeleSubmit}
      setUsername={onChangeUsername}
      setpassword={onChangePassword}
    />
  );
};
