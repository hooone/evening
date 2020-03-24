import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Row, Col, Form, Input, Select, Button } from 'antd';
import { connect } from 'umi';
import { DispatchProp } from 'react-redux';
import { IValueChange } from '@/interfaces'
import { getLocaleText, getInputValue } from '@/util'
import { pointChartStateProps, confirmPointChartConfigCommand, showPointChartConfigCommand } from '@/models/pointChartConfig'
import { IStore } from '@/store'
const { Option } = Select;

interface PointChartInfoProps extends DispatchProp {
    pointChartConfig: pointChartStateProps,
}

const PointChartInfo = (props: PointChartInfoProps) => {
    if (props.pointChartConfig.Styles.length < 1) {
        return <span></span>
    }
    function onViewConfirm() {
        let cmd: confirmPointChartConfigCommand = {
            type: 'pointChartConfig/confirm',
            ...props.pointChartConfig
        }
        props.dispatch(cmd);
    }
    function onViewCancel() {
        let cmd: showPointChartConfigCommand = {
            type: 'pointChartConfig/show',
            Fields: props.pointChartConfig.Fields,
            Styles: props.pointChartConfig.Styles,
        }
        props.dispatch(cmd);
    }
    function onChange(name: string, value: any) {
        let cmd: IValueChange = {
            type: 'pointChartConfig/dirty',
            id: 0,
            name: name,
            value: value,
        }
        props.dispatch(cmd)
    }
    let fieldOption: ReactElement[] = [];
    let numberFieldOption1: ReactElement[] = [];
    let numberFieldOption2: ReactElement[] = [];
    numberFieldOption2.push(<Option key={"field2_0"} value={0}>{"不显示"}</Option>)
    props.pointChartConfig.Fields.forEach(f => {
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
                    label="横坐标"
                >
                    <Select onChange={(value) => { onChange("xField", value) }}
                        value={props.pointChartConfig.xField} >
                        {fieldOption}
                    </Select>
                </Form.Item>
                <Form.Item label="纵坐标1">
                    <Select onChange={(value) => { onChange("y1Field", value) }}
                        value={props.pointChartConfig.y1Field} >
                        {numberFieldOption1}
                    </Select>
                </Form.Item>
                {(props.pointChartConfig.y1Field !== 0) && <Form.Item label="纵坐标1颜色">
                    <Input onChange={(value) => { onChange("y1Color", value.target.value) }}
                        value={props.pointChartConfig.y1Color} />
                </Form.Item>}
                <Form.Item label="纵坐标2">
                    <Select onChange={(value) => { onChange("y2Field", value) }}
                        value={props.pointChartConfig.y2Field} >
                        {numberFieldOption2}
                    </Select>
                </Form.Item>
                {(props.pointChartConfig.y2Field !== 0) && <Form.Item label="纵坐标2颜色">
                    <Input onChange={(value) => { onChange("y2Color", value.target.value) }}
                        value={props.pointChartConfig.y2Color} />
                </Form.Item>}
            </Col>
        </Row>
        {
            (props.pointChartConfig.dirty) && (
                <Row gutter={[16, 32]}>
                    <Col span={6} offset={8}>
                        <Button onClick={() => { onViewCancel() }}>
                            {"撤销"}
                        </Button>
                    </Col>
                    <Col span={6} offset={1}>
                        <Button type="primary"
                            onClick={() => { onViewConfirm() }}>
                            {"保存"}
                        </Button>
                    </Col>
                </Row>
            )
        }
    </Form >;
};
export default connect((state: IStore, props: any) => {
    return {
        pointChartConfig: state.pointChartConfig
    }
})(PointChartInfo);