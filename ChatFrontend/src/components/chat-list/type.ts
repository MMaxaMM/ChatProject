import { TChat } from '@utils-types';

export type TChatListProps = {
  chats: TChat[];
  isOpen: boolean;
  onClose: () => void;
  onCreateChat: () => void;
  isOpenModal: boolean;
  onCloseModal: () => void;
  onSelectChat: (selectedChat: string) => void;
};
