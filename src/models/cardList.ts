import { getLocale } from 'umi'
import { EffectsCommandMap } from 'dva'
import { ICard, IMove } from '@/interfaces'
import { getLocaleText, AJAX } from '@/util';
import { IStore } from '@/store';

interface actionProps {
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
            const data = yield handler.call(AJAX, '/api/card/updateSeq', action.move);
            if (!data.Success) {
                return;
            }
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