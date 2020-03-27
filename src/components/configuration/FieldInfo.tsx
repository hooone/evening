import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Row, Col, Form, Input, Select, Button } from 'antd';
import { connect, useIntl } from 'umi';
import { DispatchProp } from 'react-redux';
import { IStore } from '@/store'
import { fieldInfoStateProps } from '@/models/fieldInfoConfig'
import { IField } from '@/interfaces';
const { Option } = Select;

interface FieldInfoProps extends DispatchProp {
    fieldInfoConfig: fieldInfoStateProps,
}
const FieldInfo = (props: FieldInfoProps) => {
    const intl = useIntl();
    if (!props.fieldInfoConfig.visible) {
        return <span></span>
    }
    function onConfirm(field: IField) {
        props.dispatch({
            type: 'fieldInfoConfig/saveField',
            field: field,
        });
    }
    function onCancel() {
        props.dispatch({ type: 'fieldInfoConfig/close' });
        props.dispatch({ type: 'draw/subClose' });
    }
    function onChange(name: string, value: any) {
        props.dispatch({
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
                        value={props.fieldInfoConfig.Name} />
                </Form.Item>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'title',
                    })}
                >
                    <Input onChange={(value) => { onChange("Text", value.target.value) }}
                        value={props.fieldInfoConfig.Text} />
                </Form.Item>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'visible',
                    })}
                >
                    <Select onChange={(value) => { onChange("IsVisible", value === "true") }}
                        value={props.fieldInfoConfig.IsVisible ? "true" : "false"} >
                        <Option value="true">{intl.formatMessage({
                            id: 'visible',
                        })}</Option>
                        <Option value="false">{intl.formatMessage({
                            id: 'hidden',
                        })}</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'fieldtype',
                    })}
                >
                    <Select onChange={(value) => { onChange("Type", value) }}
                        value={props.fieldInfoConfig.Type} >
                        <Option value="int">int</Option>
                        <Option value="string">string</Option>
                        <Option value="float">float</Option>
                        <Option value="date">date</Option>
                        <Option value="datetime">datetime</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'filter',
                    })}
                >
                    <Select onChange={(value) => { onChange("Filter", value) }}
                        value={props.fieldInfoConfig.Filter} >
                        <Option value="ALL">{intl.formatMessage({
                            id: 'all',
                        })}</Option>
                        <Option value="EQUAL">{intl.formatMessage({
                            id: 'equal',
                        })}</Option>
                        <Option value="RANGE">{intl.formatMessage({
                            id: 'range',
                        })}</Option>
                        <Option value="MAX">{intl.formatMessage({
                            id: 'upperlimit',
                        })}</Option>
                        <Option value="MIN">{intl.formatMessage({
                            id: 'lowerlimit',
                        })}</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'default',
                    })}
                >
                    <Input onChange={(value) => { onChange("Default", value.target.value) }}
                        value={props.fieldInfoConfig.Default} />
                </Form.Item>
            </Col>
        </Row>
        {(props.fieldInfoConfig.dirty) && (
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
                        onConfirm(props.fieldInfoConfig)
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
export default connect((state: IStore) => {
    return {
        fieldInfoConfig: state.fieldInfoConfig,
    }
})(FieldInfo);