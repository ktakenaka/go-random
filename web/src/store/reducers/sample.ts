import { SUBMIT_SAMPLE_REQUEST } from "store/actionTypes";

const initialState = {
  title: null,
  list: [],
};

type State = {
  title: string | null;
  list: Array<Sample>;
};

type Sample = {
  title: string | null;
};

export default (
  state: State = initialState,
  action: { type: string; payload: string }
):State => {
  switch (action.type) {
    case SUBMIT_SAMPLE_REQUEST:
      return {
        ...state,
        list: state.list.concat([{ title: action.payload }]),
      };
    default:
      return state;
  }
};
