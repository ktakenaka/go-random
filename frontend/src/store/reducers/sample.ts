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

const initialState = {
  loading: false,
  list: [],
  item: null,
  itemError: null,
};

type State = {
  loading: boolean;
  list: Array<TypeSample>;
  item: TypeSample | null;
  itemError: any;
};

export default (
  state: State = initialState,
  action: { type: string; payload: any }
): State => {
  switch (action.type) {
    case SUBMIT_SAMPLE_REQUEST ||
      GET_SAMPLES_REQUEST ||
      UPDATE_SAMPLE_REQUEST ||
      GET_SAMPLE_REQUEST:
      return {
        ...state,
        loading: true,
      };
    case SUBMIT_SAMPLE_SUCCESS || UPDATE_SAMPLE_SUCCESS:
      return {
        ...state,
        loading: false,
      };
    case GET_SAMPLES_FAILURE || GET_SAMPLE_FAILURE:
      return {
        ...state,
        loading: false,
      };
    case GET_SAMPLES_SUCCESS:
      return {
        ...state,
        loading: false,
        list: action.payload,
      };
    case GET_SAMPLE_SUCCESS:
      return {
        ...state,
        loading: false,
        item: action.payload,
      };
    case SUBMIT_SAMPLE_FAILURE || UPDATE_SAMPLE_FAILURE:
      return {
        ...state,
        loading: false,
        itemError: action.payload,
      };
    case CLEANUP_SAMPLE:
      return initialState;
    default:
      return state;
  }
};
