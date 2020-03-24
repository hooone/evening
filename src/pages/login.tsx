import React from 'react';
import styles from './login.css';
import { Store } from 'rc-field-form/lib/interface';
import { Form, Input, Button, Checkbox, Layout, message, Row, Col } from 'antd';
import reqwest from "reqwest"
import { CommonResult } from '@/interfaces';

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
};
const Login = () => {
  const onFinish = (values: Store) => {
    console.log({
      User: values.username,
      Password: values.password,
    })
    reqwest({
      url: '/login',
      type: 'json',
      method: 'post',
      data: {
        User: values.username,
        Password: values.password,
      }
      , error: function () {
        message.error('登录失败，请检查用户名密码');
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
      <Layout.Header />
      <Layout.Content>
        <Row>
          <Col offset={8} span={8}>
            <Form
              {...layout}
              name="basic"
              initialValues={{ remember: true }}
              onFinish={(values: Store) => { onFinish(values) }}
              onFinishFailed={onFinishFailed}
            >
              <Form.Item
                label="Username"
                name="username"
                rules={[{ required: true, message: 'Please input your username!' }]}
              >
                <Input />
              </Form.Item>

              <Form.Item
                label="Password"
                name="password"
                rules={[{ required: true, message: 'Please input your password!' }]}
              >
                <Input.Password />
              </Form.Item>

              <Form.Item {...tailLayout}>
                <Button type="primary" htmlType="submit">
                  Submit
        </Button>
              </Form.Item>
            </Form>
          </Col>
        </Row>
      </Layout.Content>
      <Layout.Footer />
    </Layout>
  );
};

export default Login