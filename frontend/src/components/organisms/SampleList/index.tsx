import React, { Fragment } from "react";
import { LiWrapper } from "./style";

type Sample = {
  title: string;
};

type Props = {
  samples: Array<Sample>;
};

const SampleList = ({ samples }: Props) => {
  return (
    <Fragment>
      {samples.map((sample, index) => (
        <LiWrapper key={index}>{sample.title}</LiWrapper>
      ))}
    </Fragment>
  );
};

export default SampleList;
