import { put, call, takeLatest } from "redux-saga/effects";

import { getSamplesAPI, postSampleAPI } from "api/app/sample";
import { SUBMIT_SAMPLE_REQUEST, GET_SAMPLES_REQUEST } from "store/actionTypes";
import * as sampleAction from "store/actionCreators/sample";
import * as appAction from "store/actionCreators/app";

function* createSample(action: any) {
  try {
    yield call(postSampleAPI, action.payload);
    yield put(sampleAction.submitSampleSuccess());
    yield put(sampleAction.getSamplesRequest());
    yield put(appAction.setMessage("Succeed to create a sample", true));
    yield put(appAction.changeLocation("/samples"));
  } catch (err) {
    yield put(appAction.setMessage("failed to get submit sample", false));
    yield console.log(err);
    yield put(sampleAction.submitSampleFailure());
  }
}

function* getSamples() {
  try {
    const res = yield call(getSamplesAPI);
    yield put(sampleAction.getSamplesSuccess(res.data.data));
  } catch (err) {
    yield put(appAction.setMessage("falied to get samples", false));
    yield console.log(err);
    yield put(sampleAction.getSamplesFailure());
  }
}

export default function* actionWatcher(): Generator {
  yield takeLatest(SUBMIT_SAMPLE_REQUEST, createSample);
  yield takeLatest(GET_SAMPLES_REQUEST, getSamples);
}
