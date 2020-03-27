import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Row, Col, Form, Input, Select, Button } from 'antd';
import { connect, useIntl } from 'umi';
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
    const intl = useIntl();
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
        if (name === "Width" || name === "Pos") {
            props.dispatch({
                type: 'draw/cardChanged',
                name: name,
                value: value,
            })
        }
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
                    label={intl.formatMessage({
                        id: 'title',
                    })}
                >
                    <Input onChange={(value) => { onChange("Text", value.target.value) }}
                        value={props.cardInfoConfig.Text} />
                </Form.Item>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'chart',
                    })}
                >
                    <Select onChange={(value) => { onChange("Style", value) }}
                        value={props.cardInfoConfig.Style} >
                        <Option value="TABLE"> label={intl.formatMessage({
                            id: 'table',
                        })}</Option>
                        <Option value="RECT">{intl.formatMessage({
                            id: 'rectchart',
                        })}</Option>
                        <Option value="POINT">{intl.formatMessage({
                            id: 'pointchart',
                        })}</Option>
                        <Option value="DESC">{intl.formatMessage({
                            id: 'descriptiontbl',
                        })}</Option>
                        <Option value="STAT">{intl.formatMessage({
                            id: 'stat',
                        })}</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'width',
                    })}
                >
                    <Select onChange={(value) => { onChange("Width", value) }}
                        value={props.cardInfoConfig.Width + ""} >
                        <Option value="2">8.3%</Option>
                        <Option value="4">16.7%</Option>
                        <Option value="6">25%</Option>
                        <Option value="8">33.3%</Option>
                        <Option value="10">41.7%</Option>
                        <Option value="12">50%</Option>
                        <Option value="14">58.3%</Option>
                        <Option value="16">66.7%</Option>
                        <Option value="18">75%</Option>
                        <Option value="20">83.3%</Option>
                        <Option value="22">91.7%</Option>
                        <Option value="24">100%</Option>
                    </Select>
                </Form.Item>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'pos4left',
                    })}
                >
                    <Select onChange={(value) => { onChange("Pos", value) }}
                        value={props.cardInfoConfig.Pos + ""} >
                        <Option value="2">8.3%</Option>
                        <Option value="4">16.7%</Option>
                        <Option value="6">25%</Option>
                        <Option value="8">33.3%</Option>
                        <Option value="10">41.7%</Option>
                        <Option value="12">50%</Option>
                        <Option value="14">58.3%</Option>
                        <Option value="16">66.7%</Option>
                        <Option value="18">75%</Option>
                        <Option value="20">83.3%</Option>
                        <Option value="22">91.7%</Option>
                        <Option value="24">100%</Option>
                    </Select>
                </Form.Item>
            </Col>
        </Row>
        {(props.cardInfoConfig.dirty) && (
            <Row gutter={[16, 32]}>
                <Col span={6} offset={8}>
                    <Button onClick={() => { onViewCancel(props.cardInfoConfig.Id) }}>
                        {intl.formatMessage({
                            id: 'cancel',
                        })}
                    </Button>
                </Col>
                <Col span={6} offset={1}>
                    <Button type="primary" onClick={() => { onViewConfirm() }}>
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
        cardInfoConfig: state.cardInfoConfig
    }
})(CardInfo);