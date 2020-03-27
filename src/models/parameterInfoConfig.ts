import { getLocale } from 'umi'
import reqwest from 'reqwest'
import { EffectsCommandMap } from 'dva'

import { IValueChange, IParameter, IViewAction } from '@/interfaces'
import { getLocaleText } from '@/util';

export interface parameterStateProps {
    cardId: number,
    visible: boolean,
    dirty: boolean,
    parameters: IParameter[],
}

export default {
    namespace: 'parameterInfoConfig',
    state: {
        cardId: 0,
        visible: false,
        dirty: false,
        parameters: []
    },
    reducers: {
        show(state: parameterStateProps, action: parameterStateProps) {
            return {
                cardId: action.cardId,
                visible: true,
                parameters: action.parameters
            }
        },
        dirty(state: parameterStateProps, action: IValueChange) {
            let newState: parameterStateProps = Object.assign(state)
            newState.dirty = true
            newState.parameters.forEach(p => {
                if (p.Id == action.id) {
                    p[action.name] = action.value
                }
            })
            return newState;
        },
        close(state: parameterStateProps) {
            return {
                visible: false,
                dirty: false,
                parameters: [],
            };
        },
    },
    effects: {
        *saveParameter(action: parameterStateProps, handler: EffectsCommandMap) {
            const data = yield handler.call(reqwest, {
                url: '/api/parameter/update'
                , type: 'json'
                , method: 'post'
                , data: { CardId: action.cardId, Data: JSON.stringify(action.parameters) }
            });
            let actions: IViewAction[] = data.Data
            let lang = getLocale()
            actions.forEach(f => {
                f.Text = getLocaleText(f.Locale)
                f.Parameters.forEach(p => {
                    p.Field.Text = getLocaleText(p.Field.Locale)
                })
            })
            yield handler.put({ type: 'close' });
            yield handler.put({ type: 'draw/subClose' });
            yield handler.put({ type: 'actionList/show', actions: actions, cardId: action.cardId });
        },
    },
};