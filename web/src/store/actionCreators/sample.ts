import { SUBMIT_SAMPLE_REQUEST } from "../actionTypes";

export const submitSample = (title: string) => {
  return { type: SUBMIT_SAMPLE_REQUEST, payload: title };
};
