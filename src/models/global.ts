import { setLocale } from 'umi';

export interface globalStateProps {
    PathName: string,
    lang: string,
    User: string,
}
export default {
    namespace: 'global',
    state: { PathName: "/", lang: "", User: "" },
    reducers: {
        savePathName(state: globalStateProps, action: globalStateProps) {
            let newState = JSON.parse(JSON.stringify(state))
            newState.PathName = action.PathName
            return newState
        },
        changeLocale(state: globalStateProps, action: globalStateProps) {
            let newState = JSON.parse(JSON.stringify(state))
            newState.lang = action.lang
            setLocale(action.lang, true);
            return newState
        },
        saveUser(state: globalStateProps, action: globalStateProps) {
            let newState = JSON.parse(JSON.stringify(state))
            newState.User = action.User
            return newState
        }
    },
};