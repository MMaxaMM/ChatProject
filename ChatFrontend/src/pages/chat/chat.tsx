import { FC, useEffect, useState } from 'react';
import { ChatUI, ChatOpenUI } from '@ui-pages';
import { ChatList } from '@components';
import { ChatType, TChat, TMessage, getChatTypeFromString } from '@utils-types';
import { useSelector, useDispatch } from '@store';
import {
  sendMessage,
  createChat,
  getCurrentChatId,
  setChatId,
  getChats,
  postMessage,
  selectChatById,
  getChatHistory,
  postAudio,
  postVideo,
  postRAGMessage,
  refreshUsername
} from '@slices';
import { useParams, useNavigate } from 'react-router-dom';

export const Chat: FC = () => {
  const params = useParams();
  const [isOpen, setIsOpen] = useState(true);
  const currentChatId = useSelector(getCurrentChatId);
  const currentChat = useSelector((state) =>
    selectChatById(state, currentChatId)
  );
  const currentChatType = currentChat?.chat_type;
  const [isOpenModal, setIsOpenModal] = useState(false);
  const [isOpenUserModal, setIsOpenUserModal] = useState(false);
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

  const onOpenUserModal = () => {
    setIsOpenUserModal(true);
  };

  const onCloseUserModal = () => {
    setIsOpenUserModal(false);
  };

  const onSelectChats = (selectedChat: string) => {
    const chatType = getChatTypeFromString(selectedChat);
    dispatch(createChat(chatType));
  };

  useEffect(() => {
    dispatch(getChats());
    dispatch(refreshUsername());
  }, []);

  useEffect(() => {
    if (currentChatId !== -1 && currentChat) {
      setIndex(currentChat.chat_id);
    }
  }, [currentChatId]);

  useEffect(() => {
    if (index !== -1) {
      dispatch(getChatHistory(index));
      navigate(`/chat/${index}`);
      console.log(`ffff${index}`);
      dispatch(setChatId(index));
    }
  }, [index]);

  const onSendMessage = async (message: string) => {
    const data: TMessage = {
      role: 'user',
      content: message,
      isNew: false,
      content_type: 1
    };
    const post = ChatType.typeRAG ? postRAGMessage : postMessage;
    if (currentChatId === -1) {
      const res = await dispatch(createChat(ChatType.typeChat)).unwrap();
      const query = { chat_id: res.chat_id, message: data };
      dispatch(sendMessage(query));
      await dispatch(post(query));
    } else {
      const query = { chat_id: currentChatId, message: data };
      dispatch(sendMessage(query));
      await dispatch(post(query));
    }
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
    console.log(currentChatType);
    console.log(data);
    dispatch(sendMessage({ chat_id: currentChatId, message: data }));
    currentChatType === ChatType.typeAudio
      ? dispatch(postAudio(query))
      : dispatch(postVideo(query));
  };
  return (
    <>
      <ChatList
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
          isUserModalOpen={isOpenUserModal}
          onCloseUserModal={onCloseUserModal}
          onOpenUserModal={onOpenUserModal}
          onOpenTab={toggleOpen}
          onSendMessage={onSendMessage}
          onSendFile={onSendFile}
        />
      ) : (
        <ChatUI
          isAsideOpen={isOpen}
          isUserModalOpen={isOpenUserModal}
          onCloseUserModal={onCloseUserModal}
          onOpenUserModal={onOpenUserModal}
          onSendMessage={onSendMessage}
          onOpenTab={toggleOpen}
          onSendFile={onSendFile}
        />
      )}
    </>
  );
};
