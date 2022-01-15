// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/thermeon/boilerplate-golang/infrastructures/server/models"
)

// VoidTransactionCreatedCode is the HTTP code returned for type VoidTransactionCreated
const VoidTransactionCreatedCode int = 201

/*VoidTransactionCreated successful operation

swagger:response voidTransactionCreated
*/
type VoidTransactionCreated struct {
	/*location of the created resource

	 */
	Location string `json:"location"`
}

// NewVoidTransactionCreated creates VoidTransactionCreated with default headers values
func NewVoidTransactionCreated() *VoidTransactionCreated {

	return &VoidTransactionCreated{}
}

// WithLocation adds the location to the void transaction created response
func (o *VoidTransactionCreated) WithLocation(location string) *VoidTransactionCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the void transaction created response
func (o *VoidTransactionCreated) SetLocation(location string) {
	o.Location = location
}

// WriteResponse to the client
func (o *VoidTransactionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header location

	location := o.Location
	if location != "" {
		rw.Header().Set("location", location)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// VoidTransactionBadRequestCode is the HTTP code returned for type VoidTransactionBadRequest
const VoidTransactionBadRequestCode int = 400

/*VoidTransactionBadRequest Bad Request

swagger:response voidTransactionBadRequest
*/
type VoidTransactionBadRequest struct {
}

// NewVoidTransactionBadRequest creates VoidTransactionBadRequest with default headers values
func NewVoidTransactionBadRequest() *VoidTransactionBadRequest {

	return &VoidTransactionBadRequest{}
}

// WriteResponse to the client
func (o *VoidTransactionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// VoidTransactionUnauthorizedCode is the HTTP code returned for type VoidTransactionUnauthorized
const VoidTransactionUnauthorizedCode int = 401

/*VoidTransactionUnauthorized Unauthorized

swagger:response voidTransactionUnauthorized
*/
type VoidTransactionUnauthorized struct {
}

// NewVoidTransactionUnauthorized creates VoidTransactionUnauthorized with default headers values
func NewVoidTransactionUnauthorized() *VoidTransactionUnauthorized {

	return &VoidTransactionUnauthorized{}
}

// WriteResponse to the client
func (o *VoidTransactionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// VoidTransactionNotFoundCode is the HTTP code returned for type VoidTransactionNotFound
const VoidTransactionNotFoundCode int = 404

/*VoidTransactionNotFound Not Found

swagger:response voidTransactionNotFound
*/
type VoidTransactionNotFound struct {
}

// NewVoidTransactionNotFound creates VoidTransactionNotFound with default headers values
func NewVoidTransactionNotFound() *VoidTransactionNotFound {

	return &VoidTransactionNotFound{}
}

// WriteResponse to the client
func (o *VoidTransactionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*VoidTransactionDefault generic error response

swagger:response voidTransactionDefault
*/
type VoidTransactionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewVoidTransactionDefault creates VoidTransactionDefault with default headers values
func NewVoidTransactionDefault(code int) *VoidTransactionDefault {
	if code <= 0 {
		code = 500
	}

	return &VoidTransactionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the void transaction default response
func (o *VoidTransactionDefault) WithStatusCode(code int) *VoidTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the void transaction default response
func (o *VoidTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the void transaction default response
func (o *VoidTransactionDefault) WithPayload(payload *models.Error) *VoidTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the void transaction default response
func (o *VoidTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VoidTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
