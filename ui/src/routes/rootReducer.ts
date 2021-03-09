import {combineReducers} from 'redux';
import {reducer as searchRecipes} from './searchRecipes/index';

const rootReducer = combineReducers({
  searchRecipes,
});

export default rootReducer;