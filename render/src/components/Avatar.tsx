import React from 'react';
import { setLocale } from 'umi';
import { Menu, Dropdown, Button } from 'antd';
import { ClickParam } from 'antd/lib/menu';
import { DownOutlined } from '@ant-design/icons';

const Avatar = () => {
    function handleMenuClick(key: ClickParam) {
        if (key.key == "LOGOUT") {
            //window.open('/logout', '_self')
        }
        else {
            setLocale(key.key, false)
        }
    }
    return (
        <Dropdown overlay={<Menu onClick={(value) => { handleMenuClick(value) }}>
            <Menu.Item key="zh-CN">中文</Menu.Item>
            <Menu.Item key="en-US">English</Menu.Item>
            <Menu.Divider />
            <Menu.Item key="LOGOUT">登出</Menu.Item>
        </Menu>}>
            <Button type="ghost" style={{ marginLeft: 8 }}>
                admin <DownOutlined />
            </Button>
        </Dropdown>

    )
}
export default Avatar