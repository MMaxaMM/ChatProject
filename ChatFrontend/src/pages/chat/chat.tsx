import { FC, useEffect, useState } from 'react';
import { ChatUI, ChatOpenUI } from '@ui-pages';
import { ChatList } from '@components';
import { ChatType, TChat, TMessage, getChatTypeFromString } from '@utils-types';
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
  getChatHistory,
  postAudio,
  setChatType,
  getCurrentChatType,
  postVideo
} from '@slices';
import { useParams, useNavigate } from 'react-router-dom';

export const Chat: FC = () => {
  const params = useParams();
  const [isOpen, setIsOpen] = useState(true);
  const currentChatId = useSelector(getCurrentChatId);
  const currentChat = useSelector((state) =>
    selectChatById(state, currentChatId)
  );
  const currentChatType = currentChat?.chat_type
    ? currentChat.chat_type
    : ChatType.typeChat;
  const cT = useSelector(getCurrentChatType);
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
  };

  useEffect(() => {
    dispatch(getChats());
  }, []);

  useEffect(() => {
    if (currentChatId !== -1 && currentChat) {
      setIndex(currentChat.chat_id);
      dispatch(setChatType(currentChat.chat_type));
    }
  }, [currentChatId, currentChat]);
  console.log(currentChatId);
  console.log(currentChat);
  useEffect(() => {
    dispatch(getChatHistory(index));
    if (index !== -1) {
      navigate(`/chat/${index}`);
      dispatch(setChatId(index));
      dispatch(setChatType(currentChatType));
    }
  }, [index]);

  const onSendMessage = (message: string) => {
    const data: TMessage = {
      role: 'user',
      content: message,
      isNew: false,
      content_type: 1
    };
    if (currentChatId === -1) {
      dispatch(createChat(ChatType.typeChat));
    }
    const query = { chat_id: currentChatId, message: data };
    navigate(`/chat/${currentChatId}`);
    dispatch(sendMessage(query));
    dispatch(postMessage(query));
  };

  const onSendFile = (file: File) => {
    const formData = new FormData();
    if (currentChatType === ChatType.typeAudio) {
      formData.append('audio', file);
    } else {
      formData.append('video', file);
    }
    const query = {
      chat_id: currentChatId,
      formData: formData
    };
    const data: TMessage = {
      role: 'user',
      content: URL.createObjectURL(file),
      isNew: false,
      content_type: currentChatType === ChatType.typeAudio ? 2 : 3
    };
    dispatch(sendMessage({ chat_id: currentChatId, message: data }));
    currentChatType === ChatType.typeAudio
      ? dispatch(postAudio(query))
      : dispatch(postVideo(query));
  };
  const chats: TChat[] = useSelector(getStoreChats);
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
          onSendFile={onSendFile}
        />
      ) : (
        <ChatUI
          isAsideOpen={isOpen}
          onSendMessage={onSendMessage}
          onOpenTab={toggleOpen}
          onSendFile={onSendFile}
        />
      )}
    </>
  );
};
