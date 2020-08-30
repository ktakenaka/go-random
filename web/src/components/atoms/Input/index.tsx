import React from 'react';
import { useFormikContext, getIn } from 'formik';
import { InputWrapper, ErrorMessage } from './style';

type Props = {
  name: string
}

const Input = ({ name }: Props) => {
  const { values, handleChange, handleBlur, touched, errors } = useFormikContext();

  let isInvalid = null;

  if (touched[name] && errors[name]) {
    isInvalid = true;
  }

  return (
    <InputWrapper>
      <input
        value={getIn(values, name)}
        name={name}
        className={isInvalid && `error`}
        onChange={handleChange}
        onBlur={handleBlur}
      />
      {isInvalid && <ErrorMessage>{errors[name]}</ErrorMessage>}
    </InputWrapper>
  )
}

export default Input;
