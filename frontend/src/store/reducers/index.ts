import { combineReducers } from "redux-immutable";
import tutorial from "./tutorial";
import sample from "./sample";
import session from "./session";

export default combineReducers({
  tutorial,
  sample,
  session,
});
