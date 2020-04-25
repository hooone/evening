import React, { InputHTMLAttributes, ReactElement } from 'react';
import { Table, Popconfirm, Button, Menu, Layout } from 'antd';
import { connect, getLocale, Link, Dispatch, globalStateProps } from 'umi';
import { DispatchProp } from 'react-redux';
import { FolderOutlined, FileOutlined } from '@ant-design/icons';
import { getLocaleText, findAttribute } from '@/util';
const { SubMenu } = Menu;
const { Content, Header, Sider } = Layout;
import { navStateProps, ICollapseCommand, IMoveNavCommand } from '@/models/navigation';
import { IStore } from '@/store'

interface NavProps extends DispatchProp {
    nav: navStateProps,
    global: globalStateProps,
}

const Navigation = (props: NavProps) => {
    function onCollapse() {
        let onCollapse: ICollapseCommand = {
            type: 'nav/onCollapse',
        }
        props.dispatch(onCollapse);
    }
    function onDragStart(e: React.DragEvent) {
        let real = e.target as HTMLElement
        if (real.tagName === "path") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "svg") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "I") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "A") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "DIV" && real.classList.contains("ant-menu-submenu-title")) {
            e.dataTransfer?.setData("type", "folder");
            e.dataTransfer?.setData("id", findAttribute(real, "data-treeid"));
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-folder")) {
            e.dataTransfer?.setData("type", "folder");
            e.dataTransfer?.setData("id", findAttribute(real, "data-treeid"));
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-node")) {
            e.dataTransfer?.setData("type", "page");
            e.dataTransfer?.setData("id", findAttribute(real, "data-pageid"));
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-tree")) {
            e.dataTransfer?.setData("type", "page");
            e.dataTransfer?.setData("id", findAttribute(real, "data-pageid"));
        }
    }
    function onDrop(e: React.DragEvent) {
        e.preventDefault()
        let moveNav: IMoveNavCommand = {
            type: "nav/moveNav",
            SourceFolder: 0,
            SourcePage: 0,
            TargetFolder: 0,
            TargetPage: 0,
            Position: 0
        }
        let real = e.target as HTMLElement
        if (real.tagName === "path") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "svg") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "I") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "A") {
            real = real.parentElement as HTMLElement
        }
        if ((real.tagName === "DIV" && real.classList.contains("ant-menu-submenu-title")) ||
            (real.tagName === "DIV" && real.classList.contains("ant-layout-sider-children")) ||
            (real.tagName === "LI" && real.classList.contains("nav-node")) ||
            (real.tagName === "LI" && real.classList.contains("nav-tree"))) {
            real.classList.remove("dragEnter")
            let sourceType = e.dataTransfer?.getData("type");
            let sourceID = e.dataTransfer?.getData("id");
            if (sourceType === "page") {
                moveNav.SourcePage = parseInt(sourceID ?? "0")
            }
            else if (sourceType === "folder") {
                moveNav.SourceFolder = parseInt(sourceID ?? "0")
            }
        }
        if (real.tagName === "DIV" && real.classList.contains("ant-menu-submenu-title")) {
            moveNav.TargetFolder = parseInt(findAttribute(real, "data-treeid"))
            if (moveNav.SourcePage !== 0) {
                if (e.clientY - real.offsetTop < real.clientHeight / 4) {
                    moveNav.Position = 1
                }
                else if ((e.clientY - real.offsetTop) > (real.clientHeight * 3 / 4)) {
                    moveNav.Position = 3
                }
                else {
                    moveNav.Position = 2
                }
            }
            else if (moveNav.SourceFolder !== 0) {
                if (e.clientY - real.offsetTop < real.clientHeight / 2) {
                    moveNav.Position = 1
                }
                else {
                    moveNav.Position = 3
                }
            }
        }
        else if ((real.tagName === "LI" && real.classList.contains("nav-node")) ||
            (real.tagName === "LI" && real.classList.contains("nav-tree"))) {
            moveNav.TargetPage = parseInt(findAttribute(real, "data-pageid"))
            if (e.clientY - real.offsetTop < real.clientHeight / 2) {
                moveNav.Position = 1
            }
            else {
                moveNav.Position = 3
            }
        }
        if ((real.tagName === "DIV" && real.classList.contains("ant-menu-submenu-title")) ||
            (real.tagName === "DIV" && real.classList.contains("ant-layout-sider-children")) ||
            (real.tagName === "LI" && real.classList.contains("nav-node")) ||
            (real.tagName === "LI" && real.classList.contains("nav-tree"))) {
            props.dispatch(moveNav);
        }

    }
    function onDragOver(e: React.DragEvent) {
        let real = e.target as HTMLElement
        if (real.tagName === "path") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "svg") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "I") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "A") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "DIV" && real.classList.contains("ant-menu-submenu-title")) {
            if (e.clientY - real.offsetTop < real.clientHeight / 4) {
                real.classList.add("dragTop")
                real.classList.remove("dragInner")
                real.classList.remove("dragBottom")
            }
            else if ((e.clientY - real.offsetTop) > (real.clientHeight * 3 / 4)) {
                real.classList.remove("dragTop")
                real.classList.remove("dragInner")
                real.classList.add("dragBottom")
            }
            else {
                real.classList.remove("dragTop")
                real.classList.add("dragInner")
                real.classList.remove("dragBottom")
            }
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-node")) {
            if (e.clientY - real.offsetTop < real.clientHeight / 2) {
                real.classList.add("dragTop")
                real.classList.remove("dragInner")
                real.classList.remove("dragBottom")
            }
            else {
                real.classList.remove("dragTop")
                real.classList.remove("dragInner")
                real.classList.add("dragBottom")
            }
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-tree")) {
            if (e.clientY - real.offsetTop < real.clientHeight / 2) {
                real.classList.add("dragTop")
                real.classList.remove("dragInner")
                real.classList.remove("dragBottom")
            }
            else {
                real.classList.remove("dragTop")
                real.classList.remove("dragInner")
                real.classList.add("dragBottom")
            }
        }
        e.preventDefault()
    }
    function onDragEnter(e: React.DragEvent) {
        let real = e.target as HTMLElement
        if (real.tagName === "path") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "svg") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "I") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "A") {
            real = real.parentElement as HTMLElement
        }
        console.log("enter")
        console.log(real.tagName)
        if (real.tagName === "DIV" && real.classList.contains("ant-menu-submenu-title")) {
            if (!real.classList.contains("dragEnter")) {
                real.classList.add("dragEnter")
            }
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-node")) {
            if (!real.classList.contains("dragEnter")) {
                real.classList.add("dragEnter")
            }
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-tree")) {
            if (!real.classList.contains("dragEnter")) {
                real.classList.add("dragEnter")
            }
        }
        e.preventDefault()
    }
    function onDragLeave(e: React.DragEvent) {
        let real = e.target as HTMLElement
        if (real.tagName === "path") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "svg") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "I") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "SPAN") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "A") {
            real = real.parentElement as HTMLElement
        }
        if (real.tagName === "DIV" && real.classList.contains("ant-menu-submenu-title")) {
            real.classList.remove("dragEnter")
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-node")) {
            real.classList.remove("dragEnter")
        }
        else if (real.tagName === "LI" && real.classList.contains("nav-tree")) {
            real.classList.remove("dragEnter")
        }
        e.preventDefault()
    }
    let trees: React.ReactElement[] = [];
    let SelectedKey: string[] = []
    let OpenKey: string[] = []
    let route = props.global.PathName.split('/')
    route = route.filter((str) => str)
    props.nav.Navs.forEach(
        folder => {
            if (!folder.IsFolder) {
                if (folder.Pages !== null && folder.Pages.length === 1) {
                    if (route.length === 1) {
                        if (folder.Pages[0].Name === route[0]) {
                            SelectedKey = ["page_" + folder.Pages[0].Id]
                        }
                    }
                    trees.push((folder.Pages[0].Name === "home") ? (<Menu.Item
                        draggable="false"
                        key={"page_" + folder.Pages[0].Id} className="nav-tree"
                        data-contextmenu="treepage"
                        data-treeid={folder.Id}
                        data-pageid={folder.Pages[0].Id}
                        data-text={folder.Pages[0].Text}
                        data-name={folder.Pages[0].Name}
                    >
                        <Link to={"/" + folder.Pages[0].Name + "/"}>
                            <FileOutlined /><span >
                                {getLocaleText(folder.Pages[0].Locale)}
                            </span>
                        </Link>
                    </Menu.Item>) : (<Menu.Item
                        draggable="true"
                        onDragStart={(e) => { onDragStart(e) }}
                        onDragEnter={(e) => { onDragEnter(e) }}
                        onDragLeave={(e) => { onDragLeave(e) }}
                        onDragOver={(e) => { onDragOver(e) }}
                        key={"page_" + folder.Pages[0].Id} className="nav-tree"
                        data-contextmenu="treepage"
                        data-treeid={folder.Id}
                        data-pageid={folder.Pages[0].Id}
                        data-text={folder.Pages[0].Text}
                        data-name={folder.Pages[0].Name}
                    >
                        <Link to={"/" + folder.Pages[0].Name + "/"}>
                            <FileOutlined /><span >
                                {getLocaleText(folder.Pages[0].Locale)}
                            </span>
                        </Link>
                    </Menu.Item>)
                    )
                }
            }
            else {
                var nodes: React.ReactElement[] = [];
                if (folder.Pages !== null) {
                    folder.Pages.forEach(
                        node => {
                            if (route.length === 2) {
                                if (folder.Name === route[0] && node.Name === route[1]) {
                                    OpenKey = ["treenode_" + folder.Id]
                                    SelectedKey = ["node_" + node.Id]
                                }
                            }
                            nodes.push(<Menu.Item
                                draggable="true"
                                onDragStart={(e) => { onDragStart(e) }}
                                onDragEnter={(e) => { onDragEnter(e) }}
                                onDragLeave={(e) => { onDragLeave(e) }}
                                onDragOver={(e) => { onDragOver(e) }}
                                key={"node_" + node.Id} className="nav-node"
                                data-contextmenu="nodepage"
                                data-pageid={node.Id}
                                data-text={node.Text}
                                data-name={node.Name}

                            >
                                <Link to={"/" + folder.Name + "/" + node.Name + "/"}>
                                    <FileOutlined /><span  >
                                        {getLocaleText(node.Locale)}
                                    </span>
                                </Link>
                            </Menu.Item >)

                        }
                    )
                }
                trees.push(
                    <SubMenu
                        draggable="true"
                        onDragStart={(e: React.DragEvent) => { onDragStart(e) }}
                        onDragEnter={(e: React.DragEvent) => { onDragEnter(e) }}
                        onDragLeave={(e: React.DragEvent) => { onDragLeave(e) }}
                        onDragOver={(e: React.DragEvent) => { onDragOver(e) }}
                        key={"treenode_" + folder.Id} className="nav-folder"
                        data-contextmenu="treeview" data-treeid={folder.Id}
                        data-text={folder.Name} data-name={folder.Name} title={<span ><FolderOutlined />
                            <span>
                                {getLocaleText(folder.Locale)}
                            </span></span>}>{nodes}</SubMenu>
                )
            }
        }
    )
    return <Sider
        onDrop={(e) => { onDrop(e) }}
        onDragOver={(e) => { onDragOver(e) }}
        allowdrop={"true"}
        collapsible collapsed={props.nav.collapsed} onCollapse={() => { onCollapse() }} data-contextmenu="addtree">
        <div className="ant-pro-sider-menu-logo">
            <a href="/">
                <h1>evening</h1>
            </a>
        </div>
        <Menu theme="dark" defaultSelectedKeys={SelectedKey} defaultOpenKeys={OpenKey} key={'menu_' + Math.floor((Math.random() * 100000) + 1)} mode="inline">
            {trees}
        </Menu>
    </Sider>
};
export default connect((state: IStore) => {
    return {
        nav: state.nav,
        global: state.global,
    }
})(Navigation);