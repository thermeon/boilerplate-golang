// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Postv1tokenomniReader is a Reader for the Postv1tokenomni structure.
type Postv1tokenomniReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1tokenomniReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1tokenomniOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostv1tokenomniUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostv1tokenomniInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1tokenomniOK creates a Postv1tokenomniOK with default headers values
func NewPostv1tokenomniOK() *Postv1tokenomniOK {
	return &Postv1tokenomniOK{}
}

/* Postv1tokenomniOK describes a response with status code 200, with default header values.

Success
*/
type Postv1tokenomniOK struct {
	Payload *models.CreateOmniTokenResponse
}

func (o *Postv1tokenomniOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/token/omni][%d] postv1tokenomniOK  %+v", 200, o.Payload)
}
func (o *Postv1tokenomniOK) GetPayload() *models.CreateOmniTokenResponse {
	return o.Payload
}

func (o *Postv1tokenomniOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateOmniTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1tokenomniUnauthorized creates a Postv1tokenomniUnauthorized with default headers values
func NewPostv1tokenomniUnauthorized() *Postv1tokenomniUnauthorized {
	return &Postv1tokenomniUnauthorized{}
}

/* Postv1tokenomniUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Postv1tokenomniUnauthorized struct {
}

func (o *Postv1tokenomniUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/token/omni][%d] postv1tokenomniUnauthorized ", 401)
}

func (o *Postv1tokenomniUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostv1tokenomniInternalServerError creates a Postv1tokenomniInternalServerError with default headers values
func NewPostv1tokenomniInternalServerError() *Postv1tokenomniInternalServerError {
	return &Postv1tokenomniInternalServerError{}
}

/* Postv1tokenomniInternalServerError describes a response with status code 500, with default header values.

The token could not be created with the card used.
*/
type Postv1tokenomniInternalServerError struct {
}

func (o *Postv1tokenomniInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/token/omni][%d] postv1tokenomniInternalServerError ", 500)
}

func (o *Postv1tokenomniInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
