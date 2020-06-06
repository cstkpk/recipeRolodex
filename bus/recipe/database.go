package recipe

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cstkpk/recipeRolodex/constant"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/cstkpk/recipeRolodex/mysql"
)

// GetRecipeDetails queries the DB to find details pertaining to a specified recipe ID
func GetRecipeDetails(ctx context.Context, recipeID int64) (*models.Recipe, error) {
	db, err := mysql.Connect(ctx, constant.DBs.RecipeRolodex)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, constant.Errors.DbConnectionFailure
	}

	query := `SELECT season, title, author, link FROM ` + constant.RR.Recipes +
		` WHERE autoID=?`

	var details models.Recipe
	err = db.QueryRowContext(ctx, query, recipeID).Scan(
		&details.Season,
		&details.Title,
		&details.Author,
		&details.Link,
	)
	if err != nil {
		fmt.Println("Error:", err.Error())
		if err == sql.ErrNoRows {
			return nil, constant.Errors.NoRecipeIDFound
		}
		return nil, constant.Errors.DbQueryFailure
	}

	fmt.Printf("Info: Successfully retrieved recipe details for recipe with ID %v", recipeID)
	return &details, nil
}
