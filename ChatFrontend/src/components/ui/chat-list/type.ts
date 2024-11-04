import { TChat } from '@utils-types';

export type TChatListUIProps = {
  chats: TChat[];
  isOpen: boolean;
  onClose: () => void;
  onCreateChat: () => void;
};
