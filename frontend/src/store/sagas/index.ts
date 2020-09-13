import { all } from "redux-saga/effects";
import tutorial from "./tutorial";
import app from "./app";
import sample from "./sample";
import session from "./session";

export default function* rootSaga(): Generator {
  yield all([tutorial(), app(), sample(), session()]);
}
