// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Recipes recipes
//
// swagger:model Recipes
type Recipes struct {

	// recipe list
	RecipeList []*Recipe `json:"recipeList"`
}

// Validate validates this recipes
func (m *Recipes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipeList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Recipes) validateRecipeList(formats strfmt.Registry) error {

	if swag.IsZero(m.RecipeList) { // not required
		return nil
	}

	for i := 0; i < len(m.RecipeList); i++ {
		if swag.IsZero(m.RecipeList[i]) { // not required
			continue
		}

		if m.RecipeList[i] != nil {
			if err := m.RecipeList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipeList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Recipes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recipes) UnmarshalBinary(b []byte) error {
	var res Recipes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
