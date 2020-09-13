import {
  CREATE_SESSION_REQUEST,
  CREATE_SESSION_SUCCESS,
  CREATE_SESSION_FAILURE,
} from "store/actionTypes";

const initialState = {
  loading: false,
  errors: [],
};

type State = {
  loading: boolean;
  errors: Array<Error>;
};

// TODO: commonize the location of types
type Error = {
  source: Record<string, unknown>;
  detail: string;
};

export default (
  state: State = initialState,
  action: { type: string; payload: any }
): State => {
  switch (action.type) {
    case CREATE_SESSION_REQUEST:
      return {
        ...state,
        loading: true,
        errors: [],
      };
    case CREATE_SESSION_SUCCESS:
      return {
        ...state,
        loading: false,
        errors: [],
      };
    case CREATE_SESSION_FAILURE:
      return {
        ...state,
        loading: false,
        errors: action.payload,
      };
    default:
      return state;
  }
};
