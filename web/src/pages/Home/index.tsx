import React, { Fragment, useEffect } from "react";
import { connect } from "react-redux";

import { MainTemplate, SampleList, FormWrapper, Counter } from "components";
import {
  submitSampleRequest,
  getSamplesRequest,
} from "store/actionCreators/sample";
import {
  countIncrement,
  countDecrement,
  countIncrementAsync,
} from "store/actionCreators/tutorial";

const HomePage = ({
  count,
  title,
  samples,
  countIncrement,
  countDecrement,
  countIncrementAsync,
  getSamplesRequest,
  submitSampleRequest,
}: Props) => {
  useEffect(() => {
    getSamplesRequest();
  }, []);

  const onChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    title = e.target.value;
  };

  const onSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    submitSampleRequest(title);
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
        onIncrement={() => countIncrement()}
        onDecrement={() => countDecrement()}
        onIncrementAsync={() => countIncrementAsync()}
      />
    </MainTemplate>
  );
};

interface Props {
  count: number;
  title: string;
  samples: Array<Sample>;
  countIncrement: () => void;
  countDecrement: () => void;
  countIncrementAsync: () => void;
  getSamplesRequest: () => void;
  submitSampleRequest: (title: string) => void;
}

type Sample = {
  title: string;
};

const mapStateToProps = (state: Readonly<any>) => ({
  count: state.get("tutorial").count,
  title: state.get("sample").title,
  samples: state.get("sample").list,
});

const mapDispatchToProps = {
  countIncrement: countIncrement,
  countDecrement: countDecrement,
  countIncrementAsync: countIncrementAsync,
  getSamplesRequest: getSamplesRequest,
  submitSampleRequest: submitSampleRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
