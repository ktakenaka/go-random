import React from "react";
import { connect } from "react-redux";

import { Wrapper, Header, SideBar, Body, Message, Content } from "./styles";
import { MainHeader, ActionMessage } from "components";
import { changeLocation } from "store/actionCreators/app";

type Props = {
  children: React.ReactNode;
  header?: JSX.Element;
  actionMessage?: JSX.Element;
  changeLocation: (location: string) => void;
};

const MainTemplate = ({
  children,
  header = <MainHeader />,
  actionMessage = <ActionMessage />,
  changeLocation,
}: Props) => {
  return (
    <Wrapper>
      <Header>{header}</Header>
      <Body>
        <SideBar>
          <li onClick={() => changeLocation("/")}>Home</li>
          <li onClick={() => changeLocation("/samples")}>sample</li>
        </SideBar>
        <Message>{actionMessage}</Message>
        <Content>{children}</Content>
      </Body>
    </Wrapper>
  );
};

const mapDispatchToProps = {
  changeLocation: changeLocation,
};

export default connect(null, mapDispatchToProps)(MainTemplate);
