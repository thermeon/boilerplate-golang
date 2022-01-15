// Code generated by go-swagger; DO NOT EDIT.

package pin_pad_selection_prompt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Getv1pinPadSelectionPromptReader is a Reader for the Getv1pinPadSelectionPrompt structure.
type Getv1pinPadSelectionPromptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Getv1pinPadSelectionPromptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetv1pinPadSelectionPromptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetv1pinPadSelectionPromptInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetv1pinPadSelectionPromptOK creates a Getv1pinPadSelectionPromptOK with default headers values
func NewGetv1pinPadSelectionPromptOK() *Getv1pinPadSelectionPromptOK {
	return &Getv1pinPadSelectionPromptOK{}
}

/* Getv1pinPadSelectionPromptOK describes a response with status code 200, with default header values.

Success
*/
type Getv1pinPadSelectionPromptOK struct {
	Payload *models.PinPadSelectionResponse
}

func (o *Getv1pinPadSelectionPromptOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/pinPadSelectionPrompt][%d] getv1pinPadSelectionPromptOK  %+v", 200, o.Payload)
}
func (o *Getv1pinPadSelectionPromptOK) GetPayload() *models.PinPadSelectionResponse {
	return o.Payload
}

func (o *Getv1pinPadSelectionPromptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PinPadSelectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetv1pinPadSelectionPromptInternalServerError creates a Getv1pinPadSelectionPromptInternalServerError with default headers values
func NewGetv1pinPadSelectionPromptInternalServerError() *Getv1pinPadSelectionPromptInternalServerError {
	return &Getv1pinPadSelectionPromptInternalServerError{}
}

/* Getv1pinPadSelectionPromptInternalServerError describes a response with status code 500, with default header values.

Returned if the cardholder doesn't respond within the specified amount of time.
*/
type Getv1pinPadSelectionPromptInternalServerError struct {
}

func (o *Getv1pinPadSelectionPromptInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/pinPadSelectionPrompt][%d] getv1pinPadSelectionPromptInternalServerError ", 500)
}

func (o *Getv1pinPadSelectionPromptInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}