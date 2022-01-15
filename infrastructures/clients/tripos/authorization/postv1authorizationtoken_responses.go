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

// Postv1authorizationtokenReader is a Reader for the Postv1authorizationtoken structure.
type Postv1authorizationtokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1authorizationtokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1authorizationtokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostv1authorizationtokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1authorizationtokenOK creates a Postv1authorizationtokenOK with default headers values
func NewPostv1authorizationtokenOK() *Postv1authorizationtokenOK {
	return &Postv1authorizationtokenOK{}
}

/* Postv1authorizationtokenOK describes a response with status code 200, with default header values.

Success
*/
type Postv1authorizationtokenOK struct {
	Payload *models.TokenAuthorizationResponse
}

func (o *Postv1authorizationtokenOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/authorization/token][%d] postv1authorizationtokenOK  %+v", 200, o.Payload)
}
func (o *Postv1authorizationtokenOK) GetPayload() *models.TokenAuthorizationResponse {
	return o.Payload
}

func (o *Postv1authorizationtokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenAuthorizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1authorizationtokenUnauthorized creates a Postv1authorizationtokenUnauthorized with default headers values
func NewPostv1authorizationtokenUnauthorized() *Postv1authorizationtokenUnauthorized {
	return &Postv1authorizationtokenUnauthorized{}
}

/* Postv1authorizationtokenUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Postv1authorizationtokenUnauthorized struct {
}

func (o *Postv1authorizationtokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/authorization/token][%d] postv1authorizationtokenUnauthorized ", 401)
}

func (o *Postv1authorizationtokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}