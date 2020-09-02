import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";

import { createStore, applyMiddleware } from "redux";
import createSagaMiddleware from "redux-saga";

import "./assets/index.css";
import App from "./App";
// import * as serviceWorker from './serviceWorker';
import rootSaga from "./store/saga";
import reducer from "./store/reducer";

const sagaMiddleware = createSagaMiddleware();

// re-consider naming when refactoring, combine reducer?
export const store = createStore(reducer, applyMiddleware(sagaMiddleware));

sagaMiddleware.run(rootSaga);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById("root")
);
