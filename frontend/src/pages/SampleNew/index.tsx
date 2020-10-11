import React from "react";
import { connect } from "react-redux";

import MainTemplate from "components/templates/MainTemplate";
import SampleForm from "components/organisms/SampleForm";
import { submitSampleRequest } from "store/actionCreators/sample";

interface Props {
  submitSampleRequest: typeof submitSampleRequest;
}

const SampleNewPage = ({ submitSampleRequest }: Props) => {
  return (
    <MainTemplate>
      <h1>New Sample</h1>
      <SampleForm
        onSubmit={submitSampleRequest}
        sample={{ title: null, content: null }}
      />
    </MainTemplate>
  );
};

const mapDispatchToProps = {
  submitSampleRequest,
};

export default connect(null, mapDispatchToProps)(SampleNewPage);
