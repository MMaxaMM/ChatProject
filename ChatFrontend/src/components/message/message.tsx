import { FC, memo } from 'react';
import { MessageUI } from '@ui';
import { TMessageProps } from './type';

export const Message: FC<TMessageProps> = memo(({ message }) => (
  <MessageUI message={message} />
));
