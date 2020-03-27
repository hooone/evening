import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Row, Col, Form, Input, Select, Button } from 'antd';
import { connect, useIntl } from 'umi';
import { DispatchProp } from 'react-redux';
import { IValueChange } from '@/interfaces'
import { getLocaleText, getInputValue } from '@/util'
import { rectChartStateProps, confirmRectChartConfigCommand, showRectChartConfigCommand } from '@/models/rectChartConfig'
import { IStore } from '@/store'
const { Option } = Select;

interface RectChartInfoProps extends DispatchProp {
    rectChartConfig: rectChartStateProps,
}

const RectChartInfo = (props: RectChartInfoProps) => {
    const intl = useIntl();
    if (props.rectChartConfig.Styles.length < 1) {
        return <span></span>
    }
    function onViewConfirm() {
        let cmd: confirmRectChartConfigCommand = {
            type: 'rectChartConfig/confirm',
            ...props.rectChartConfig
        }
        props.dispatch(cmd);
    }
    function onViewCancel() {
        let cmd: showRectChartConfigCommand = {
            type: 'rectChartConfig/show',
            Fields: props.rectChartConfig.Fields,
            Styles: props.rectChartConfig.Styles,
        }
        props.dispatch(cmd);
    }
    function onChange(name: string, value: any) {
        let cmd: IValueChange = {
            type: 'rectChartConfig/dirty',
            id: 0,
            name: name,
            value: value,
        }
        props.dispatch(cmd)
    }
    let fieldOption: ReactElement[] = [];
    let numberFieldOption1: ReactElement[] = [];
    let numberFieldOption2: ReactElement[] = [];
    numberFieldOption2.push(<Option key={"field2_0"} value={0}>{intl.formatMessage({
        id: 'none',
    })}</Option>)
    props.rectChartConfig.Fields.forEach(f => {
        fieldOption.push(<Option key={"xfield_" + f.Id} value={f.Id}>{getLocaleText(f.Locale)}</Option>)
        if (f.Type === 'int' || f.Type === 'float') {
            numberFieldOption1.push(<Option key={"field1_" + f.Id} value={f.Id}>{getLocaleText(f.Locale)}</Option>)
            numberFieldOption2.push(<Option key={"field2_" + f.Id} value={f.Id}>{getLocaleText(f.Locale)}</Option>)
        }
    })
    return <Form layout="vertical" hideRequiredMark>
        <Row gutter={48}>
            <Col span={24} style={{ paddingLeft: 0 }}>
                <Form.Item
                    label={intl.formatMessage({
                        id: 'horcoor',
                    })}
                >
                    <Select onChange={(value) => { onChange("xField", value) }}
                        value={props.rectChartConfig.xField} >
                        {fieldOption}
                    </Select>
                </Form.Item>
                <Form.Item label={intl.formatMessage({
                    id: 'vercoor',
                }) + " 1"}>
                    <Select onChange={(value) => { onChange("y1Field", value) }}
                        value={props.rectChartConfig.y1Field} >
                        {numberFieldOption1}
                    </Select>
                </Form.Item>
                {(props.rectChartConfig.y1Field !== 0) && < Form.Item
                    label={intl.formatMessage({
                        id: 'vercoor',
                    }) + "1" + intl.formatMessage({
                        id: 'type',
                    })}>
                    <Select onChange={(value) => { onChange("y1Type", value) }}
                        value={props.rectChartConfig.y1Type} >
                        <Option value={"BAR"}>{intl.formatMessage({
                            id: 'bar',
                        })}</Option>
                        <Option value={"LINE"}>{intl.formatMessage({
                            id: 'line',
                        })}</Option>
                    </Select>
                </Form.Item>}
                {(props.rectChartConfig.y1Field !== 0) &&
                    <Form.Item label={intl.formatMessage({
                        id: 'vercoor',
                    }) + "1" + intl.formatMessage({
                        id: 'color',
                    })}>
                        <Input onChange={(value) => { onChange("y1Color", value.target.value) }}
                            value={props.rectChartConfig.y1Color} />
                    </Form.Item>}
                <Form.Item label={intl.formatMessage({
                    id: 'vercoor',
                }) + " 2"}>
                    <Select onChange={(value) => { onChange("y2Field", value) }}
                        value={props.rectChartConfig.y2Field} >
                        {numberFieldOption2}
                    </Select>
                </Form.Item>
                {(props.rectChartConfig.y2Field !== 0) && < Form.Item
                    label={intl.formatMessage({
                        id: 'vercoor',
                    }) + "2" + intl.formatMessage({
                        id: 'type',
                    })}>
                    <Select onChange={(value) => { onChange("y2Type", value) }}
                        value={props.rectChartConfig.y2Type} >
                        <Option value={"BAR"}>{intl.formatMessage({
                            id: 'bar',
                        })}</Option>
                        <Option value={"LINE"}>{intl.formatMessage({
                            id: 'line',
                        })}</Option>
                    </Select>
                </Form.Item>}
                {(props.rectChartConfig.y2Field !== 0) &&
                    <Form.Item label={intl.formatMessage({
                        id: 'vercoor',
                    }) + "2" + intl.formatMessage({
                        id: 'color',
                    })}>
                        <Input onChange={(value) => { onChange("y2Color", value.target.value) }}
                            value={props.rectChartConfig.y2Color} />
                    </Form.Item>}
            </Col>
        </Row>
        {
            (props.rectChartConfig.dirty) && (
                <Row gutter={[16, 32]}>
                    <Col span={6} offset={8}>
                        <Button onClick={() => { onViewCancel() }}>
                            {intl.formatMessage({
                                id: 'cancel',
                            })}
                        </Button>
                    </Col>
                    <Col span={6} offset={1}>
                        <Button type="primary"
                            onClick={() => { onViewConfirm() }}>
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
export default connect((state: IStore, props: any) => {
    return {
        rectChartConfig: state.rectChartConfig
    }
})(RectChartInfo);