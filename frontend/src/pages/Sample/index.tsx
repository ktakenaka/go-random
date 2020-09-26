import React, { useEffect, useState } from "react";
import { connect } from "react-redux";

import { MainTemplate, SampleList, SampleForm } from "components";
import {
  submitSampleRequest,
  typeSubmitSampleRequest,
  getSamplesRequest,
  typeGetSampleRequest,
} from "store/actionCreators/sample";
import { TypeSample } from "constants/type";

const SamplePage = ({
  samples,
  getSamplesRequest,
  submitSampleRequest,
}: Props) => {
  useEffect(getSamplesRequest, []);
  const [sample, setSample] = useState<any>(null);

  const onChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setSample({ ...sample, [e.target.name]: e.target.value });
  };

  const onSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    if (sample) {
      submitSampleRequest(sample);
    } else {
      console.log("sample must be defined");
    }
  };

  return (
    <MainTemplate>
      <h1>Sample</h1>
      <h2>List</h2>
      <SampleList samples={samples} />

      <h2>New Sample</h2>
      <SampleForm onChange={onChange} onSubmit={onSubmit} />
    </MainTemplate>
  );
};

interface Props {
  samples: Array<TypeSample>;
  getSamplesRequest: typeGetSampleRequest;
  submitSampleRequest: typeSubmitSampleRequest;
}

const mapStateToProps = (state: Readonly<any>) => ({
  samples: state.getIn(["sample", "list"]),
});

const mapDispatchToProps = {
  getSamplesRequest: getSamplesRequest,
  submitSampleRequest: submitSampleRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(SamplePage);
