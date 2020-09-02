import { SAMPLE_INCREMENT_ASYNC, SAMPLE_INCREMENT } from "./actions";

import { put, takeEvery, all } from "redux-saga/effects";

const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));

// TODO: refactor the location
// saga is just a generator function.
// When a Promise is yielded to the middleware, the middleware will suspend the Saga until the Promise completes
function* incrementAsync() {
  yield delay(1000);
  yield put({ type: SAMPLE_INCREMENT });
}

function* watchIncrementAsync() {
  yield takeEvery(SAMPLE_INCREMENT_ASYNC, incrementAsync);
}

export default function* rootSaga() {
  yield all([watchIncrementAsync()]);
}
