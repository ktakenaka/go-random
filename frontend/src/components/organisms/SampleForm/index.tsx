import React, { useState } from "react";
import FormInput from "components/atoms/FormInput";
import Button from "components/atoms/Button";
import Form from "components/molecules/Form";
import { TypeSample } from "constants/type";

type Props = {
  item: TypeSample;
  onSubmit: (sample: TypeSample) => void;
};

const SampleForm = ({ onSubmit, item }: Props) => {
  const [sample, setSample] = useState<any>(item);

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSample({ ...sample, [e.target.name]: e.target.value });
  };

  const onOK = () => {
    onSubmit(sample);
  };

  return (
    <Form onSubmit={onOK}>
      <FormInput name="title" onChange={onChange} />
      <FormInput name="content" onChange={onChange} />
      <Button
        type="submit"
        size="medium"
        color="blue"
        disabled={false}
        value="submit"
      />
    </Form>
  );
};

export default SampleForm;
