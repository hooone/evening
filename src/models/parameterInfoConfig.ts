import { getLocale } from 'umi'
import { EffectsCommandMap } from 'dva'
import { Datetime2Offset, Date2Offset, Offset2Date, Offset2Datetime } from '@/util'
import { IValueChange, IParameter, IViewAction, CommonResult } from '@/interfaces'
import { getLocaleText, AJAX } from '@/util';

export interface parameterStateProps {
    cardId: number,
    actionId: number,
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
            action.parameters.forEach(param => {
                if (param.Field.Type == "datetime") {
                    let sp = param.Default.split("||");
                    let ofSp = sp.map(f => Offset2Datetime(f))
                    param.Default = ofSp.join("||")
                }
                else if (param.Field.Type == "date") {
                    let sp = param.Default.split("||");
                    let ofSp = sp.map(f => Offset2Date(f))
                    param.Default = ofSp.join("||")
                }
            })
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
        *load(action: parameterStateProps, handler: EffectsCommandMap) {
            yield handler.put({ type: 'close' });

            const data = yield handler.call(AJAX, '/api/parameter', { ActionId: action.actionId });
            if (!data.Success) {
                return;
            }

            let params: IParameter[] = data.Data;
            params.forEach(param => {
                param.Field.Text = getLocaleText(param.Field.Locale)
            })
            if (data.Success) {
                yield handler.put({
                    type: 'show',
                    parameters: data.Data,
                    cardId: action.cardId,
                });
            }
        },
        *saveParameter(action: parameterStateProps, handler: EffectsCommandMap) {
            action.parameters.forEach(param => {
                if (param.Field.Type == "datetime") {
                    let sp = param.Default.split("||");
                    let ofSp = sp.map(f => Datetime2Offset(f))
                    param.Default = ofSp.join("||")
                }
                else if (param.Field.Type == "date") {
                    let sp = param.Default.split("||");
                    let ofSp = sp.map(f => Date2Offset(f))
                    param.Default = ofSp.join("||")
                }
            })

            const data = yield handler.call(AJAX, '/api/parameter/update', { CardId: action.cardId, Data: JSON.stringify(action.parameters) });
            if (!data.Success) {
                return;
            }

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