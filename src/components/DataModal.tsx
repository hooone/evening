import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Modal, Button, Form, InputNumber, Input, Select } from 'antd';
import { connect, useIntl, getLocale, Dispatch } from 'umi';
import { getLocaleText, getInputValue } from '@/util'
import { ExclamationCircleOutlined } from '@ant-design/icons';
import { DispatchProp } from 'react-redux';
import { IStore } from '@/store'
import { modalStateProps, ICancelCommand, IHandleOkCommand } from '@/models/modal'
const { confirm } = Modal;
const { Option } = Select;

interface ModalProps extends DispatchProp {
    modal: modalStateProps,
}

const DataModal = (props: ModalProps) => {
    let rd = Math.floor((Math.random() * 100000) + 1);
    if (!props.modal.visible || props.modal.action === null || props.modal.action === undefined) {
        return (<span></span>)
    }
    const intl = useIntl();
    function onChange(id: string, value: string) {
        let dom = document.getElementById(id);
        if (dom) {
            dom.setAttribute('data-value', value);
        }
    }
    function handleCancel() {
        let modalCancel: ICancelCommand = {
            type: 'modal/cancel',
        }
        props.dispatch(modalCancel);
    }

    function handleOk() {
        let formdata = {
            __Key: props.modal.record.__Key
        };
        props.modal.action.Parameters.forEach(
            (parameter, idx) => {
                let paName = parameter.Field.Name
                if (props.modal.action.Type === "READ") {
                    paName = "param" + idx;
                }
                if ((parameter.IsVisible && parameter.IsEditable) || typeof (props.modal.record[paName]) === 'undefined') {
                    if (parameter.Field.Type === "int") {
                        formdata[paName] = parseInt(getInputValue('param_' + idx))
                    }
                    else if (parameter.Field.Type === "bool") {
                        let v = getInputValue('param_' + idx)
                        formdata[paName] = v && v.toUpperCase() === 'TRUE'
                    }
                    // else if (parameter.Field.Type === "datetime") {
                    //     formdata[paName] = new DataMethods().Offset2Datetime(inputValue('param_' + idx));
                    // }
                    // else if (parameter.Field.Type === "date") {
                    //     formdata[paName] = new DataMethods().Offset2Date(inputValue('param_' + idx));
                    // }
                    else {
                        formdata[paName] = getInputValue('param_' + idx)
                    }
                }
                else {
                    // if (parameter.Field.Type === "datetime") {
                    //     formdata[paName] = new DataMethods().Offset2Datetime(this.state.record[paName]);
                    // }
                    // else if (parameter.Field.Type === "date") {
                    //     formdata[paName] = new DataMethods().Offset2Date(this.state.record[paName]);
                    // }
                    // else 
                    {
                        formdata[paName] = props.modal.record[paName]
                    }
                }
            }
        )
        // let action = this.state.action
        // action.Parameters.forEach(
        //     parameter => {
        //         parameter.Default = formdata[parameter.Field.Name]
        //     }
        // )
        // this.setState({
        //     confirmLoading: true,
        //     record: formdata,
        //     action: action
        // })

        if (props.modal.action.DoubleCheck) {

            confirm({
                title:
                    (<span><span>
                        {intl.formatMessage(
                            {
                                id: 'confirm',
                            }
                        )}
                    </span>
                        <span>
                            {getLocaleText(props.modal.action.Locale) + "?"}
                        </span>
                    </span>)
                ,
                icon: <ExclamationCircleOutlined />,
                onOk() {
                    let handleOk: IHandleOkCommand = {
                        type: 'modal/handleOk',
                        action: props.modal.action,
                        formdata: formdata
                    }
                    props.dispatch(handleOk);
                },
                onCancel() {

                },
            });
        }
        else {
            let handleOk: IHandleOkCommand = {
                type: 'modal/handleOk',
                action: props.modal.action,
                formdata: formdata
            }
            props.dispatch(handleOk);
        }
    }
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

    let formitems: React.ReactElement[] = [];
    props.modal.action.Parameters.forEach(
        (parameter, idx) => {
            let value = parameter.Default;
            if (props.modal.action.Type === "UPDATE" || props.modal.action.Type === "DELETE") {
                value = props.modal.record[parameter.Field.Name];
            }
            else {
                // if (parameter.Field.Type === "datetime") {
                //     value = new DataMethods().Offset2Datetime(value);
                // }
                // else if (parameter.Field.Type === "date") {
                //     value = new DataMethods().Offset2Date(value);
                // }
            }
            if (parameter.IsVisible || parameter.IsEditable) {
                let label = (<span>
                    {getLocaleText(parameter.Field.Locale)}
                </span>);
                if (props.modal.action.Type === "READ" && (parameter.Compare === "lt" || parameter.Compare === "le")) {
                    label = (<span>{getLocaleText(parameter.Field.Locale) + ' ' +
                        intl.formatMessage(
                            {
                                id: 'max',
                            }
                        )}</span>)
                } else if (props.modal.action.Type === "READ" && (parameter.Compare === "gt" || parameter.Compare === "ge")) {
                    label = (<span>{getLocaleText(parameter.Field.Locale) + ' ' +
                        intl.formatMessage(
                            {
                                id: 'min',
                            }
                        )}</span>)
                }
                if (parameter.Field.Type === 'int') {
                    formitems.push(
                        <Form.Item label={label} key={'paramItem_' + idx} style={{ marginBottom: '22px' }}>
                            <InputNumber id={'param_' + idx} key={'param_' + idx + '_' + rd}
                                style={{ width: '100%' }} defaultValue={parseInt(value)} disabled={!(parameter.IsEditable)} />
                        </ Form.Item>
                    )
                }
                // else if (parameter.Field.Type === 'datetime') {
                //     formitems.push(
                //         <Form.Item label={label} key={'paramItem_' + idx} style={{ marginBottom: '22px' }}>
                //             <DatePicker showTime id={'param_' + idx} key={'param_' + idx + '_' + rd} style={{ width: '100%' }}
                //                 defaultValue={moment(value, datetimeFormat)} format={datetimeFormat} disabled={!(parameter.IsEditable)}
                //             />
                //         </Form.Item>
                //     )
                // }
                // else if (parameter.Field.Type === 'date') {
                //     formitems.push(
                //         <Form.Item label={label} key={'paramItem_' + idx} style={{ marginBottom: '22px' }}>
                //             <DatePicker id={'param_' + idx} key={'param_' + idx + '_' + rd} style={{ width: '100%' }}
                //                 defaultValue={moment(value, dateFormat)} format={dateFormat} disabled={!(parameter.IsEditable)} />
                //         </Form.Item>
                //     )
                // }
                else if (parameter.Field.Type === 'bool') {
                    formitems.push(
                        <Form.Item label={label} key={'paramItem_' + idx} style={{ marginBottom: '22px' }}>
                            <Select id={'param_' + idx} key={'param_' + idx + '_' + rd}
                                disabled={!(parameter.IsEditable)}
                                defaultValue={value + ''} data-value={value + ''} style={{ width: '100%' }}
                                onChange={(value) => onChange('param_' + idx, value)}>
                                <Option value="true">{intl.formatMessage(
                                    {
                                        id: 'yes',
                                    }
                                )}</Option>
                                <Option value="false">{intl.formatMessage(
                                    {
                                        id: 'no',
                                    }
                                )}</Option>
                            </Select>
                        </Form.Item>
                    )
                }
                else {
                    formitems.push(
                        <Form.Item label={label} key={'paramItem_' + idx} style={{ marginBottom: '22px' }}>
                            <Input id={'param_' + idx} key={'param_' + idx + '_' + rd} defaultValue={value}
                                disabled={!(parameter.IsEditable)} />
                        </Form.Item>
                    )
                }
            }
        }
    )
    return (<Modal
        title={getLocaleText(props.modal.action.Locale)}
        visible={props.modal.visible}
        onOk={() => { handleOk() }}
        onCancel={() => { handleCancel() }}
        footer={[
            <Button key="back" onClick={() => { handleCancel() }}>
                {intl.formatMessage(
                    {
                        id: 'cancel',
                    }
                )}
            </Button>,
            <Button key="submit" type="primary" loading={props.modal.confirmLoading} onClick={() => { handleOk() }}>
                {intl.formatMessage(
                    {
                        id: 'confirm',
                    }
                )}
            </Button>,
        ]}
    >
        <Form {...formItemLayout} labelAlign='left' >
            {formitems}
        </Form>
    </Modal >)
};
export default connect((state: IStore) => {
    return {
        modal: state.modal,
    }
})(DataModal);