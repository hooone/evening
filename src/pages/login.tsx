import React from 'react';
import { useIntl } from 'umi'
import styles from './login.css';
import { Store } from 'rc-field-form/lib/interface';
import { Form, Input, Button, Checkbox, Layout, message, Row, Col } from 'antd';
import reqwest from "reqwest"
import { CommonResult } from '@/interfaces';
import { RightSquareTwoTone } from '@ant-design/icons';

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
};
const Login = () => {
  const intl = useIntl();
  const onFinish = (values: Store) => {
    reqwest({
      url: '/login',
      type: 'json',
      method: 'post',
      data: {
        User: values.username,
        Password: values.password,
      }
      , error: function () {
        message.error(intl.formatMessage(
          {
            id: 'loginfailed',
          }
        ));
      }
      , success: function (data: CommonResult) {
        window.location.href = "/home"
      }
    })
  };

  const onFinishFailed = () => {
  };

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Layout.Content>
        <div style={{
          overflow: 'auto',
          height: '300px',
          width: '100%',
          position: 'absolute',
          margin: 'auto',
          top: '0',
          bottom: '0',
          left: '0',
          right: '0',
          verticalAlign: 'middle',
          fontSize: '60px',
        }}>
          <Row>
            <Col span={8}  offset={8} >
              <Form
                {...layout}
                name="basic"
                initialValues={{ remember: true }}
                onFinish={(values: Store) => { onFinish(values) }}
                onFinishFailed={onFinishFailed}
              >
                <Form.Item
                  label={intl.formatMessage(
                    {
                      id: 'username',
                    }
                  )}
                  name="username"
                  rules={[{ required: true, message: 'Please input your username!' }]}
                >
                  <Input />
                </Form.Item>

                <Form.Item
                  label={intl.formatMessage(
                    {
                      id: 'password',
                    }
                  )}
                  name="password"
                  rules={[{ required: true, message: 'Please input your password!' }]}
                >
                  <Input.Password />
                </Form.Item>

                <Form.Item {...tailLayout}>
                  <Button type="primary" htmlType="submit">
                    {intl.formatMessage(
                      {
                        id: 'confirm',
                      }
                    )}
                  </Button>
                </Form.Item>
              </Form>
            </Col>
          </Row>
        </div>
      </Layout.Content>
    </Layout>
  );
};

export default Login