import React, { useState } from 'react';

import { MainTemplate, SampleList, FormWrapper } from '../../components';

const HomePage: React.FC = () => {
  const [samples, setSamples] = useState<string[]>(['sample1', 'sample2', 'sample3']);
  const [value, setvalue] = useState<string>('');

  const onChange = (e: React.ChangeEvent<HTMLInputElement>):void => {
    const value = e.target.value;
    setvalue(value);
  };

  const onSubmit = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    setSamples([...samples.concat(value)]);
  }
  
  return (
    <MainTemplate>
      <h2>HOME</h2>

      <SampleList samples={samples} />
      <FormWrapper onChange={onChange} onSubmit={onSubmit} />
    </MainTemplate>
  );
}

export default HomePage;
