import { getLocale } from 'umi';
import { EffectsCommandMap } from 'dva'
import reqwest from 'reqwest'
import { getLocaleText } from '@/util'
import { ICard } from '@/interfaces';
import { showRectChartConfigCommand } from './rectChartConfig'
import { showPointChartConfigCommand } from './pointChartConfig'

export interface ILoadDrawCommand {
    type: string,
    cardId: number,
}

export interface drawStateProps {
    cardId: number,
    cardChanged: boolean,
    visible: boolean,
    subVisible: boolean,
    styleVisible: boolean,
    title: string,
    subTitle: string,
}
export default {
    namespace: 'draw',
    state: {
        cardId: 0,
        cardChanged: false,
        visible: false,
        subVisible: false,
        styleVisible: false,
        title: "",
        subTitle: "",
    },
    reducers: {
        open(state: drawStateProps, action: drawStateProps) {
            return {
                cardId: action.cardId,
                cardChanged: false,
                visible: true,
                subVisible: false,
                styleVisible: action.styleVisible,
                title: action.title,
                subTitle: "",
            }
        },
        subOpen(state: drawStateProps, action: drawStateProps) {
            return {
                cardId: state.cardId,
                cardChanged: state.cardChanged,
                visible: state.visible,
                subVisible: true,
                title: state.title,
                styleVisible: state.styleVisible,
                subTitle: action.title,
            }
        },
        close(state: drawStateProps, action: drawStateProps) {
            return {
                cardId: 0,
                cardChanged: false,
                visible: false,
                subVisible: false,
                styleVisible: false,
                title: "",
                subTitle: "",
            }
        },
        subClose(state: drawStateProps, action: drawStateProps) {
            return {
                cardId: state.cardId,
                cardChanged: state.cardChanged,
                visible: state.visible,
                subVisible: false,
                styleVisible: state.styleVisible,
                title: state.title,
                subTitle: "",
            }
        },
        cardChanged(state: drawStateProps, action: drawStateProps) {
            return {
                cardId: state.cardId,
                cardChanged: true,
                visible: state.visible,
                subVisible: state.subVisible,
                title: state.title,
                styleVisible: state.styleVisible,
                subTitle: state.subTitle,
            }
        },
        styleChanged(state: drawStateProps, action: drawStateProps) {
            return {
                cardId: state.cardId,
                cardChanged: state.cardChanged,
                visible: state.visible,
                subVisible: state.subVisible,
                title: state.title,
                styleVisible: action.styleVisible,
                subTitle: state.subTitle,
            }
        }
    },
    effects: {
        *loadDraw(action: ILoadDrawCommand, handler: EffectsCommandMap) {
            const data = yield handler.call(reqwest, {
                url: '/api/card/getById'
                , type: 'json'
                , method: 'post'
                , data: { lang: getLocale(), CardId: action.cardId }
            });
            let card: ICard = data.Data;
            card.Text = getLocaleText(card.Locale)
            card.Fields.forEach(f => {
                f.Text = getLocaleText(f.Locale)
            })
            card.Actions.forEach(a => {
                a.Text = getLocaleText(a.Locale)
                if (a.Parameters !== null) {
                    a.Parameters.forEach(p => {
                        p.Field.Text = getLocaleText(p.Field.Locale)
                    })
                }
            })
            yield handler.put({ type: 'cardInfoConfig/show', card: card });
            let showStyle: showRectChartConfigCommand = {
                type: 'rectChartConfig/show',
                Fields: card.Fields,
                Styles: card.Styles,
            }
            yield handler.put(showStyle);
            let showPointStyle: showPointChartConfigCommand = {
                type: 'pointChartConfig/show',
                Fields: card.Fields,
                Styles: card.Styles,
            }
            yield handler.put(showPointStyle);
            yield handler.put({ type: 'fieldList/show', cardId: card.Id, fields: card.Fields });
            yield handler.put({ type: 'actionList/show', cardId: card.Id, actions: card.Actions });
            yield handler.put({ type: 'open', title: card.Text, cardId: card.Id, styleVisible: card.Styles ? card.Styles.length > 0 : false });
        }
    },
};