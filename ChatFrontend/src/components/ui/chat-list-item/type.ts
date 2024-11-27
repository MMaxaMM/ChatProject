import { TChat } from '@utils-types';
import React from 'react';

export type TChatListItemUIProps = {
  chat: TChat;
  onClick: () => void;
  onDelete: React.MouseEventHandler<HTMLButtonElement>;
};
