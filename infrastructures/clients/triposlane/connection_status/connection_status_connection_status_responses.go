// Code generated by go-swagger; DO NOT EDIT.

package connection_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/triposlane/models"
)

// ConnectionStatusConnectionStatusReader is a Reader for the ConnectionStatusConnectionStatus structure.
type ConnectionStatusConnectionStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectionStatusConnectionStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectionStatusConnectionStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewConnectionStatusConnectionStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewConnectionStatusConnectionStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewConnectionStatusConnectionStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewConnectionStatusConnectionStatusServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectionStatusConnectionStatusOK creates a ConnectionStatusConnectionStatusOK with default headers values
func NewConnectionStatusConnectionStatusOK() *ConnectionStatusConnectionStatusOK {
	return &ConnectionStatusConnectionStatusOK{}
}

/* ConnectionStatusConnectionStatusOK describes a response with status code 200, with default header values.

Valid request.
*/
type ConnectionStatusConnectionStatusOK struct {
	Payload *models.ConnectionStatusDto
}

func (o *ConnectionStatusConnectionStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/lanes/{laneId}/connectionstatus][%d] connectionStatusConnectionStatusOK  %+v", 200, o.Payload)
}
func (o *ConnectionStatusConnectionStatusOK) GetPayload() *models.ConnectionStatusDto {
	return o.Payload
}

func (o *ConnectionStatusConnectionStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConnectionStatusDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionStatusConnectionStatusBadRequest creates a ConnectionStatusConnectionStatusBadRequest with default headers values
func NewConnectionStatusConnectionStatusBadRequest() *ConnectionStatusConnectionStatusBadRequest {
	return &ConnectionStatusConnectionStatusBadRequest{}
}

/* ConnectionStatusConnectionStatusBadRequest describes a response with status code 400, with default header values.

Invalid fields on request.
*/
type ConnectionStatusConnectionStatusBadRequest struct {
	Payload *models.ExceptionDto
}

func (o *ConnectionStatusConnectionStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/lanes/{laneId}/connectionstatus][%d] connectionStatusConnectionStatusBadRequest  %+v", 400, o.Payload)
}
func (o *ConnectionStatusConnectionStatusBadRequest) GetPayload() *models.ExceptionDto {
	return o.Payload
}

func (o *ConnectionStatusConnectionStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExceptionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionStatusConnectionStatusUnauthorized creates a ConnectionStatusConnectionStatusUnauthorized with default headers values
func NewConnectionStatusConnectionStatusUnauthorized() *ConnectionStatusConnectionStatusUnauthorized {
	return &ConnectionStatusConnectionStatusUnauthorized{}
}

/* ConnectionStatusConnectionStatusUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, and AccountToken provided in the HTTP request header are invalid.
*/
type ConnectionStatusConnectionStatusUnauthorized struct {
	Payload *models.ExceptionDto
}

func (o *ConnectionStatusConnectionStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/lanes/{laneId}/connectionstatus][%d] connectionStatusConnectionStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *ConnectionStatusConnectionStatusUnauthorized) GetPayload() *models.ExceptionDto {
	return o.Payload
}

func (o *ConnectionStatusConnectionStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExceptionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionStatusConnectionStatusNotFound creates a ConnectionStatusConnectionStatusNotFound with default headers values
func NewConnectionStatusConnectionStatusNotFound() *ConnectionStatusConnectionStatusNotFound {
	return &ConnectionStatusConnectionStatusNotFound{}
}

/* ConnectionStatusConnectionStatusNotFound describes a response with status code 404, with default header values.

Lane not found.
*/
type ConnectionStatusConnectionStatusNotFound struct {
	Payload *models.ExceptionDto
}

func (o *ConnectionStatusConnectionStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/lanes/{laneId}/connectionstatus][%d] connectionStatusConnectionStatusNotFound  %+v", 404, o.Payload)
}
func (o *ConnectionStatusConnectionStatusNotFound) GetPayload() *models.ExceptionDto {
	return o.Payload
}

func (o *ConnectionStatusConnectionStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExceptionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionStatusConnectionStatusServiceUnavailable creates a ConnectionStatusConnectionStatusServiceUnavailable with default headers values
func NewConnectionStatusConnectionStatusServiceUnavailable() *ConnectionStatusConnectionStatusServiceUnavailable {
	return &ConnectionStatusConnectionStatusServiceUnavailable{}
}

/* ConnectionStatusConnectionStatusServiceUnavailable describes a response with status code 503, with default header values.

Returned if a resource is not availabe.
*/
type ConnectionStatusConnectionStatusServiceUnavailable struct {
	Payload *models.ExceptionDto
}

func (o *ConnectionStatusConnectionStatusServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/lanes/{laneId}/connectionstatus][%d] connectionStatusConnectionStatusServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ConnectionStatusConnectionStatusServiceUnavailable) GetPayload() *models.ExceptionDto {
	return o.Payload
}

func (o *ConnectionStatusConnectionStatusServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExceptionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}