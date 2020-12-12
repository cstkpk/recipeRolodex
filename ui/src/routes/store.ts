import {applyMiddleware, compose, createStore} from 'redux';
import createSagaMiddleware from 'redux-saga';

import rootReducer from './rootReducer';
import rootSaga from './rootSaga';

declare global {
  interface Window {
    __REDUX_DEVTOOLS_EXTENSION_COMPOSE__: any;
  }
}

const sagaMiddleware = createSagaMiddleware();

const composeEnhancers =
  typeof window === 'object' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ // eslint-disable-line
    ? /* istanbul ignore next */ window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({}) // eslint-disable-line
    : compose;

const enhancer = composeEnhancers(applyMiddleware(sagaMiddleware /* other middleware*/));
const store = createStore(rootReducer, enhancer);

sagaMiddleware.run(rootSaga);

export type GlobalState = ReturnType<typeof store.getState>;

export default store;