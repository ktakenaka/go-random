export const SAMPLE_INCREMENT = "SAMPLE_INCREMENT";
export const SAMPLE_DECREMENT = "SAMPLE_DECREMENT";
export const SAMPLE_INCREMENT_ASYNC = "SAMPLE_INCREMENT_ASYNC";
export const SUBMIT_SAMPLE_REQUEST = "SUBMIT_SAMPLE_REQUEST";

export const actionCreate = (type: string) => {
  return { type: type };
};

export const submitSample = (title: string) => {
  return { type: SUBMIT_SAMPLE_REQUEST, payload: title };
};
