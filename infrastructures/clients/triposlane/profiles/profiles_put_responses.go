// Code generated by go-swagger; DO NOT EDIT.

package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/triposlane/models"
)

// ProfilesPutReader is a Reader for the ProfilesPut structure.
type ProfilesPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilesPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProfilesPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProfilesPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProfilesPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewProfilesPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProfilesPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewProfilesPutOK creates a ProfilesPutOK with default headers values
func NewProfilesPutOK() *ProfilesPutOK {
	return &ProfilesPutOK{}
}

/* ProfilesPutOK describes a response with status code 200, with default header values.

Valid request.
*/
type ProfilesPutOK struct {
}

func (o *ProfilesPutOK) Error() string {
	return fmt.Sprintf("[PUT /v1/lanes/{laneId}/profiles/idle][%d] profilesPutOK ", 200)
}

func (o *ProfilesPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProfilesPutBadRequest creates a ProfilesPutBadRequest with default headers values
func NewProfilesPutBadRequest() *ProfilesPutBadRequest {
	return &ProfilesPutBadRequest{}
}

/* ProfilesPutBadRequest describes a response with status code 400, with default header values.

Invalid fields on request.
*/
type ProfilesPutBadRequest struct {
	Payload *models.ExceptionDto
}

func (o *ProfilesPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/lanes/{laneId}/profiles/idle][%d] profilesPutBadRequest  %+v", 400, o.Payload)
}
func (o *ProfilesPutBadRequest) GetPayload() *models.ExceptionDto {
	return o.Payload
}

func (o *ProfilesPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExceptionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfilesPutUnauthorized creates a ProfilesPutUnauthorized with default headers values
func NewProfilesPutUnauthorized() *ProfilesPutUnauthorized {
	return &ProfilesPutUnauthorized{}
}

/* ProfilesPutUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, and AccountToken provided in the HTTP request header are invalid.
*/
type ProfilesPutUnauthorized struct {
	Payload *models.ExceptionDto
}

func (o *ProfilesPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/lanes/{laneId}/profiles/idle][%d] profilesPutUnauthorized  %+v", 401, o.Payload)
}
func (o *ProfilesPutUnauthorized) GetPayload() *models.ExceptionDto {
	return o.Payload
}

func (o *ProfilesPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExceptionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfilesPutNotFound creates a ProfilesPutNotFound with default headers values
func NewProfilesPutNotFound() *ProfilesPutNotFound {
	return &ProfilesPutNotFound{}
}

/* ProfilesPutNotFound describes a response with status code 404, with default header values.

Lane not found.
*/
type ProfilesPutNotFound struct {
	Payload *models.ExceptionDto
}

func (o *ProfilesPutNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/lanes/{laneId}/profiles/idle][%d] profilesPutNotFound  %+v", 404, o.Payload)
}
func (o *ProfilesPutNotFound) GetPayload() *models.ExceptionDto {
	return o.Payload
}

func (o *ProfilesPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExceptionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfilesPutInternalServerError creates a ProfilesPutInternalServerError with default headers values
func NewProfilesPutInternalServerError() *ProfilesPutInternalServerError {
	return &ProfilesPutInternalServerError{}
}

/* ProfilesPutInternalServerError describes a response with status code 500, with default header values.

Returned if an error occurred processing the request.
*/
type ProfilesPutInternalServerError struct {
	Payload *models.ExceptionDto
}

func (o *ProfilesPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/lanes/{laneId}/profiles/idle][%d] profilesPutInternalServerError  %+v", 500, o.Payload)
}
func (o *ProfilesPutInternalServerError) GetPayload() *models.ExceptionDto {
	return o.Payload
}

func (o *ProfilesPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExceptionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
