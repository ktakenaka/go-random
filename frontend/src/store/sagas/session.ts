import { put, call, takeLatest } from "redux-saga/effects";

import { postSessionAPI } from "api/app/session";
import { CREATE_SESSION_REQUEST } from "store/actionTypes";
import * as sessionAction from "store/actionCreators/session";
import * as appAction from "store/actionCreators/app";
import { CSRF_TOKEN_KEY } from "constants/auth";

function* createSession(action: any) {
  try {
    const res = yield call(postSessionAPI, action.payload);
    yield localStorage.setItem(CSRF_TOKEN_KEY, res.data.data.csrf_token);
    yield put(sessionAction.createSessionSuccess());
    yield put(appAction.setMessage("succeeded to sign in"));
    yield put(appAction.changeLocastion("/"));
  } catch (err) {
    yield put(appAction.setMessage(err));
    yield put(sessionAction.createSessionFailure());
    yield put(appAction.changeLocastion("/"));
  }
}

export default function* actionWatcher(): Generator {
  yield takeLatest(CREATE_SESSION_REQUEST, createSession);
}
