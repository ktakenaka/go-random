import {
  SAMPLE_INCREMENT,
  SAMPLE_DECREMENT,
  SAMPLE_INCREMENT_ASYNC,
} from "store/actionTypes";

type Action = {
  type: string;
};

export const countIncrement = (): Action => {
  return { type: SAMPLE_INCREMENT };
};

export const countDecrement = (): Action => {
  return { type: SAMPLE_DECREMENT };
};

export const countIncrementAsync = (): Action => {
  return { type: SAMPLE_INCREMENT_ASYNC };
};
