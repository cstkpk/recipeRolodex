import {all} from 'redux-saga/effects';
import searchRecipes from './searchRecipes';

function* rootSaga() {
  yield all([
    searchRecipes.sagas.default(),
  ]);
};

export default rootSaga;