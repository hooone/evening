import React from 'react';
import { useIntl, history } from 'umi'
import { Store } from 'rc-field-form/lib/interface';
import { Typography, Form, Input, Button, Checkbox, Layout, message, Row, Col } from 'antd';
import reqwest from "reqwest"
import { CommonResult } from '@/interfaces';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

const { Title, Paragraph, Text } = Typography;
const layout = {
  wrapperCol: { span: 24 },
};
const SignUp = () => {
  const intl = useIntl();
  const onFinish = (values: Store) => {
    if (values.passwordrepeat !== values.password) {
      message.error(intl.formatMessage(
        {
          id: 'passwordnotequal',
        }
      ));
      return;
    }
    reqwest({
      url: '/api/signup',
      type: 'json',
      method: 'post',
      data: {
        User: values.username,
        Password: values.password,
      }
      , error: function (rst: CommonResult) {
        message.error(rst.Message);
      }
      , success: function (rst: CommonResult) {
        if (rst.Success) {
          window.location.href = "/home"
        } else {
          message.error(rst.Message);
        }
      }
    })
  };

  const onFinishFailed = () => {
  };

  return (
    <Layout style={{ minHeight: '100vh' }} className={"login-layout"}>
      <Layout.Content>
        <Row>
          <Col span={6} offset={9} style={{ marginTop: '48px' }}>
            <Typography style={{ padding: '12px' }}>
              <Title>
                evening
            </Title>
              <Paragraph>
                新用户注册
              </Paragraph>
            </Typography>
            <Form
              {...layout}
              name="basic"
              initialValues={{ remember: true }}
              onFinish={(values: Store) => { onFinish(values) }}
              onFinishFailed={onFinishFailed}
            >
              <Form.Item
                name="username"
                style={{ marginBottom: 0 }}
                rules={[{ required: true, message: 'Please input your username!' }]}
              >
                <Input prefix={<UserOutlined />} placeholder={intl.formatMessage(
                  {
                    id: 'username',
                  }
                )} />
              </Form.Item>

              <Form.Item
                name="password"
                style={{ marginBottom: 0 }}
                rules={[{ required: true, message: 'Please input your password!' }]}
              >
                <Input.Password prefix={<LockOutlined />} placeholder={intl.formatMessage(
                  {
                    id: 'password',
                  }
                )} />
              </Form.Item>

              <Form.Item
                name="passwordrepeat"
                style={{ marginBottom: 0 }}
                rules={[{ required: true, message: 'Please input your password!' }]}
              >
                <Input.Password prefix={<LockOutlined />} placeholder={intl.formatMessage(
                  {
                    id: 'passwordrepeat',
                  }
                )} />
              </Form.Item>

              <Form.Item name="remember" style={{ marginBottom: 0 }}>
                <div>
                  <Button type="link" style={{ float: 'right' }} onClick={() => { history.push("/login") }}>{intl.formatMessage(
                    {
                      id: 'backlogin',
                    }
                  )}</Button>
                </div>
              </Form.Item>

              <Form.Item>
                <Button type="primary" htmlType="submit" style={{ width: "100%" }}>
                  {intl.formatMessage(
                    {
                      id: 'signup',
                    }
                  )}
                </Button>
              </Form.Item>
            </Form>
          </Col>
        </Row>
      </Layout.Content>
    </Layout>
  );
};

export default SignUp