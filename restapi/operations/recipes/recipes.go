package recipes

import (
	busrecipes "github.com/cstkpk/recipeRolodex/bus/recipes"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/cstkpk/recipeRolodex/rrerror"
	"github.com/go-openapi/runtime/middleware"
)

// Get calls GetRecipes and returns a recipe list based on user queries
func Get(params GetRecipesParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	list, err := busrecipes.GetRecipes(ctx, *params.Ingredient1, *params.Ingredient2, *params.Ingredient3, *params.Season)

	if err != nil {
		switch e := err.(type) {
		case *rrerror.RRError:
			switch e.Code() {
			// 404
			case GetRecipesNotFoundCode:
				status := models.ReturnCode{Code: int64(GetRecipesNotFoundCode), Message: e.Error()}
				return NewGetRecipesNotFound().WithPayload(&status)
			}
		}
		// 500 / catch-all
		status := models.ReturnCode{Code: int64(GetRecipesInternalServerErrorCode), Message: err.Error()}
		return NewGetRecipesInternalServerError().WithPayload(&status)
	}

	return NewGetRecipesOK().WithPayload(list)
}
