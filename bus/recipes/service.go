package busrecipes

import (
	"context"

	"github.com/cstkpk/recipeRolodex/logger"
	"github.com/cstkpk/recipeRolodex/models"
)

// GetRecipes gets a list of recipes that match a user's input
func GetRecipes(ctx context.Context, ing1, ing2, ing3, season string) (*models.Recipes, error) {

	list, err := GetRecipesList(ctx, ing1, ing2, ing3, season)
	if err != nil {
		logger.Error.Println(logger.GetCallInfo(), err.Error())
		return nil, err
	}

	return list, nil
}
