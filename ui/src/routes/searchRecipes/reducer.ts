import produce from 'immer';
import {Actions} from './actions';
import * as actionType from './actionTypes';

interface State {
  getRecipesRequest: {
    error: RecipeRolodex.ReturnCode;
    fetching: boolean;
    success: boolean;
  };
  recipeList: RecipeRolodex.Recipe[];
  searchQuery: RecipeRolodex.SearchRecipesQuery;
};

export const INITIAL_STATE: State = {
  getRecipesRequest: {
    error: {
      code: 0,
      id: '',
      message: '',
    },
    fetching: false,
    success: false,
  },
  recipeList: [
    {
      author: '',
      id: 0,
      link: '',
      season: '',
      title: '',
    },
  ],
  searchQuery: {
    season: '',
    ingredient1: '',
    ingredient2: '',
    ingredient3: '',
  },
};

const reducer = (state: State = INITIAL_STATE, action: Actions): State =>
  produce(state, (draft) => {
    switch(action.type) {
      case actionType.GET_RECIPES:
        draft.getRecipesRequest.error = INITIAL_STATE.getRecipesRequest.error;
        draft.getRecipesRequest.fetching = true;
        draft.getRecipesRequest.success = INITIAL_STATE.getRecipesRequest.success;
        draft.searchQuery = action.searchQuery;
        return;
      case actionType.GET_RECIPES_SUCCESS:
        draft.getRecipesRequest.error = INITIAL_STATE.getRecipesRequest.error;
        draft.getRecipesRequest.fetching = INITIAL_STATE.getRecipesRequest.fetching;
        draft.getRecipesRequest.success = true;;
        draft.recipeList = action.recipeList;
        return;
      case actionType.GET_RECIPES_FAILURE:
        draft.getRecipesRequest.error = action.error;
        draft.getRecipesRequest.fetching = INITIAL_STATE.getRecipesRequest.fetching;
        draft.getRecipesRequest.success = INITIAL_STATE.getRecipesRequest.success;
        return;
      default:
        return state;
    }
  });

export default reducer;
