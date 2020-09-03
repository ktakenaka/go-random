import {
  SAMPLE_INCREMENT,
  SAMPLE_DECREMENT,
  SUBMIT_SAMPLE_REQUEST,
} from "./actions";

const initialState = {
  count: 0,
  sample: { title: null },
  samples: [],
};

type State = {
  count: number;
  sample: Sample;
  samples: Array<Sample>;
};

type Sample = {
  title: string | null;
};

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
      return {
        ...state,
        samples: state.samples.concat([{ title: action.payload }]),
      };
    default:
      return state;
  }
};
