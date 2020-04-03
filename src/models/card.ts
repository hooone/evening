import { getDvaApp, getLocale } from 'umi';
import reqwest from 'reqwest'
import { EffectsCommandMap, SubscriptionAPI } from 'dva'
import { ICard } from '@/interfaces'
import { IStore } from '@/store'
import { Datetime2Offset, Date2Offset, Offset2Date, Offset2Datetime } from '@/util'

export interface cardStateProps extends ICard {
}
export interface CardContentProps {
    onAction: Function,
    card: ICard,
}

interface loadProps {
    cardId: number,
    card: ICard,
    type: string,
}

const CardModel = {
    namespace: 'card',
    state: {},
    reducers: {
        saveData(state: cardStateProps, action: ICard) {
            let newState = Object.assign(state)
            newState.data = action.data
            return newState
        },
        saveCard(state: cardStateProps, action: loadProps) {
            action.card.data = []
            return action.card
        }
    },
    effects: {
        *loadCard(action: loadProps, handler: EffectsCommandMap) {
            const data = yield handler.call(reqwest, {
                url: '/api/card/getById'
                , type: 'json'
                , method: 'post'
                , data: { CardId: action.cardId }
            });
            let card: ICard = data.Data
            if (card !== null && card.Actions !== undefined) {
                card.Actions.forEach(
                    action => {
                        if (action.Type === "READ") {
                            card.Reader = action;
                        }
                    }
                )
            }
            yield handler.put({ type: 'saveCard', card: card });
            yield handler.put({ type: 'loadData', card: card });
        },
        *loadData(action: loadProps, handler: EffectsCommandMap) {
            let card = action.card
            if (card === undefined) {
                let namespace = action.type.split('/')[0]
                let app = getDvaApp();
                card = yield handler.select((state: IStore) => state[namespace]);
            }
            if (!card.Id) {
                return;
            }
            //load card data
            let params = {
                __CardId: card.Id
            }
            if (card.Reader === undefined) {
                return
            }
            card.Reader.Parameters.forEach((parameter, idx) => {
                if (parameter.Field.Type === 'datetime') {
                    params['param' + idx] = Offset2Datetime(parameter.Default)
                }
                else if(parameter.Field.Type==='date'){
                    params['param' + idx] = Offset2Date(parameter.Default)
                }
                else {
                    params['param' + idx] = parameter.Default
                }
            })
            const data = yield handler.call(reqwest, {
                url: '/data/read'
                , type: 'json'
                , method: 'post'
                , data: { lang: getLocale(), ...params }
            });
            yield handler.put({ type: 'saveData', data: data.Data });
        }
    },
    subscriptions: {
        setup(handler: SubscriptionAPI) {
            handler.dispatch({
                type: 'loadData',
            });
        },
    },
}



export default CardModel