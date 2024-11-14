import { TChat } from '@utils-types';

export type TChatOpenUIProps = {
  isAsideOpen: boolean;
  chat: TChat;
  onOpenTab: () => void;
  onSendMessage: (message: string) => void;
};
