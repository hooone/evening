
import { getDvaApp, getLocale } from 'umi';
import modelExtend from 'dva-model-extend';
import { EffectsCommandMap } from 'dva'
import CardModel from "./card"
import { IPage, ICard, IFolder } from '@/interfaces'
import { IStore } from '@/store'
import { AJAX } from '@/util'

export interface pageStateProps extends IPage {

}
export default {
    namespace: 'page',
    state: { Id: 0, Name: "", Cards: [] },
    reducers: {
        savePage(state: pageStateProps, action: pageStateProps) {
            return {
                Id: action.Id,
                Name: action.Name,
                Cards: action.Cards
            }
        },
    },
    effects: {
        *loadPage(action: pageStateProps, handler: EffectsCommandMap) {
            let cards: ICard[] = []

            let pageName = ''

            let app = getDvaApp();
            let pathname = app._history.location.pathname
            let route = pathname.split('/')
            route = route.filter((str: string) => str)
            pageName = route.pop() ?? ""
            if (!pageName) {
                return;
            }
            yield handler.put({ type: 'global/savePathName', PathName: pathname, });

            const data = yield handler.call(AJAX, '/api/card', {
                PageID: 0,
                PageName: pageName
            });
            if (!data.Success) {
                yield handler.put({ type: 'savePage', Cards: [], Id: 0, Name: '' });
                return;
            }
            cards = data.Data
            const states = yield handler.select((state: IStore) => state);
            let lang = getLocale();
            for (var i = 0; i < cards.length; i++) {
                let card = cards[i]
                if (card !== null && card.Actions !== undefined) {
                    card.Actions.forEach(
                        action => {
                            if (action.Type === "READ") {
                                card.Reader = action;
                            }
                        }
                    )
                }
                if (card.Locale[lang] === undefined || card.Locale[lang] === "") {
                    if (card.Locale["default"] !== undefined) {
                        card.Text = card.Locale["default"]
                    }
                    else {
                        card.Text = card.Name
                    }
                }
                else {
                    card.Text = card.Locale[lang]
                }
                if (states['card_' + card.Id] === undefined) {
                    let newModel = modelExtend(CardModel, { namespace: 'card_' + card.Id })
                    newModel.state = card
                    app.model(newModel)
                }
            };

            yield handler.put({ type: 'cardList/show', cards: cards });
            yield handler.put({ type: 'savePage', Cards: cards, Id: 0, Name: pageName });
        },
    },
};

