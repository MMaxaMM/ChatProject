import { SyntheticEvent, ChangeEventHandler } from 'react';

export type TRegisterUIProps = {
  user: string;
  error: string | undefined;
  password: string;
  handeleSubmit: (e: SyntheticEvent) => void;
  setUsername: ChangeEventHandler<HTMLInputElement>;
  setpassword: ChangeEventHandler<HTMLInputElement>;
};
