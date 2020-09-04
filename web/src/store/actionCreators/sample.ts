import { SUBMIT_SAMPLE_REQUEST } from "store/actionTypes";

type Action = {
  type: string,
  payload: string
}

export const submitSample = (title: string):Action => {
  return { type: SUBMIT_SAMPLE_REQUEST, payload: title };
};
