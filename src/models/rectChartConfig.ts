
import { getLocale } from 'umi';
import { ICard, IValueChange, IField, IStyle } from '@/interfaces';
import { getLocaleText, AJAX } from '@/util';
import { EffectsCommandMap, SubscriptionAPI } from 'dva'

export interface rectChartStateProps {
    dirty: boolean,
    xField: number,
    y1Field: number,
    y1Type: 'BAR' | 'LINE',
    y1Color: string,
    y2Field: number,
    y2Type: 'BAR' | 'LINE',
    y2Color: string,
    Fields: IField[],
    Styles: IStyle[],
}

export interface showRectChartConfigCommand {
    type: string,
    Fields: IField[],
    Styles: IStyle[],
}
export interface confirmRectChartConfigCommand extends rectChartStateProps {
    type: string,
}

export default {
    namespace: 'rectChartConfig',
    state: {
        dirty: false,
        xField: 0,
        y1Field: 0,
        y1Type: 'BAR',
        y1Color: '#fad248',
        y2Field: 0,
        y2Type: 'LINE',
        y2Color: 'yellow',
        Fields: [],
        Styles: [],
    },
    reducers: {
        show(state: rectChartStateProps, action: showRectChartConfigCommand) {
            let newState = {
                dirty: false,
                xField: 0,
                y1Field: 0,
                y1Type: 'BAR',
                y1Color: '#fad248',
                y2Field: 0,
                y2Type: 'LINE',
                y2Color: 'yellow',
                Fields: action.Fields ? action.Fields : [],
                Styles: action.Styles ? action.Styles.filter(f => f.Type === "RECT") : [],
            }
            if (!action.Styles) {
                return newState;
            }
            action.Styles.forEach(f => {
                if (f.Type === "RECT" && f.Property === "XAXIS") {
                    newState.xField = f.FieldId
                }
                else if (f.Type === "RECT" && f.Property === "Y1AXIS") {
                    newState.y1Field = f.FieldId
                    newState.y1Type = f.Value
                }
                else if (f.Type === "RECT" && f.Property === "Y2AXIS") {
                    newState.y2Field = f.FieldId
                    newState.y2Type = f.Value
                }
                else if (f.Type === "RECT" && f.Property === "Y1COLOR") {
                    newState.y1Color = f.Value
                }
                else if (f.Type === "RECT" && f.Property === "Y2COLOR") {
                    newState.y2Color = f.Value
                }
            })
            return newState;
        },
        dirty(state: rectChartStateProps, action: IValueChange) {
            let newState = Object.assign(state)
            newState.dirty = true
            newState[action.name] = action.value
            return newState;
        },
        save(state: rectChartStateProps, action: IValueChange) {
            let newState = Object.assign(state)
            newState.dirty = false
            return newState;
        },
    },
    effects: {
        *confirm(action: confirmRectChartConfigCommand, handler: EffectsCommandMap) {
            action.Styles.forEach(f => {
                if (f.Type === "RECT" && f.Property === "XAXIS") {
                    f.FieldId = action.xField
                }
                else if (f.Type === "RECT" && f.Property === "Y1AXIS") {
                    f.FieldId = action.y1Field
                    f.Value = action.y1Type
                }
                else if (f.Type === "RECT" && f.Property === "Y2AXIS") {
                    f.FieldId = action.y2Field
                    f.Value = action.y2Type
                }
                else if (f.Type === "RECT" && f.Property === "Y1COLOR") {
                    f.Value = action.y1Color
                }
                else if (f.Type === "RECT" && f.Property === "Y2COLOR") {
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