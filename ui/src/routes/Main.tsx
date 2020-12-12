import React from 'react';
import {Route, Switch} from 'react-router-dom';

import Home from './home/components';

const Main = () => {
  return (
    <Switch>
      <Route exact={true} path="/" component={Home}/>
    </Switch>
  )
}

export default Main;