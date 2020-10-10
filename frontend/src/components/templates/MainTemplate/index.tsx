import React from "react";
import { Link } from "react-router-dom";
import { List } from "antd";

import { Wrapper, Header, SideBar, Body, Message, Content } from "./styles";
import MainHeader from "components/organisms/MainHeader";
import ActionMessage from "components/organisms/ActionMessage";

type Props = {
  children: React.ReactNode;
};
const MainTemplate = ({ children }: Props) => {
  return (
    <Wrapper>
      <Header>
        <MainHeader />
      </Header>
      <Body>
        <SideBar>
          <List>
            <Link to="/home">HOME</Link>
          </List>
          <List>
            <Link to="/samples">Sample</Link>
          </List>
        </SideBar>
        <Message>
          <ActionMessage />
        </Message>
        <Content>{children}</Content>
      </Body>
    </Wrapper>
  );
};

export default MainTemplate;
