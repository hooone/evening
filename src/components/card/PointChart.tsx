import React, { ReactElement } from 'react';
import { Table, Popconfirm, Button, Divider } from 'antd';
import { connect, useIntl, getLocale } from 'umi';
import { getLocaleText } from '@/util'
import { Chart, Geom, Axis, Tooltip, Legend, Coord } from 'bizcharts';
import { ICard, IStyle } from '@/interfaces';
import { CardContentProps } from '@/models/card';


interface ISeriesData {
    xAxis: number,
    y1Axis: number,
    y2Axis: number,
    type: string,
}
interface setting {
    xAxis?: IStyle,
    y1Axis?: IStyle,
    y2Axis?: IStyle,
}
const RectChart = (props: CardContentProps) => {
    const intl = useIntl();
    // convert styles to setting
    let setting: setting = {
    }
    props.card.Styles.forEach(st => {
        if (st.Property === "XAXIS") {
            setting.xAxis = st
        }
        else if (st.Property === "Y1AXIS") {
            setting.y1Axis = st
        }
        else if (st.Property === "Y2AXIS") {
            setting.y2Axis = st
        }
    })
    if (!setting.xAxis?.Field || (!setting.y1Axis?.Field && !setting.y2Axis?.Field)) {
        return (<p> {intl.formatMessage(
            {
                id: 'emptyChart',
            }
        )}</p>)
    }
    let chartData: ISeriesData[] = [];
    if (props.card.data !== undefined) {
        props.card.data.forEach(record => {
            let xData;
            if (setting.xAxis && setting.xAxis.FieldId != 0) {
                xData = record[setting.xAxis.Field.Name]
            }
            let y1Data: number = 0;
            if (setting.y1Axis && setting.y1Axis.FieldId != 0 &&
                (setting.y1Axis.Field.Type === "int" || setting.y1Axis.Field.Type === "float")) {
                if (setting.y1Axis.Field.Type === "int") {
                    y1Data = parseInt(record[setting.y1Axis.Field.Name])
                }
                else if (setting.y1Axis.Field.Type === "float") {
                    y1Data = parseFloat(record[setting.y1Axis.Field.Name])
                }
            }
            let y2Data: number = 0;
            if (setting.y2Axis && setting.y2Axis.FieldId != 0 &&
                (setting.y2Axis.Field.Type === "int" || setting.y2Axis.Field.Type === "float")) {
                if (setting.y2Axis.Field.Type === "int") {
                    y2Data = parseInt(record[setting.y2Axis.Field.Name])
                }
                else if (setting.y2Axis.Field.Type === "float") {
                    y2Data = parseFloat(record[setting.y2Axis.Field.Name])
                }
            }
            chartData.push({
                xAxis: xData,
                y1Axis: y1Data,
                y2Axis: y2Data,
                type: setting.y1Axis && setting.y1Axis.Field ? getLocaleText(setting.y1Axis.Field.Locale) : "y1Axis"
            })
        })
    }

    let cols = {
        xAxis: { alias: setting.xAxis && setting.xAxis.Field ? getLocaleText(setting.xAxis.Field.Locale) : "" },
        y1Axis: { alias: setting.y1Axis && setting.y1Axis.Field ? getLocaleText(setting.y1Axis.Field.Locale) : "" },
        y2Axis: { alias: setting.y2Axis && setting.y2Axis.Field ? getLocaleText(setting.y2Axis.Field.Locale) : "" },
    }
    console.log(cols)
    console.log(chartData)
    return (<Chart forceFit={true} height={400} data={chartData} scale={cols} padding={['15%', '10%']}>
        <Tooltip />
        {(setting.xAxis) && (setting.xAxis.Field) && <Axis name="xAxis" title />}
        {(setting.y1Axis) && (setting.y1Axis.Field) && <Axis name="y1Axis" position="left" title />}
        {(setting.y2Axis) && (setting.y2Axis.Field) && <Axis name="y2Axis" position="right" title />}

        {(setting.y1Axis) && (setting.y1Axis.Field) &&
            <Geom type={'point'} shape={'circle'} position="xAxis*y1Axis" color={setting.y1Axis?.Value ?? "#fad248"} size={3} />}


        {(setting.y2Axis) && (setting.y2Axis.Field) &&
            <Geom type={'point'} shape={'circle'} position="xAxis*y2Axis" color={setting.y2Axis?.Value ?? "blue"} size={3} />}


    </Chart>)
};
export default RectChart;