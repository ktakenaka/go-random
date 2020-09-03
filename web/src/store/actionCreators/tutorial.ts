import {
  SAMPLE_INCREMENT,
  SAMPLE_DECREMENT,
  SAMPLE_INCREMENT_ASYNC,
} from "../actionTypes";

export const countIncrement = () => {
  return { type: SAMPLE_INCREMENT };
};

export const countDecrement = () => {
  return { type: SAMPLE_DECREMENT };
};

export const countIncrementAsync = () => {
  return { type: SAMPLE_INCREMENT_ASYNC };
};
