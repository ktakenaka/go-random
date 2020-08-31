import React from 'react';
import { useFormikContext, getIn } from 'formik';
import { InputWrapper, ErrorMessage } from './style';

type Props = {
  name: string
}

const Input = ({ name }: Props) => {
  const { values, handleChange } = useFormikContext();

  return (
    <InputWrapper>
      <input
        value={getIn(values, name)}
        name={name}
        onChange={handleChange}
      />
    </InputWrapper>
  )
}

export default Input;
