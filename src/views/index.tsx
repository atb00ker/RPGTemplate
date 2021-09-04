import React from 'react';
import ReactDOM from 'react-dom';
import { HashRouter as Router, Route, Switch } from 'react-router-dom';
import RPGTemplate from './components/RPGTemplate/RPRTemplate';

const App = () => (
  <Router>
    <Switch>
      <Route exact path="/" component={RPGTemplate} />
    </Switch>
  </Router>
);

ReactDOM.render(<App />, document.getElementById('react-init'));
