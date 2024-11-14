import { ChangeEventHandler, KeyboardEventHandler } from 'react';

export type TChatUIProps = {
  isAsideOpen: boolean;
  onOpenTab: () => void;
  onSendMessage: (message: string) => void;
};
