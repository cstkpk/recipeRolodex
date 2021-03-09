import {all, put, takeLatest} from 'redux-saga/effects';

import * as actions from './actions';
import * as actionType from './actionTypes';

export function* getRecipes(action: ReturnType<typeof actions.getRecipes>) {

  // TODO: create a util function to put params together instead of doing it here
  // ^^ maybe a util function for the whole fetch (options, response, etc.)
  interface QueryParams {
    [key: string]: string | number | boolean;
  }
  const queryParams: QueryParams = {
    ingredient1: action.searchQuery.ingredient1,
    ingredient2: action.searchQuery.ingredient2,
    ingredient3: action.searchQuery.ingredient3,
    season: action.searchQuery.season
  }
  
  let query: string = '';
  const params: string[] = [];
  for (const p in queryParams) {
    if (queryParams.hasOwnProperty(p)) {
      params.push(`${p}=${queryParams[p] ?? ''}`);
    }
  }
  query = `?${params.join('&')}`;
  console.log(query)

  const response = yield fetch("http://localhost:6010/rolodex/recipes"+query, {
    headers: {
      "Content-Type": "application/json"
    },
    method: "GET",
  }).then(response => response.json());

  if (!response) {
    return;
  }

  if (response.recipeList) {
    yield put(actions.getRecipesSuccess(response.recipeList));
  } else {
    yield put(actions.getRecipesFailure(response));
  }
  return;
}

export default function* root() {
  yield all([
    takeLatest(actionType.GET_RECIPES, getRecipes),
  ]);
}