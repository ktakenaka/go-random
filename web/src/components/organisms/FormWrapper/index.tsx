import React from 'react';
import FormInput from '../../atoms/FormInput';
import Form from '../../molecules/Form';

type Props = {
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void,
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void,
}

const FormWrapper = ({onSubmit, onChange}:Props) => {
  return (
    <Form onSubmit={onSubmit}>
      <FormInput onChange={onChange} />
      <input type='submit' value='submit' />
    </Form>
  );
};

export default FormWrapper;
