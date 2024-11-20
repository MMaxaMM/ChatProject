export type TMessage = {
  role: string;
  content: string;
};

export type TChat = {
  userId: number;
  chatId: number;
  chatType: ChatType;
  messages: TMessage[];
};

export type TUser = {
  username: string;
  password: string;
};

export enum ChatType {
  typeChat = 'CHAT',
  typeRAG = 'RAG',
  typeAudio = 'AUDIO',
  typeVideo = 'VIDEO'
}
