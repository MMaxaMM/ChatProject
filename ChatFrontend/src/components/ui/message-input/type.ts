import { ChangeEventHandler, KeyboardEventHandler } from 'react';

export type TMessageInputProps = {
  message: string;
  handleChange: ChangeEventHandler<HTMLTextAreaElement>;
  handleKeyDown: KeyboardEventHandler<HTMLTextAreaElement>;
};
