import React, { InputHTMLAttributes, ReactElement } from 'react';
import { List, Divider, Button, Modal } from 'antd';
import { connect, useIntl } from 'umi';
import { DispatchProp } from 'react-redux';
import { EyeOutlined, EyeInvisibleOutlined, TagOutlined, CodeOutlined, ExclamationCircleOutlined } from '@ant-design/icons';
import { actionListStateProps } from '@/models/actionList'
import { IViewAction } from '@/interfaces';
import { findAttribute } from '@/util';
import { IStore } from '@/store'

const { confirm } = Modal;

interface ActionListProps extends DispatchProp {
    actionList: actionListStateProps,
}
const ActionsList = (props: ActionListProps) => {
    const intl = useIntl();
    function createAction() {
        props.dispatch({
            type: 'actionInfoConfig/show',
            action: {
                Id: 0,
                CardId: props.actionList.cardId,
                Name: "",
                Type: "UPDATE",
                Seq: 0,
                DoubleCheck: false,
            },
        });
        props.dispatch({
            type: 'draw/subOpen',
            title: intl.formatMessage(
                {
                    id: 'create',
                }
            ),
        });
    }
    function updateAction(action: IViewAction) {
        props.dispatch({
            type: 'actionInfoConfig/show',
            action: action,
        });
        props.dispatch({
            type: 'draw/subOpen',
            title: action.Text,
        });
    }
    function updateParameter(action: IViewAction) {
        props.dispatch({
            type: 'parameterInfoConfig/load',
            actionId: action.Id,
            cardId: props.actionList.cardId,
        });
        props.dispatch({
            type: 'draw/subOpen',
            title: action.Text,
        });
    }
    function deleteAction(action: IViewAction) {
        confirm({
            title:
                (<span><span>
                    {intl.formatMessage(
                        {
                            id: 'confirm',
                        }
                    )}
                </span>
                    <span>
                        {intl.formatMessage(
                            {
                                id: 'delete',
                            }
                        )}
                    ?
                    </span>
                </span>)
            ,
            icon: <ExclamationCircleOutlined />,
            onOk() {
                props.dispatch({
                    type: 'actionList/deleteAction',
                    action: action,
                });
            },
            onCancel() {

            },
        });

    }

    function onDragStart(e: React.DragEvent) {
        let real = e.target as HTMLElement;
        if (real.tagName === "LI" && real.classList.contains("ant-list-item")) {
            e.dataTransfer.setData("id", findAttribute(real, "data-id"));
            e.dataTransfer.setData("type", findAttribute(real, "data-type"));
        }
    }
    function getRealDOM(s: HTMLElement) {
        let real = s;
        if (real.tagName === "path") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "svg") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "EM") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "A") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "I") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "LI" && !real.classList.contains("ant-list-item")) {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "UL") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "H4") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "DIV") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "DIV") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "DIV") {
            real = real.parentElement as HTMLElement
        }
        return real
    }
    function onDrop(e: React.DragEvent) {
        e.preventDefault()
        let formdata = {
            Source: 0,
            Target: 0,
            Position: 0,
            CardId: props.actionList.cardId,
        }
        let real = getRealDOM(e.target as HTMLElement)
        real.classList.add("dragEnterList")
        if (real.tagName === "LI" && real.classList.contains("ant-list-item")) {
            real.classList.remove("dragEnterList")
            formdata.Source = parseInt(e.dataTransfer.getData("id"));
            formdata.Target = parseInt(findAttribute(real, "data-id"));
            let sourceType = e.dataTransfer.getData("type");
            if (e.clientY - 71 - real.offsetTop < real.clientHeight / 2) {
                formdata.Position = 1;
            }
            else {
                formdata.Position = 2;
            }
            props.dispatch({
                type: 'actionList/updateSeq',
                move: formdata,
                cardId: props.actionList.cardId,
            });
        }

    }
    function onDragOver(e: React.DragEvent) {
        let real = getRealDOM(e.target as HTMLElement);
        if (real.tagName === "LI" && real.classList.contains("drawer-list")) {
            if (e.clientY - 71 - real.offsetTop < real.clientHeight / 2) {
                real.classList.add("dragTop")
                real.classList.remove("dragInner")
                real.classList.remove("dragBottom")
            }
            else {
                real.classList.remove("dragTop")
                real.classList.remove("dragInner")
                real.classList.add("dragBottom")
            }
        }
        e.preventDefault()
    }
    function onDragEnter(e: React.DragEvent) {
        let real = getRealDOM(e.target as HTMLElement);
        if (real.tagName === "LI" && real.classList.contains("drawer-list")) {
            real.classList.add("dragEnterList")
        }
        e.preventDefault()
    }
    function onDragLeave(e: React.DragEvent) {
        let real = e.target as HTMLElement;
        if (real.tagName === "LI" && real.classList.contains("drawer-list")) {
            real.classList.remove("dragEnterList")
        }
        e.preventDefault()
    }

    return <List
        dataSource={props.actionList.actions}
        footer={<Button type="dashed" block onClick={() => { createAction() }}>新增</Button>}
        renderItem={item => (
            <List.Item
                className="drawer-list"
                draggable="true"
                onDragStart={(event) => { onDragStart(event) }}
                onDrop={(event) => { onDrop(event) }}
                onDragEnter={(event) => { onDragEnter(event) }}
                onDragLeave={(event) => { onDragLeave(event) }}
                onDragOver={(event) => { onDragOver(event) }}
                allowdrop="true"
                key={"actList_" + item.Id} data-id={item.Id} data-type={"action"}
                actions={[
                    <Button type="link" onClick={() => { updateAction(item) }} key={`am-${item.Id}`}>
                        {intl.formatMessage({
                            id: 'update',
                        })}
                    </Button>,
                    <Button type="link" onClick={() => { updateParameter(item) }} key={`ap-${item.Id}`}>
                        {intl.formatMessage({
                            id: 'actionparam',
                        })}
                    </Button>,
                    <Button type="link" danger onClick={() => { deleteAction(item) }} key={`ad-${item.Id}`}>
                        {intl.formatMessage({
                            id: 'delete',
                        })}
                    </Button>,
                ]}
            >
                <List.Item.Meta
                    title={(item.IsVisible) ? (<span>
                        {item.Name}
                        <Divider type="vertical" />
                        <EyeOutlined style={{ marginRight: 6 }} />
                    </span>) : (<span>
                        {item.Name}
                        <Divider type="vertical" />
                        <EyeInvisibleOutlined style={{ marginRight: 6 }} />
                    </span>)}
                    description={<span>
                        <TagOutlined style={{ marginRight: 6 }} />
                        {item.Text}
                        <Divider type="vertical" />
                        <CodeOutlined style={{ marginRight: 6 }} />
                        {item.Type}

                    </span>}
                />
            </List.Item>
        )}
    />
};
export default connect((state: IStore) => {
    return {
        actionList: state.actionList,
    }
})(ActionsList);