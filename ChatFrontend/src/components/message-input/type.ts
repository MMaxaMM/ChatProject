export type TMessageInputProps = {
  onSendMessage: (message: string) => void; // Функция для отправки сообщения
};

export type MultiRefHandle = {
  fileRef: HTMLInputElement | null;
  textRef: HTMLTextAreaElement | null;
};
