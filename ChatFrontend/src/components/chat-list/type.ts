import { TChat } from '@utils-types';

export type TChatListProps = {
  chats: TChat[];
  isOpen: boolean;
  onClose: () => void;
  onCreateChat: () => void;
};
