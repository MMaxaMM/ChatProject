import { FC, SyntheticEvent, useState } from 'react';
import { LoginUI } from '@ui-pages';
import { loginUser } from '@slices';
import { useDispatch } from '../../services/store';
import { useNavigate } from 'react-router-dom';

export const Login: FC = () => {
  const navigate = useNavigate();
  const [user, setUser] = useState('');
  const [password, setPassword] = useState('');
  const dispatch = useDispatch();

  const handeleSubmit = (e: SyntheticEvent) => {
    e.preventDefault();
    dispatch(loginUser({ username: user, password: password }))
      .unwrap()
      .then(() => {
        navigate('/chat');
      })
      .catch((err) => console.log(err));
  };

  const onChangePassword = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };

  const onChangeUsername = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUser(e.target.value);
  };
  return (
    <LoginUI
      user={user}
      password={password}
      handeleSubmit={handeleSubmit}
      setUsername={onChangeUsername}
      setpassword={onChangePassword}
    />
  );
};
