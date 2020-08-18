import React, { Fragment } from 'react';
import { Router, Route, Switch } from 'react-router-dom';
import history from './browserHistory';

import './App.css';

import {
  HomePage,
} from './components';

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
