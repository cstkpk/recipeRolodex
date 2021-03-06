// Code generated by go-swagger; DO NOT EDIT.

package recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetRecipesParams creates a new GetRecipesParams object
// with the default values initialized.
func NewGetRecipesParams() GetRecipesParams {

	var (
		// initialize parameters with default values

		ingredient1Default = string("")
		ingredient2Default = string("")
		ingredient3Default = string("")
		seasonDefault      = string("")
	)

	return GetRecipesParams{
		Ingredient1: &ingredient1Default,

		Ingredient2: &ingredient2Default,

		Ingredient3: &ingredient3Default,

		Season: &seasonDefault,
	}
}

// GetRecipesParams contains all the bound params for the get recipes operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetRecipes
type GetRecipesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*an ingredient to filter recipes by
	  In: query
	  Default: ""
	*/
	Ingredient1 *string
	/*an ingredient to filter recipes by
	  In: query
	  Default: ""
	*/
	Ingredient2 *string
	/*an ingredient to filter recipes by
	  In: query
	  Default: ""
	*/
	Ingredient3 *string
	/*a season to filter recipes by
	  In: query
	  Default: ""
	*/
	Season *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetRecipesParams() beforehand.
func (o *GetRecipesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qIngredient1, qhkIngredient1, _ := qs.GetOK("ingredient1")
	if err := o.bindIngredient1(qIngredient1, qhkIngredient1, route.Formats); err != nil {
		res = append(res, err)
	}

	qIngredient2, qhkIngredient2, _ := qs.GetOK("ingredient2")
	if err := o.bindIngredient2(qIngredient2, qhkIngredient2, route.Formats); err != nil {
		res = append(res, err)
	}

	qIngredient3, qhkIngredient3, _ := qs.GetOK("ingredient3")
	if err := o.bindIngredient3(qIngredient3, qhkIngredient3, route.Formats); err != nil {
		res = append(res, err)
	}

	qSeason, qhkSeason, _ := qs.GetOK("season")
	if err := o.bindSeason(qSeason, qhkSeason, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIngredient1 binds and validates parameter Ingredient1 from query.
func (o *GetRecipesParams) bindIngredient1(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRecipesParams()
		return nil
	}

	o.Ingredient1 = &raw

	return nil
}

// bindIngredient2 binds and validates parameter Ingredient2 from query.
func (o *GetRecipesParams) bindIngredient2(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRecipesParams()
		return nil
	}

	o.Ingredient2 = &raw

	return nil
}

// bindIngredient3 binds and validates parameter Ingredient3 from query.
func (o *GetRecipesParams) bindIngredient3(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRecipesParams()
		return nil
	}

	o.Ingredient3 = &raw

	return nil
}

// bindSeason binds and validates parameter Season from query.
func (o *GetRecipesParams) bindSeason(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetRecipesParams()
		return nil
	}

	o.Season = &raw

	return nil
}
