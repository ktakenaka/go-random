import React, { useState, Fragment } from "react";
import { connect } from "react-redux";

import {
  MainTemplate,
  SampleList,
  FormWrapper,
  Counter,
} from "../../components";
import {
  actionCreate,
  SAMPLE_DECREMENT,
  SAMPLE_INCREMENT,
  SAMPLE_INCREMENT_ASYNC,
} from "../../store/actions";

const HomePage = ({ count, actionCreate }: Props) => {
  const [samples, setSamples] = useState<string[]>([
    "sample1",
    "sample2",
    "sample3",
  ]);
  const [value, setvalue] = useState<string>("");

  const onChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const title = e.target.value;
    setvalue(title);
  };

  const onSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    setSamples([...samples.concat(value)]);
  };

  return (
    <MainTemplate>
      <h2>HOME</h2>

      <Fragment>
        <SampleList samples={samples} />
        <FormWrapper onChange={onChange} onSubmit={onSubmit} />
      </Fragment>

      <Counter
        value={count}
        onIncrement={() => actionCreate(SAMPLE_INCREMENT)}
        onDecrement={() => actionCreate(SAMPLE_DECREMENT)}
        onIncrementAsync={() => actionCreate(SAMPLE_INCREMENT_ASYNC)}
      />
    </MainTemplate>
  );
};

type Props = {
  count: number;
  actionCreate: (type: string) => void;
};

interface State {
  count: number;
}

const mapStateToProps = (state: Readonly<State>) => ({
  count: state.count,
});

const mapDispatchToProps = {
  actionCreate: actionCreate,
};

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
