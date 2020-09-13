import React, { Fragment } from "react";
import { Router, Route, Switch } from "react-router-dom";
import history from "../browserHistory";

import "./index.css";

import { HomePage, SignInPage, CallbackPage } from "../pages";

function App() {
  return (
    <Fragment>
      <Router history={history}>
        <Switch>
          <Route path="/" exact component={HomePage} />
          <Route path="/google/sign-in" component={SignInPage} />
          <Route path="/google/callback" component={CallbackPage} />
        </Switch>
      </Router>
    </Fragment>
  );
}

export default App;
