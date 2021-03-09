declare namespace RecipeRolodex {
  
  interface ReturnCode {
    code?: number;
    id?: string;
    message?: string;
  }

  // SearchRecipesQuery is the search query used to filter recipes
  interface SearchRecipesQuery {
    ingredient1?: string;
    ingredient2?: string;
    ingredient3?: string;
    season?: string;
  }

  // Recipe is a single recipe item 
  interface Recipe {
    author: string;
    id: number;
    link: string;
    season: string;
    title: string;
  }
}