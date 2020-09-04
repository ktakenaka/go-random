import React from "react";

import { Wrapper, Header, Body, Content } from "./styles";
import { MainHeader } from "components";

type Props = {
  header?: JSX.Element;
  children: React.ReactNode;
};

const MainTemplate = ({ header = <MainHeader />, children }: Props) => {
  return (
    <Wrapper>
      <Header>{header}</Header>
      <Body>
        <Content>{children}</Content>
      </Body>
    </Wrapper>
  );
};

export default MainTemplate;
