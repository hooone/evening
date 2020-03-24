import { Row, Col, Form, Input, Select, Button } from 'antd';
import { connect } from 'umi';
const { Option } = Select;

const FieldInfo = ({ fieldInfoConfig, dispatch }) => {
    if (!fieldInfoConfig.visible) {
        return <span></span>
    }
    function onConfirm(field) {
        dispatch({
            type: 'fieldInfoConfig/saveField',
            field: field,
        });
    }
    function onCancel() {
        dispatch({ type: 'fieldInfoConfig/close' });
        dispatch({ type: 'draw/subClose' });
    }
    function onChange(name, value) {
        dispatch({
            type: 'fieldInfoConfig/dirty',
            name: name,
            value: value,
        })
    }
    return <Form layout="vertical" hideRequiredMark>
        <Row gutter={48}>
            <Col span={24}>
                <Form.Item
                    label="Name"
                >
                    <Input onChange={(value) => { onChange("Name", value.target.value) }}
                        value={fieldInfoConfig.Name} />
                </Form.Item>
                <Form.Item
                    label="标题"
                >
                    <Input onChange={(value) => { onChange("Text", value.target.value) }}
                        value={fieldInfoConfig.Text} />
                </Form.Item>
                <Form.Item
                    label="用户可见"
                >
                    <Select onChange={(value) => { onChange("IsVisible", value === "true") }}
                        value={fieldInfoConfig.IsVisible ? "true" : "false"} >
                        <Option value="true">可见</Option>
                        <Option value="false">不可见</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label="数据类型"
                >
                    <Select onChange={(value) => { onChange("Type", value) }}
                        value={fieldInfoConfig.Type} >
                        <Option value="int">int</Option>
                        <Option value="string">string</Option>
                        <Option value="float">float</Option>
                        <Option value="date">date</Option>
                        <Option value="datetime">datetime</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label="筛选条件"
                >
                    <Select onChange={(value) => { onChange("Filter", value) }}
                        value={fieldInfoConfig.Filter} >
                        <Option value="ALL">全选</Option>
                        <Option value="EQUAL">准确值</Option>
                        <Option value="RANGE">输入范围</Option>
                        <Option value="MAX">输入上限</Option>
                        <Option value="MIN">输入下限</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label="默认值"
                >
                    <Input onChange={(value) => { onChange("Default", value.target.value) }}
                        value={fieldInfoConfig.Default} />
                </Form.Item>
            </Col>
        </Row>
        {(fieldInfoConfig.dirty) && (
            <Row gutter={[16, 32]}>
                <Col span={6} offset={8}>
                    <Button onClick={() => { onCancel() }}>
                        取消
        </Button>
                </Col>
                <Col span={6} offset={1}>
                    <Button type="primary" onClick={() => {
                        onConfirm(fieldInfoConfig)
                    }}>
                        保存
        </Button>
                </Col>
            </Row>
        )}
    </Form>;
};
export default connect(({ fieldInfoConfig }) => ({
    fieldInfoConfig,
}))(FieldInfo);