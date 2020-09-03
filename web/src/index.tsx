import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";

import { createStore, applyMiddleware, compose } from "redux";
import createSagaMiddleware from "redux-saga";

import "./assets/index.css";
import App from "./App";
// import * as serviceWorker from './serviceWorker';
import rootSaga from "./store/saga";
import reducer from "./store/reducer";

const sagaMiddleware = createSagaMiddleware();

const composeEnhancers = (window as any).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

// re-consider naming when refactoring, combine reducer?
export const store = createStore(reducer, composeEnhancers(applyMiddleware(sagaMiddleware)));

sagaMiddleware.run(rootSaga);

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById("root")
);
