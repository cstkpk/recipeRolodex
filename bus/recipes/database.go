package busrecipes

import (
	"context"
	"database/sql"

	"github.com/cstkpk/recipeRolodex/constant"
	"github.com/cstkpk/recipeRolodex/logger"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/cstkpk/recipeRolodex/mysql"
)

// GetRecipesList queries the DB for recipes that match the requested parameters and returns a list
func GetRecipesList(ctx context.Context, ing1, ing2, ing3, season string) (*models.Recipes, error) {
	db, err := mysql.Connect(ctx, constant.DBs.RecipeRolodex)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return nil, constant.Errors.DbConnectionFailure
	}

	var recipeIDs []int64
	// If the search includes at least one ingredient, get ingredient information first
	// Otherwise go straight to query based on season
	if ing1 != "" || ing2 != "" || ing3 != "" {
		// First find IDs associated with requested ingredients
		ingredientIDs, err := getIngredientIDs(ctx, ing1, ing2, ing3, db)
		if err != nil {
			logger.Error.Println(logger.GetCallInfo(), err.Error())
			return nil, err
		}

		// Then find recipeIDs associated with ingredientIDs
		ids, err := getRecipeIDs(ctx, ingredientIDs, db)
		if err != nil {
			logger.Error.Println(logger.GetCallInfo(), err.Error())
			return nil, err
		}
		recipeIDs = ids
	}

	// Get recipe details associated with recipeIDs and/or season
	recipes, err := getRecipes(ctx, recipeIDs, season, db)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return nil, err
	}

	logger.Info.Println(logger.GetCallInfo(), "Successfully returned recipe list")
	return recipes, nil
}

func getIngredientIDs(ctx context.Context, ing1, ing2, ing3 string, db *sql.DB) ([]int64, error) {

	ingredientQuery := `SELECT id FROM ` + constant.RR.Ingredients +
		` WHERE (name LIKE ? OR name LIKE ? OR name LIKE ?)`

	logger.Info.Println(ingredientQuery)
	rows, err := db.QueryContext(ctx, ingredientQuery, ing1+"%", ing2+"%", ing3+"%")
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return nil, constant.Errors.DbQueryFailure
	}
	defer rows.Close()

	var ingredientIDs []int64
	for rows.Next() {
		var ingredientID int64
		err = rows.Scan(
			&ingredientID,
		)
		if err != nil {
			logger.Error.Println(logger.GetCallInfo(), err.Error())
			return nil, constant.Errors.InternalServer
		}
		ingredientIDs = append(ingredientIDs, ingredientID)
	}

	if ingredientIDs == nil {
		logger.Error.Println(logger.GetCallInfo(), constant.Errors.NoRecipesFound.Error())
		return nil, constant.Errors.NoRecipesFound
	}

	logger.Info.Printf("%s Successfully returned ingredientIDs: %v", logger.GetCallInfo(), ingredientIDs)
	return ingredientIDs, nil
}

func getRecipeIDs(ctx context.Context, ingredientIDs []int64, db *sql.DB) ([]int64, error) {

	linkQuery := `SELECT recipeID FROM ` + constant.RR.Link +
		` WHERE 1=1`
	var args []interface{}
	for i, id := range ingredientIDs {
		if i == 0 {
			args = append(args, id)
			linkQuery += ` AND (ingredientID=?`
		} else {
			args = append(args, id)
			linkQuery += ` OR ingredientID=?`
		}
	}
	if len(ingredientIDs) != 0 {
		linkQuery += `)`
	}

	rows, err := db.QueryContext(ctx, linkQuery, args...)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return nil, constant.Errors.DbQueryFailure
	}
	defer rows.Close()

	var recipeIDs []int64
	for rows.Next() {
		var recipeID int64
		err = rows.Scan(
			&recipeID,
		)
		if err != nil {
			logger.Error.Println(logger.GetCallInfo(), err.Error())
			return nil, constant.Errors.InternalServer
		}
		recipeIDs = append(recipeIDs, recipeID)
	}

	if len(recipeIDs) == 0 {
		logger.Error.Println(logger.GetCallInfo(), constant.Errors.NoRecipesFound.Error())
		return nil, constant.Errors.NoRecipesFound
	}

	logger.Info.Printf("%s Successfully returned recipeIDs: %v", logger.GetCallInfo(), recipeIDs)
	return recipeIDs, nil
}

func getRecipes(ctx context.Context, recipeIDs []int64, season string, db *sql.DB) (*models.Recipes, error) {

	query := `SELECT id, season, title, author, link FROM ` + constant.RR.Recipes +
		` WHERE 1=1`
	var args2 []interface{}
	for i, id := range recipeIDs {
		if i == 0 {
			args2 = append(args2, id)
			query += ` AND (id=?`
		} else {
			args2 = append(args2, id)
			query += ` OR id=?`
		}
	}
	if len(recipeIDs) != 0 {
		query += `)`
	}

	if season != "" && season != "any" && season != "Any" {
		args2 = append(args2, season)
		query += ` AND season=?`
	}

	rows, err := db.QueryContext(ctx, query, args2...)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return nil, constant.Errors.DbQueryFailure
	}
	defer rows.Close()

	var recipeList []*models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		err = rows.Scan(
			&recipe.ID,
			&recipe.Season,
			&recipe.Title,
			&recipe.Author,
			&recipe.Link,
		)
		if err != nil {
			logger.Error.Println(logger.GetCallInfo(), err.Error())
			return nil, err
		}
		recipeList = append(recipeList, &recipe)
	}
	if recipeList == nil {
		logger.Error.Println(logger.GetCallInfo(), constant.Errors.NoRecipesFound.Error())
		return nil, constant.Errors.NoRecipesFound
	}

	var recipes *models.Recipes
	recipes = &models.Recipes{}
	recipes.RecipeList = recipeList

	logger.Info.Printf("%s Successfully returned %v recipe(s)", logger.GetCallInfo(), len(recipes.RecipeList))
	return recipes, nil
}
