import { getLocale } from 'umi';
import reqwest from 'reqwest'
import { EffectsCommandMap } from 'dva'
import { IViewAction } from '@/interfaces';

export interface modalStateProps {
    visible: boolean,
    confirmLoading: boolean,
    action: IViewAction,
    record: any,
}
export interface IShowActionCommand {
    type: string,
    action: IViewAction,
    record: any,
}
export interface IHandleOkCommand {
    type: string,
    action: IViewAction,
    formdata: any,
}
export interface ICancelCommand {
    type: 'modal/cancel',
}
interface testDataProps {
    __CardId: number,
    __Key: number,
    __ActionId: number,
}

export default {
    namespace: 'modal',
    state: {
        visible: false,
        confirmLoading: false,
        action: {},
        record: {},
    },
    reducers: {
        showAction(state: modalStateProps, action: modalStateProps) {
            return {
                visible: true,
                confirmLoading: false,
                action: action.action,
                record: action.record,
            };
        },
        loading(state: modalStateProps, action: modalStateProps) {
            let newState = Object.assign(state)
            newState.confirmLoading = action.confirmLoading
            newState.record = action.record
            return newState
        },
        cancel(state: modalStateProps) {
            return {
                visible: false,
                confirmLoading: false,
                action: {},
                record: {},
            };
        },
    },
    effects: {
        *handleOk(action: IHandleOkCommand, handler: EffectsCommandMap) {
            // if (this.state.action.Type === "MULTIUPDATE") {
            //     this.props.handleCancel();
            //     return;
            // }
            yield handler.put({ type: 'loading', confirmLoading: true, record: action.formdata });
            if (action.action.Type === "READ") {
                yield handler.put({
                    type: 'cancel',
                });
                action.action.Parameters.forEach(
                    (parameter, idx) => {
                        parameter.Default = action.formdata["param" + idx]
                    }
                )
                yield handler.put({
                    type: 'card_' + action.action.CardId + '/loadData',
                    card: {
                        Id: action.action.CardId,
                        Reader: action.action
                    }
                });
            }
            else {
                //执行动作
                let actionData: testDataProps = {
                    __CardId: 0,
                    __Key: 0,
                    __ActionId: 0,
                }
                if (action.action.CardId !== undefined && !isNaN(action.action.CardId)) {
                    actionData.__CardId = action.action.CardId
                }
                if (action.formdata.__Key !== undefined && !isNaN(action.formdata.__Key)) {
                    actionData.__Key = action.formdata.__Key
                }
                if (action.action.Id !== undefined && !isNaN(action.action.Id)) {
                    actionData.__ActionId = action.action.Id
                }
                let url = "";
                if (action.action.Type === "NAV" || action.action.Type === "CARD") {
                    url += action.action.URL
                }
                else {
                    url += ("/data/" + action.action.Type.toLowerCase())
                }
                const data = yield handler.call(reqwest, {
                    url: url,
                    type: 'json',
                    method: 'post',
                    data: { lang: getLocale(), ...actionData, ...action.formdata }
                });
                if (data.Success) {
                    yield handler.put({
                        type: 'cancel',
                    });
                    if (action.action.Type === "NAV") {
                        yield handler.put({
                            type: 'nav/loadNav',
                        });
                    }
                    else if (action.action.Type === "CARD") {
                        yield handler.put({
                            type: 'page/loadPage',
                            Id: action.formdata["PageID"],
                            Name: action.formdata["PageName"],
                        });
                    }
                    else if (action.action.CardId) {
                        yield handler.put({
                            type: 'card_' + action.action.CardId + '/loadData',
                        });
                    }
                }
                else {
                    yield handler.put({ type: 'loading', loading: false, record: action.formdata });
                }
            }

        }
    }
};