import {
  SUBMIT_SAMPLE_REQUEST,
  GET_SAMPLES_REQUEST,
  SUBMIT_SAMPLE_SUCCESS,
  SUBMIT_SAMPLE_FAILURE,
  GET_SAMPLES_SUCCESS,
  GET_SAMPLES_FAILURE,
  CLEANUP_SAMPLE,
} from "store/actionTypes";
import { TypeSample } from "constants/type";

export const submitSampleRequest = (sample: TypeSample) => {
  return { type: SUBMIT_SAMPLE_REQUEST, payload: sample };
};

export const submitSampleSuccess = () => {
  return { type: SUBMIT_SAMPLE_SUCCESS };
};

export const submitSampleFailure = () => {
  return { type: SUBMIT_SAMPLE_FAILURE };
};

export const getSamplesRequest = () => {
  return { type: GET_SAMPLES_REQUEST };
};

export const getSamplesSuccess = (sampleList: Array<TypeSample>) => {
  return { type: GET_SAMPLES_SUCCESS, payload: sampleList };
};

export const getSamplesFailure = () => {
  return { type: GET_SAMPLES_FAILURE };
};

export const cleanupSample = () => {
  return { type: CLEANUP_SAMPLE };
};
