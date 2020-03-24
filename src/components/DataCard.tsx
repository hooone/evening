import React, { ReactElement } from 'react';
import { Card, Button, Modal, Divider, Upload, Row, Col } from 'antd';
import { connect, getLocale, useIntl, Dispatch } from 'umi';
import { DispatchProp } from 'react-redux';
import { getLocaleText } from '@/util';
import { UploadOutlined, PlusOutlined, EditOutlined, SettingOutlined, SearchOutlined, ExclamationCircleOutlined } from '@ant-design/icons';
import DataTable from '@/components/card/DataTable';
import RectChart from '@/components/card/RectChart';
import PointChart from '@/components/card/PointChart';
import DescriptionsCard from '@/components/card/DescriptionsCard';
import StatisticCard from '@/components/card/StatisticCard';
import { cardStateProps } from '@/models/card';
import { ILoadDrawCommand } from '@/models/draw';
import { IShowActionCommand, IHandleOkCommand, ICancelCommand } from '@/models/modal';
import { IViewAction } from '@/interfaces'
import { IStore } from '@/store'
const { confirm } = Modal;

interface DataCardProps extends DispatchProp {
    card: cardStateProps,
}
const DataCard = (props: DataCardProps) => {
    const intl = useIntl();
    function showDrawer(cardId: number) {
        let loadDraw: ILoadDrawCommand = {
            type: 'draw/loadDraw',
            cardId: cardId,
        }
        props.dispatch(loadDraw);
    }
    function handleAction(action: IViewAction, record: any) {
        //生成record  
        if (action.Type === "CREATE" || action.Type === "READ") {
            action.Parameters.forEach(
                parameter => {
                    record[parameter.Field.Name] = parameter.Default
                }
            )
        }
        if (action.Type === "READ") {
            action.Parameters.forEach(
                (parameter, idx) => {
                    record["param" + idx] = parameter.Default
                }
            )
        }
        //判断是否有需要输入
        let hasInputValue = false;
        action.Parameters.forEach(
            parameter => {
                hasInputValue = hasInputValue || parameter.IsEditable;
            }
        )
        if (hasInputValue) {
            let showAction: IShowActionCommand = {
                type: 'modal/showAction',
                action: action,
                record: record,
            }
            props.dispatch(showAction);
        }
        else {
            if (action.DoubleCheck) {
                confirm({
                    title: (<span>
                        <span>{intl.formatMessage(
                            {
                                id: 'confirm',
                            }
                        )}</span>
                        <span>{getLocaleText(action.Locale) + "?"} </span>
                    </span>)
                    ,
                    icon: <ExclamationCircleOutlined />,
                    onOk() {
                        let handleOk: IHandleOkCommand = {
                            type: 'modal/handleOk',
                            action: action,
                            formdata: record,
                        }
                        props.dispatch(handleOk);
                    },
                    onCancel() {
                        let modalCancel: ICancelCommand = {
                            type: 'modal/cancel',
                        }
                        props.dispatch(modalCancel);
                    },
                });
            }
            else {
                let handleOk: IHandleOkCommand = {
                    type: 'modal/handleOk',
                    action: action,
                    formdata: record,
                }
                props.dispatch(handleOk);
            }
        }
    }

    let headActions: React.ReactElement[] = [];
    props.card.Actions.forEach(
        action => {
            if (action.Type === "CREATE") {
                headActions.push(<Button key={"action_" + action.Id} icon={<PlusOutlined />}
                    onClick={() => { handleAction(action, {}) }} >
                    {getLocaleText(action.Locale)}</Button >);
            }
        }
    )
    props.card.Actions.forEach(
        action => {
            if (action.Type === "MULTIUPDATE") {
                headActions.push(<Button key={"action_" + action.Id} icon={<EditOutlined />}  >
                    {getLocaleText(action.Locale)}
                </Button >);
            }
        }
    )
    props.card.Actions.forEach(
        action => {
            if (action.Type === "IMPORT" || action.Type === "SELECTIMPORT" || action.Type === "MULTIIMPORT") {

                headActions.push(<Upload
                    key={"action_" + action.Id}>
                    <Button>
                        <UploadOutlined />
                        {getLocaleText(action.Locale)}
                    </Button>
                </Upload>);
            }
        }
    )
    props.card.Actions.forEach(
        action => {
            if (action.Type === "EXPORT" || action.Type === "SELECTEXPORT" || action.Type === "MULTIEXPORT") {
                headActions.push(<Button key={"action_" + action.Id} icon={<UploadOutlined />} href="/" target="view_window" >
                    {getLocaleText(action.Locale)}
                </Button >);
            }
        }
    )
    if (props.card.Fields === undefined || props.card.Fields === null || props.card.Fields.length === 0) {

        return (<div data-contextmenu="card" data-cardid={props.card.Id} >
            <Card key={"tablecard_"} title={getLocaleText(props.card.Locale)}
                extra={
                    [
                        <Button key={"card_" + props.card.Id + "_setting"}
                            icon={<SettingOutlined />} type="dashed" shape="circle" size="small"
                            onClick={() => { showDrawer(props.card.Id) }}>
                        </Button>
                    ]
                }>
                <p> {intl.formatMessage(
                    {
                        id: 'emptyField',
                    }
                )}</p>
            </Card>
        </div>)
    }
    else if (props.card.Style === "RECT") {
        return (<div data-contextmenu="card" data-cardid={props.card.Id}  >
            <Card key={"tablecard_" + props.card.Id} title={getLocaleText(props.card.Locale)} className="tableCard"
                extra={[
                    <Button key={"card_" + props.card.Id + "_read"} icon={<SearchOutlined />}
                        onClick={() => { handleAction(props.card.Reader, {}) }}>
                        {intl.formatMessage(
                            {
                                id: 'read',
                            }
                        )}
                    </Button>,
                    <Divider type="vertical" key={"Divider_setting"} />,
                    <Button key={"card_" + props.card.Id + "_setting"}
                        icon={<SettingOutlined />} type="dashed" shape="circle" size="small"
                        onClick={() => { showDrawer(props.card.Id) }}>
                    </Button>
                ]}>
                <RectChart
                    card={props.card}
                />
            </Card>
        </div >)
    }
    else if (props.card.Style === "POINT") {
        return (<div data-contextmenu="card" data-cardid={props.card.Id}  >
            <Card key={"tablecard_" + props.card.Id} title={getLocaleText(props.card.Locale)} className="tableCard"
                extra={[
                    <Button key={"card_" + props.card.Id + "_read"} icon={<SearchOutlined />}
                        onClick={() => { handleAction(props.card.Reader, {}) }}>
                        {intl.formatMessage(
                            {
                                id: 'read',
                            }
                        )}
                    </Button>,
                    <Divider type="vertical" key={"Divider_setting"} />,
                    <Button key={"card_" + props.card.Id + "_setting"}
                        icon={<SettingOutlined />} type="dashed" shape="circle" size="small"
                        onClick={() => { showDrawer(props.card.Id) }}>
                    </Button>
                ]}>
                <PointChart
                    card={props.card}
                />
            </Card>
        </div >)
    }
    else if (props.card.Style === "DESC") {
        return (<div data-contextmenu="card" data-cardid={props.card.Id}  >
            <Card key={"tablecard_" + props.card.Id} title={getLocaleText(props.card.Locale)} className="tableCard"
                extra={[
                    <Button key={"card_" + props.card.Id + "_read"} icon={<SearchOutlined />}
                        onClick={() => { handleAction(props.card.Reader, {}) }}>
                        {intl.formatMessage(
                            {
                                id: 'read',
                            }
                        )}
                    </Button>,
                    <Divider type="vertical" key={"Divider_setting"} />,
                    <Button key={"card_" + props.card.Id + "_setting"}
                        icon={<SettingOutlined />} type="dashed" shape="circle" size="small"
                        onClick={() => { showDrawer(props.card.Id) }}>
                    </Button>
                ]}>
                <DescriptionsCard
                    card={props.card}
                />
            </Card>
        </div >)
    }
    else if (props.card.Style === "STAT") {
        return (<div data-contextmenu="card" data-cardid={props.card.Id}  >
            <Card key={"tablecard_" + props.card.Id} className="tableCard">
                <Row>
                    <Col span={23}>
                        <StatisticCard card={props.card} />
                    </Col>
                    <Col span={1}>
                        <Button key={"card_" + props.card.Id + "_setting"}
                            icon={<SettingOutlined />} type="dashed" shape="circle" size="small"
                            onClick={() => { showDrawer(props.card.Id) }}>
                        </Button>
                    </Col>
                </Row>
            </Card>
        </div >)
    }
    else {
        return (<div data-contextmenu="card" data-cardid={props.card.Id}  >
            <Card key={"tablecard_" + props.card.Id} title={getLocaleText(props.card.Locale)} className="tableCard"
                extra={headActions.concat(
                    [
                        <Button key={"card_" + props.card.Id + "_read"} icon={<SearchOutlined />}
                            onClick={() => { handleAction(props.card.Reader, {}) }}>
                            {intl.formatMessage(
                                {
                                    id: 'read',
                                }
                            )}
                        </Button>,
                        <Divider type="vertical" key={"Divider_setting"} />,
                        <Button key={"card_" + props.card.Id + "_setting"}
                            icon={<SettingOutlined />} type="dashed" shape="circle" size="small"
                            onClick={() => { showDrawer(props.card.Id) }}>
                        </Button>
                    ])
                }>
                <DataTable
                    card={props.card}
                    onAction={handleAction}
                />
            </Card>
        </div >)
    }
};


export default connect((state: IStore, props: any) => {
    return {
        card: state['card_' + props["cardInfo"].Id]
    }
})(DataCard);