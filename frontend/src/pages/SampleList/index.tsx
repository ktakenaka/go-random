import React, { useEffect, useState } from "react";
import { connect } from "react-redux";
import { Link } from "react-router-dom";
import { Radio, PageHeader, Button, Upload } from "antd";
import { UploadFile } from "antd/lib/upload/interface";
import { UploadOutlined } from "@ant-design/icons";

import MainTemplate from "components/templates/MainTemplate";
import SampleList from "components/organisms/SampleList";
import {
  getSamplesRequest,
  deleteSampleRequest,
  importSamplesRequest,
  cleanupSample,
} from "store/actionCreators/sample";
import { exportURL as exportSamplesURL } from "api/app/sample";
import { TypeSample } from "constants/type";

const charsets = [
  { label: "UTF-8", value: "utf8" },
  { label: "Shift_JIS", value: "sjis" },
];

interface Props {
  samples: Array<TypeSample>;
  getSamplesRequest: typeof getSamplesRequest;
  deleteSampleRequest: typeof deleteSampleRequest;
  importSamplesRequest: typeof importSamplesRequest;
  cleanupSample: typeof cleanupSample;
}

const SamplePage = ({
  samples,
  getSamplesRequest,
  deleteSampleRequest,
  importSamplesRequest,
}: Props) => {
  const [charset, setCharset] = useState<"utf8" | "sjis">("utf8");
  const [file, setFile] = useState<UploadFile | null>();

  useEffect(() => {
    getSamplesRequest();
  }, [getSamplesRequest]);

  const beforeUpload = (file: File) => {
    // TODO: make it after users' confirmation
    importSamplesRequest(file);
    return false;
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
        <Upload
          key="2"
          name="file"
          fileList={file ? [file] : []}
          beforeUpload={beforeUpload}
          onRemove={() => setFile(null)}
        >
          <Button icon={<UploadOutlined />}>Upload</Button>
        </Upload>,
        //<Button onClick={handleUpload} disabled={!Boolean(file)}>Import</Button>
      ]}
    ></PageHeader>
  );

  const onCharsetSelected = (e: any) => {
    setCharset(e.target.value);
  };

  const handleDelete = (id: number) => {
    deleteSampleRequest(id);
  };

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
      <a href={exportSamplesURL(charset)} target="_blank" rel="noopener noreferrer" download>
        CSV Export
      </a>
    </>
  );

  return (
    <MainTemplate>
      {pageHeader}
      <SampleList
        header="Sample List"
        samples={samples}
        footer={csvExport}
        onDelete={handleDelete}
      />
    </MainTemplate>
  );
};

const mapStateToProps = (state: Readonly<any>) => ({
  samples: state.getIn(["sample", "list"]),
});

const mapDispatchToProps = {
  getSamplesRequest,
  deleteSampleRequest,
  importSamplesRequest,
  cleanupSample,
};

export default connect(mapStateToProps, mapDispatchToProps)(SamplePage);
