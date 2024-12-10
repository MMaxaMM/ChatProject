export type TChatUIProps = {
  isAsideOpen: boolean;
  isUserModalOpen: boolean;
  onOpenUserModal: () => void;
  onCloseUserModal: () => void;
  onOpenTab: () => void;
  onSendMessage: (message: string) => void;
  onSendFile: (file: File) => void;
};
