export type TMessage = {
  role: string;
  content: string;
};

export type TChat = {
  userId: number;
  chatId: number;
  messages: TMessage[];
};
