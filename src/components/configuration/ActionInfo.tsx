import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Row, Col, Form, Input, Select, Button } from 'antd';
import { connect, Dispatch } from 'umi';
import { DispatchProp } from 'react-redux';
import { IModal, IViewAction } from '@/interfaces'
import { getLocaleText, getInputValue } from '@/util'
import { actionInfoStateProps } from '@/models/actionInfoConfig'
import { IStore } from '@/store'
const { Option } = Select;

interface ActionInfoProps extends DispatchProp {
    actionInfoConfig: actionInfoStateProps,
}

const ActionInfo = (props: ActionInfoProps) => {
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
                <Form.Item label="标题">
                    <Input onChange={(value) => { onChange("Text", value.target.value) }}
                        value={props.actionInfoConfig.Text} />
                </Form.Item>
                <Form.Item label="操作类型">
                    <Select
                        onChange={(value) => { onChange("Type", value) }}
                        value={props.actionInfoConfig.Type} >
                        <Option value="CREATE">添加</Option>
                        <Option value="UPDATE">修改</Option>
                        <Option value="DELETE">删除</Option>
                        <Option value="IMPORT">导入</Option>
                        <Option value="EXPORT">导出</Option>
                        <Option value="SELECTIMPORT">选中行导入</Option>
                        <Option value="SELECTEXPORT">选中行导出</Option>
                        <Option value="MULTIUPDATE">批量修改</Option>
                        <Option value="MULTIEXPORT">批量导出</Option>
                    </Select>
                </Form.Item>
                <Form.Item label="二次确认" >
                    <Select
                        onChange={(value) => { onChange("DoubleCheck", value === "true") }}
                        value={props.actionInfoConfig.DoubleCheck ? "true" : "false"} >
                        <Option value="true">是</Option>
                        <Option value="false">否</Option>
                    </Select>
                </Form.Item>
            </Col>
        </Row>
        {
            (props.actionInfoConfig.dirty) && (
                <Row gutter={[16, 32]}>
                    <Col span={6} offset={8}>
                        <Button onClick={() => { onCancel() }}>
                            取消
        </Button>
                    </Col>
                    <Col span={6} offset={1}>
                        <Button type="primary" onClick={() => {
                            onConfirm(props.actionInfoConfig)
                        }}>
                            保存
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