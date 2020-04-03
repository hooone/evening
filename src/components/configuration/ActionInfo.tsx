import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Row, Col, Form, Input, Select, Button } from 'antd';
import { connect, useIntl } from 'umi';
import { DispatchProp } from 'react-redux';
import { IModal, IViewAction } from '@/interfaces'
import { actionInfoStateProps } from '@/models/actionInfoConfig'
import { IStore } from '@/store'
const { Option } = Select;

interface ActionInfoProps extends DispatchProp {
    actionInfoConfig: actionInfoStateProps,
}

const ActionInfo = (props: ActionInfoProps) => {
    const intl = useIntl();
    if (!props.actionInfoConfig.visible) {
        return <span></span>
    }
    function onConfirm(action: IViewAction) {
        props.dispatch({
            type: 'actionInfoConfig/saveAction',
            action: action,
        });
    }
    function onCancel() {
        props.dispatch({ type: 'draw/subClose' });
    }
    function onChange(name: string, value: any) {
        props.dispatch({
            type: 'actionInfoConfig/dirty',
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
                        value={props.actionInfoConfig.Name} />
                </Form.Item>
                <Form.Item label={intl.formatMessage({
                    id: 'title',
                })}>
                    <Input onChange={(value) => { onChange("Text", value.target.value) }}
                        value={props.actionInfoConfig.Text} />
                </Form.Item>
                <Form.Item label={intl.formatMessage({
                    id: 'actiontype',
                })}>
                    <Select
                        onChange={(value) => { onChange("Type", value) }}
                        value={props.actionInfoConfig.Type} >
                        <Option value="CREATE">{intl.formatMessage({
                            id: 'create',
                        })}</Option>
                        <Option value="UPDATE">{intl.formatMessage({
                            id: 'update',
                        })}</Option>
                        <Option value="DELETE">{intl.formatMessage({
                            id: 'delete',
                        })}</Option>
                    </Select>
                </Form.Item>
                <Form.Item label={intl.formatMessage({
                    id: 'doublecheck',
                })} >
                    <Select
                        onChange={(value) => { onChange("DoubleCheck", value === "true") }}
                        value={props.actionInfoConfig.DoubleCheck ? "true" : "false"} >
                        <Option value="true">{intl.formatMessage({
                            id: 'yes',
                        })}</Option>
                        <Option value="false">{intl.formatMessage({
                            id: 'no',
                        })}</Option>
                    </Select>
                </Form.Item>
            </Col>
        </Row>
        {
            (props.actionInfoConfig.dirty) && (
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
                            onConfirm(props.actionInfoConfig)
                        }}>
                            {intl.formatMessage({
                                id: 'confirm',
                            })}
                        </Button>
                    </Col>
                </Row>
            )
        }
    </Form >;
};

export default connect((state: IStore) => {
    return {
        actionInfoConfig: state.actionInfoConfig,
    }
})(ActionInfo);