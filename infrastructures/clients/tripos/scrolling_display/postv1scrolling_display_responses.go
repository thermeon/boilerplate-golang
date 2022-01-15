// Code generated by go-swagger; DO NOT EDIT.

package scrolling_display

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Postv1scrollingDisplayReader is a Reader for the Postv1scrollingDisplay structure.
type Postv1scrollingDisplayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1scrollingDisplayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1scrollingDisplayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostv1scrollingDisplayBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1scrollingDisplayOK creates a Postv1scrollingDisplayOK with default headers values
func NewPostv1scrollingDisplayOK() *Postv1scrollingDisplayOK {
	return &Postv1scrollingDisplayOK{}
}

/* Postv1scrollingDisplayOK describes a response with status code 200, with default header values.

Success
*/
type Postv1scrollingDisplayOK struct {
	Payload *models.DisplayResponse
}

func (o *Postv1scrollingDisplayOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/scrollingDisplay][%d] postv1scrollingDisplayOK  %+v", 200, o.Payload)
}
func (o *Postv1scrollingDisplayOK) GetPayload() *models.DisplayResponse {
	return o.Payload
}

func (o *Postv1scrollingDisplayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DisplayResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1scrollingDisplayBadRequest creates a Postv1scrollingDisplayBadRequest with default headers values
func NewPostv1scrollingDisplayBadRequest() *Postv1scrollingDisplayBadRequest {
	return &Postv1scrollingDisplayBadRequest{}
}

/* Postv1scrollingDisplayBadRequest describes a response with status code 400, with default header values.

Returned if parameter values are missing or invalid.
*/
type Postv1scrollingDisplayBadRequest struct {
}

func (o *Postv1scrollingDisplayBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/scrollingDisplay][%d] postv1scrollingDisplayBadRequest ", 400)
}

func (o *Postv1scrollingDisplayBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
