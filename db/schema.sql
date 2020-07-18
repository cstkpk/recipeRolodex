USE RecipeRolodex;

DROP TABLE IF EXISTS Link;
DROP TABLE IF EXISTS Recipes;
DROP TABLE IF EXISTS Ingredients;

CREATE TABLE Ingredients (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(50) NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

CREATE TABLE Recipes (
  id INT NOT NULL AUTO_INCREMENT,
  season VARCHAR(50) NOT NULL,
  title VARCHAR(100) NOT NULL UNIQUE,
  author VARCHAR(100) NOT NULL,
  link VARCHAR(150) NOT NULL UNIQUE,
  PRIMARY KEY (id)
);

/* When you click a link to a self-written recipe, it'll take you to a new page with a new
endpoint request that will load data from another db table holding all of the recipe info */
/* E.g. /recipe will be the endpoint for self-written recipe details, and otherwise to view 
details the link will take you to an external website */

/* Entering a recipe through the ui will happen through a form
Maybe each step is a textarea box with max characters?
Ingredients will have to be entered so that the ingredient name is search accessible
OR as you enter your recipe, you separately enter main ingredients somewhere */

/* Table to link Ingredients and Recipes (two foreign keys) */
CREATE TABLE Link (
  id INT NOT NULL AUTO_INCREMENT,
  recipeID INT NOT NULL,
  ingredientID INT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (recipeID) REFERENCES Recipes (id),
  FOREIGN KEY (ingredientID) REFERENCES Ingredients (id)
);