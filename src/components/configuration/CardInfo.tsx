import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Row, Col, Form, Input, Select, Button } from 'antd';
import { connect } from 'umi';
import { DispatchProp } from 'react-redux';
import { IModal, IViewAction } from '@/interfaces'
import { getLocaleText, getInputValue } from '@/util'
import { cardInfoStateProps } from '@/models/cardInfoConfig'
import { IStore } from '@/store'
const { Option } = Select;

interface CardInfoProps extends DispatchProp {
    cardInfoConfig: cardInfoStateProps,
}

const CardInfo = (props: CardInfoProps) => {
    function onViewConfirm() {
        props.dispatch({
            type: 'cardInfoConfig/confirm',
            card: props.cardInfoConfig,
        });
    }
    function onViewCancel(cardId: number) {
        props.dispatch({
            type: 'draw/loadDraw',
            cardId: cardId,
        });
    }
    function onChange(name: string, value: any) {
        props.dispatch({
            type: 'cardInfoConfig/dirty',
            name: name,
            value: value,
        })
    }
    return <Form layout="vertical" hideRequiredMark>
        <Row gutter={48}>
            <Col span={24} style={{ paddingLeft: 0 }}>
                <Form.Item
                    label="Name"
                >
                    <Input onChange={(value) => { onChange("Name", value.target.value) }}
                        value={props.cardInfoConfig.Name} />
                </Form.Item>
                <Form.Item
                    label="标题"
                >
                    <Input onChange={(value) => { onChange("Text", value.target.value) }}
                        value={props.cardInfoConfig.Text} />
                </Form.Item>
                <Form.Item
                    label="图表"
                >
                    <Select onChange={(value) => { onChange("Style", value) }}
                        value={props.cardInfoConfig.Style} >
                        <Option value="TABLE">表格</Option>
                        <Option value="RECT">x-y图</Option>
                        <Option value="POINT">散点图</Option>
                        <Option value="DESC">详细描述</Option>
                        <Option value="STAT">统计数据</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label="宽度 (1-12)"
                >
                    <Select onChange={(value) => { onChange("Width", value) }}
                        value={props.cardInfoConfig.Width + ""} >
                        <Option value="2">1</Option>
                        <Option value="4">2</Option>
                        <Option value="6">3</Option>
                        <Option value="8">4</Option>
                        <Option value="10">5</Option>
                        <Option value="12">6</Option>
                        <Option value="14">7</Option>
                        <Option value="16">8</Option>
                        <Option value="18">9</Option>
                        <Option value="20">10</Option>
                        <Option value="22">11</Option>
                        <Option value="24">12</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label="位置 (基于左侧)"
                >
                    <Select onChange={(value) => { onChange("Pos", value) }}
                        value={props.cardInfoConfig.Pos + ""} >
                        <Option value="0">0</Option>
                        <Option value="2">1</Option>
                        <Option value="4">2</Option>
                        <Option value="6">3</Option>
                        <Option value="8">4</Option>
                        <Option value="10">5</Option>
                        <Option value="12">6</Option>
                        <Option value="14">7</Option>
                        <Option value="16">8</Option>
                        <Option value="18">9</Option>
                        <Option value="20">10</Option>
                        <Option value="22">11</Option>
                    </Select>
                </Form.Item>
            </Col>
        </Row>
        {(props.cardInfoConfig.dirty) && (
            <Row gutter={[16, 32]}>
                <Col span={6} offset={8}>
                    <Button onClick={() => { onViewCancel(props.cardInfoConfig.Id) }}>
                        撤销
                                    </Button>
                </Col>
                <Col span={6} offset={1}>
                    <Button type="primary" onClick={() => { onViewConfirm() }}>
                        保存
                                    </Button>
                </Col>
            </Row>
        )}
    </Form>;
};
export default connect((state: IStore, props: any) => {
    return {
        cardInfoConfig: state.cardInfoConfig
    }
})(CardInfo);