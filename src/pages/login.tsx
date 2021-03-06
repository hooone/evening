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
    <Layout style={{ minHeight: '100vh' }} className={"login-layout"}>
      <Layout.Content>
        <Row>
          <Col span={6} offset={9} style={{ marginTop: '48px' }}>
            <Typography style={{ padding: '12px' }}>
              <Title>
                evening
            </Title>
              <Paragraph>
                中后台数据管理网页快速开发工具
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
              <Form.Item name="remember" style={{ marginBottom: 0 }}>
                <div>
                  <Checkbox style={{ lineHeight: '32px' }}>{intl.formatMessage(
                    {
                      id: 'remember',
                    }
                  )}</Checkbox>
                  <Button type="link" style={{ float: 'right' }} onClick={() => { history.push("/signup") }}>{intl.formatMessage(
                    {
                      id: 'register',
                    }
                  )}</Button>
                </div>
              </Form.Item>

              <Form.Item>
                <Button type="primary" htmlType="submit" style={{ width: "100%" }}>
                  {intl.formatMessage(
                    {
                      id: 'login',
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

export default Login