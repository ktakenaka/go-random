import React, { Fragment } from 'react';

type Props = {
  samples: string[]
}

const SampleList = ({samples}: Props) => {
  return (
    <Fragment>
      {samples.map((sample, index) => (
        <li key={index}>{sample}</li>
      ))}
    </Fragment>
  )
}

export default SampleList;
