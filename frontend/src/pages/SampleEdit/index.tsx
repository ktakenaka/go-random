import React, { useEffect } from "react";
import { connect } from "react-redux";
import { useParams } from "react-router-dom";

import MainTemplate from "components/templates/MainTemplate";
import SampleForm from "components/organisms/SampleForm";
import {
  getSampleRequest,
  updateSampleRequest,
} from "store/actionCreators/sample";
import { TypeSample } from "constants/type";

interface Props {
  sample: TypeSample;
  updateSampleRequest: typeof updateSampleRequest;
  getSampleRequest: typeof getSampleRequest;
}

const SampleEditPage = ({
  sample,
  getSampleRequest,
  updateSampleRequest,
}: Props) => {
  const { id } = useParams<any>();

  useEffect(() => {
    getSampleRequest(id);
  }, [getSampleRequest, id]);

  const onSubmit = (values: any) => {
    updateSampleRequest(id, values);
  };

  return (
    <MainTemplate>
      <h1>Edit Sample</h1>
      <SampleForm onSubmit={onSubmit} sample={sample} />
    </MainTemplate>
  );
};

const mapStateToProps = (state: Readonly<any>) => ({
  sample: state.getIn(["sample", "item"]),
});

const mapDispatchToProps = {
  getSampleRequest,
  updateSampleRequest,
};

export default connect(mapStateToProps, mapDispatchToProps)(SampleEditPage);
