package ready

import (
	busready "github.com/cstkpk/recipeRolodex/bus/ready"
	"github.com/cstkpk/recipeRolodex/models"
	"github.com/go-openapi/runtime/middleware"
)

// Get calls GetReady to make sure that a db connection can be established
func Get(params GetReadyParams) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	err := busready.GetReady(ctx)

	if err != nil {
		status := models.ReturnCode{Code: int64(GetReadyInternalServerErrorCode), Message: err.Error()}
		return NewGetReadyInternalServerError().WithPayload(&status)
	}
	status := models.ReturnCode{Code: int64(GetReadyOKCode)}
	return NewGetReadyOK().WithPayload(&status)
}
