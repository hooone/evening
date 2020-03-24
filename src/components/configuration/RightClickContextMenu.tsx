import React, { Component } from 'react';
import { connect, getLocale, useIntl } from 'umi';
import { Modal } from 'antd';
const { confirm } = Modal;
import { ExclamationCircleOutlined } from '@ant-design/icons';

const RightClickContextMenu = ({ contextMenu, dispatch }) => {
    function getLocaleText(Name, Locale) {
        let lang = getLocale();
        if (Locale[lang] === undefined || Locale[lang] === "") {
            if (Locale["default"] !== undefined) {
                return Locale["default"]
            }
            return Name;
        }
        else {
            return Locale[lang];
        }
    }
    if (contextMenu.visible) {
        const intl = useIntl();
        return (<div id="contextmenu" className="skin" style={{ left: contextMenu.left, top: contextMenu.top }}>
            <ul className="dropdown-menu">
                {(contextMenu.menu === "addtree") &&
                    <li onClick={() => {
                        dispatch({
                            type: 'contextMenu/createfolder',
                            record: contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'createfolder',
                            }
                        )}
                    </li>}
                {(contextMenu.menu === "addtree") &&
                    <li onClick={() => {
                        dispatch({
                            type: 'contextMenu/createpage',
                            record: contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'createpage',
                            }
                        )}
                    </li>}
                {(contextMenu.menu === "treeview") &&
                    <li onClick={() => {
                        dispatch({
                            type: 'contextMenu/createNodePage',
                            record: contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'createsubpage',
                            }
                        )}
                    </li>}
                {(contextMenu.menu === "treeview") &&
                    <li onClick={() => {
                        dispatch({
                            type: 'contextMenu/updateFolder',
                            record: contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'update',
                            }
                        )}
                    </li>}
                {(contextMenu.menu === "treeview") &&
                    <li onClick={() => {
                        let locale = {
                            "zh-CN": '删除文件夹',
                            "en-US": "delete folder"
                        }
                        confirm({
                            title: (<span>
                                <span>{intl.formatMessage(
                                    {
                                        id: 'confirm',
                                    }
                                )}</span>
                                <span>{getLocaleText("Delete", locale) + "?"} </span>
                            </span>)
                            ,
                            icon: <ExclamationCircleOutlined />,
                            onOk() {
                                dispatch({
                                    type: 'contextMenu/deleteFolder',
                                    record: contextMenu.record,
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
                {(contextMenu.menu === "treepage") &&
                    <li onClick={() => {
                        dispatch({
                            type: 'contextMenu/updatePage',
                            record: contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'update',
                            }
                        )}
                    </li>}
                {(contextMenu.menu === "treepage") &&
                    <li onClick={() => {
                        let locale = {
                            "zh-CN": '删除页面',
                            "en-US": "delete page"
                        }
                        confirm({
                            title: (<span>
                                <span>{intl.formatMessage(
                                    {
                                        id: 'confirm',
                                    }
                                )}</span>
                                <span>{getLocaleText("Delete", locale) + "?"} </span>
                            </span>)
                            ,
                            icon: <ExclamationCircleOutlined />,
                            onOk() {
                                dispatch({
                                    type: 'contextMenu/deletePage',
                                    record: contextMenu.record,
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
                {(contextMenu.menu === "nodepage") &&
                    <li onClick={() => {
                        dispatch({
                            type: 'contextMenu/updatePage',
                            record: contextMenu.record,
                        })
                    }}>
                        {intl.formatMessage(
                            {
                                id: 'update',
                            }
                        )}
                    </li>}
                {(contextMenu.menu === "nodepage") &&
                    <li onClick={() => {
                        let locale = {
                            "zh-CN": '删除页面',
                            "en-US": "delete page"
                        }
                        confirm({
                            title: (<span>
                                <span>{intl.formatMessage(
                                    {
                                        id: 'confirm',
                                    }
                                )}</span>
                                <span>{getLocaleText("Delete", locale) + "?"} </span>
                            </span>)
                            ,
                            icon: <ExclamationCircleOutlined />,
                            onOk() {
                                dispatch({
                                    type: 'contextMenu/deletePage',
                                    record: contextMenu.record,
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
                {(contextMenu.menu === "card") &&
                    <li onClick={() => {
                        let locale = {
                            "zh-CN": '删除卡片',
                            "en-US": "delete card"
                        }
                        confirm({
                            title: (<span>
                                <span>{intl.formatMessage(
                                    {
                                        id: 'confirm',
                                    }
                                )}</span>
                                <span>{getLocaleText("Delete", locale) + "?"} </span>
                            </span>)
                            ,
                            icon: <ExclamationCircleOutlined />,
                            onOk() {
                                dispatch({
                                    type: 'contextMenu/deleteCard',
                                    record: contextMenu.record,
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
                {(contextMenu.menu === "content") &&
                    <li onClick={() => {
                        dispatch({
                            type: 'contextMenu/createCard',
                            record: contextMenu.record,
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

export default connect(({ contextMenu }) => ({
    contextMenu,
}))(RightClickContextMenu);