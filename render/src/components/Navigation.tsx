import React from 'react';
import { Menu, Layout } from 'antd';
import { connect, Link, useIntl } from 'umi';
import { DispatchProp } from 'react-redux';
import { FolderOutlined, FileOutlined } from '@ant-design/icons';
const { SubMenu } = Menu;
const { Sider } = Layout;
interface NavProps extends DispatchProp {
    navigation: {
        path: string,
        collapsed: boolean,
    },
}

const Navigation = (props: NavProps) => {
    const intl = useIntl();
    function onCollapse() {
        props.dispatch({ type: 'navigation/onCollapse' });
    }
    let trees: React.ReactElement[] = [];
    let SelectedKey: string[] = []
    let OpenKey: string[] = []
    let route = props.navigation.path.split('/')
    route = route.filter((str) => str)
    if (route.length === 1) {
        SelectedKey = ["page_" + route[0]]
    }
    else if (route.length === 2) {
        OpenKey = ["folder_" + route[0]]
        SelectedKey = ["page_" + route[1]]
    }
    return <Sider
        collapsible collapsed={props.navigation.collapsed} onCollapse={() => { onCollapse() }}>
        <div className="ant-pro-sider-menu-logo">
            <a href="/">
                <h1>evening</h1>
            </a>
        </div>
        <Menu theme="dark" defaultSelectedKeys={SelectedKey} defaultOpenKeys={OpenKey} key={'menu_' + Math.floor((Math.random() * 100000) + 1)} mode="inline">
            <% folders.forEach(folder => {%>
            <%  if (!folder.IsFolder) {
            if (folder.Pages !== null && folder.Pages.length === 1) {%>
            <Menu.Item key={"page_<%- folder.Pages[0].Name %>"} className="nav-tree">
                <Link to={"/<%- folder.Pages[0].Name %>"}>
                    <FileOutlined />
                    <span>
                        <%- '{intl.formatMessage({ id: "'+folder.Pages[0].Locale.Name+'" })}' %>
                    </span>
                </Link>
            </Menu.Item>
            <%  }
            }else {
            if(folder.Pages !== null) {%>
            <SubMenu key={"folder_<%- folder.Name %>"} className="nav-folder"
                title={<span ><FolderOutlined /><span><%- '{intl.formatMessage({ id: "'+folder.Locale.Name+'" })}' %></span></span>}>
                    <% folder.Pages.forEach(
                    node => {%>
                    <Menu.Item key={"page_<%- node.Name %>"} className="nav-node">
                    <Link to={"/<%- folder.Name %>/<%- node.Name %>"}>
                        <FileOutlined /><span  >
                        <%- '{intl.formatMessage({ id: "'+node.Locale.Name+'" })}' %>
                        </span>
                    </Link>
                </Menu.Item >
                    <% }
                    )%>
            </SubMenu>
           
            <%}
            } %>
            <% }); %>
        </Menu>
    </Sider>
};
export default connect((state: any) => {
    return {
        navigation: state.navigation,
    }
})(Navigation);