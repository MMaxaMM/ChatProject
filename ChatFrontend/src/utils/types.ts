export type TMessage = {
  role: string;
  content: string;
};

export type TChat = {
  user_id: number;
  chat_id: number;
  chat_type: ChatType;
  date: string;
  content: string;
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

export function getChatTypeFromString(ChatTypeStr: string): ChatType {
  return Object.values(ChatType).includes(ChatTypeStr as ChatType)
    ? (ChatTypeStr as ChatType)
    : ChatType.typeChat;
}

export function getIndexByChatType(chatType: ChatType): number {
  const chatTypes = Object.values(ChatType);
  return chatTypes.indexOf(chatType);
}

export function getChatTypeByIndex(chatIndex: number): ChatType {
  const chatTypes = Object.values(ChatType);
  return chatTypes[chatIndex];
}
