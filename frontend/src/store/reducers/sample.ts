import {
  SUBMIT_SAMPLE_REQUEST,
  SUBMIT_SAMPLE_SUCCESS,
  SUBMIT_SAMPLE_FAILURE,
  GET_SAMPLES_REQUEST,
  GET_SAMPLES_SUCCESS,
  GET_SAMPLES_FAILURE,
  CLEANUP_SAMPLE,
} from "store/actionTypes";
import { TypeSample } from "constants/type";

const initialState = {
  postLoading: false,
  list: [],
  listLoading: false,
};

type State = {
  postLoading: boolean;
  list: Array<TypeSample>;
  listLoading: boolean;
};

export default (
  state: State = initialState,
  action: { type: string; payload: any }
): State => {
  switch (action.type) {
    case SUBMIT_SAMPLE_REQUEST:
      return {
        ...state,
        postLoading: true,
      };
    case SUBMIT_SAMPLE_SUCCESS:
      return {
        ...state,
        postLoading: false,
      };
    case SUBMIT_SAMPLE_FAILURE:
      return {
        ...state,
        postLoading: false,
      };
    case GET_SAMPLES_REQUEST:
      return {
        ...state,
        listLoading: true,
      };
    case GET_SAMPLES_SUCCESS:
      return {
        ...state,
        listLoading: false,
        list: action.payload,
      };
    case GET_SAMPLES_FAILURE:
      return {
        ...state,
        listLoading: false,
      };
    case CLEANUP_SAMPLE:
      return initialState;
    default:
      return state;
  }
};
