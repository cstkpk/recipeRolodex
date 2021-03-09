import * as actionType from './actionTypes';

import {SearchRecipesForm} from './formValidation';

export const getRecipes = (searchQuery: SearchRecipesForm) => 
({
  searchQuery,
  type: actionType.GET_RECIPES,
} as const);
export const getRecipesSuccess = (recipeList: RecipeRolodex.Recipe[]) =>
({
  recipeList,
  type: actionType.GET_RECIPES_SUCCESS,
} as const);
export const getRecipesFailure = (error: RecipeRolodex.ReturnCode) =>
({
  error,
  type: actionType.GET_RECIPES_FAILURE,
} as const); 

export type Actions = ReturnType<typeof getRecipes | typeof getRecipesSuccess | typeof getRecipesFailure>;
