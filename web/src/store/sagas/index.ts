import { all } from "redux-saga/effects";
import watchIncrementAsync from "./tutorial";

export default function* rootSaga():Generator {
  yield all([watchIncrementAsync()]);
}
