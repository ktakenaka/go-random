import { put, call, takeLatest } from "redux-saga/effects";

import * as sampleAPI from "api/app/sample";
import {
  SUBMIT_SAMPLE_REQUEST,
  GET_SAMPLES_REQUEST,
  UPDATE_SAMPLE_REQUEST,
  GET_SAMPLE_REQUEST,
  DELETE_SAMPLE_REQUEST,
} from "store/actionTypes";
import * as sampleAction from "store/actionCreators/sample";
import * as appAction from "store/actionCreators/app";

function* createSample(action: any) {
  try {
    yield call(sampleAPI.postSampleAPI, action.payload);
    yield put(sampleAction.submitSampleSuccess());
    yield put(sampleAction.getSamplesRequest());
    yield put(appAction.setMessage("Succeed to create a sample", true));
    yield put(appAction.changeLocation("/samples"));
  } catch (err) {
    yield put(appAction.setMessage("Failed to submit sample", false));
    yield console.log(err);
    yield put(sampleAction.submitSampleFailure(err));
  }
}

function* getSamples() {
  try {
    const res = yield call(sampleAPI.list);
    yield put(sampleAction.getSamplesSuccess(res.data.data));
  } catch (err) {
    yield put(appAction.setMessage("Falied to get samples", false));
    yield console.log(err);
    yield put(sampleAction.getSamplesFailure());
  }
}

function* updateSample({ id, payload }: any) {
  try {
    yield call(sampleAPI.putSampleAPI, id, payload);
    yield put(sampleAction.updateSampleSuccess());
    yield put(appAction.setMessage("Succeed to update sample", true));
    yield put(appAction.changeLocation("/samples"));
  } catch (err) {
    yield put(appAction.setMessage("Falied to update sample", false));
    yield console.log(err);
    yield put(sampleAction.updateSampleFailure(err));
  }
}

function* getSample({ id }: any) {
  try {
    const res = yield call(sampleAPI.get, id);
    yield put(sampleAction.getSampleSuccess(res.data.data));
  } catch (err) {
    yield console.log(err);
    yield put(sampleAction.getSampleFailure());
  }
}

function* deleteSample({ id }: any) {
  try {
    yield call(sampleAPI.destroy, id);
    yield put(sampleAction.deleteSampleSuccess());
    yield put(appAction.setMessage("Succeed to delete a sample", true));
    yield put(sampleAction.getSamplesRequest());
  } catch (err) {
    yield console.log(err);
    yield put(sampleAction.deleteSampleFailure(err));
  }
}

export default function* actionWatcher(): Generator {
  yield takeLatest(SUBMIT_SAMPLE_REQUEST, createSample);
  yield takeLatest(GET_SAMPLES_REQUEST, getSamples);
  yield takeLatest(UPDATE_SAMPLE_REQUEST, updateSample);
  yield takeLatest(GET_SAMPLE_REQUEST, getSample);
  yield takeLatest(DELETE_SAMPLE_REQUEST, deleteSample);
}
