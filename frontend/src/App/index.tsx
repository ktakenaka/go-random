import React, { Fragment } from "react";
import { Router, Route, Switch } from "react-router-dom";
import history from "../browserHistory";

import "./index.css";

import HomePage from "pages/Home";
import SignInPage from "pages/Google/SignIn";
import CallbackPage from "pages/Google/Callback";
import SampleListPage from "pages/SampleList";

function App() {
  return (
    <Fragment>
      <Router history={history}>
        <Switch>
          <Route path="/home" exact component={HomePage} />
          <Route path="/google/sign-in" component={SignInPage} />
          <Route path="/google/callback" component={CallbackPage} />
          <Route path="/samples" component={SampleListPage} />
        </Switch>
      </Router>
    </Fragment>
  );
}

export default App;
