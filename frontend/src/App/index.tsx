import React, { Fragment } from "react";
import { Router, Route, Switch } from "react-router-dom";
import history from "../browserHistory";

import "./index.css";

import HomePage from "pages/Home";
import SignInPage from "pages/Google/SignIn";
import CallbackPage from "pages/Google/Callback";
import SampleListPage from "pages/SampleList";
import SampleNewPage from "pages/SampleNew";
import SampleEditPage from "pages/SampleEdit";

function App() {
  return (
    <Fragment>
      <Router history={history}>
        <Switch>
          <Route path="/home" exact component={HomePage} />
          <Route path="/google/sign-in" exact component={SignInPage} />
          <Route path="/google/callback" exact component={CallbackPage} />
          <Route path="/samples" exact component={SampleListPage} />
          <Route path="/samples/new" exact component={SampleNewPage} />
          <Route path="/samples/:id/edit" exact component={SampleEditPage} />
        </Switch>
      </Router>
    </Fragment>
  );
}

export default App;
