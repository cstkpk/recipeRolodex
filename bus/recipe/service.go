package recipe

import (
	"context"
	"fmt"

	"github.com/cstkpk/recipeRolodex/models"
)

// GetRecipe calls a DB query function to find recipe details
func GetRecipe(ctx context.Context, recipeID int64) (*models.Recipe, error) {

	details, err := GetRecipeDetails(ctx, recipeID)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, err
	}

	return details, nil
}

// PostRecipe calls a DB query function to insert the new recipe details
func PostRecipe(ctx context.Context, newRecipe *models.NewRecipe) error {

	err := PostRecipeDetails(ctx, newRecipe)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return err
	}

	return nil
}
