export type TMessageInputProps = {
  onSendMessage: (message: string) => void; // Функция для отправки сообщения
  onSendFile: (file: File) => void;
};

export type MultiRefHandle = {
  fileRef: HTMLInputElement | null;
  textRef: HTMLTextAreaElement | null;
};
