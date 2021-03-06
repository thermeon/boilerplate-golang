// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/thermeon/boilerplate-golang/infrastructures/server/models"
)

// AuthorizationTransactionCreatedCode is the HTTP code returned for type AuthorizationTransactionCreated
const AuthorizationTransactionCreatedCode int = 201

/*AuthorizationTransactionCreated successful operation

swagger:response authorizationTransactionCreated
*/
type AuthorizationTransactionCreated struct {
	/*location of the created resource

	 */
	Location string `json:"location"`
}

// NewAuthorizationTransactionCreated creates AuthorizationTransactionCreated with default headers values
func NewAuthorizationTransactionCreated() *AuthorizationTransactionCreated {

	return &AuthorizationTransactionCreated{}
}

// WithLocation adds the location to the authorization transaction created response
func (o *AuthorizationTransactionCreated) WithLocation(location string) *AuthorizationTransactionCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the authorization transaction created response
func (o *AuthorizationTransactionCreated) SetLocation(location string) {
	o.Location = location
}

// WriteResponse to the client
func (o *AuthorizationTransactionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header location

	location := o.Location
	if location != "" {
		rw.Header().Set("location", location)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AuthorizationTransactionBadRequestCode is the HTTP code returned for type AuthorizationTransactionBadRequest
const AuthorizationTransactionBadRequestCode int = 400

/*AuthorizationTransactionBadRequest Bad Request

swagger:response authorizationTransactionBadRequest
*/
type AuthorizationTransactionBadRequest struct {
}

// NewAuthorizationTransactionBadRequest creates AuthorizationTransactionBadRequest with default headers values
func NewAuthorizationTransactionBadRequest() *AuthorizationTransactionBadRequest {

	return &AuthorizationTransactionBadRequest{}
}

// WriteResponse to the client
func (o *AuthorizationTransactionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AuthorizationTransactionUnauthorizedCode is the HTTP code returned for type AuthorizationTransactionUnauthorized
const AuthorizationTransactionUnauthorizedCode int = 401

/*AuthorizationTransactionUnauthorized Unauthorized

swagger:response authorizationTransactionUnauthorized
*/
type AuthorizationTransactionUnauthorized struct {
}

// NewAuthorizationTransactionUnauthorized creates AuthorizationTransactionUnauthorized with default headers values
func NewAuthorizationTransactionUnauthorized() *AuthorizationTransactionUnauthorized {

	return &AuthorizationTransactionUnauthorized{}
}

// WriteResponse to the client
func (o *AuthorizationTransactionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// AuthorizationTransactionNotFoundCode is the HTTP code returned for type AuthorizationTransactionNotFound
const AuthorizationTransactionNotFoundCode int = 404

/*AuthorizationTransactionNotFound Not Found

swagger:response authorizationTransactionNotFound
*/
type AuthorizationTransactionNotFound struct {
}

// NewAuthorizationTransactionNotFound creates AuthorizationTransactionNotFound with default headers values
func NewAuthorizationTransactionNotFound() *AuthorizationTransactionNotFound {

	return &AuthorizationTransactionNotFound{}
}

// WriteResponse to the client
func (o *AuthorizationTransactionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*AuthorizationTransactionDefault generic error response

swagger:response authorizationTransactionDefault
*/
type AuthorizationTransactionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAuthorizationTransactionDefault creates AuthorizationTransactionDefault with default headers values
func NewAuthorizationTransactionDefault(code int) *AuthorizationTransactionDefault {
	if code <= 0 {
		code = 500
	}

	return &AuthorizationTransactionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the authorization transaction default response
func (o *AuthorizationTransactionDefault) WithStatusCode(code int) *AuthorizationTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the authorization transaction default response
func (o *AuthorizationTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the authorization transaction default response
func (o *AuthorizationTransactionDefault) WithPayload(payload *models.Error) *AuthorizationTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the authorization transaction default response
func (o *AuthorizationTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AuthorizationTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
