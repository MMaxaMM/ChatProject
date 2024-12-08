import { TChat } from '@utils-types';

export type TChatListProps = {
  isOpen: boolean;
  onClose: () => void;
  onCreateChat: () => void;
  isOpenModal: boolean;
  onCloseModal: () => void;
  onSelectChat: (selectedChat: string) => void;
};
