// Code generated by go-swagger; DO NOT EDIT.

package recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cstkpk/recipeRolodex/models"
)

// GetRecipesOKCode is the HTTP code returned for type GetRecipesOK
const GetRecipesOKCode int = 200

/*GetRecipesOK successful

swagger:response getRecipesOK
*/
type GetRecipesOK struct {

	/*
	  In: Body
	*/
	Payload *models.Recipes `json:"body,omitempty"`
}

// NewGetRecipesOK creates GetRecipesOK with default headers values
func NewGetRecipesOK() *GetRecipesOK {

	return &GetRecipesOK{}
}

// WithPayload adds the payload to the get recipes o k response
func (o *GetRecipesOK) WithPayload(payload *models.Recipes) *GetRecipesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recipes o k response
func (o *GetRecipesOK) SetPayload(payload *models.Recipes) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecipesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRecipesNotFoundCode is the HTTP code returned for type GetRecipesNotFound
const GetRecipesNotFoundCode int = 404

/*GetRecipesNotFound not found

swagger:response getRecipesNotFound
*/
type GetRecipesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewGetRecipesNotFound creates GetRecipesNotFound with default headers values
func NewGetRecipesNotFound() *GetRecipesNotFound {

	return &GetRecipesNotFound{}
}

// WithPayload adds the payload to the get recipes not found response
func (o *GetRecipesNotFound) WithPayload(payload *models.ReturnCode) *GetRecipesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recipes not found response
func (o *GetRecipesNotFound) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecipesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRecipesInternalServerErrorCode is the HTTP code returned for type GetRecipesInternalServerError
const GetRecipesInternalServerErrorCode int = 500

/*GetRecipesInternalServerError internal service error

swagger:response getRecipesInternalServerError
*/
type GetRecipesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewGetRecipesInternalServerError creates GetRecipesInternalServerError with default headers values
func NewGetRecipesInternalServerError() *GetRecipesInternalServerError {

	return &GetRecipesInternalServerError{}
}

// WithPayload adds the payload to the get recipes internal server error response
func (o *GetRecipesInternalServerError) WithPayload(payload *models.ReturnCode) *GetRecipesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recipes internal server error response
func (o *GetRecipesInternalServerError) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecipesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetRecipesDefault unexpected error

swagger:response getRecipesDefault
*/
type GetRecipesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ReturnCode `json:"body,omitempty"`
}

// NewGetRecipesDefault creates GetRecipesDefault with default headers values
func NewGetRecipesDefault(code int) *GetRecipesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetRecipesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get recipes default response
func (o *GetRecipesDefault) WithStatusCode(code int) *GetRecipesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get recipes default response
func (o *GetRecipesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get recipes default response
func (o *GetRecipesDefault) WithPayload(payload *models.ReturnCode) *GetRecipesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recipes default response
func (o *GetRecipesDefault) SetPayload(payload *models.ReturnCode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecipesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
