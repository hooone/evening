import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Radio, Layout, PageHeader, Breadcrumb, Menu, Dropdown, Button } from 'antd';
import { connect, getLocale, Dispatch } from 'umi';
import { DispatchProp } from 'react-redux';
import { FolderOutlined, FileOutlined, DownOutlined } from '@ant-design/icons';
import { ClickParam } from 'antd/lib/menu';
import { navStateProps } from '@/models/navigation'
import { globalStateProps } from '@/models/global'
import { IStore } from '@/store'
import { getLocaleText } from '@/util'
import { ExportUmi } from '@/render'

const { Content, Header, Sider } = Layout;
interface PageHeadProps extends DispatchProp {
    global: globalStateProps,
    nav: navStateProps,
}

const PageHead = (props: PageHeadProps) => {

    function handleLocale(value: string) {
        props.dispatch({
            type: 'global/changeLocale',
            lang: value,
        });
    }
    function handleMenuClick(key: ClickParam) {
        if (key.key == "LOGOUT") {
            window.open('/logout', '_self')
        }
        else if (key.key === "export") {
            ExportUmi();
        }
        else {
            handleLocale(key.key)
        }
    }
    const menu = (
        <Menu onClick={(value) => { handleMenuClick(value) }}>
            <Menu.Item key="zh-CN">中文</Menu.Item>
            <Menu.Item key="en-US">English</Menu.Item>
            <Menu.Divider />
            <Menu.Item key="export">导出</Menu.Item>
            <Menu.Divider />
            <Menu.Item key="LOGOUT">登出</Menu.Item>
        </Menu>
    );
    let route: string[] = []
    if (props.global.PathName) {
        route = props.global.PathName.split("/").filter((str) => (str))
    }
    let title = ""
    let bread: React.ReactElement[] = []
    props.nav.Navs.forEach(
        folder => {

            if (!folder.IsFolder) {
                if (folder.Pages !== null && folder.Pages.length === 1) {
                    if (route.length === 1) {
                        if (folder.Pages[0].Name === route[0]) {
                            title = getLocaleText(folder.Pages[0].Locale)

                            bread.push(<Breadcrumb.Item key={"Bread_" + folder.Pages[0].Id}>
                                <FileOutlined /><span>
                                    {getLocaleText(folder.Pages[0].Locale)}
                                </span></Breadcrumb.Item>)
                        }
                    }
                }
            }
            else {
                if (folder.Pages !== null) {
                    folder.Pages.forEach(
                        node => {
                            if (route.length === 2) {
                                if (folder.Name === route[0] && node.Name === route[1]) {
                                    title = getLocaleText(node.Locale)
                                    bread.push(<Breadcrumb.Item key={"BreadFolder_" + folder.Id}>
                                        <FolderOutlined />
                                        <span>
                                            {getLocaleText(folder.Locale)}
                                        </span></Breadcrumb.Item>)
                                    bread.push(<Breadcrumb.Item key={"Bread_" + node.Id}>
                                        <FileOutlined />
                                        <span>
                                            {getLocaleText(node.Locale)}
                                        </span></Breadcrumb.Item>)
                                }
                            }
                        }
                    )
                }
            }
        }
    )
    return <Header style={{ background: '#fff', padding: 0 }} >
        <PageHeader
            style={{
                border: '1px solid rgb(235, 237, 240)',
            }}
            onBack={() => null}
            title={title}
            subTitle={<Breadcrumb separator="/" className="breadcrumb">
                {bread}
            </Breadcrumb>}
            extra={
                <Dropdown overlay={menu}>
                    <Button type="ghost" style={{ marginLeft: 8 }}>
                        {props.global.User} <DownOutlined />
                    </Button>
                </Dropdown>
            }
        />
    </Header>;
};

export default connect((state: IStore) => {
    return {
        global: state.global,
        nav: state.nav,
    }
})(PageHead);
