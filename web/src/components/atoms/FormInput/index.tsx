import React from 'react';

type Props = {
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

const FormInput = ({onChange}:Props) => {
  return <input type='text' onChange={onChange} />;
}

export default FormInput;
