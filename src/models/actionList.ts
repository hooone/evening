import reqwest from 'reqwest'
import { message } from 'antd'
import { getLocale } from 'umi'
import { EffectsCommandMap } from 'dva'
import { IViewAction, IMove, CommonResult } from '@/interfaces'


export interface actionListStateProps {
    cardId: number,
    actions: IViewAction[],
}

interface actionProps {
    action: IViewAction,
}

interface moveActionProps {
    cardId: number,
    move: IMove,
}

export default {
    namespace: 'actionList',
    state: {
        cardId: 0,
        actions: []
    },
    reducers: {
        show(state: actionListStateProps, action: actionListStateProps) {
            return {
                cardId: action.cardId,
                actions: action.actions.filter(f => f.Type !== "READ")
            }
        },
    },
    effects: {
        *deleteAction(action: actionProps, handler: EffectsCommandMap) {
            if (action.action.Id !== 0) {
                //create action
                const data: CommonResult = yield handler.call(reqwest, {
                    url: '/api/action/delete'
                    , type: 'json'
                    , method: 'post'
                    , data: action.action
                });
                if (!data.Success) {
                    if (data.Message) {
                        message.error(data.Message);
                    }
                    return;
                }
                let actions: IViewAction[] = []
                actions = data.Data;
                //flash action list
                let lang = getLocale()
                actions.forEach(f => {
                    if (f.Locale[lang] === undefined || f.Locale[lang] === "") {
                        if (f.Locale["default"] !== undefined) {
                            f.Text = f.Locale["default"]
                        }
                        else {
                            f.Text = f.Name
                        }
                    }
                    else {
                        f.Text = f.Locale[lang]
                    }

                    f.Parameters.forEach(p => {
                        if (p.Field.Locale[lang] === undefined || p.Field.Locale[lang] === "") {
                            if (p.Field.Locale["default"] !== undefined) {
                                p.Field.Text = f.Locale["default"]
                            }
                            else {
                                p.Field.Text = f.Name
                            }
                        }
                        else {
                            p.Field.Text = f.Locale[lang]
                        }
                    })
                })
                yield handler.put({ type: 'show', cardId: action.action.CardId, actions: actions });
            }
        },
        *updateSeq(action: moveActionProps, handler: EffectsCommandMap) {
            //move action
            const data = yield handler.call(reqwest, {
                url: '/api/action/updateSeq'
                , type: 'json'
                , method: 'post'
                , data: action.move
            });
            //flash action list
            let lang = getLocale()
            let actions: IViewAction[] = []
            actions = data.Data
            actions.forEach(f => {
                if (f.Locale[lang] === undefined || f.Locale[lang] === "") {
                    if (f.Locale["default"] !== undefined) {
                        f.Text = f.Locale["default"]
                    }
                    else {
                        f.Text = f.Name
                    }
                }
                else {
                    f.Text = f.Locale[lang]
                }
                f.Parameters.forEach(p => {
                    if (p.Field.Locale[lang] === undefined || p.Field.Locale[lang] === "") {
                        if (p.Field.Locale["default"] !== undefined) {
                            p.Field.Text = f.Locale["default"]
                        }
                        else {
                            p.Field.Text = f.Name
                        }
                    }
                    else {
                        p.Field.Text = f.Locale[lang]
                    }
                })
            })
            yield handler.put({ type: 'show', cardId: action.cardId, actions: actions });
        },
    }
};