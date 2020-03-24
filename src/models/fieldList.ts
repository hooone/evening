import reqwest from 'reqwest'
import { getLocale } from 'umi'
import { EffectsCommandMap } from 'dva'
import { fieldListStateProps, IField, IValueChange, IMove } from '@/interfaces';
import { getLocaleText } from '@/util';

interface showProps {
    field: IField,
}
interface updateSeqProps {
    move: IMove,
    cardId: number,
}

export default {
    namespace: 'fieldList',
    state: {
        cardId: 0,
        fieldList: []
    },
    reducers: {
        show(state: fieldListStateProps, action: fieldListStateProps) {
            return {
                cardId: action.cardId,
                fieldList: action.fields,
            }
        },
    },
    effects: {
        *deleteField(action: showProps, handler: EffectsCommandMap) {
            if (action.field.Id !== 0) {
                //create field
                const data = yield handler.call(reqwest, {
                    url: '/api/field/delete'
                    , type: 'json'
                    , method: 'post'
                    , data: action.field
                });
                //flash field list
                let fields: IField[] = data.Data
                fields.forEach(f => {
                    f.Text = getLocaleText(f.Locale)
                })
                yield handler.put({ type: 'show', cardId: action.field.CardId, fields: fields });
            }
        },
        *updateSeq(action: updateSeqProps, handler: EffectsCommandMap) {
            //move field
            const data = yield handler.call(reqwest, {
                url: '/api/field/updateSeq'
                , type: 'json'
                , method: 'post'
                , data: action.move
            });
            //flash field list
            let fields: IField[] = data.Data
            fields.forEach(f => {
                f.Text = getLocaleText(f.Locale)
            })
            yield handler.put({ type: 'show', cardId: action.cardId, fields: data.Data });
        },
    }
};