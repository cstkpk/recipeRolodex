package busrecipes

import (
	"context"
	"fmt"

	"github.com/cstkpk/recipeRolodex/models"
)

// GetRecipes gets a list of recipes that match a user's input
func GetRecipes(ctx context.Context, ing1, ing2, ing3, season string) (*models.Recipes, error) {

	list, err := GetRecipesList(ctx, ing1, ing2, ing3, season)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, err
	}

	return list, nil
}
