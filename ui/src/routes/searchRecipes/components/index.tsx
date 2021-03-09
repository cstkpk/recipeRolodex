import * as React from 'react';

import SearchRecipesForm from  './SearchRecipesForm';

import styles from './index.scss';

const SearchRecipes = () => {

  return (
    <>
      <h3 className={styles.title}>Search Recipes</h3>
      <SearchRecipesForm />
    </>
  )
};

export default SearchRecipes;