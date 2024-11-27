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

export const MessageInput: FC<TMessageInputProps> = ({ onSendMessage }) => {
  const [message, setMessage] = useState<string>(''); // Состояние для текста

  const multiRef = useRef<MultiRefHandle>(null);

  // Функция, которая обновляет состояние при изменении textarea
  const handleChange = (event: ChangeEvent<HTMLTextAreaElement>) => {
    setMessage(event.target.value);
  };

  const handleKeyDown = (event: KeyboardEvent<HTMLTextAreaElement>) => {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault(); // Отключаем перенос строки на Enter
      handleSend();
    }
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
  return (
    <MessageInputUI
      ref={multiRef}
      message={message}
      handleChange={handleChange}
      handleKeyDown={handleKeyDown}
      handleSend={handleSend}
      handleClickFile={handleClick}
      handleFileChange={handleFileChange}
    />
  );
};
