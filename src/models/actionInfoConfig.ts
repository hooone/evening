import reqwest from 'reqwest'
import { getLocale, getDvaApp } from 'umi'
import { EffectsCommandMap } from 'dva'
import { IViewAction, IValueChange } from '@/interfaces';

export interface actionInfoStateProps extends IViewAction {
    visible: boolean,
    dirty: boolean,
}
interface showProps {
    action: IViewAction,
}

export default {
    namespace: 'actionInfoConfig',
    state: {},
    reducers: {
        show(state: actionInfoStateProps, action: showProps) {
            return { visible: true, dirty: false, ...action.action }
        },
        dirty(state: actionInfoStateProps, action: IValueChange) {
            let newState = Object.assign(state)
            newState.dirty = true
            newState[action.name] = action.value
            return newState;
        },
        close(state: actionInfoStateProps) {
            return {
                visible: false,
                dirty: false,
                Id: 0,
            };
        },
    },
    effects: {
        *saveAction(action: showProps, handler: EffectsCommandMap) {
            let actions: IViewAction[] = []
            if (action.action.Id === 0) {
                //create action
                const data = yield handler.call(reqwest, {
                    url: '/api/action/create'
                    , type: 'json'
                    , method: 'post'
                    , data: { lang: getLocale(), ...action.action }
                });
                actions = data.Data
            }
            else {
                //update action
                const data = yield handler.call(reqwest, {
                    url: '/api/action/update'
                    , type: 'json'
                    , method: 'post'
                    , data: { lang: getLocale(), ...action.action }
                });
                actions = data.Data
            }
            //close sub drawer
            yield handler.put({ type: 'close' });
            yield handler.put({ type: 'draw/subClose' });
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
            })
            yield handler.put({ type: 'actionList/show', cardId: action.action.CardId, actions: actions });
        },
    }
};