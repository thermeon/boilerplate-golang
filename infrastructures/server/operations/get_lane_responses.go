// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/thermeon/boilerplate-golang/infrastructures/server/models"
)

// GetLaneOKCode is the HTTP code returned for type GetLaneOK
const GetLaneOKCode int = 200

/*GetLaneOK Lane retrieved

swagger:response getLaneOK
*/
type GetLaneOK struct {

	/*
	  In: Body
	*/
	Payload *models.Lane `json:"body,omitempty"`
}

// NewGetLaneOK creates GetLaneOK with default headers values
func NewGetLaneOK() *GetLaneOK {

	return &GetLaneOK{}
}

// WithPayload adds the payload to the get lane o k response
func (o *GetLaneOK) WithPayload(payload *models.Lane) *GetLaneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lane o k response
func (o *GetLaneOK) SetPayload(payload *models.Lane) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLaneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetLaneBadRequestCode is the HTTP code returned for type GetLaneBadRequest
const GetLaneBadRequestCode int = 400

/*GetLaneBadRequest Bad Request

swagger:response getLaneBadRequest
*/
type GetLaneBadRequest struct {
}

// NewGetLaneBadRequest creates GetLaneBadRequest with default headers values
func NewGetLaneBadRequest() *GetLaneBadRequest {

	return &GetLaneBadRequest{}
}

// WriteResponse to the client
func (o *GetLaneBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetLaneUnauthorizedCode is the HTTP code returned for type GetLaneUnauthorized
const GetLaneUnauthorizedCode int = 401

/*GetLaneUnauthorized Unauthorized

swagger:response getLaneUnauthorized
*/
type GetLaneUnauthorized struct {
}

// NewGetLaneUnauthorized creates GetLaneUnauthorized with default headers values
func NewGetLaneUnauthorized() *GetLaneUnauthorized {

	return &GetLaneUnauthorized{}
}

// WriteResponse to the client
func (o *GetLaneUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetLaneNotFoundCode is the HTTP code returned for type GetLaneNotFound
const GetLaneNotFoundCode int = 404

/*GetLaneNotFound Not Found

swagger:response getLaneNotFound
*/
type GetLaneNotFound struct {
}

// NewGetLaneNotFound creates GetLaneNotFound with default headers values
func NewGetLaneNotFound() *GetLaneNotFound {

	return &GetLaneNotFound{}
}

// WriteResponse to the client
func (o *GetLaneNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*GetLaneDefault generic error response

swagger:response getLaneDefault
*/
type GetLaneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLaneDefault creates GetLaneDefault with default headers values
func NewGetLaneDefault(code int) *GetLaneDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLaneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get lane default response
func (o *GetLaneDefault) WithStatusCode(code int) *GetLaneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get lane default response
func (o *GetLaneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get lane default response
func (o *GetLaneDefault) WithPayload(payload *models.Error) *GetLaneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get lane default response
func (o *GetLaneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLaneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}