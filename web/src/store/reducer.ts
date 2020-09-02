import {
  SAMPLE_INCREMENT,
  SAMPLE_DECREMENT,
  SUBMIT_SAMPLE_REQUEST,
} from "./actions";

const initialState = {
  count: 0,
  samples: [],
};

interface State {
  count: number;
  samples: Array<Sample>;
}

interface Sample {
  title: string;
}

export default (
  state: State = initialState,
  action: { type: string; payload: any }
) => {
  switch (action.type) {
    case SAMPLE_INCREMENT:
      return { ...state, count: state.count + 1 };
    case SAMPLE_DECREMENT:
      return { ...state, count: state.count - 1 };
    case SUBMIT_SAMPLE_REQUEST:
      state.samples.push({ title: action.payload });
      return state;
    default:
      return state;
  }
};
