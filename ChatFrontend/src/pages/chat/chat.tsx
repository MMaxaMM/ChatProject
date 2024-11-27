import { FC, useEffect, useState } from 'react';
import { ChatUI, ChatOpenUI } from '@ui-pages';
import { ChatList } from '@components';
import { TChat, TMessage, getChatTypeFromString } from '@utils-types';
import { useSelector, useDispatch } from '@store';
import {
  getStoreChats,
  sendMessage,
  createChat,
  getCurrentChatId,
  setChatId,
  getChats,
  postMessage,
  selectChatById,
  getChatHistory
} from '@slices';
import { useParams, useNavigate } from 'react-router-dom';

export const Chat: FC = () => {
  const params = useParams();
  const [isOpen, setIsOpen] = useState(true);
  const currentChatId = useSelector(getCurrentChatId);
  const currentChat = useSelector((state) =>
    selectChatById(state, currentChatId)
  );
  const [isOpenModal, setIsOpenModal] = useState(false);
  const [index, setIndex] = useState(parseInt(params.id ? params.id : '-1'));
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const toggleOpen = () => {
    setIsOpen(!isOpen);
  };

  const onCloseModal = () => {
    setIsOpenModal(false);
  };

  const onCreateChat = () => {
    setIsOpenModal(true);
  };

  const onSelectChats = (selectedChat: string) => {
    const chatType = getChatTypeFromString(selectedChat);
    dispatch(createChat(chatType));
    // dispatch(setChatId(-1));
  };

  useEffect(() => {
    dispatch(getChats());
    if (currentChatId !== -1) {
      setIndex(currentChatId);
    }
  }, [currentChatId]);

  useEffect(() => {
    if (index !== -1) {
      navigate(`/chat/${index}`);
      dispatch(setChatId(index));
      dispatch(getChatHistory(index));
    }
  }, [index]);

  const onSendMessage = (message: string) => {
    const data: TMessage = {
      role: 'user',
      content: message,
      isNew: false
    };
    navigate(`/chat/${currentChatId}`);
    const query = { chat_id: currentChatId, message: data };
    dispatch(sendMessage(query));
    dispatch(postMessage(query));
  };
  const chats: TChat[] = useSelector(getStoreChats);
  console.log(`chat ${currentChat}`);
  console.log(`chat id ${currentChatId}`);
  console.log(`index ${index}`);
  console.log(chats);
  return (
    <>
      <ChatList
        chats={chats}
        isOpenModal={isOpenModal}
        onCloseModal={onCloseModal}
        isOpen={isOpen}
        onClose={toggleOpen}
        onCreateChat={onCreateChat}
        onSelectChat={onSelectChats}
      />
      {index >= 0 && currentChat?.messages?.length ? (
        <ChatOpenUI
          isAsideOpen={isOpen}
          chat={currentChat}
          onOpenTab={toggleOpen}
          onSendMessage={onSendMessage}
        />
      ) : (
        <ChatUI
          isAsideOpen={isOpen}
          onSendMessage={onSendMessage}
          onOpenTab={toggleOpen}
        />
      )}
    </>
  );
};
