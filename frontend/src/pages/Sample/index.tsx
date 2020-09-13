import React, { Fragment, useEffect } from "react";
import { connect } from "react-redux";

import { MainTemplate, SampleList, FormWrapper } from "components";
import {
  submitSampleRequest,
  getSamplesRequest,
} from "store/actionCreators/sample";

const SamplePage = ({
  title,
  samples,
  getSamplesRequest,
  submitSampleRequest,
}: Props) => {
  useEffect(getSamplesRequest, []);

  const onChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    title = e.target.value;
  };

  const onSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    submitSampleRequest(title);
  };

  return (
    <MainTemplate>
      <h2>Sample</h2>

      <Fragment>
        <SampleList samples={samples} />
        <FormWrapper onChange={onChange} onSubmit={onSubmit} />
      </Fragment>
    </MainTemplate>
  );
};

interface Props {
  title: string;
  samples: Array<Sample>;
  getSamplesRequest: () => void;
  submitSampleRequest: (title: string) => void;
}

type Sample = {
  title: string;
};

const mapStateToProps = (state: Readonly<any>) => ({
  title: state.get("sample").title,
  samples: state.get("sample").list,
});

const mapDispatchToProps = {
  getSamplesRequest: getSamplesRequest,
  submitSampleRequest: submitSampleRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(SamplePage);
