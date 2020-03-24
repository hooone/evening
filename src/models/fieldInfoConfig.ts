import reqwest from 'reqwest'
import { getLocale } from 'umi'
import { EffectsCommandMap } from 'dva'
import { fieldInfoStateProps, IField, IValueChange } from '@/interfaces';

interface showProps {
    field: IField,
}


export default {
    namespace: 'fieldInfoConfig',
    state: {},
    reducers: {
        show(state: fieldInfoStateProps, action: showProps) {
            return { visible: true, dirty: false, ...action.field }
        },
        dirty(state: fieldInfoStateProps, action: IValueChange) {
            let newState = Object.assign(state)
            newState.dirty = true
            newState[action.name] = action.value
            return newState;
        },
        close(state: fieldInfoStateProps) {
            return {
                visible: false,
                dirty: false,
                Id: 0,
            };
        },
    },
    effects: {
        *saveField(action: showProps, handler: EffectsCommandMap) {
            let fields: IField[] = []
            if (action.field.Id === 0) {
                //create field
                const data = yield handler.call(reqwest, {
                    url: '/api/field/create'
                    , type: 'json'
                    , method: 'post'
                    , data: { lang: getLocale(), ...action.field }
                });
                fields = data.Data
            }
            else {
                //update field
                const data = yield handler.call(reqwest, {
                    url: '/api/field/update'
                    , type: 'json'
                    , method: 'post'
                    , data: { lang: getLocale(), ...action.field }
                });
                fields = data.Data
            }
            //close sub drawer
            yield handler.put({ type: 'close' });
            yield handler.put({ type: 'draw/subClose' });
            //flash field list
            let lang = getLocale()
            fields.forEach(f => {
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
            yield handler.put({ type: 'fieldList/show', cardId: action.field.CardId, fields: fields });
        },
    }
};