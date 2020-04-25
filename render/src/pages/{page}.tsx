import React, { ReactElement } from 'react';
import { connect, useIntl } from 'umi';
import { Layout, PageHeader, Breadcrumb, Row, Col } from 'antd';
import { FolderOutlined, FileOutlined } from '@ant-design/icons';;
import Navigation from '@/components/Navigation';
import Avatar from '@/components/Avatar';
<% page.Cards.forEach(card=>{ %>
<%- "import "+card.Name.replace(card.Name[0],card.Name[0].toUpperCase())+"Card from '@/components/"+card.Name.replace(card.Name[0],card.Name[0].toUpperCase())+"Card'" %>;
<% })%>
const { Content, Header } = Layout;
const <%= page.Name %> = () => {
  const intl = useIntl();
  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Navigation />
      <Layout>
        <Header style={{ background: '#fff', padding: 0 }} >
          <PageHeader
            style={{
              border: '1px solid rgb(235, 237, 240)',
            }}
            onBack={() => null}
            title=<%- '{intl.formatMessage({ id: "'+page.Locale.Name+'" })}' %>
            subTitle={<Breadcrumb separator="/" className="breadcrumb">
              <% if(page.FolderLocale){ %>
                <Breadcrumb.Item>
                  <FolderOutlined />
                  <span>
                    <%- '{intl.formatMessage({ id: "'+page.FolderLocale.Name+'" })}' %>
                </span>
                </Breadcrumb.Item>
                <% } %>
              <Breadcrumb.Item>
                <FileOutlined />
                <span>
                  <%- '{intl.formatMessage({ id: "'+page.Locale.Name+'" })}' %>
                </span>
              </Breadcrumb.Item>
            </Breadcrumb>}
            extra={
              <Avatar />
            }
          />
        </Header>
        <Content style={{ margin: '0 16px', paddingBottom: '100px' }}>
          <% page.Rows.forEach(
        row => {%>
          <%- "<Row>" %>


          <% let offset = 0;
            row.Cols.forEach(col => {%>
              
            <%- "<Col span={" + col.Width + "} offset={" + (col.Pos - offset) + "}>" %>
            <% col.Children.forEach(card => {%>
              <%- "<"+card.Name.replace(card.Name[0],card.Name[0].toUpperCase())+"Card />" %>
            <% })%>
            <%- "</Col>" %>

          <% offset = col.Pos + col.Width;}); %>


          <%- "</Row>" %>
      <%}
            )%>
        </Content>
      </Layout >
    </Layout >
  );
};

export default <%= page.Name %> 