package recipes

import (
	busrecipes "github.com/cstkpk/recipeRolodex/bus/recipes"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/go-openapi/runtime/middleware"
)

// Get calls GetRecipes and returns a recipe list based on user queries
func Get(params GetRecipesParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	list, err := busrecipes.GetRecipes(ctx, *params.Ingredient1, *params.Ingredient2, *params.Ingredient3, *params.Season)

	if err != nil {
		status := models.ReturnCode{Code: int64(GetRecipesInternalServerErrorCode), Message: err.Error()}
		return NewGetRecipesInternalServerError().WithPayload(&status)
	}

	return NewGetRecipesOK().WithPayload(list)
}
