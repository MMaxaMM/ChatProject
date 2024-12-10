import {
  FC,
  useEffect,
  useRef,
  ChangeEvent,
  KeyboardEvent,
  useState
} from 'react';
import { TMessageInputProps, MultiRefHandle } from './type';
import { MessageInputUI } from '@ui';
import { useSelector } from '@store';
import { getCurrentChatType, getProgress } from '@slices';
import { ChatType } from '@utils-types';

export const MessageInput: FC<TMessageInputProps> = ({
  onSendMessage,
  onSendFile
}) => {
  const [message, setMessage] = useState<string>(''); // Состояние для текста
  const multiRef = useRef<MultiRefHandle>(null);
  const progress = useSelector(getProgress);
  const chatType = useSelector(getCurrentChatType);
  useEffect(() => {}, [chatType]);
  console.log(chatType);
  // Функция, которая обновляет состояние при изменении textarea
  const handleChange = (event: ChangeEvent<HTMLTextAreaElement>) => {
    setMessage(event.target.value);
  };

  const handleSend = () => {
    if (message.trim()) {
      // Проверка на пустой ввод
      onSendMessage(message); // Отправка сообщения
      setMessage(''); // Очистка после отправки
    }
  };
  // Подстраивает высоту textarea под содержимое
  useEffect(() => {
    if (multiRef.current?.textRef) {
      multiRef.current?.textRef
        ? (multiRef.current.textRef.style.height = 'auto')
        : undefined;
      multiRef.current?.textRef
        ? (multiRef.current.textRef.style.height = `${multiRef.current.textRef.scrollHeight}px`)
        : undefined;
    }
  }, [message]);

  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    setSelectedFile(file || null);
  };

  const handleClick = () => {
    if (multiRef.current?.fileRef) {
      multiRef.current?.fileRef.click(); // Триггерим клик по скрытому инпуту
    }
  };

  const handleUpload = () => {
    if (selectedFile) {
      onSendFile(selectedFile);
    }
    setSelectedFile(null); // Очистка после отправки
  };
  const onSend = chatType === ChatType.typeChat ? handleSend : handleUpload;
  const handleKeyDown = (event: KeyboardEvent<HTMLTextAreaElement>) => {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault(); // Отключаем перенос строки на Enter
      onSend();
    }
  };

  return (
    <MessageInputUI
      ref={multiRef}
      chatType={chatType}
      selectedFile={selectedFile}
      progress={progress}
      message={message}
      handleChange={handleChange}
      handleKeyDown={handleKeyDown}
      handleSend={onSend}
      handleClickFile={handleClick}
      handleFileChange={handleFileChange}
    />
  );
};
