// Code generated by go-swagger; DO NOT EDIT.

package display

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Postv1displayReader is a Reader for the Postv1display structure.
type Postv1displayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1displayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1displayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostv1displayBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1displayOK creates a Postv1displayOK with default headers values
func NewPostv1displayOK() *Postv1displayOK {
	return &Postv1displayOK{}
}

/* Postv1displayOK describes a response with status code 200, with default header values.

Success
*/
type Postv1displayOK struct {
	Payload *models.DisplayResponse
}

func (o *Postv1displayOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/display][%d] postv1displayOK  %+v", 200, o.Payload)
}
func (o *Postv1displayOK) GetPayload() *models.DisplayResponse {
	return o.Payload
}

func (o *Postv1displayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DisplayResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1displayBadRequest creates a Postv1displayBadRequest with default headers values
func NewPostv1displayBadRequest() *Postv1displayBadRequest {
	return &Postv1displayBadRequest{}
}

/* Postv1displayBadRequest describes a response with status code 400, with default header values.

Returned if parameter values are missing or invalid.
*/
type Postv1displayBadRequest struct {
}

func (o *Postv1displayBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/display][%d] postv1displayBadRequest ", 400)
}

func (o *Postv1displayBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
