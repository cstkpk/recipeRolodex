import * as React from 'react';
import {Helmet} from 'react-helmet';
import {Provider} from 'react-redux';
import {BrowserRouter} from 'react-router-dom';

import store from './store';
import Main from './Main';


const App = () => (
  <Provider store={store}>
    <BrowserRouter>
      <Helmet>
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <meta httpEquiv="X-UA-Compatible" content="ie=edge" />
        <title>Recipe Rolodex</title>
      </Helmet>
      <Main />
    </BrowserRouter>
  </Provider>
)

export default App;