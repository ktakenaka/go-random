import React, { useEffect } from "react";
import { Form as FormAnt, Input as InputAnt, Button as ButtonAnt } from "antd";

import { TypeSample } from "constants/type";

type Props = {
  sample: TypeSample;
  onSubmit: (sample: TypeSample) => void;
};

const SampleForm = ({ onSubmit, sample }: Props) => {
  const [form] = FormAnt.useForm();

  useEffect(() => {
    form.setFieldsValue(sample as any);
  }, [form, sample]);

  const onFinish = (values: TypeSample) => {
    onSubmit(values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <FormAnt
      name="basic"
      form={form}
      labelCol={{ span: 3 }}
      wrapperCol={{ span: 20 }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
    >
      <FormAnt.Item
        label="Title"
        name="title"
        rules={[{ required: true, message: "Please input title" }]}
      >
        <InputAnt />
      </FormAnt.Item>

      <FormAnt.Item
        label="Content"
        name="content"
        rules={[{ required: true, message: "Please input content" }]}
      >
        <InputAnt />
      </FormAnt.Item>

      <FormAnt.Item wrapperCol={{ offset: 3 }}>
        <ButtonAnt type="primary" htmlType="submit">
          Submit
        </ButtonAnt>
      </FormAnt.Item>
    </FormAnt>
  );
};

export default SampleForm;
