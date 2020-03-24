import { Drawer, Tabs } from 'antd';
import { connect } from 'umi';
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


const ConfigDrawer = ({ draw, dispatch }) => {
    function onClose() {
        dispatch({
            type: 'draw/close',
        });
        dispatch({
            type: 'card_' + draw.cardId + '/loadCard',
            cardId: draw.cardId,
        });
        if (draw.cardChanged) {
            dispatch({
                type: 'page/loadPage',
            });
        }
    }
    function onSubClose() {
        dispatch({
            type: 'draw/subClose',
        });
        dispatch({
            type: 'fieldInfoConfig/close',
        });
        dispatch({
            type: 'actionInfoConfig/close',
        });
        dispatch({
            type: 'parameterInfoConfig/close',
        });
    }
    return <Drawer
        title={draw.title}
        placement="right"
        closable={false}
        width={412}
        onClose={() => { onClose() }}
        visible={draw.visible}>
        <Tabs defaultActiveKey="view" tabPosition={"left"} style={{ height: "100%" }}>
            <TabPane tab="属性" key={"view"}>
                <CardInfo />
            </TabPane>
            <TabPane tab="字段" key={"property"}>
                <FieldsList />
            </TabPane>
            <TabPane tab="操作" key={"action"}>
                <ActionList />
            </TabPane>
            {(draw.styleVisible) && <TabPane tab="样式" key={"relation"}>
                <RectChartInfo />
                <PointChartInfo />
            </TabPane>}
            <TabPane tab="卡片排序" key={"order"}>
                <CardsList />
            </TabPane>
        </Tabs>
        <Drawer
            title={draw.subTitle}
            width={297}
            onClose={() => { onSubClose() }}
            visible={draw.subVisible}
        >
            <FieldInfo />
            <ActionInfo />
            <ParameterInfo />
        </Drawer>
    </Drawer>
};
export default connect(({ draw }) => ({
    draw,
}))(ConfigDrawer);