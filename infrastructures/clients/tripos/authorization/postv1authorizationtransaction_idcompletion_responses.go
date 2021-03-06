// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Postv1authorizationtransactionIdcompletionReader is a Reader for the Postv1authorizationtransactionIdcompletion structure.
type Postv1authorizationtransactionIdcompletionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1authorizationtransactionIdcompletionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1authorizationtransactionIdcompletionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostv1authorizationtransactionIdcompletionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostv1authorizationtransactionIdcompletionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1authorizationtransactionIdcompletionOK creates a Postv1authorizationtransactionIdcompletionOK with default headers values
func NewPostv1authorizationtransactionIdcompletionOK() *Postv1authorizationtransactionIdcompletionOK {
	return &Postv1authorizationtransactionIdcompletionOK{}
}

/* Postv1authorizationtransactionIdcompletionOK describes a response with status code 200, with default header values.

Success
*/
type Postv1authorizationtransactionIdcompletionOK struct {
	Payload *models.AuthorizationCompletionResponse
}

func (o *Postv1authorizationtransactionIdcompletionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/authorization/{transactionId}/completion][%d] postv1authorizationtransactionIdcompletionOK  %+v", 200, o.Payload)
}
func (o *Postv1authorizationtransactionIdcompletionOK) GetPayload() *models.AuthorizationCompletionResponse {
	return o.Payload
}

func (o *Postv1authorizationtransactionIdcompletionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationCompletionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1authorizationtransactionIdcompletionBadRequest creates a Postv1authorizationtransactionIdcompletionBadRequest with default headers values
func NewPostv1authorizationtransactionIdcompletionBadRequest() *Postv1authorizationtransactionIdcompletionBadRequest {
	return &Postv1authorizationtransactionIdcompletionBadRequest{}
}

/* Postv1authorizationtransactionIdcompletionBadRequest describes a response with status code 400, with default header values.

1) Returned if MarketCode is not HotelLodging but the request includes lodging parameters on a valid lodging request. 2) Returned if the lodging checkInDate or checkOutDate is not in the ISO 8601 format of YYYY-MM-DD on a valid lodging request. 3) Returned if the lodging roomAmount is less than 0.00 on a valid lodging request. 4) Returned if the lodging duration is less than 0 on a valid lodging request.
*/
type Postv1authorizationtransactionIdcompletionBadRequest struct {
}

func (o *Postv1authorizationtransactionIdcompletionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/authorization/{transactionId}/completion][%d] postv1authorizationtransactionIdcompletionBadRequest ", 400)
}

func (o *Postv1authorizationtransactionIdcompletionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostv1authorizationtransactionIdcompletionUnauthorized creates a Postv1authorizationtransactionIdcompletionUnauthorized with default headers values
func NewPostv1authorizationtransactionIdcompletionUnauthorized() *Postv1authorizationtransactionIdcompletionUnauthorized {
	return &Postv1authorizationtransactionIdcompletionUnauthorized{}
}

/* Postv1authorizationtransactionIdcompletionUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Postv1authorizationtransactionIdcompletionUnauthorized struct {
}

func (o *Postv1authorizationtransactionIdcompletionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/authorization/{transactionId}/completion][%d] postv1authorizationtransactionIdcompletionUnauthorized ", 401)
}

func (o *Postv1authorizationtransactionIdcompletionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
