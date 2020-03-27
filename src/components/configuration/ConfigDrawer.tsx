import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Drawer, Tabs } from 'antd';
import { connect, useIntl } from 'umi';
import { DispatchProp } from 'react-redux';
import { drawStateProps } from '@/models/draw'
import { IStore } from '@/store'
import CardInfo from '@/components/configuration/CardInfo';
import FieldsList from '@/components/configuration/FieldsList';
import FieldInfo from '@/components/configuration/FieldInfo';
import ActionList from '@/components/configuration/ActionsList';
import ActionInfo from '@/components/configuration/ActionInfo';
import ParameterInfo from '@/components/configuration/ParameterInfo';
import CardsList from '@/components/configuration/CardsList';
import RectChartInfo from '@/components/configuration/RectChartInfo';
import PointChartInfo from '@/components/configuration/PointChartInfo';
const { TabPane } = Tabs;

interface ConfigDrawerProps extends DispatchProp {
    draw: drawStateProps,
}

const ConfigDrawer = (props: ConfigDrawerProps) => {
    const intl = useIntl();
    function onClose() {
        props.dispatch({
            type: 'draw/close',
        });
        props.dispatch({
            type: 'card_' + props.draw.cardId + '/loadCard',
            cardId: props.draw.cardId,
        });
        if (props.draw.cardChanged) {
            props.dispatch({
                type: 'page/loadPage',
            });
        }
    }
    function onSubClose() {
        props.dispatch({
            type: 'draw/subClose',
        });
        props.dispatch({
            type: 'fieldInfoConfig/close',
        });
        props.dispatch({
            type: 'actionInfoConfig/close',
        });
        props.dispatch({
            type: 'parameterInfoConfig/close',
        });
    }
    return <Drawer
        title={props.draw.title}
        placement="right"
        closable={false}
        width={412}
        onClose={() => { onClose() }}
        visible={props.draw.visible}>
        <Tabs defaultActiveKey="view" tabPosition={"left"} style={{ height: "100%" }}>
            <TabPane tab={intl.formatMessage({
                id: 'property',
            })} key={"view"}>
                <CardInfo />
            </TabPane>
            <TabPane tab={intl.formatMessage({
                id: 'field',
            })} key={"field"}>
                <FieldsList />
            </TabPane>
            <TabPane tab={intl.formatMessage({
                id: 'actions',
            })} key={"action"}>
                <ActionList />
            </TabPane>
            {(props.draw.styleVisible) && <TabPane tab={intl.formatMessage({
                id: 'style',
            })} key={"relation"}>
                <RectChartInfo />
                <PointChartInfo />
            </TabPane>}
            <TabPane tab={intl.formatMessage({
                id: 'cardorder',
            })} key={"order"}>
                <CardsList />
            </TabPane>
        </Tabs>
        <Drawer
            title={props.draw.subTitle}
            width={297}
            onClose={() => { onSubClose() }}
            visible={props.draw.subVisible}
        >
            <FieldInfo />
            <ActionInfo />
            <ParameterInfo />
        </Drawer>
    </Drawer>
};
export default connect((state: IStore) => {
    return {
        draw: state.draw,
    }
})(ConfigDrawer);