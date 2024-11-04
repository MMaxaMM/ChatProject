import { FC, useState } from 'react';
import { ChatOpenUI } from '@ui-pages';
import { ChatList } from '@components';
import { TChat } from '@utils-types';
import { useParams } from 'react-router-dom';

export const ChatOpen: FC = () => {
  const [isOpen, setIsOpen] = useState(true);
  const params = useParams();
  const index = parseInt(params.id ? params.id : '0');

  const toggleOpen = () => {
    setIsOpen(!isOpen);
  };

  const onCreateChat = () => void 0;

  const chats: TChat[] = [
    {
      userId: 0,
      chatId: 0,
      messages: [
        {
          role: 'user',
          content: 'расскажи сказку'
        },
        {
          role: 'assistent',
          content: `Жила-была однажды в далёком лесу маленькая лисичка по имени Ляля. Она была не как все — хоть и была очень ловкая, резвая и хитрая, но её сердечко было доброе и совсем не любило хитрить. В отличие от других лесных жителей, Ляля всегда старалась помогать тем, кто попадал в беду.

Однажды весной, когда снег уже почти сошёл, а деревья зазеленели, Ляля услышала в лесу странный звук — будто кто-то жалобно пищал. Она побежала на звук и вскоре увидела птенца, который сидел под кустом. Он упал из гнезда, и теперь был один-одинёшенек, испуганный и дрожащий.

Лисичка тихонько подошла к птенцу, чтобы не напугать его ещё сильнее, и спросила: — Ты потерялся?

Птенец поднял головку и тихо пискнул: — Да, я выпал из гнезда, а летать ещё не умею. Как же мне вернуться домой?

Ляля задумалась. Она понимала, что ей самой не добраться до гнезда на высокой сосне, но у неё была идея. Она помчалась по лесу, чтобы найти своих друзей. Сперва Ляля заглянула к белке Соне, которая прекрасно прыгала по деревьям. Белка охотно согласилась помочь.

Ляля вернулась к птенцу с Соней, и та ловко взяла птенчика, прыгая с ветки на ветку, пока не добралась до его гнезда. Птенец был так счастлив, что теперь был в безопасности.

— Спасибо тебе, Ляля! — пискнул он, — И тебе, Соня!

Когда Ляля и Соня спустились на землю, они почувствовали себя настоящими героями. С тех пор, если кому-то в лесу нужна была помощь, жители знали, к кому обратиться. Ляля и её друзья
Жила-была однажды в далёком лесу маленькая лисичка по имени Ляля. Она была не как все — хоть и была очень ловкая, резвая и хитрая, но её сердечко было доброе и совсем не любило хитрить. В отличие от других лесных жителей, Ляля всегда старалась помогать тем, кто попадал в беду.

Однажды весной, когда снег уже почти сошёл, а деревья зазеленели, Ляля услышала в лесу странный звук — будто кто-то жалобно пищал. Она побежала на звук и вскоре увидела птенца, который сидел под кустом. Он упал из гнезда, и теперь был один-одинёшенек, испуганный и дрожащий.

Лисичка тихонько подошла к птенцу, чтобы не напугать его ещё сильнее, и спросила: — Ты потерялся?

Птенец поднял головку и тихо пискнул: — Да, я выпал из гнезда, а летать ещё не умею. Как же мне вернуться домой?

Ляля задумалась. Она понимала, что ей самой не добраться до гнезда на высокой сосне, но у неё была идея. Она помчалась по лесу, чтобы найти своих друзей. Сперва Ляля заглянула к белке Соне, которая прекрасно прыгала по деревьям. Белка охотно согласилась помочь.

Ляля вернулась к птенцу с Соней, и та ловко взяла птенчика, прыгая с ветки на ветку, пока не добралась до его гнезда. Птенец был так счастлив, что теперь был в безопасности.

— Спасибо тебе, Ляля! — пискнул он, — И тебе, Соня!

Когда Ляля и Соня спустились на землю, они почувствовали себя настоящими героями. С тех пор, если кому-то в лесу нужна была помощь, жители знали, к кому обратиться. Ляля и её друзья`
        }
      ]
    },
    {
      userId: 1,
      chatId: 1,
      messages: [
        {
          role: 'user',
          content: 'hello'
        }
      ]
    }
  ];

  return (
    <>
      <ChatList
        chats={chats}
        isOpen={isOpen}
        onClose={toggleOpen}
        onCreateChat={onCreateChat}
      />
      <ChatOpenUI
        isAsideOpen={isOpen}
        chat={chats[index]}
        onOpenTab={toggleOpen}
      />
    </>
  );
};
