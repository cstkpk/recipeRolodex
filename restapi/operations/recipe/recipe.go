package recipe

import (
	busrecipe "github.com/cstkpk/recipeRolodex/bus/recipe"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/cstkpk/recipeRolodex/rrerror"
	"github.com/go-openapi/runtime/middleware"
)

// Get calls GetRecipe and returns recipe details based on a recipe ID
func Get(params GetRecipeParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	details, err := busrecipe.GetRecipe(ctx, params.RecipeID)

	if err != nil {
		switch e := err.(type) {
		case *rrerror.RRError:
			switch e.Code() {
			// 400
			case GetRecipeBadRequestCode:
				status := models.ReturnCode{Code: int64(GetRecipeBadRequestCode), Message: e.Error()}
				return NewGetRecipeBadRequest().WithPayload(&status)
			// 404
			case GetRecipeNotFoundCode:
				status := models.ReturnCode{Code: int64(GetRecipeNotFoundCode), Message: e.Error()}
				return NewGetRecipeNotFound().WithPayload(&status)
			}
		}
		// 500 / catch-all
		status := models.ReturnCode{Code: int64(GetRecipeInternalServerErrorCode), Message: err.Error()}
		return NewGetRecipeInternalServerError().WithPayload(&status)
	}
	return NewGetRecipeOK().WithPayload(details)
}

// Post inserts the new recipe details into the DB and returns an error or success code
func Post(params PostRecipeParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	err := busrecipe.PostRecipe(ctx, params.NewRecipe)

	if err != nil {
		switch e := err.(type) {
		case *rrerror.RRError:
			switch e.Code() {
			// 400
			case PostRecipeBadRequestCode:
				status := models.ReturnCode{Code: int64(PostRecipeBadRequestCode), Message: e.Error()}
				return NewPostRecipeBadRequest().WithPayload(&status)
			}
		}
		// 500 / catch-all
		status := models.ReturnCode{Code: int64(PostRecipeInternalServerErrorCode), Message: err.Error()}
		return NewPostRecipeInternalServerError().WithPayload(&status)
	}
	status := models.ReturnCode{Code: int64(PostRecipeOKCode), Message: "Successfully added new recipe"}
	return NewPostRecipeOK().WithPayload(&status)
}
