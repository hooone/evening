import { getLocale, addLocale } from 'umi';
import { EffectsCommandMap, SubscriptionAPI } from 'dva'
import reqwest from 'reqwest'
import { navMoveProps, IFolder } from '@/interfaces'
import { IStore } from '@/store'

export interface navStateProps {
    Navs: IFolder[],
    collapsed: boolean,
}
export interface ICollapseCommand {
    type: string,
}
export interface IMoveNavCommand {
    type: string,
    SourceFolder: number,
    SourcePage: number,
    TargetFolder: number,
    TargetPage: number,
    Position: number,
}
export default {
    namespace: 'nav',
    state: { Path: "/", Navs: [], collapsed: false },
    reducers: {
        onCollapse(state: navStateProps, action: navStateProps) {
            return {
                Path: state.Path,
                Navs: state.Navs,
                collapsed: !state.collapsed,
            };
        },
        saveNav(state: navStateProps, action: navStateProps) {
            return {
                Path: action.Path,
                Navs: action.Navs,
            }
        },
    },
    effects: {
        *loadNav(action: navStateProps, handler: EffectsCommandMap) {
            let pathname = action.Path
            if (!pathname) {
                let nav: navStateProps = yield handler.select(((state: IStore) => state.nav));
                pathname = nav.Path
            }
            const data = yield handler.call(reqwest, {
                url: '/api/navigation'
                , type: 'json'
                , method: 'post'
                , data: {}
            });
            yield handler.put({ type: 'saveNav', Navs: data.Data, Path: pathname, });
            console.log(data)
            yield handler.put({
                type: 'global/saveUser',
                User: data.Message,
            });
        },

        *moveNav(action: IMoveNavCommand, handler: EffectsCommandMap) {
            let nav = yield handler.select((state: IStore) => state.nav)
            const move = yield handler.call(reqwest, {
                url: '/api/navigation/move'
                , type: 'json'
                , method: 'post'
                , data: action
            });
            const data = yield handler.call(reqwest, {
                url: '/api/navigation'
                , type: 'json'
                , method: 'post'
                , data: {}
            });
            yield handler.put({ type: 'loadNav', Navs: data.Data, Path: nav.Path, });
        },
    },
    subscriptions: {
        setup(handler: SubscriptionAPI) {
            console.log('setup')
            console.log(window.location.href)
            if (handler.history.location.pathname.indexOf("/login") >= 0) {
                return;
            }
            handler.dispatch({
                type: 'loadNav',
                Path: handler.history.location.pathname,
            });
            console.log(getLocale())
            handler.dispatch({
                type: 'global/changeLocale',
                lang: getLocale(),
            });
            handler.history.listen(location => {
                console.log("history listem")
                console.log(handler.history.location.pathname)
                handler.dispatch({
                    type: 'page/loadPage',
                    Name: handler.history.location.pathname,
                });
            });
        },
    },
};