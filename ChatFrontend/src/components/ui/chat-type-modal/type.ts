import { ChangeEventHandler } from 'react';

export type TChatTypeModalUIProps = {
  onClose: () => void;
  handleConfirmSelection: () => void;
  handleCheckboxChange: ChangeEventHandler<HTMLInputElement>;
};
