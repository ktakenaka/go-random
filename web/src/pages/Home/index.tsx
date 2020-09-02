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
  submitSample,
} from "../../store/actions";

const HomePage = ({ count, samples, actionCreate, submitSample }: Props) => {
  const [samplers, setSamples] = useState<string[]>([
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
    submitSample(value);
    setSamples([...samplers.concat(value)]);
  };

  return (
    <MainTemplate>
      <h2>HOME</h2>

      <Fragment>
        <SampleList samples={samplers} />
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

interface Props extends State {
  actionCreate: (type: string) => void;
  submitSample: (title: string) => void;
};

interface Sample {
  title: string;
}

interface State {
  count: number;
  samples: Array<Sample>;
}

const mapStateToProps = (state: Readonly<State>) => ({
  count: state.count,
  samples: state.samples,
});

const mapDispatchToProps = {
  actionCreate: actionCreate,
  submitSample: submitSample,
};

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
