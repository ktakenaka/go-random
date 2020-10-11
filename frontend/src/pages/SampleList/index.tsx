import React, { useEffect, useState } from "react";
import { connect } from "react-redux";
import { Link } from "react-router-dom";
import { Radio, PageHeader, Button } from "antd";

import MainTemplate from "components/templates/MainTemplate";
import SampleList from "components/organisms/SampleList";
import { getSamplesRequest, cleanupSample } from "store/actionCreators/sample";
import { TypeSample } from "constants/type";

const charsets = [
  { label: "UTF-8", value: "utf8" },
  { label: "Shift_JIS", value: "sjis" },
];

interface Props {
  samples: Array<TypeSample>;
  getSamplesRequest: typeof getSamplesRequest;
  cleanupSample: typeof cleanupSample;
}

const SamplePage = ({ samples, getSamplesRequest }: Props) => {
  const [charset, setCharset] = useState<"utf8" | "sjis">("utf8");

  useEffect(() => {
    getSamplesRequest();
  }, [getSamplesRequest]);

  const onCharsetSelected = (e: any) => {
    setCharset(e.target.value);
  };

  // TODO: define organism
  const pageHeader = (
    <PageHeader
      ghost={false}
      title="Sample"
      subTitle="This is a sample to practice coding"
      extra={[
        <Button key="1" type="primary">
          <Link to="/samples/new">New</Link>
        </Button>,
      ]}
    ></PageHeader>
  );

  // TODO: define organism
  const csvExport = (
    <>
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
        CSV Export
      </a>
    </>
  );

  return (
    <MainTemplate>
      {pageHeader}
      <SampleList header="Sample List" samples={samples} footer={csvExport} />
    </MainTemplate>
  );
};

const mapStateToProps = (state: Readonly<any>) => ({
  samples: state.getIn(["sample", "list"]),
});

const mapDispatchToProps = {
  getSamplesRequest,
  cleanupSample,
};

export default connect(mapStateToProps, mapDispatchToProps)(SamplePage);
