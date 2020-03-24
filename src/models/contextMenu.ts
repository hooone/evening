import { EffectsCommandMap } from 'dva'
import { contextMenuStateProps } from "@/interfaces"

export default {
    namespace: 'contextMenu',
    state: { visible: false, menu: "", left: 0, top: 0, record: {} },
    reducers: {
        open(state: contextMenuStateProps, action: contextMenuStateProps) {
            return {
                visible: true,
                menu: action.menu,
                left: action.left,
                top: action.top,
                record: action.record,
            }
        },
        close(state: contextMenuStateProps) {
            return {
                visible: false,
                menu: "",
                left: 0,
                top: 0,
                record: {},
            }
        },
    },
    effects: {
        *createfolder(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Locale: {
                    "zh-CN": '增加菜单栏文件夹',
                    "en-US": "create folder"
                },
                DoubleCheck: false,
                URL: '/api/navigation/createFolder',
                Type: "NAV",
                Parameters: [
                    {
                        ParameterID: 'Name_addTree',
                        Default: "",
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Name',
                            Text: '名称',
                            Locale: {
                                "zh-CN": '名称',
                                "en-US": "Name"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'Text_addTree',
                        Default: "",
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Text',
                            Text: '文本',
                            Locale: {
                                "zh-CN": '文本',
                                "en-US": "Text"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'IsFolder_addTree',
                        Default: "true",
                        IsEditable: false,
                        IsVisible: false,
                        Field: {
                            Name: 'IsFolder',
                            Text: '是文件夹',
                            Locale: {
                                "zh-CN": '是文件夹',
                                "en-US": "IsFolder"
                            },
                            Type: 'string'
                        },
                    }
                ]
            };
            yield handler.put({
                type: 'modal/showAction',
                action: constAction,
                record: action.record,
            });
        },
        * createpage(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Text: '增加页面',
                Locale: {
                    "zh-CN": '增加页面',
                    "en-US": "create page"
                },
                DoubleCheck: false,
                URL: '/api/navigation/createTreePage',
                Type: "NAV",
                Parameters: [
                    {
                        ParameterID: 'Name_addTreePage',
                        Default: "",
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Name',
                            Text: '名称',
                            Locale: {
                                "zh-CN": '名称',
                                "en-US": "Name"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'Text_addTreePage',
                        Default: "",
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Text',
                            Text: '文本',
                            Locale: {
                                "zh-CN": '文本',
                                "en-US": "Text"
                            },
                            Type: 'string'
                        },
                    }
                ]
            };
            yield handler.put({
                type: 'modal/showAction',
                action: constAction,
                record: action.record,
            });
        },
        * updateFolder(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Text: '修改文件夹属性',
                Locale: {
                    "zh-CN": '修改文件夹属性',
                    "en-US": "update folder"
                },
                DoubleCheck: false,
                URL: '/api/navigation/updateFolder',
                Type: "NAV",
                Parameters: [
                    {
                        ParameterID: 'Name_addTree',
                        Default: action.record.name,
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Name',
                            Text: '名称',
                            Locale: {
                                "zh-CN": '名称',
                                "en-US": "Name"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'Text_addTree',
                        Default: action.record.text,
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Text',
                            Text: '文本',
                            Locale: {
                                "zh-CN": '文本',
                                "en-US": "Text"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'FolderId_deleteID',
                        Default: action.record.treeid,
                        IsEditable: false,
                        IsVisible: false,
                        Field: {
                            Name: 'FolderId',
                            Text: '菜单ID',
                            Locale: {
                                "zh-CN": '菜单ID',
                                "en-US": "FolderId"
                            },
                            Type: 'int'
                        },
                    }
                ]
            };
            yield handler.put({
                type: 'modal/showAction',
                action: constAction,
                record: {
                    FolderId: action.record.treeid,
                    Text: action.record.text,
                    Name: action.record.name,
                },
            });
        },
        * deleteFolder(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Text: '删除文件夹',
                Locale: {
                    "zh-CN": '删除文件夹',
                    "en-US": "delete folder"
                },
                DoubleCheck: true,
                URL: '/api/navigation/deleteFolder',
                Type: "NAV",
                Parameters: [
                    {
                        ParameterID: 'FolderId_deleteID',
                        Default: "0",
                        IsEditable: false,
                        IsVisible: false,
                        Field: {
                            Name: 'FolderId',
                            Text: '菜单ID',
                            Locale: {
                                "zh-CN": '菜单ID',
                                "en-US": "FolderId"
                            },
                            Type: 'int'
                        },
                    }
                ]
            };
            yield handler.put({
                type: 'modal/handleOk',
                action: constAction,
                formdata: {
                    FolderId: action.record.treeid,
                },
            });
        },
        * deletePage(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Text: '删除页面',
                Locale: {
                    "zh-CN": '删除页面',
                    "en-US": "delete page"
                },
                DoubleCheck: true,
                URL: '/api/navigation/deletePage',
                Type: "NAV",
                Parameters: [
                    {
                        ParameterID: 'PageID_deleteID',
                        Default: action.record.pageid,
                        IsEditable: false,
                        IsVisible: false,
                        Field: {
                            Name: 'PageID',
                            Text: '页面ID',
                            Locale: {
                                "zh-CN": '页面ID',
                                "en-US": "PageID"
                            },
                            Type: 'int'
                        },
                    }
                ]
            };
            yield handler.put({
                type: 'modal/handleOk',
                action: constAction,
                formdata: {
                    PageID: action.record.pageid,
                },
            });
        },
        * createNodePage(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Text: '增加页面',
                Locale: {
                    "zh-CN": '增加页面',
                    "en-US": "create page"
                },
                DoubleCheck: false,
                URL: '/api/navigation/createNodePage',
                Type: "NAV",
                Parameters: [
                    {
                        ParameterID: 'Name_addNode',
                        Default: "",
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Name',
                            Text: '名称',
                            Locale: {
                                "zh-CN": '名称',
                                "en-US": "Name"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'Text_addNode',
                        Default: "",
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Text',
                            Text: '文本',
                            Locale: {
                                "zh-CN": '文本',
                                "en-US": "Text"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'FolderId_addNode',
                        Default: action.record.treeid,
                        IsEditable: false,
                        IsVisible: false,
                        Field: {
                            Name: 'FolderId',
                            Text: '菜单ID',
                            Locale: {
                                "zh-CN": '菜单ID',
                                "en-US": "FolderId"
                            },
                            Type: 'int'
                        },
                    }
                ]
            };
            yield handler.put({
                type: 'modal/showAction',
                action: constAction,
                record: {
                    FolderId: action.record.treeid,
                },
            });
        },
        * updatePage(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Text: '修改页面',
                Locale: {
                    "zh-CN": '修改页面',
                    "en-US": "update page"
                },
                DoubleCheck: false,
                URL: '/api/navigation/updatePage',
                Type: "NAV",
                Parameters: [
                    {
                        ParameterID: 'Name_updateNode',
                        Default: action.record.name,
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Name',
                            Text: '名称',
                            Locale: {
                                "zh-CN": '名称',
                                "en-US": "Name"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'Text_updateNode',
                        Default: action.record.text,
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Text',
                            Text: '文本',
                            Locale: {
                                "zh-CN": '文本',
                                "en-US": "Text"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'PageID_addNode',
                        Default: action.record.pageid,
                        IsEditable: false,
                        IsVisible: false,
                        Field: {
                            Name: 'PageID',
                            Text: '页面ID',
                            Locale: {
                                "zh-CN": '页面ID',
                                "en-US": "PageID"
                            },
                            Type: 'int'
                        },
                    }
                ]
            };
            yield handler.put({
                type: 'modal/showAction',
                action: constAction,
                record: {
                    PageID: action.record.pageid,
                },
            });
        },
        * createCard(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Text: '添加卡片',
                Locale: {
                    "zh-CN": '添加卡片',
                    "en-US": "create card"
                },
                DoubleCheck: false,
                URL: '/api/card/create',
                Type: "CARD",
                Parameters: [
                    {
                        ParameterID: 'Name_createCard',
                        Default: "",
                        IsEditable: true,
                        IsVisible: true,
                        Field: {
                            Name: 'Name',
                            Text: '名称',
                            Locale: {
                                "zh-CN": '名称',
                                "en-US": "Name"
                            },
                            Type: 'string'
                        },
                    },
                    {
                        ParameterID: 'Text_createCard',
                        Default: "",
                        IsVisible: true,
                        IsEditable: true,
                        Field: {
                            Name: 'Text',
                            Text: '文本',
                            Locale: {
                                "zh-CN": '文本',
                                "en-US": "Text"
                            },
                            Type: 'string'
                        }
                    },
                    {
                        ParameterID: 'PageID_createCard',
                        Default: action.record.pageid,
                        IsVisible: false,
                        IsEditable: false,
                        Field: {
                            Name: 'PageID',
                            Text: 'PageID',
                            Locale: {
                                "zh-CN": 'PageID',
                                "en-US": "PageID"
                            },
                            Type: 'int'
                        }
                    },
                    {
                        ParameterID: 'PageName_createCard',
                        Default: action.record.pagename,
                        IsVisible: false,
                        IsEditable: false,
                        Field: {
                            Name: 'PageName',
                            Text: 'PageName',
                            Locale: {
                                "zh-CN": 'PageName',
                                "en-US": "PageName"
                            },
                            Type: 'string'
                        }
                    }
                ]
            };
            yield handler.put({
                type: 'modal/showAction',
                action: constAction,
                record: {
                    PageID: action.record.pageid,
                    PageName: action.record.pagename,
                },
            });
        },
        * deleteCard(action: contextMenuStateProps, handler: EffectsCommandMap) {
            let constAction = {
                Text: '删除卡片',
                Locale: {
                    "zh-CN": '删除卡片',
                    "en-US": "delete card"
                },
                DoubleCheck: true,
                URL: '/api/card/delete',
                Type: "CARD",
                Parameters: [
                    {
                        ParameterID: 'CardID_deleteCard',
                        Default: action.record.cardid,
                        IsVisible: false,
                        IsEditable: false,
                        Field: {
                            Name: 'CardId',
                            Text: 'CardId',
                            Locale: {
                                "zh-CN": 'CardId',
                                "en-US": "CardId"
                            },
                            Type: 'int'
                        }
                    }
                ]
            };
            yield handler.put({
                type: 'modal/handleOk',
                action: constAction,
                formdata: {
                    CardId: action.record.cardid,
                },
            });
        },
    },
};