import React, { useMemo, useCallback, useEffect } from "react";
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
  const nonMemorizedObj = {hello: "non-memo"};
  useEffect(() => {
    // React uses referential equal to compare options of useEffect,
    // that's why it's evaluated everytime
    console.log("this is re-rendered everytime to click count");
    console.log(nonMemorizedObj);
  }, [nonMemorizedObj]);

  const memorizedObj = useMemo(() => [1,2,3], []);
  useEffect(() => {
    // when using `useMemo`, memorizedObj is referential equal everytime
    // because it's exactly the same object
    console.log("this is called once");
    console.log(memorizedObj);
  }, [memorizedObj]);

  const defineOnceFunc = useCallback(() => {
    console.log("hello useCallback");
  }, []);
  useEffect(() => {
    // "defineOnceFunc" is always referencial equal
    // because I don't pass any args to `useCallback`
    console.log("this is called once");
    defineOnceFunc();
  }, [defineOnceFunc]);

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
