import React, { Fragment } from 'react';
import { LiWrapper } from './style';

type Props = {
  samples: string[]
}

const SampleList = ({samples}: Props) => {
  return (
    <Fragment>
      {samples.map((sample, index) => (
        <LiWrapper key={index}>{sample}</LiWrapper>
      ))}
    </Fragment>
  )
}

export default SampleList;
