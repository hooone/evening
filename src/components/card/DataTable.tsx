import React, { ReactElement } from 'react';
import { Table, Popconfirm, Button, Divider } from 'antd';
import { connect, useIntl, getLocale } from 'umi';
import { CardContentProps } from '@/models/card';
import { getLocaleText } from '@/util';
import { ICard, IViewAction } from '@/interfaces';


const DataTable = (props: CardContentProps) => {
    const intl = useIntl();

    let columns = [];
    props.card.Fields.forEach(
        field => {
            if (field.IsVisible) {
                let column = {
                    id: field.Id,
                    title: getLocaleText(field.Locale),
                    dataIndex: field.Name,
                }
                columns.push(column)
            }
        }
    )
    // let selectAble = (card.model.Reverse)
    let selectAble = false
    let actions: IViewAction[] = [];
    props.card.Actions.forEach(
        action => {
            if (action.Type === "MULTIUPDATE" || action.Type === "MULTIIMPORT" || action.Type === "MULTIEXPORT") {
                selectAble = true;
            }
            else if (action.Type === "SELECTIMPORT" || action.Type === "SELECTEXPORT") {
                selectAble = true;
            }
            if (action.Type === "UPDATE") {
                actions.push(action)
            }
            if (action.Type === "DELETE") {
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
            render: (text: any, record: any) => (
                <span>
                    {actions.map(
                        (action, idx) => {
                            return [(idx > 0) && (<Divider key={'actionDivider_' + action.Id} type="vertical" />),
                            (< span style={{ color: 'rgb(64,144,255)', cursor: 'pointer' }} key={'action_' + action.Id}
                                data-acionid={action.Id}
                                onClick={() => { props.onAction(action, record) }}>
                                {getLocaleText(action.Locale)}
                            </span>)]
                        }
                    )}
                </span >
            )
        })

    return (<Table
        columns={columns}
        dataSource={props.card.data}
        rowKey="__Key">
    </Table>)
};
export default DataTable;