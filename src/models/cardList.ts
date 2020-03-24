import reqwest from 'reqwest'
import { getLocale } from 'umi'
import { EffectsCommandMap } from 'dva'
import { ICard, IStore, IMove } from '@/interfaces'
import { getLocaleText } from '@/util';

interface actionProps {
    cardId: number,
    move: IMove,
    cards: ICard[],
}

export default {
    namespace: 'cardList',
    state: [],
    reducers: {
        show(state: ICard[], action: actionProps) {
            return action.cards
        },
    },
    effects: {
        *updateSeq(action: actionProps, handler: EffectsCommandMap) {
            yield handler.put({ type: 'draw/cardChanged' });
            //move card
            const data = yield handler.call(reqwest, {
                url: '/api/card/updateSeq'
                , type: 'json'
                , method: 'post'
                , data: action.move
            });
            //flash card list
            let cards: ICard[] = []
            cards = data.Data
            cards.forEach(f => {
                f.Text = getLocaleText(f.Locale)
            })
            yield handler.put({ type: 'show', cards: cards });
        },
    }
};