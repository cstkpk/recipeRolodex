package recipe

import (
	"context"
	"database/sql"

	"github.com/cstkpk/recipeRolodex/constant"
	"github.com/cstkpk/recipeRolodex/logger"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/cstkpk/recipeRolodex/mysql"
)

// GetRecipeDetails queries the DB to find details pertaining to a specified recipe ID
func GetRecipeDetails(ctx context.Context, recipeID int64) (*models.Recipe, error) {
	db, err := mysql.Connect(ctx, constant.DBs.RecipeRolodex)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return nil, constant.Errors.DbConnectionFailure
	}

	query := `SELECT season, title, author, link FROM ` + constant.RR.Recipes +
		` WHERE id=?`

	var details models.Recipe
	err = db.QueryRowContext(ctx, query, recipeID).Scan(
		&details.Season,
		&details.Title,
		&details.Author,
		&details.Link,
	)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		if err == sql.ErrNoRows {
			return nil, constant.Errors.NoRecipeIDFound
		}
		return nil, constant.Errors.DbQueryFailure
	}

	logger.Info.Printf("%s Successfully retrieved recipe details for recipe with ID %v", logger.GetCallInfo(), recipeID)
	return &details, nil
}

// PostRecipeDetails inserts the new recipe details into the DB
func PostRecipeDetails(ctx context.Context, newRecipe *models.NewRecipe) error {
	db, err := mysql.Connect(ctx, constant.DBs.RecipeRolodex)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return constant.Errors.DbConnectionFailure
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return constant.Errors.InternalServer
	}

	// Add recipe details to Recipes table
	err = InsertRecipeDetails(ctx, tx, *newRecipe.Season, *newRecipe.Title, *newRecipe.Author, *newRecipe.Link)
	if err != nil {
		tx.Rollback()
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return err
	}

	// Add ingredients to Ingredients table if they don't already exist
	if len(newRecipe.IngredientList) != 0 {
		err = InsertIngredients(ctx, tx, newRecipe.IngredientList)
		if err != nil {
			tx.Rollback()
			logger.Error.Println(logger.GetCallInfo(), err.Error())
			return err
		}
	}

	err = InsertLink(ctx, tx, *newRecipe.Title, newRecipe.IngredientList)
	if err != nil {
		tx.Rollback()
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return constant.Errors.DbInsertFailure
	}

	return nil
}
