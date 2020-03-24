import React, { ReactElement } from 'react';
import { Table, Popconfirm, Button, Divider } from 'antd';
import { connect, useIntl, getLocale } from 'umi';

const DataTable = ({ onAction, card }) => {
    const intl = useIntl();
    const rowSelection = {
        onChange: (selectedRowKeys, selectedRows) => {
            if (selectedRows.length > 0)
                this.props.SelectionChange(card.Id + '', selectedRows[0])
        },
        type: 'radio'
    };
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
    let columns = [];
    card.Fields.forEach(
        prop => {
            if (prop.IsVisible) {
                let column = {
                    id: prop.ID,
                    title: getLocaleText(prop.Name, prop.Locale),
                    dataIndex: prop.Name,
                }
                columns.push(column)
            }
        }
    )
    // let selectAble = (card.model.Reverse)
    let selectAble = false
    let actions = [];
    card.Actions.forEach(
        action => {
            if (action.Type === "MULTIUPDATE" || action.Type === "MULTIIMPORT" || action.Type === "MULTIEXPORT") {
                rowSelection.type = "checkbox";
                selectAble = true;
            }
            else if (action.Type === "SELECTIMPORT" || action.Type === "SELECTEXPORT") {
                selectAble = true;
            }
            if (action.Type === "UPDATE") {
                if (actions.length > 0) {
                    actions.push(
                        { Id: action.Id, Type: "DIVIDER" }
                    )
                }
                actions.push(action)
            }
        }
    )
    card.Actions.forEach(
        action => {
            if (action.Type === "DELETE") {
                if (actions.length > 0) {
                    actions.push(
                        { Id: action.Id, Type: "DIVIDER" }
                    )
                }
                actions.push(action)
            }
        }
    )
    columns.push(
        {
            title: (<div>{intl.formatMessage(
                {
                    id: 'actions',
                }
            )}</div>),
            key: 'action',
            render: (text, record) => (
                <span>
                    {actions.map(
                        action => {
                            if (action.Type === "DIVIDER") {
                                return <Divider key={'actionDivider_' + action.Id} type="vertical" />;
                            }
                            else {
                                return <span style={{ color: 'rgb(64,144,255)', cursor: 'pointer' }} key={'action_' + action.Id}
                                    data-acionid={action.Id}
                                    onClick={() => { onAction(action, record) }}>
                                    {getLocaleText(action.Name, action.Locale)}
                                </span>
                            }
                        }
                    )}
                </span>
            )
        })

    return (<Table
        rowSelection={selectAble ? rowSelection : null}
        columns={columns}
        dataSource={card.data}
        rowKey="__Key">
    </Table>)
};
export default DataTable;