import { ChatType } from '@utils-types';
import { ChangeEventHandler, KeyboardEventHandler } from 'react';

export type TMessageInputProps = {
  message: string;
  chatType: ChatType;
  selectedFile: File | null;
  handleChange: ChangeEventHandler<HTMLTextAreaElement>;
  handleKeyDown: KeyboardEventHandler<HTMLTextAreaElement>;
  handleSend: () => void;
  handleFileChange: ChangeEventHandler<HTMLInputElement>;
  handleClickFile: () => void;
  progress: number | null;
};

export type MultiRefHandle = {
  fileRef: HTMLInputElement | null;
  textRef: HTMLTextAreaElement | null;
};
