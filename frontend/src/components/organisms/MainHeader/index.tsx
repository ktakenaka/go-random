import React from "react";

import { Wrapper, LeftContent, RightContent, Divide, Item } from "./styles";

const MainHeader = () => {
  return (
    <Wrapper>
      <LeftContent>
        <Item>Left</Item>
      </LeftContent>
      <Divide />
      <RightContent>
        <Item>Right</Item>
      </RightContent>
    </Wrapper>
  );
};

export default MainHeader;
