import { all } from "redux-saga/effects";
import tutorialSaga from "./tutorial";
import sampleSaga from "./sample";

export default function* rootSaga(): Generator {
  yield all([tutorialSaga(), sampleSaga()]);
}
