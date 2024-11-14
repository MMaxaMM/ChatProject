import { SyntheticEvent, ChangeEventHandler } from 'react';

export type TLoginUIProps = {
  user: string;
  password: string;
  handeleSubmit: (e: SyntheticEvent) => void;
  setUsername: ChangeEventHandler<HTMLInputElement>;
  setpassword: ChangeEventHandler<HTMLInputElement>;
};
