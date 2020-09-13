import { put, takeLatest } from "redux-saga/effects";
import history from "browserHistory";

import { CHANGE_LOCATION, SET_MESSAGE } from "store/actionTypes";
import { unsetMessage } from "store/actionCreators/app";

function* changeLocation(action: any) {
  yield history.push(action.payload);
}

const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));

function* setMessage() {
  yield delay(3000);
  yield put(unsetMessage());
}

export default function* actionWatcher(): Generator {
  yield takeLatest(CHANGE_LOCATION, changeLocation);
  yield takeLatest(SET_MESSAGE, setMessage);
}
