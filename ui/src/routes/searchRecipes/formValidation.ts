export interface SearchRecipesForm {
  season?: string;
  ingredient1?: string;
  ingredient2?: string;
  ingredient3?: string;
}

const alpha = /[^a-zA-Z\s]+/;

export const validateSearchRecipesForm = (values: SearchRecipesForm): Partial<SearchRecipesForm> => {
  const errors: Partial<SearchRecipesForm> = {};

  // ingredients must be a-zA-Z strings, no symbols or numbers
  if (values.ingredient1 && alpha.test(values.ingredient1)) {
    errors.ingredient1 = 'Ingredient must not contain any numbers or symbols'
  }
  if (values.ingredient2 && alpha.test(values.ingredient2)) {
    errors.ingredient2 = 'Ingredient must not contain any numbers or symbols'
  }
  if (values.ingredient3 && alpha.test(values.ingredient3)) {
    errors.ingredient3 = 'Ingredient must not contain any numbers or symbols'
  }

  return errors;
};