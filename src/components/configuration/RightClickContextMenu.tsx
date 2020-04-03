import React, { Component } from 'react';
import { connect, getLocale, useIntl } from 'umi';
import { Modal, message } from 'antd';
import { DispatchProp } from 'react-redux';
import { ExclamationCircleOutlined } from '@ant-design/icons';
import { contextMenuStateProps } from '@/models/contextMenu';
import { getLocaleText } from '@/util';
import { IStore } from '@/store'
import { ILocale } from '@/interfaces';
const { confirm } = Modal;


interface RightClickContextMenuProps extends DispatchProp {
    contextMenu: contextMenuStateProps,
}
const RightClickContextMenu = (props: RightClickContextMenuProps) => {
    if (props.contextMenu.visible) {
        const intl = useIntl();
        return (<div id="contextmenu" className="skin"
            style={{ left: props.contextMenu.left, top: props.contextMenu.top }}>
            <ul className="dropdown-menu">
                {(props.contextMenu.menu === "addtree") &&
                    <li onClick={() => {
                        props.dispatch({
                            type: 'contextMenu/createfolder',
                            record: props.contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage({
                            id: 'createfolder',
                        })}
                    </li>}
                {(props.contextMenu.menu === "addtree") &&
                    <li onClick={() => {
                        props.dispatch({
                            type: 'contextMenu/createpage',
                            record: props.contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'createpage',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "treeview") &&
                    <li onClick={() => {
                        props.dispatch({
                            type: 'contextMenu/createNodePage',
                            record: props.contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'createsubpage',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "treeview") &&
                    <li onClick={() => {
                        props.dispatch({
                            type: 'contextMenu/updateFolder',
                            record: props.contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'update',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "treeview") &&
                    <li onClick={() => {
                        let locale: ILocale = {
                            "zh-CN": '删除文件夹',
                            "en-US": "delete folder",
                            Default: "delete folder",
                        }
                        confirm({
                            title: (<span>
                                <span>{intl.formatMessage(
                                    {
                                        id: 'confirm',
                                    }
                                )}</span>
                                <span>{getLocaleText(locale) + "?"} </span>
                            </span>)
                            ,
                            icon: <ExclamationCircleOutlined />,
                            onOk() {
                                props.dispatch({
                                    type: 'contextMenu/deleteFolder',
                                    record: props.contextMenu.record,
                                })
                            },
                            onCancel() {

                            },
                        });
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'delete',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "treepage") &&
                    <li onClick={() => {
                        props.dispatch({
                            type: 'contextMenu/updatePage',
                            record: props.contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'update',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "treepage") &&
                    <li onClick={() => {
                        if (props.contextMenu.record.name === "home") {
                            message.error('不允许删除主页');
                            return;
                        }
                        let locale: ILocale = {
                            "zh-CN": '删除页面',
                            "en-US": "delete page",
                            Default: "delete page",
                        }
                        confirm({
                            title: (<span>
                                <span>{intl.formatMessage(
                                    {
                                        id: 'confirm',
                                    }
                                )}</span>
                                <span>{getLocaleText(locale) + "?"} </span>
                            </span>)
                            ,
                            icon: <ExclamationCircleOutlined />,
                            onOk() {
                                props.dispatch({
                                    type: 'contextMenu/deletePage',
                                    record: props.contextMenu.record,
                                })
                            },
                            onCancel() {

                            },
                        });
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'delete',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "nodepage") &&
                    <li onClick={() => {
                        props.dispatch({
                            type: 'contextMenu/updatePage',
                            record: props.contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'update',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "nodepage") &&
                    <li onClick={() => {
                        let locale: ILocale = {
                            "zh-CN": '删除页面',
                            "en-US": "delete page",
                            "Default": "delete page",
                        }
                        let deleteHome: ILocale = {
                            "zh-CN": '不允许删除主页',
                            "en-US": "can not delete home",
                            "Default": "can not delete home",
                        }
                        if (props.contextMenu.record.name === "home") {
                            message.error(getLocaleText(deleteHome));
                            return;
                        }
                        confirm({
                            title: (<span>
                                <span>{intl.formatMessage(
                                    {
                                        id: 'confirm',
                                    }
                                )}</span>
                                <span>{getLocaleText(locale) + "?"} </span>
                            </span>)
                            ,
                            icon: <ExclamationCircleOutlined />,
                            onOk() {
                                props.dispatch({
                                    type: 'contextMenu/deletePage',
                                    record: props.contextMenu.record,
                                })
                            },
                            onCancel() {

                            },
                        });
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'delete',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "card") &&
                    <li onClick={() => {
                        let locale: ILocale = {
                            "zh-CN": '删除窗体',
                            "en-US": "delete card",
                            "Default": "delete card",
                        }
                        confirm({
                            title: (<span>
                                <span>{intl.formatMessage(
                                    {
                                        id: 'confirm',
                                    }
                                )}</span>
                                <span>{getLocaleText(locale) + "?"} </span>
                            </span>)
                            ,
                            icon: <ExclamationCircleOutlined />,
                            onOk() {
                                props.dispatch({
                                    type: 'contextMenu/deleteCard',
                                    record: props.contextMenu.record,
                                })
                            },
                            onCancel() {

                            },
                        });
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'deletebox',
                            }
                        )}
                    </li>}
                {(props.contextMenu.menu === "content") &&
                    <li onClick={() => {
                        props.dispatch({
                            type: 'contextMenu/createCard',
                            record: props.contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'createbox',
                            }
                        )}
                    </li>}
            </ul>
        </div >
        )
    }
    else {
        return <span></span>
    }
}

export default connect((state: IStore, props: any) => {
    return {
        contextMenu: state.contextMenu
    }
})(RightClickContextMenu);