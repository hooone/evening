import React, { Component } from 'react';
import { getDvaApp } from 'umi';
import styles from './index.less';
import { ConfigProvider, PageHeader, Breadcrumb, Switch, Divider, Layout, Menu, Radio } from 'antd';

import Navigation from '@/components/Navigation';
import PageHead from '@/components/PageHead';
import PageContent from '@/components/PageContent';
import DataModal from '@/components/DataModal';
import RightClickContextMenu from '@/components/configuration/RightClickContextMenu';
import ConfigDrawer from '@/components/configuration/ConfigDrawer';
const { Content, Header, Sider } = Layout;

class Index extends Component {
  findAttribute = (ele, attr) => {
    if (!ele)
      return '';
    while (!ele.getAttribute(attr)) {
      if (ele.id === 'root')
        return '';
      else {
        ele = ele.parentElement;
      }
      if (!ele)
        return '';
    }
    return ele.getAttribute(attr);
  }
  getAllData = (ele, attr) => {
    if (!ele)
      return '';
    while (!ele.getAttribute(attr)) {
      if (ele.id === 'root')
        return '';
      else {
        ele = ele.parentElement;
      }
      if (!ele)
        return '';
    }
    let record = {}
    for (var i = 0; i < ele.attributes.length; i++) {
      let attr = ele.attributes[i]
      if (attr.localName.indexOf('data-') >= 0) {
        record[attr.localName.replace('data-', '')] = attr.value;
      }
    }
    return record
  }

  handleContextMenu = (event) => {
    event.preventDefault()
    let drawing = document.getElementsByTagName('HTML')[0];
    let rightedge = drawing.clientWidth - event.clientX;
    let bottomedge = drawing.clientHeight - event.clientY;
    let left = 0
    let top = 0
    if (rightedge < 158)
      left = drawing.scrollLeft + event.clientX - 160
    else
      left = drawing.scrollLeft + event.clientX
    top = drawing.scrollTop + event.clientY
    let v = this.findAttribute(event.target, "data-contextmenu")
    let record = this.getAllData(event.target, "data-contextmenu")
    let app = getDvaApp();
    app._store.dispatch({
      type: 'contextMenu/open',
      left: left,
      top: top,
      menu: v,
      record: record,
    })
  }
  handleClick = (event) => {
    let app = getDvaApp();
    app._store.dispatch({
      type: 'contextMenu/close',
    })
  }
  componentDidMount() {
    // 添加右键点击、点击事件监听
    document.addEventListener('contextmenu', this.handleContextMenu)
    document.addEventListener('click', this.handleClick)
  }

  componentWillUnmount() {
    // 移除事件监听
    document.removeEventListener('contextmenu', this.handleContextMenu)
    document.removeEventListener('click', this.handleClick)
  }

  render() {
    return (
      <Layout style={{ minHeight: '100vh' }}>
        <RightClickContextMenu />
        <Navigation />
        <Layout>
          <PageHead />
          <PageContent />
        </Layout>
        <DataModal></DataModal>
        <ConfigDrawer></ConfigDrawer>
      </Layout>
    );
  }
}

export default Index