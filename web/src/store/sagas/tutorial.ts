import { SAMPLE_INCREMENT_ASYNC, SAMPLE_INCREMENT } from "../actionTypes";

import { put, takeEvery } from "redux-saga/effects";

const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));

function* incrementAsync() {
  yield delay(1000);
  yield put({ type: SAMPLE_INCREMENT });
}

export default function* actionWatcher() {
  yield takeEvery(SAMPLE_INCREMENT_ASYNC, incrementAsync);
}
