import React, { Fragment } from "react";
import { Router, Route, Switch } from "react-router-dom";
import history from "../browserHistory";

import "./index.css";

import { HomePage } from "../pages";

function App() {
  return (
    <Fragment>
      <Router history={history}>
        <Switch>
          <Route path="/" exact component={HomePage} />
        </Switch>
      </Router>
    </Fragment>
  );
}

export default App;
