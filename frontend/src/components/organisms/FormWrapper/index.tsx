import React from "react";
import FormInput from "components/atoms/FormInput";
import Button from "components/atoms/Button";
import Form from "components/molecules/Form";

type Props = {
  onSubmit: any;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const FormWrapper = ({ onSubmit, onChange }: Props) => {
  return (
    <Form onSubmit={onSubmit}>
      <FormInput onChange={onChange} />
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

export default FormWrapper;
