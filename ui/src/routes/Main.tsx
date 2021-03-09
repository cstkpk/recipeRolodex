import React from 'react';
import {Route, Switch} from 'react-router-dom';

import Home from './home/components';
import SearchRecipes from './searchRecipes/components';

const Main = () => {
  return (
    <>
    <Switch>
      <Route exact={true} path="/" component={Home} />
      <Route exact={true} path="/search" component={SearchRecipes} />
    </Switch>
    </>
  )
}

export default Main;