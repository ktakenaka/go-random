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

const HomePage = ({
  count,
  sample,
  samples,
  actionCreate,
  submitSample,
}: Props) => {
  const onChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    const title = e.target.value;
    sample.title = title;
  };

  const onSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    submitSample(sample.title);
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

interface Props extends State {
  actionCreate: (type: string) => void;
  submitSample: (title: string) => void;
}

type Sample = {
  title: string;
};

type State = {
  count: number;
  sample: Sample;
  samples: Array<Sample>;
};

const mapStateToProps = (state: Readonly<State>) => ({
  count: state.count,
  sample: state.sample,
  samples: state.samples,
});

const mapDispatchToProps = {
  actionCreate: actionCreate,
  submitSample: submitSample,
};

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
