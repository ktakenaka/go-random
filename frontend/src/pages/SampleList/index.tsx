import React, { useEffect, useState } from "react";
import { connect } from "react-redux";

import MainTemplate from "components/templates/MainTemplate";
import SampleList from "components/organisms/SampleList";
import SampleForm from "components/organisms/SampleForm";
import {
  submitSampleRequest,
  typeSubmitSampleRequest,
  getSamplesRequest,
  typeGetSampleRequest,
} from "store/actionCreators/sample";
import { TypeSample } from "constants/type";
import { Radio } from "antd";

const charsets = [
  { label: "UTF-8", value: "utf8" },
  { label: "Shift_JIS", value: "sjis" },
];

const SamplePage = ({
  samples,
  getSamplesRequest,
  submitSampleRequest,
}: Props) => {
  useEffect(getSamplesRequest, []);
  const [sample, setSample] = useState<any>(null);
  const [charset, setCharset] = useState<"utf8" | "sjis">("utf8");

  const onChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
    setSample({ ...sample, [e.target.name]: e.target.value });
  };

  const onSubmit = (): void => {
    submitSampleRequest(sample);
  };

  const onCharsetSelected = (e: any) => {
    setCharset(e.target.value);
  };

  return (
    <MainTemplate>
      <h1>Sample</h1>
      <h2>List</h2>
      <SampleList samples={samples} />

      <Radio.Group onChange={onCharsetSelected} value={charset}>
        {charsets.map((item, i) => (
          <Radio key={i} value={item.value}>
            {item.label}
          </Radio>
        ))}
      </Radio.Group>
      <a
        href={`http://127.0.0.1:8080/api/v1/export/samples?charset=${charset}`}
        download
      >
        csv
      </a>

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
