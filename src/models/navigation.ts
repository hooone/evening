import { getLocale, addLocale } from 'umi';
import { EffectsCommandMap, SubscriptionAPI } from 'dva'
import { navMoveProps, IFolder } from '@/interfaces'
import { IStore } from '@/store'
import { getLocaleText, AJAX } from '@/util';

export interface navStateProps {
    Navs: IFolder[],
    Path: string,
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

            const data = yield handler.call(AJAX, '/api/navigation', {});
            if (!data.Success) {
                return;
            }

            let folders: IFolder[] = data.Data
            folders.forEach(folder => {
                if (folder.IsFolder) {
                    folder.Text = getLocaleText(folder.Locale)
                }
                if (folder.Pages) {
                    folder.Pages.forEach(page => {
                        page.Text = getLocaleText(page.Locale)
                    })
                }
            })
            yield handler.put({ type: 'saveNav', Navs: folders, Path: pathname, });
            yield handler.put({
                type: 'global/saveUser',
                User: data.Message,
            });
        },

        *moveNav(action: IMoveNavCommand, handler: EffectsCommandMap) {
            let nav = yield handler.select((state: IStore) => state.nav)

            const move = yield handler.call(AJAX, '/api/navigation/move', action);
            if (!move.Success) {
                return;
            }
            const data = yield handler.call(AJAX, '/api/navigation', {});
            if (!move.Success) {
                return;
            }

            let folders: IFolder[] = data.Data
            folders.forEach(folder => {
                if (folder.IsFolder) {
                    folder.Text = getLocaleText(folder.Locale)
                }
                if (folder.Pages) {
                    folder.Pages.forEach(page => {
                        page.Text = getLocaleText(page.Locale)
                    })
                }
            })
            yield handler.put({ type: 'loadNav', Navs: folders, Path: nav.Path, });
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
                console.log("history listen")
                console.log(handler.history.location.pathname)
                handler.dispatch({
                    type: 'page/loadPage',
                    Name: handler.history.location.pathname,
                });
            });
        },
    },
};