import {
  SUBMIT_SAMPLE_REQUEST,
  GET_SAMPLES_REQUEST,
  SUBMIT_SAMPLE_SUCCESS,
  SUBMIT_SAMPLE_FAILURE,
  GET_SAMPLES_SUCCESS,
  GET_SAMPLES_FAILURE,
} from "store/actionTypes";

type Action = {
  type: string;
  payload?: any;
};

export const submitSampleRequest = (title: string): Action => {
  return { type: SUBMIT_SAMPLE_REQUEST, payload: title };
};

export const submitSampleSuccess = (): Action => {
  return { type: SUBMIT_SAMPLE_SUCCESS };
};

export const submitSampleFailure = (): Action => {
  return { type: SUBMIT_SAMPLE_FAILURE };
};

export const getSamplesRequest = (): Action => {
  return { type: GET_SAMPLES_REQUEST };
};

export const getSamplesSuccess = (sampleList: Array<any>): Action => {
  return { type: GET_SAMPLES_SUCCESS, payload: sampleList };
};

export const getSamplesFailure = (): Action => {
  return { type: GET_SAMPLES_FAILURE };
};
