import { getLocale } from 'umi'
import { EffectsCommandMap } from 'dva'
import { IField, IValueChange, IMove } from '@/interfaces';
import { getLocaleText, AJAX } from '@/util';

interface showProps {
    field: IField,
}
interface updateSeqProps {
    move: IMove,
    cardId: number,
}
export interface fieldListStateProps {
    cardId: number,
    fields: IField[],
}

export default {
    namespace: 'fieldList',
    state: {
        cardId: 0,
        fields: [],
    },
    reducers: {
        show(state: fieldListStateProps, action: fieldListStateProps) {
            return {
                cardId: action.cardId,
                fields: action.fields,
            }
        },
    },
    effects: {
        *deleteField(action: showProps, handler: EffectsCommandMap) {
            if (action.field.Id !== 0) {
                //create field
                const data = yield handler.call(AJAX, '/api/field/delete', action.field);
                if (!data.Success) {
                    return;
                }
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
            const data = yield handler.call(AJAX, '/api/field/updateSeq', action.move);
            if (!data.Success) {
                return;
            }
            //flash field list
            let fields: IField[] = data.Data
            fields.forEach(f => {
                f.Text = getLocaleText(f.Locale)
            })
            yield handler.put({ type: 'show', cardId: action.cardId, fields: data.Data });
        },
    }
};