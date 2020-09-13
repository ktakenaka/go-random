import React from "react";
import { connect } from "react-redux";

import { Wrapper, LeftContent, RightContent, Divide, Item } from "./styles";
import { changeLocation } from "store/actionCreators/app";

const MainHeader = ({
  changeLocation,
}: {
  changeLocation: (location: string) => void;
}) => {
  return (
    <Wrapper>
      <LeftContent>
        <Item>Left</Item>
        <button onClick={() => changeLocation("/")}>Home</button>

        <button onClick={() => changeLocation("/samples")}>sample</button>
      </LeftContent>
      <Divide />
      <RightContent>
        <Item>Right</Item>
      </RightContent>
    </Wrapper>
  );
};

const mapDispatchToProps = {
  changeLocation: changeLocation,
};

export default connect(null, mapDispatchToProps)(MainHeader);
