import React from 'react';
import { Card, Table, Divider, Descriptions, Empty, Row, Col, Statistic, Modal, Button, Form, InputNumber, Input, Select, DatePicker } from 'antd';
import { connect, useIntl } from 'umi';
import { Chart, Geom, Axis, Tooltip, Legend, Coord } from 'bizcharts';
import { UploadOutlined, PlusOutlined, EditOutlined, SettingOutlined, SearchOutlined, ExclamationCircleOutlined } from '@ant-design/icons';
const { confirm } = Modal;

const <%- card.Name.replace(card.Name[0],card.Name[0].toUpperCase()) %>Card = (props: any) => {
    const intl = useIntl();
    const dateFormat = 'YYYY-MM-DD';
    const datetimeFormat = 'YYYY-MM-DD HH:mm:ss';
    const formItemLayout = {
        labelCol: {
            xs: { span: 24 },
            sm: { span: 6 },
        },
        wrapperCol: {
            xs: { span: 24 },
            sm: { span: 18 },
        },
    };
<% if(card.Style==='RECT'){ %>
interface ISeriesData {
    xAxis: number,
    y1Axis: number,
    y2Axis: number,
    type: string,
}
    let chartData: ISeriesData[] = [];
   
    chartData = props.<%- card.Name +"model" %>.Data.map((record: any) => {
        return {
        <% card.Styles.forEach(st => {%>
            <%if (st.Property === "XAXIS"&&st.Field) {%>
            <%- '"xAxis": record.'+ st.Field.Name +',' %>
            <%}else if (st.Property === "Y1AXIS"&&st.Field) {%>
            <%- '"y1Axis": '+(st.Field.Type==="float"?"parseFloat(":"parseInt(")+'record.'+ st.Field.Name +'),' %>
            <%- '"type": intl.formatMessage({ id: "'+ st.Field.Name +'" }),' %>
            <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
            <%- '"y2Axis": '+(st.Field.Type==="float"?"parseFloat(":"parseInt(")+'record.'+ st.Field.Name +'),' %>
            <%}%>
        <% }) %>
            
        }
    })
    
    let cols = {
    <% card.Styles.forEach(st => {%>
        <%if (st.Property === "XAXIS"&&st.Field) {%>
        <%- 'xAxis: { alias: intl.formatMessage({ id: "'+ st.Field.Name +'" }) },' %>
        <%}else if (st.Property === "Y1AXIS"&&st.Field) {%>
        <%- 'y1Axis: { alias: intl.formatMessage({ id: "'+ st.Field.Name +'" }) },' %>
        <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
        <%- 'y2Axis: { alias: intl.formatMessage({ id: "'+ st.Field.Name +'" }) },' %>
        <%}%>
        
    <% }) %>
    }
    <% let setting = {
    }
    card.Styles.forEach(st => {
        if (st.Property === "XAXIS") {
            setting.xAxis = st
        }
        else if (st.Property === "Y1AXIS") {
            setting.y1Axis = st
        }
        else if (st.Property === "Y2AXIS") {
            setting.y2Axis = st
        }
        else if (st.Property === "Y1COLOR") {
            setting.y1Color = st.Value
        }
        else if (st.Property === "Y2COLOR") {
            setting.y2Color = st.Value
        }
    }) 
    if (setting.y1Axis?.FieldId != 0 && setting.y1Axis?.Value === "BAR" &&
            setting.y2Axis?.FieldId != 0 && setting.y2Axis?.Value === "BAR") {%>
    // if y1Axis and y2Axis both are BAR
    props.<%- card.Name +"model" %>.Data.forEach((record: any) => {
        chartData.push({
        <% card.Styles.forEach(st => {%>
            <%if (st.Property === "XAXIS"&&st.Field) {%>
            <%- '"xAxis": record.'+ st.Field.Name +',' %>
            <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
            <%- '"y1Axis": '+(st.Field.Type==="float"?"parseFloat(":"parseInt(")+'record.'+ st.Field.Name +'),' %>
            <%- '"type": intl.formatMessage({ id: "'+ st.Field.Name +'" }),' %>
            <%}%>
            <% }) %>
            <%- '"y2Axis": 0,' %>
                    
        })
    })

    return (
        <div>
            <Card title={<%- 'intl.formatMessage({ id: "'+card.Locale.Name+'" })' %>} className="tableCard"
                extra={<div>
                    <Button icon={<SearchOutlined />}
                        onClick={() => { props.dispatch({ type: "<%- card.Name %>model/ReadModalShow" }) }}>
                        {intl.formatMessage({ id: "read", })}
                    </Button>
                </div>}>
                <Chart forceFit={true} height={400} data={chartData} scale={cols} padding={['15%', '10%']}>
                    <Tooltip />
                <% card.Styles.forEach(st => {%>
                    <%if (st.Property === "XAXIS"&&st.Field) {%>
                    <%- '<Axis name="xAxis" title />' %>
                    <%}else if (st.Property === "Y1AXIS"&&st.Field) {%>
                    <%- '<Axis name="y1Axis" position="left" title />' %>
                    <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
                    <%- '<Axis name="y2Axis" position="right" title />' %>
                    <%}%>
                <% }) %>
                    <Geom type={'interval'} position="xAxis*y1Axis"
                        color={['type', (value: string) => {
                            if (value === cols.y1Axis.alias) {
                                return "<%- setting.y1Color ?? '#fad248' %>";
                            }
                            if (value === cols.y2Axis.alias) {
                                return "<%- setting.y2Color ?? '#blue' %>";
                            }
                            return "red";
                        }]}
                        adjust={[{
                            type: 'dodge',
                            marginRatio: 1 / 32,
                        }]}
                    />
                </Chart>
            </Card>
    <%} else {%>

    return (
        <div>
            <Card title={<%- 'intl.formatMessage({ id: "'+card.Locale.Name+'" })' %>} className="tableCard"
                extra={<div>
                    <Button icon={<SearchOutlined />}
                        onClick={() => { props.dispatch({ type: "<%- card.Name %>model/ReadModalShow" }) }}>
                        {intl.formatMessage({ id: "read", })}
                    </Button>
                </div>}>
                <Chart forceFit={true} height={400} data={chartData} scale={cols} padding={['15%', '10%']}>
                    <Tooltip />
                <% card.Styles.forEach(st => {%>
                    <%if (st.Property === "XAXIS"&&st.Field) {%>
                    <%- '<Axis name="xAxis" title />' %>
                    <%}else if (st.Property === "Y1AXIS"&&st.Field) {%>
                    <%- '<Axis name="y1Axis" position="left" title />' %>
                    <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
                    <%- '<Axis name="y2Axis" position="right" title />' %>
                    <%}%>
                <% }) %>

                
                <% card.Styles.forEach(st => {%>
                    <%if (st.Property === "Y1AXIS"&&st.Field&&setting.y1Axis.Value=="BAR") {%>
                    <%- '<Geom type={"interval"} position="xAxis*y1Axis" color={"'+(setting.y1Color ?? "#fad248")+'"} />' %>
                    <%}else if (st.Property === "Y1AXIS"&&st.Field&&st.Field&&setting.y1Axis.Value=="LINE") {%>
                    <%- '<Geom type={"line"} position="xAxis*y1Axis" color={"'+(setting.y1Color ?? "#fad248")+'"} size={3} />' %>
                    <%}else if (st.Property === "Y2AXIS"&&st.Field&&setting.y2Axis.Value=="BAR") {%>
                    <%- '<Geom type={"interval"} position="xAxis*y2Axis" color={"'+(setting.y2Color  ?? "#blue")+'"} />' %>
                    <%}else if (st.Property === "Y2AXIS"&&st.Field&&setting.y2Axis.Value=="LINE") {%>
                    <%- '<Geom type={"line"} position="xAxis*y2Axis" color={"'+(setting.y2Color ?? "#blue")+'"} size={3} />' %>
                    <%}%>
                <% }) %>
                </Chart>
            </Card>
    <%}%>
<%}%>


<% if(card.Style==='POINT'){ %>
interface ISeriesData {
    xAxis: number,
    y1Axis: number,
    y2Axis: number,
}
    let chartData: ISeriesData[] = [];
    chartData = props.<%- card.Name +"model" %>.Data.map((record: any) => {
        return {
        <% card.Styles.forEach(st => {%>
            <%if (st.Property === "XAXIS"&&st.Field) {%>
            <%- '"xAxis": record.'+ st.Field.Name +',' %>
            <%}else if (st.Property === "Y1AXIS"&&st.Field) {%>
            <%- '"y1Axis": '+(st.Field.Type==="float"?"parseFloat(":"parseInt(")+'record.'+ st.Field.Name +'),' %>
            <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
            <%- '"y2Axis": '+(st.Field.Type==="float"?"parseFloat(":"parseInt(")+'record.'+ st.Field.Name +'),' %>
            <%}%>
        <% }) %>
        }
    })
    let cols = {
    <% card.Styles.forEach(st => {%>
        <%if (st.Property === "XAXIS"&&st.Field) {%>
        <%- 'xAxis: { alias: intl.formatMessage({ id: "'+ st.Field.Name +'" }) },' %>
        <%}else if (st.Property === "Y1AXIS"&&st.Field) {%>
        <%- 'y1Axis: { alias: intl.formatMessage({ id: "'+ st.Field.Name +'" }) },' %>
        <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
        <%- 'y2Axis: { alias: intl.formatMessage({ id: "'+ st.Field.Name +'" }) },' %>
        <%}%>
       
    <% }) %>
    }
    return (
        <div>
            <Card title={<%- 'intl.formatMessage({ id: "'+card.Locale.Name+'" })' %>} className="tableCard"
                extra={<div>
                    <Button icon={<SearchOutlined />}
                        onClick={() => { props.dispatch({ type: "<%- card.Name %>model/ReadModalShow" }) }}>
                        {intl.formatMessage({ id: "read", })}
                    </Button>
                </div>}>
                <Chart forceFit={true} height={400} data={chartData} scale={cols} padding={['15%', '10%']}>
                    <Tooltip />
                <% card.Styles.forEach(st => {%>
                    <%if (st.Property === "XAXIS"&&st.Field) {%>
                    <%- '<Axis name="xAxis" title />' %>
                    <%}else if (st.Property === "Y1AXIS"&&st.Field) {%>
                    <%- '<Axis name="y1Axis" position="left" title />' %>
                    <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
                    <%- '<Axis name="y2Axis" position="right" title />' %>
                    <%}%>
                <% }) %>

                
                <% card.Styles.forEach(st => {%>
                    <%if (st.Property === "Y1AXIS"&&st.Field) {%>
                    <%- '<Geom type={"point"} shape={"circle"} position="xAxis*y1Axis" color={"'+(st.Value ?? "#fad248")+'"} size={3} />' %>
                    <%}else if (st.Property === "Y2AXIS"&&st.Field) {%>
                    <%- '<Geom type={"point"} shape={"circle"} position="xAxis*y2Axis" color={"'+(st.Value ?? "#blue")+'"} size={3} />' %>
                    <%}%>
                <% }) %>
                </Chart>
            </Card>
<%}%>



<% if(card.Style==='STAT'){ %>
    if (!props.<%- card.Name +"model" %>.Data || props.<%- card.Name +"model" %>.Data.length < 1) {
        return (<Empty />)
    }
    return (
        <div>
            <Card className="tableCard">
                <Row>
                    <Col span={24}>
                        <Row gutter={16}>
                            <% card.Fields.forEach(
                                field => {
                                    if (field.IsVisible) {%>
                            <%- '<Col span={'+Math.max(3, Math.floor(24 / card.Fields.length))+'}>' %>
                                <Statistic title={<%- 'intl.formatMessage({ id: "'+field.Locale.Name+'" })' %>}
                                    value={props.<%- card.Name +"model" %>.Data[0]["<%- field.Name %>"]} />
                            </Col>
                            <% }}) %>
                        </Row>
                    </Col>
                </Row>
            </Card>
<%}%>

<% if(card.Style==='DESC'){ %>
    if (!props.<%- card.Name +"model" %>.Data || props.<%- card.Name +"model" %>.Data.length < 1) {
        return (<Empty />)
    }
    return (
        <div>
            <Card title={<%- 'intl.formatMessage({ id: "'+card.Locale.Name+'" })' %>} className="tableCard">
                <Descriptions bordered size={'default'}>
                    <% card.Fields.forEach(
                        field => {
                            if (field.IsVisible) {%>
                    <%- '<Descriptions.Item label={intl.formatMessage({ id: "'+field.Locale.Name+'" })}>' %>
                        {props.<%- card.Name +"model" %>.Data[0]["<%- field.Name %>"]}
                    </Descriptions.Item>
                    <% }}) %>
                </Descriptions>
            </Card>
<%}%>



<% if(card.Style==='TABLE'){ %>
    const columns: any[] = [
        <% card.Fields.forEach(
            field => {
                if (field.IsVisible) {%>
        <%- "{"%>
            <%- 'title: intl.formatMessage({ id: "'+field.Locale.Name+'" }),' %>
            <%- 'dataIndex: "'+ field.Name + '",'%>
            <%- 'key: "'+ field.Name+ '",' %>
        <%- "}," %>
                <%}
            }
        ) %>
        {
            title: (<div>{intl.formatMessage({ id: 'actions' })}</div>),
            key: 'action',
            render: (text: any, record: any) => (
                <div>
                    <% card.Actions.filter(act=>(act.Type=="UPDATE"||act.Type=="DELETE")).forEach(
                    (action, idx) => {
                    let needInput=false;
                    action.Parameters.forEach(param=>{
                        needInput=needInput||(param.IsEditable&&param.IsVisible)
                    })
                    %>
                <% if(idx>0) {%>
                    <%-'<Divider type="vertical" />' %>
                <%}%>
                    <% if (!needInput){%>
                    <% if (!action.DoubleCheck){%>
                    <%-'<span onClick={() => { props.dispatch({ type: "'+card.Name+'model/'+action.Name+'", "record": record }) }} style={{ color: "rgb(64,144,255)", cursor: "pointer" }}>' %>
                    <%} else {%>
                    <%-'<span onClick={() => { '%>
                        <%-'confirm({'%>
                            <%-'title: (<span>'%>
                                <%-'<span>{intl.formatMessage({ id: "confirm", })}'%>
                                <%-'</span>'%>
                                <%-'<span>{intl.formatMessage({ id: "'+action.Name+'", })} </span>'%>
                            <%-'</span>),'%>
                            <%-'icon: <ExclamationCircleOutlined />,'%>
                            <%-'onOk() {'%>
                                <%-'props.dispatch({ type: "'+card.Name+'model/'+action.Name+'", "record": record })'%>
                            <%-'},'%>
                        <%-'})'%>
                    <%-'}}'%>
                        <%-'style={{ color: "rgb(64,144,255)", cursor: "pointer" }}>' %>
                  
                    <%}%>    
                    <%} else {%>
                    <%-'<span onClick={() => { props.dispatch({ type: "'+card.Name+'model/'+action.Name+'ModalShow", "record": record }) }} style={{ color: "rgb(64,144,255)", cursor: "pointer" }}>' %>
                    <%}%>
                        <%- '{intl.formatMessage({ id: "'+action.Locale.Name+'" })}' %>
                    <%-"</span>" %>
                        <%}
                    )%>
                </div>
            )
        }
    ]
    return (
        <div>
            <Card title={<%- 'intl.formatMessage({ id: "'+card.Locale.Name+'" })' %>} className="tableCard"
                extra={<div>
                <% card.Actions.forEach(action=>{ 
                    if(action.Type === "CREATE"){
                    let needInput=false;
                    action.Parameters.forEach(param=>{
                        needInput=needInput||(param.IsEditable&&param.IsVisible)
                    })%>
                    <%- "<Button icon={<PlusOutlined />} " %>
                        <% if(true) { %>
                        <%- 'onClick={() => { props.dispatch({ type: "'+card.Name+'model/'+action.Name+'ModalShow" }) }} >' %>
                        <% }else{%>
                        <% if(!action.DoubleCheck){%>
                        <%- 'onClick={() => { props.dispatch({ type: "'+card.Name+'model/'+action.Name+'" }) }} >' %>
                        <% }else{%>
                        <%- 'onClick={() => { ' %>
                            <%-'confirm({'%>
                                <%-'title: (<span>'%>
                                    <%-'<span>{intl.formatMessage({ id: "confirm", })}'%>
                                    <%-'</span>'%>
                                    <%-'<span>{intl.formatMessage({ id: "'+action.Name+'", })} </span>'%>
                                <%-'</span>),'%>
                                <%-'icon: <ExclamationCircleOutlined />,'%>
                                <%-'onOk() {'%>
                                    <%-'props.dispatch({ type: "'+card.Name+'model/'+action.Name+'" })'%>
                                <%-'},'%>
                            <%-'})'%>
                        <%- '}}>' %>
                        <%}%>
                        <%}%>
                        <%- '{intl.formatMessage({ id: "'+action.Locale.Name+'", })}' %>
                    <%- "</Button>" %>
                <%  }}) %>

                    <Button icon={<SearchOutlined />}
                    <% card.Actions.forEach(action=>{ 
                        if(action.Type === "READ"){
                        let needInput=false;
                        action.Parameters.forEach(param=>{
                            needInput=needInput||(param.IsEditable&&param.IsVisible)
                        })%>
                        <% if(needInput) {%>
                        <%- 'onClick={() => { props.dispatch({ type: "'+ card.Name +'model/ReadModalShow" }) }}>'  %>
                        <% }else {%>
                        <%- 'onClick={() => { props.dispatch({ type: "'+ card.Name +'model/Read" }) }}>'  %>
                        <% } %>
                        <% }}) %>
                        {intl.formatMessage({ id: "read", })}
                    </Button>
                </div>}>
                <Table
                    columns={columns}
                    dataSource={props.<%- card.Name +"model" %>.Data}
                    rowKey="__Key">
                </Table>
            </Card>
<%}%>
<% card.Actions.forEach(action=>{
    let needInput=false;
    action.Parameters.forEach(param=>{
        needInput=needInput||(param.IsEditable&&param.IsVisible)
    })
    if(needInput){%>
            <%- '<Modal' %>
                <%- 'title={intl.formatMessage({ id: "'+action.Locale.Name+'" })}' %>
                <%- 'onCancel={() => { props.dispatch({ type: "'+card.Name+'model/modalcancel" }) }}'%>
                <% if(action.DoubleCheck) {%>
                <%- 'onOk={() => { '%>
                    <%-'confirm({'%>
                        <%-'title: (<span>'%>
                            <%-'<span>{intl.formatMessage({ id: "confirm", })}'%>
                            <%-'</span>'%>
                            <%-'<span>{intl.formatMessage({ id: "'+action.Name+'", })} </span>'%>
                        <%-'</span>),'%>
                        <%-'icon: <ExclamationCircleOutlined />,'%>
                        <%-'onOk() {'%>
                            <%-'props.dispatch({ type: "'+card.Name+'model/'+action.Name+'" })'%>
                        <%-'},'%>
                    <%-'})'%>
                <%-'}}' %>
                <% }else{%>
                <%- 'onOk={() => { props.dispatch({ type: "'+card.Name+'model/'+action.Name+'" }) }}' %>
                <% }%>
                <%- 'cancelText={intl.formatMessage({ id: "cancel", })}' %>
                <%- 'okText={intl.formatMessage({ id: "confirm", })}' %>
                <%- 'visible={props.'+card.Name+'model.'+action.Name+'ModalVisible}>' %>
                <%- '<Form {...formItemLayout} labelAlign="left" fields={props.'+card.Name+'model.'+action.Name+'FormData}'%>
                    <%- 'onFieldsChange={(changedFields, allFields) => {'%>
                        <%- 'props.dispatch({ type: "'+card.Name+'model/'+action.Name+'ModalChange", Fields: allFields });'%>
                    <%- '}}>'%>
                <% action.Parameters.forEach(
                    (parameter, idx) => {
                        if (parameter.IsVisible || parameter.IsEditable) {%>
                            <%if (parameter.Field.Type === 'int') {%>
                    <Form.Item label={intl.formatMessage({ id: "<%- parameter.Field.Locale.Name %>", })<%- (parameter.Compare === "gt" || parameter.Compare === "ge")?' + intl.formatMessage({ id: "min", })':((parameter.Compare === "lt" || parameter.Compare === "le")?' + intl.formatMessage({ id: "max", })':'')%>} 
                        name={"<%- parameter.Field.Name + ((parameter.Compare === 'gt' || parameter.Compare === 'ge')?'_min':((parameter.Compare === 'lt' || parameter.Compare === 'le')?'_max':'')) %>"} style={{ marginBottom: '22px' }}>
                        <InputNumber style={{ width: '100%' }} disabled={<%- (parameter.IsEditable?"false":"true") %>} />
                    </ Form.Item>
                            <%} else if (parameter.Field.Type === 'datetime') {%>
                    <Form.Item label={intl.formatMessage({ id: "<%- parameter.Field.Locale.Name %>", })<%- (parameter.Compare === "gt" || parameter.Compare === "ge")?' + intl.formatMessage({ id: "min", })':((parameter.Compare === "lt" || parameter.Compare === "le")?' + intl.formatMessage({ id: "max", })':'')%>} 
                        name={"<%- parameter.Field.Name + ((parameter.Compare === 'gt' || parameter.Compare === 'ge')?'_min':((parameter.Compare === 'lt' || parameter.Compare === 'le')?'_max':'')) %>"} style={{ marginBottom: '22px' }}>
                        <DatePicker style={{ width: '100%' }} disabled={<%- (parameter.IsEditable?"false":"true") %>} showTime format={datetimeFormat} />
                    </ Form.Item>
                            <%} else if (parameter.Field.Type === 'date') {%>
                    <Form.Item label={intl.formatMessage({ id: "<%- parameter.Field.Locale.Name %>", })<%- (parameter.Compare === "gt" || parameter.Compare === "ge")?' + intl.formatMessage({ id: "min", })':((parameter.Compare === "lt" || parameter.Compare === "le")?' + intl.formatMessage({ id: "max", })':'')%>} 
                        name={"<%- parameter.Field.Name + ((parameter.Compare === 'gt' || parameter.Compare === 'ge')?'_min':((parameter.Compare === 'lt' || parameter.Compare === 'le')?'_max':'')) %>"} style={{ marginBottom: '22px' }}>
                        <DatePicker style={{ width: '100%' }} disabled={<%- (parameter.IsEditable?"false":"true") %>} format={dateFormat} />
                    </ Form.Item>
                            <%}else if (parameter.Field.Type === 'bool') {%>
                    <Form.Item label={intl.formatMessage({ id: "<%- parameter.Field.Locale.Name %>", })<%- (parameter.Compare === "gt" || parameter.Compare === "ge")?' + intl.formatMessage({ id: "min", })':((parameter.Compare === "lt" || parameter.Compare === "le")?' + intl.formatMessage({ id: "max", })':'')%>} 
                        name={"<%- parameter.Field.Name + ((parameter.Compare === 'gt' || parameter.Compare === 'ge')?'_min':((parameter.Compare === 'lt' || parameter.Compare === 'le')?'_max':'')) %>"} style={{ marginBottom: '22px' }}>
                        <Select disabled={<%- (parameter.IsEditable?"false":"true") %>} style={{ width: '100%' }}>
                            <Option value="true">{intl.formatMessage({ id: 'yes'})}</Option>
                            <Option value="false">{intl.formatMessage({ id: 'no'})}</Option>
                        </Select>
                    </ Form.Item>
                            <%}else {%>
                    <Form.Item label={intl.formatMessage({ id: "<%- parameter.Field.Locale.Name %>", })<%- (parameter.Compare === "gt" || parameter.Compare === "ge")?' + intl.formatMessage({ id: "min", })':((parameter.Compare === "lt" || parameter.Compare === "le")?' + intl.formatMessage({ id: "max", })':'')%>} 
                        name={"<%- parameter.Field.Name + ((parameter.Compare === 'gt' || parameter.Compare === 'ge')?'_min':((parameter.Compare === 'lt' || parameter.Compare === 'le')?'_max':'')) %>"} style={{ marginBottom: '22px' }}>
                        <Input disabled={<%- (parameter.IsEditable?"false":"true") %>} />
                    </ Form.Item>
                            <%}%>
                            <%}
                    }
                )%>
                <%- '</Form>' %>
            <%- '</Modal>' %>
<%}})%>
        </div>
    )
}
export default connect((state: any) => {
    return {
        <%- card.Name +"model" %>: state.<%- card.Name +"model" %>,
    }
})(<%- card.Name.replace(card.Name[0],card.Name[0].toUpperCase()) +"Card" %>);