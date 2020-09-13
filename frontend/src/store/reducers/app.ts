import { SET_MESSAGE, UNSET_MESSAAGE } from "store/actionTypes";

const initialState = {
  message: null,
  success: true,
  showMessage: false,
};

type State = {
  message: string | null;
  success: boolean;
  showMessage: boolean;
};

export default (
  state: State = initialState,
  action: { type: string; payload: any }
): State => {
  switch (action.type) {
    case SET_MESSAGE:
      return {
        ...state,
        message: action.payload.message,
        success: action.payload.success,
        showMessage: true,
      };
    case UNSET_MESSAAGE:
      return {
        ...state,
        message: null,
        showMessage: false,
      };
    default:
      return state;
  }
};
