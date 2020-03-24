import { Row, Col, Form, Input, Select, Button, Switch } from 'antd';
import { connect } from 'umi';
const { Option } = Select;

const ParameterInfo = ({ parameterInfoConfig, dispatch }) => {
    console.log(parameterInfoConfig)
    if (!parameterInfoConfig.visible) {
        return <span></span>
    }
    function onConfirm(parameters, cardId) {
        dispatch({
            type: 'parameterInfoConfig/saveParameter',
            parameters: parameters,
            cardId: cardId,
        });
    }
    function onCancel() {
        dispatch({ type: 'parameterInfoConfig/close' });
        dispatch({ type: 'draw/subClose' });
    }
    function onChange(name, id, value) {
        dispatch({
            type: 'parameterInfoConfig/dirty',
            name: name,
            id: id,
            value: value,
        })
    }
    let parameterItems = []
    if (parameterInfoConfig && parameterInfoConfig.parameters) {
        parameterInfoConfig.parameters.forEach(param => {
            parameterItems.push(<Form.Item
                key={"paramItem_" + param.Id}
                label={[
                    <span key={"paramN_" + param.Id} >  {param.Field.Text} </span>,
                    <Switch key={"paramV_" + param.Id}
                        onChange={(v) => onChange("IsVisible", param.Id, v)}
                        checked={param.IsVisible}
                        checkedChildren={"可见"}
                        unCheckedChildren={"隐藏"} style={{ marginLeft: "8px" }} />,
                    <Switch key={"paramE_" + param.Id}
                        onChange={(v) => onChange("IsEditable", param.Id, v)}
                        checked={param.IsEditable}
                        checkedChildren={"输入"}
                        unCheckedChildren={"只读"}
                        style={{ marginLeft: "8px" }} />]}
            >
                <Input key={"default_" + param.Id}
                    onChange={(v) => onChange("Default", param.Id, v.target.value)}
                    value={param.Default} />
            </Form.Item>)
        })
    }

    return <Form layout="vertical" hideRequiredMark>
        <Row gutter={48}>
            <Col span={24}>
                {parameterItems}
            </Col>
        </Row>
        {(parameterInfoConfig) && (parameterInfoConfig.dirty) && (
            <Row gutter={[16, 32]}>
                <Col span={6} offset={8}>
                    <Button onClick={() => { onCancel() }}>
                        取消
        </Button>
                </Col>
                <Col span={6} offset={1}>
                    <Button type="primary" onClick={() => {
                        onConfirm(parameterInfoConfig.parameters, parameterInfoConfig.cardId)
                    }}>
                        保存
        </Button>
                </Col>
            </Row>
        )}
    </Form>;
};
export default connect(({ parameterInfoConfig }) => ({
    parameterInfoConfig,
}))(ParameterInfo);