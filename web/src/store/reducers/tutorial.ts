import { SAMPLE_INCREMENT, SAMPLE_DECREMENT } from "store/actionTypes";

const initialState = {
  count: 0,
};

type State = {
  count: number;
};

export default (
  state: State = initialState,
  action: { type: string }
): State => {
  switch (action.type) {
    case SAMPLE_INCREMENT:
      return { ...state, count: state.count + 1 };
    case SAMPLE_DECREMENT:
      return { ...state, count: state.count - 1 };
    default:
      return state;
  }
};
