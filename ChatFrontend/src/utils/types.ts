export type TMessage = {
  role: string;
  content: string;
};

export type TChat = {
  userId: number;
  chatId: number;
  messages: TMessage[];
};

export type TUser = {
  username: string;
  password: string;
}
