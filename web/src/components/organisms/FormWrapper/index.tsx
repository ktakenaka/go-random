import React from "react";
import FormInput from "../../atoms/FormInput";
import Button from "../../atoms/Button";
import Form from "../../molecules/Form";

type Props = {
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
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
