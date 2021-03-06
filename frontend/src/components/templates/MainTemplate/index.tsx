import React from "react";
import { Link } from "react-router-dom";

import { Layout, Menu } from "antd";

import { Message } from "./styles";
import ActionMessage from "components/organisms/ActionMessage";
import { moveLocation } from "utils/changeLocation";

type Props = {
  children: React.ReactNode;
};
const MainTemplate = ({ children }: Props) => {
  const { Header, Content, Sider, Footer } = Layout;

  // TODO: define organism
  const header = (
    <Header className="header">
      <div className="logo" />
      <Menu theme="dark" mode="horizontal" defaultSelectedKeys={["2"]}>
        <Menu.Item key="1">
          <Link to="/home">Sample App</Link>
        </Menu.Item>
      </Menu>
    </Header>
  );

  // TODO: define organism
  const sidebar = (
    <Sider width={200} className="site-layout-background">
      <Menu
        mode="inline"
        // defaultSelectedKeys={["1"]}
        // defaultOpenKeys={["sub1"]}
        style={{ height: "100%", borderRight: 0 }}
      >
        <Menu.Item onClick={() => moveLocation("/samples")}>Sample</Menu.Item>
      </Menu>
    </Sider>
  );

  return (
    <Layout>
      {header}
      <Message>
        <ActionMessage />
      </Message>
      <Layout>
        {sidebar}
        <Layout style={{ padding: "24px 24px 24px" }}>
          <Content
            className="site-layout-background"
            style={{
              background: "white",
              padding: 24,
              minHeight: 280,
            }}
          >
            {children}
          </Content>
        </Layout>
      </Layout>
      <Footer className="footer">footer</Footer>
    </Layout>
  );
};

export default MainTemplate;
