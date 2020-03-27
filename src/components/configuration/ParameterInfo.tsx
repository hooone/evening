import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Row, Col, Form, Input, Select, Button, Switch } from 'antd';
import { connect, useIntl } from 'umi';
import { DispatchProp } from 'react-redux';
import { parameterStateProps } from '@/models/parameterInfoConfig';
import { IParameter } from '@/interfaces';
import { IStore } from '@/store'
const { Option } = Select;

interface ParameterInfoProps extends DispatchProp {
    parameterInfoConfig: parameterStateProps,
}
const ParameterInfo = (props: ParameterInfoProps) => {
    const intl = useIntl();
    if (!props.parameterInfoConfig.visible) {
        return <span></span>
    }
    function onConfirm(parameters: IParameter[], cardId: number) {
        props.dispatch({
            type: 'parameterInfoConfig/saveParameter',
            parameters: parameters,
            cardId: cardId,
        });
    }
    function onCancel() {
        props.dispatch({ type: 'parameterInfoConfig/close' });
        props.dispatch({ type: 'draw/subClose' });
    }
    function onChange(name: string, id: number, value: any) {
        props.dispatch({
            type: 'parameterInfoConfig/dirty',
            name: name,
            id: id,
            value: value,
        })
    }
    let parameterItems: ReactElement[] = []
    if (props.parameterInfoConfig && props.parameterInfoConfig.parameters) {
        props.parameterInfoConfig.parameters.forEach(param => {
            parameterItems.push(<Form.Item
                key={"paramItem_" + param.Id}
                label={[
                    <span key={"paramN_" + param.Id} >  {param.Field.Text} </span>,
                    <Switch key={"paramV_" + param.Id}
                        onChange={(v) => onChange("IsVisible", param.Id, v)}
                        checked={param.IsVisible}
                        checkedChildren={intl.formatMessage({
                            id: 'visible',
                        })}
                        unCheckedChildren={intl.formatMessage({
                            id: 'hidden',
                        })} style={{ marginLeft: "8px" }} />,
                    <Switch key={"paramE_" + param.Id}
                        onChange={(v) => onChange("IsEditable", param.Id, v)}
                        checked={param.IsEditable}
                        checkedChildren={intl.formatMessage({
                            id: 'editable',
                        })}
                        unCheckedChildren={intl.formatMessage({
                            id: 'readonly',
                        })}
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
        {(props.parameterInfoConfig) && (props.parameterInfoConfig.dirty) && (
            <Row gutter={[16, 32]}>
                <Col span={6} offset={8}>
                    <Button onClick={() => { onCancel() }}>
                        {intl.formatMessage({
                            id: 'cancel',
                        })}
                    </Button>
                </Col>
                <Col span={6} offset={1}>
                    <Button type="primary" onClick={() => {
                        onConfirm(props.parameterInfoConfig.parameters, props.parameterInfoConfig.cardId)
                    }}>
                        {intl.formatMessage({
                            id: 'confirm',
                        })}
                    </Button>
                </Col>
            </Row>
        )}
    </Form>;
};

export default connect((state: IStore, props: any) => {
    return {
        parameterInfoConfig: state.parameterInfoConfig
    }
})(ParameterInfo);