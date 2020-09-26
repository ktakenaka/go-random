import React, { Fragment } from "react";
import { LiWrapper } from "./style";
import { TypeSample } from "constants/type";

type Props = {
  samples: Array<TypeSample>;
};

const SampleList = ({ samples }: Props) => {
  return (
    <Fragment>
      {samples.map((sample, index) => (
        <Fragment key={index}>
          <h3>{index + 1}</h3>
          <LiWrapper>{sample.title}</LiWrapper>
          <LiWrapper>{sample.content}</LiWrapper>
        </Fragment>
      ))}
    </Fragment>
  );
};

export default SampleList;
