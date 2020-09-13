import {
  CHANGE_LOCATION,
  SET_MESSAGE,
  UNSET_MESSAAGE,
} from "store/actionTypes";

type Action = {
  type: string;
  payload?: any;
};

export const changeLocation = (location: string): Action => {
  return { type: CHANGE_LOCATION, payload: location };
};

export const setMessage = (message: string, success = true): Action => {
  return { type: SET_MESSAGE, payload: { message: message, success: success } };
};

export const unsetMessage = (): Action => {
  return { type: UNSET_MESSAAGE };
};
