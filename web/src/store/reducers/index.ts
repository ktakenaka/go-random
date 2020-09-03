import { combineReducers } from "redux-immutable";
import tutorial from "./tutorial";
import sample from "./sample";

export default combineReducers({
  tutorial,
  sample,
});
