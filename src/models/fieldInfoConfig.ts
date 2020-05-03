import { getLocale } from 'umi'
import { EffectsCommandMap } from 'dva'
import { IField, IValueChange } from '@/interfaces';
import { AJAX, Datetime2Offset, Date2Offset, Offset2Date, Offset2Datetime } from '@/util'

interface showProps {
    field: IField,
}

export interface fieldInfoStateProps extends IField {
    visible: boolean,
    dirty: boolean,
}

export default {
    namespace: 'fieldInfoConfig',
    state: {},
    reducers: {
        show(state: fieldInfoStateProps, action: showProps) {
            if (action.field.Type == "datetime") {
                let sp = action.field.Default.split("||");
                let ofSp = sp.map(f => Offset2Datetime(f))
                action.field.Default = ofSp.join("||")
            }
            else if (action.field.Type == "date") {
                let sp = action.field.Default.split("||");
                let ofSp = sp.map(f => Offset2Date(f))
                action.field.Default = ofSp.join("||")
            }
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
            if (action.field.Type == "datetime") {
                let sp = action.field.Default.split("||");
                let ofSp = sp.map(f => Datetime2Offset(f))
                action.field.Default = ofSp.join("||")
            }
            else if (action.field.Type == "date") {
                let sp = action.field.Default.split("||");
                let ofSp = sp.map(f => Date2Offset(f))
                action.field.Default = ofSp.join("||")
            }
            if (action.field.Id === 0) {
                //create field          
                const data = yield handler.call(AJAX, '/api/field/create', action.field);
                if (!data.Success) {
                    return;
                }
                fields = data.Data
            }
            else {
                //update field   
                const data = yield handler.call(AJAX, '/api/field/update', action.field);
                if (!data.Success) {
                    return;
                }
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