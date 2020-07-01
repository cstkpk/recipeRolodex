package busrecipes

import (
	"context"
	"fmt"

	"github.com/cstkpk/recipeRolodex/constant"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/cstkpk/recipeRolodex/mysql"
)

// GetRecipesList queries the DB for recipes that match the requested parameters and returns a list
func GetRecipesList(ctx context.Context, ing1, ing2, ing3, season string) (*models.Recipes, error) {
	db, err := mysql.Connect(ctx, constant.DBs.RecipeRolodex)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, constant.Errors.DbConnectionFailure
	}

	// First find IDs associated with requested ingredients
	ingredientQuery := `SELECT autoID FROM ` + constant.RR.Ingredients +
		` WHERE (name=? OR name=? OR name=?)`

	rows, err := db.QueryContext(ctx, ingredientQuery, ing1, ing2, ing3)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, constant.Errors.DbQueryFailure
	}
	defer rows.Close()

	var ingredientIDs []string
	for rows.Next() {
		var ingredientID string
		err = rows.Scan(
			&ingredientID,
		)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return nil, constant.Errors.InternalServer
		}
		ingredientIDs = append(ingredientIDs, ingredientID)
	}

	// Then find recipeIDs associated with ingredientIDs
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

	rows, err = db.QueryContext(ctx, linkQuery, args...)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, constant.Errors.DbQueryFailure
	}
	defer rows.Close()

	var recipeIDs []string
	for rows.Next() {
		var recipeID string
		err = rows.Scan(
			&recipeID,
		)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return nil, constant.Errors.InternalServer
		}
		recipeIDs = append(recipeIDs, recipeID)
	}

	if len(recipeIDs) == 0 {
		fmt.Println("Error:", constant.Errors.NoRecipesFound.Error())
		return nil, constant.Errors.NoRecipesFound
	}

	// Then get recipe details associated with recipeIDs (and season if included in query)
	query := `SELECT autoID, season, title, author, link FROM ` + constant.RR.Recipes +
		` WHERE 1=1`
	var args2 []interface{}
	for i, id := range recipeIDs {
		if i == 0 {
			args2 = append(args2, id)
			query += ` AND (autoID=?`
		} else {
			args2 = append(args2, id)
			query += ` OR autoID=?`
		}
	}
	if len(recipeIDs) != 0 {
		query += `)`
	}

	if season != "" && season != "any" && season != "Any" {
		args2 = append(args2, season)
		query += ` AND season=?`
	}

	rows, err = db.QueryContext(ctx, query, args2...)
	if err != nil {
		fmt.Println("Error:", err.Error())
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
			fmt.Println("Error: ", err.Error())
			return nil, err
		}
		recipeList = append(recipeList, &recipe)
	}
	if recipeList == nil {
		fmt.Println("Error: recipeList is empty")
		return nil, constant.Errors.InternalServer
	}

	var recipes *models.Recipes
	recipes = &models.Recipes{}
	recipes.RecipeList = recipeList

	fmt.Println("Info: Successfully returned recipe list")
	return recipes, nil
}
