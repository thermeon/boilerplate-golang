// Code generated by go-swagger; DO NOT EDIT.

package input

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Getv1inputlaneIDReader is a Reader for the Getv1inputlaneID structure.
type Getv1inputlaneIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Getv1inputlaneIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetv1inputlaneIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetv1inputlaneIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetv1inputlaneIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetv1inputlaneIDOK creates a Getv1inputlaneIDOK with default headers values
func NewGetv1inputlaneIDOK() *Getv1inputlaneIDOK {
	return &Getv1inputlaneIDOK{}
}

/* Getv1inputlaneIDOK describes a response with status code 200, with default header values.

Success
*/
type Getv1inputlaneIDOK struct {
	Payload *models.InputResponse
}

func (o *Getv1inputlaneIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/input/{laneId}][%d] getv1inputlaneIdOK  %+v", 200, o.Payload)
}
func (o *Getv1inputlaneIDOK) GetPayload() *models.InputResponse {
	return o.Payload
}

func (o *Getv1inputlaneIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InputResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetv1inputlaneIDBadRequest creates a Getv1inputlaneIDBadRequest with default headers values
func NewGetv1inputlaneIDBadRequest() *Getv1inputlaneIDBadRequest {
	return &Getv1inputlaneIDBadRequest{}
}

/* Getv1inputlaneIDBadRequest describes a response with status code 400, with default header values.

Returned if parameter values are missing or invalid.
*/
type Getv1inputlaneIDBadRequest struct {
}

func (o *Getv1inputlaneIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/input/{laneId}][%d] getv1inputlaneIdBadRequest ", 400)
}

func (o *Getv1inputlaneIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetv1inputlaneIDInternalServerError creates a Getv1inputlaneIDInternalServerError with default headers values
func NewGetv1inputlaneIDInternalServerError() *Getv1inputlaneIDInternalServerError {
	return &Getv1inputlaneIDInternalServerError{}
}

/* Getv1inputlaneIDInternalServerError describes a response with status code 500, with default header values.

Returned if the cardholder allows the PIN pad to time out before submitting input.
*/
type Getv1inputlaneIDInternalServerError struct {
}

func (o *Getv1inputlaneIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/input/{laneId}][%d] getv1inputlaneIdInternalServerError ", 500)
}

func (o *Getv1inputlaneIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
