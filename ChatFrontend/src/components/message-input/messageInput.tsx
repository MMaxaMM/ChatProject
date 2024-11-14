import {
  FC,
  useEffect,
  useRef,
  ChangeEvent,
  KeyboardEvent,
  useState
} from 'react';
import { TMessageInputProps } from './type';
import { MessageInputUI } from '@ui';

export const MessageInput: FC<TMessageInputProps> = ({ onSendMessage }) => {
  const [message, setMessage] = useState<string>(''); // Состояние для текста
  const textareaRef = useRef<HTMLTextAreaElement>(null); // Реф для textarea

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
    if (textareaRef.current) {
      textareaRef.current.style.height = 'auto'; // Сбрасываем высоту перед перерасчетом
      textareaRef.current.style.height = `${textareaRef.current.scrollHeight}px`; // Устанавливаем высоту в зависимости от контента
    }
  }, [message]);
  return (
    <MessageInputUI
      ref={textareaRef}
      message={message}
      handleChange={handleChange}
      handleKeyDown={handleKeyDown}
      handleSend={handleSend}
    />
  );
};
