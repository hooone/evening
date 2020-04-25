
import { SubscriptionAPI } from 'dva'

export interface NavigationModel {
    path: string,
    collapsed: boolean,
}


export default {
    namespace: 'navigation',
    state: { path: "/", collapsed: false },
    reducers: {
        savePath(state: NavigationModel, action: any) {
            return {
                path: action.Path,
                collapsed: state.collapsed,
            }
        },
        onCollapse(state: NavigationModel, action: any) {
            return {
                path: state.path,
                collapsed: !state.collapsed,
            }
        },
    },
    effects: {
    },
    subscriptions: {
        setup(handler: SubscriptionAPI) {
            handler.history.listen(location => {
                handler.dispatch({
                    type: 'savePath',
                    Path: handler.history.location.pathname,
                });
            });
        },
    },
};