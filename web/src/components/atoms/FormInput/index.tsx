import React from "react";
import { FormInputWrapper } from "./style";

type Props = {
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const FormInput = ({ onChange }: Props) => {
  return (
    <FormInputWrapper>
      <input type="text" onChange={onChange} />
    </FormInputWrapper>
  );
};

export default FormInput;
