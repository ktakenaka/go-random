import { all } from "redux-saga/effects";
import watchIncrementAsync from "./tutorial";

export default function* rootSaga() {
  yield all([watchIncrementAsync()]);
}
