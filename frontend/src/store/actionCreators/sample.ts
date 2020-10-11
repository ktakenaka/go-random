import {
  SUBMIT_SAMPLE_REQUEST,
  SUBMIT_SAMPLE_SUCCESS,
  SUBMIT_SAMPLE_FAILURE,
  UPDATE_SAMPLE_REQUEST,
  UPDATE_SAMPLE_SUCCESS,
  UPDATE_SAMPLE_FAILURE,
  GET_SAMPLES_REQUEST,
  GET_SAMPLES_SUCCESS,
  GET_SAMPLES_FAILURE,
  GET_SAMPLE_REQUEST,
  GET_SAMPLE_SUCCESS,
  GET_SAMPLE_FAILURE,
  CLEANUP_SAMPLE,
} from "store/actionTypes";
import { TypeSample } from "constants/type";

export const submitSampleRequest = (sample: TypeSample) => {
  return { type: SUBMIT_SAMPLE_REQUEST, payload: sample };
};

export const submitSampleSuccess = () => {
  return { type: SUBMIT_SAMPLE_SUCCESS };
};

export const submitSampleFailure = (err: any) => {
  return { type: SUBMIT_SAMPLE_FAILURE, payload: err };
};

export const updateSampleRequest = (id: number, sample: any) => {
  return { type: UPDATE_SAMPLE_REQUEST, id: id, payload: sample };
};

export const updateSampleSuccess = () => {
  return { type: UPDATE_SAMPLE_SUCCESS };
};

export const updateSampleFailure = (err: any) => {
  return { type: UPDATE_SAMPLE_FAILURE, payload: err };
};

export const getSamplesRequest = () => {
  return { type: GET_SAMPLES_REQUEST };
};

export const getSamplesSuccess = (sampleList: Array<TypeSample>) => {
  return { type: GET_SAMPLES_SUCCESS, payload: sampleList };
};

export const getSampleFailure = () => {
  return { type: GET_SAMPLE_FAILURE };
};

export const getSampleRequest = (id: number) => {
  return { type: GET_SAMPLE_REQUEST, id: id };
};

export const getSampleSuccess = (sample: TypeSample) => {
  return { type: GET_SAMPLE_SUCCESS, payload: sample };
};

export const getSamplesFailure = () => {
  return { type: GET_SAMPLES_FAILURE };
};

export const cleanupSample = () => {
  return { type: CLEANUP_SAMPLE };
};
