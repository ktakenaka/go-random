import React, { useMemo, useCallback } from "react";
import { connect } from "react-redux";
import { Divider, Button } from "antd";

import MainTemplate from "components/templates/MainTemplate";
import Counter from "components/organisms/Counter";
import {
  countIncrement,
  countDecrement,
  countIncrementAsync,
} from "store/actionCreators/tutorial";
import { changeLocation } from "store/actionCreators/app";

interface Props {
  count: number;
  countIncrement: typeof countIncrement;
  countDecrement: typeof countDecrement;
  countIncrementAsync: typeof countIncrementAsync;
  changeLocation: typeof changeLocation;
}

const HomePage = ({
  count,
  countIncrement,
  countDecrement,
  countIncrementAsync,
  changeLocation,
}: Props) => {
  return (
    <MainTemplate>
      <h2>HOME</h2>
      <Button shape="round" onClick={() => changeLocation("/google/sign-in")}>
        Sign In
      </Button>

      <Divider plain>[Demo] Counter</Divider>

      <Counter
        value={count}
        onIncrement={() => countIncrement()}
        onDecrement={() => countDecrement()}
        onIncrementAsync={() => countIncrementAsync()}
      />
    </MainTemplate>
  );
};

const mapStateToProps = (state: Readonly<any>) => ({
  count: state.get("tutorial").count,
});

const mapDispatchToProps = {
  countIncrement,
  countDecrement,
  countIncrementAsync,
  changeLocation,
};

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
