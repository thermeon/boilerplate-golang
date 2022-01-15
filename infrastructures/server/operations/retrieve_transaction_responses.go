// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/thermeon/boilerplate-golang/infrastructures/server/models"
)

// RetrieveTransactionOKCode is the HTTP code returned for type RetrieveTransactionOK
const RetrieveTransactionOKCode int = 200

/*RetrieveTransactionOK Transaction found and retrieved

swagger:response retrieveTransactionOK
*/
type RetrieveTransactionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Transaction `json:"body,omitempty"`
}

// NewRetrieveTransactionOK creates RetrieveTransactionOK with default headers values
func NewRetrieveTransactionOK() *RetrieveTransactionOK {

	return &RetrieveTransactionOK{}
}

// WithPayload adds the payload to the retrieve transaction o k response
func (o *RetrieveTransactionOK) WithPayload(payload *models.Transaction) *RetrieveTransactionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve transaction o k response
func (o *RetrieveTransactionOK) SetPayload(payload *models.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTransactionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RetrieveTransactionBadRequestCode is the HTTP code returned for type RetrieveTransactionBadRequest
const RetrieveTransactionBadRequestCode int = 400

/*RetrieveTransactionBadRequest Bad Request

swagger:response retrieveTransactionBadRequest
*/
type RetrieveTransactionBadRequest struct {
}

// NewRetrieveTransactionBadRequest creates RetrieveTransactionBadRequest with default headers values
func NewRetrieveTransactionBadRequest() *RetrieveTransactionBadRequest {

	return &RetrieveTransactionBadRequest{}
}

// WriteResponse to the client
func (o *RetrieveTransactionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// RetrieveTransactionUnauthorizedCode is the HTTP code returned for type RetrieveTransactionUnauthorized
const RetrieveTransactionUnauthorizedCode int = 401

/*RetrieveTransactionUnauthorized Unauthorized

swagger:response retrieveTransactionUnauthorized
*/
type RetrieveTransactionUnauthorized struct {
}

// NewRetrieveTransactionUnauthorized creates RetrieveTransactionUnauthorized with default headers values
func NewRetrieveTransactionUnauthorized() *RetrieveTransactionUnauthorized {

	return &RetrieveTransactionUnauthorized{}
}

// WriteResponse to the client
func (o *RetrieveTransactionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// RetrieveTransactionNotFoundCode is the HTTP code returned for type RetrieveTransactionNotFound
const RetrieveTransactionNotFoundCode int = 404

/*RetrieveTransactionNotFound Not Found

swagger:response retrieveTransactionNotFound
*/
type RetrieveTransactionNotFound struct {
}

// NewRetrieveTransactionNotFound creates RetrieveTransactionNotFound with default headers values
func NewRetrieveTransactionNotFound() *RetrieveTransactionNotFound {

	return &RetrieveTransactionNotFound{}
}

// WriteResponse to the client
func (o *RetrieveTransactionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*RetrieveTransactionDefault generic error response

swagger:response retrieveTransactionDefault
*/
type RetrieveTransactionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewRetrieveTransactionDefault creates RetrieveTransactionDefault with default headers values
func NewRetrieveTransactionDefault(code int) *RetrieveTransactionDefault {
	if code <= 0 {
		code = 500
	}

	return &RetrieveTransactionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the retrieve transaction default response
func (o *RetrieveTransactionDefault) WithStatusCode(code int) *RetrieveTransactionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the retrieve transaction default response
func (o *RetrieveTransactionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the retrieve transaction default response
func (o *RetrieveTransactionDefault) WithPayload(payload *models.Error) *RetrieveTransactionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the retrieve transaction default response
func (o *RetrieveTransactionDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RetrieveTransactionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
