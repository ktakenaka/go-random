import {
  CREATE_SESSION_REQUEST,
  CREATE_SESSION_SUCCESS,
  CREATE_SESSION_FAILURE,
} from "store/actionTypes";

type Action = {
  type: string;
  payload?: any;
};

export const createSessionRequest = (code: string, nonce: string): Action => {
  return {
    type: CREATE_SESSION_REQUEST,
    payload: { code: code, nonce: nonce },
  };
};

export const createSessionSuccess = (): Action => {
  return { type: CREATE_SESSION_SUCCESS };
};

export const createSessionFailure = (): Action => {
  return { type: CREATE_SESSION_FAILURE };
};
