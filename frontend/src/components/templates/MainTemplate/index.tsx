import React from "react";

import { Wrapper, Header, Body, Content } from "./styles";
import { MainHeader, ActionMessage } from "components";

type Props = {
  children: React.ReactNode;
  header?: JSX.Element;
  actionMessage?: JSX.Element;
};

const MainTemplate = ({
  children,
  header = <MainHeader />,
  actionMessage = <ActionMessage />,
}: Props) => {
  return (
    <Wrapper>
      <Header>{header}</Header>
      <Body>
        {actionMessage}
        <Content>{children}</Content>
      </Body>
    </Wrapper>
  );
};

export default MainTemplate;
