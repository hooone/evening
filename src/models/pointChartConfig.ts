
import { getLocale } from 'umi';
import { ICard, IValueChange, IField, IStyle } from '@/interfaces';
import { getLocaleText, AJAX } from '@/util';
import { EffectsCommandMap, SubscriptionAPI } from 'dva'

export interface pointChartStateProps {
    dirty: boolean,
    xField: number,
    y1Field: number,
    y1Color: string,
    y2Field: number,
    y2Color: string,
    Fields: IField[],
    Styles: IStyle[],
}

export interface showPointChartConfigCommand {
    type: string,
    Fields: IField[],
    Styles: IStyle[],
}
export interface confirmPointChartConfigCommand extends pointChartStateProps {
    type: string,
}

export default {
    namespace: 'pointChartConfig',
    state: {
        dirty: false,
        xField: 0,
        y1Field: 0,
        y1Color: '#fad248',
        y2Field: 0,
        y2Color: 'yellow',
        Fields: [],
        Styles: [],
    },
    reducers: {
        show(state: pointChartStateProps, action: showPointChartConfigCommand) {
            let newState = {
                dirty: false,
                xField: 0,
                y1Field: 0,
                y1Color: '#fad248',
                y2Field: 0,
                y2Color: 'yellow',
                Fields: action.Fields ? action.Fields : [],
                Styles: action.Styles ? action.Styles.filter(f => f.Type === "POINT") : [],
            }
            if (!action.Styles) {
                return newState;
            }
            action.Styles.forEach(f => {
                if (f.Type === "POINT" && f.Property === "XAXIS") {
                    newState.xField = f.FieldId
                }
                else if (f.Type === "POINT" && f.Property === "Y1AXIS") {
                    newState.y1Field = f.FieldId
                    newState.y1Color = f.Value
                }
                else if (f.Type === "POINT" && f.Property === "Y2AXIS") {
                    newState.y2Field = f.FieldId
                    newState.y2Color = f.Value
                }
            })
            return newState;
        },
        dirty(state: pointChartStateProps, action: IValueChange) {
            let newState = Object.assign(state)
            newState.dirty = true
            newState[action.name] = action.value
            return newState;
        },
        save(state: pointChartStateProps, action: IValueChange) {
            let newState = Object.assign(state)
            newState.dirty = false
            return newState;
        },
    },
    effects: {
        *confirm(action: confirmPointChartConfigCommand, handler: EffectsCommandMap) {
            action.Styles.forEach(f => {
                if (f.Type === "POINT" && f.Property === "XAXIS") {
                    f.FieldId = action.xField
                }
                else if (f.Type === "POINT" && f.Property === "Y1AXIS") {
                    f.FieldId = action.y1Field
                    f.Value = action.y1Color
                }
                else if (f.Type === "POINT" && f.Property === "Y2AXIS") {
                    f.FieldId = action.y2Field
                    f.Value = action.y2Color
                }
            })
            const data = yield handler.call(AJAX, '/api/style/update', { Data: JSON.stringify(action.Styles) });
            if (!data.Success) {
                return;
            }
            if (data.Data === 0) {
                yield handler.put({ type: 'save' });
            }

        }
    },
};