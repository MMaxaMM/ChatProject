import { ChangeEventHandler } from 'react';

export type TChatTypeModalUIProps = {
  selectedChat: string;
  onClose: () => void;
  handleConfirmSelection: () => void;
  handleCheckboxChange: ChangeEventHandler<HTMLInputElement>;
};
