import { takeLatest } from "redux-saga/effects";
import history from "browserHistory";

import { CHANGE_LOCATION, SET_MESSAGE } from "store/actionTypes";

function* changeLocation(action: any) {
  yield history.push(action.payload);
}

function* setMessage(action: any) {
  // TODO: show modal or banner
  yield console.log(action.payload);
}

export default function* actionWatcher(): Generator {
  yield takeLatest(CHANGE_LOCATION, changeLocation);
  yield takeLatest(SET_MESSAGE, setMessage);
}
