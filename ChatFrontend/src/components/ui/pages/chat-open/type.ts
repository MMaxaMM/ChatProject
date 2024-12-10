import { TChat } from '@utils-types';

export type TChatOpenUIProps = {
  isAsideOpen: boolean;
  chat: TChat;
  isUserModalOpen: boolean;
  onOpenUserModal: () => void;
  onCloseUserModal: () => void;
  onOpenTab: () => void;
  onSendMessage: (message: string) => void;
  onSendFile: (file: File) => void;
};
