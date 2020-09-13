import { CHANGE_LOCATION, SET_MESSAGE } from "store/actionTypes";

type Action = {
  type: string;
  payload: string;
};

export const changeLocastion = (location: string): Action => {
  return { type: CHANGE_LOCATION, payload: location };
};

export const setMessage = (message: any): Action => {
  return { type: SET_MESSAGE, payload: message };
};
