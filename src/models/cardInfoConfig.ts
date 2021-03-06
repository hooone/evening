
import { getLocale } from 'umi';
import { ICard, IValueChange } from '@/interfaces';
import { getLocaleText, AJAX } from '@/util';
import { EffectsCommandMap, SubscriptionAPI } from 'dva'
import { showRectChartConfigCommand } from './rectChartConfig'
import { showPointChartConfigCommand } from './pointChartConfig'

export interface cardInfoStateProps extends ICard {
  dirty: boolean,
}
interface actionProps {
  card: ICard,
}
export default {
  namespace: 'cardInfoConfig',
  state: { dirty: false },
  reducers: {
    show(state: cardInfoStateProps, action: actionProps) {
      return { dirty: false, ...action.card };
    },
    dirty(state: cardInfoStateProps, action: IValueChange) {
      let newState = Object.assign(state)
      newState.dirty = true
      newState[action.name] = action.value
      return newState;
    },
  },
  effects: {
    *confirm(action: actionProps, handler: EffectsCommandMap) {
      const data = yield handler.call(AJAX, '/api/card/update', {
        Id: action.card.Id,
        Name: action.card.Name,
        Style: action.card.Style,
        Text: action.card.Text,
        Width: action.card.Width,
        Pos: action.card.Pos,
      });
      if (!data.Success) {
        return;
      }
      let card: ICard = data.Data;
      card.Text = getLocaleText(card.Locale)

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
      yield handler.put({ type: 'draw/styleChanged', styleVisible: card.Styles ? card.Styles.length > 0 : false });
      yield handler.put({ type: 'show', card: card });

    }
  },
};