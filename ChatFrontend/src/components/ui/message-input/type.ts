import { ChangeEventHandler, KeyboardEventHandler } from 'react';

export type TMessageInputProps = {
  message: string;
  handleChange: ChangeEventHandler<HTMLTextAreaElement>;
  handleKeyDown: KeyboardEventHandler<HTMLTextAreaElement>;
  handleSend: () => void;
  handleFileChange: ChangeEventHandler<HTMLInputElement>;
  handleClickFile: () => void;
};

export type MultiRefHandle = {
  fileRef: HTMLInputElement | null;
  textRef: HTMLTextAreaElement | null;
};
