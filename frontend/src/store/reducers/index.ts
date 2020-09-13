import { combineReducers } from "redux-immutable";
import tutorial from "./tutorial";
import app from "./app";
import sample from "./sample";
import session from "./session";

export default combineReducers({
  tutorial,
  app,
  sample,
  session,
});
