import React from "react";
import { FormInputWrapper } from "./style";

type Props = {
  name: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const FormInput = ({ name, onChange }: Props) => {
  return (
    <FormInputWrapper>
      <input name={name} type="text" onChange={onChange} />
    </FormInputWrapper>
  );
};

export default FormInput;
