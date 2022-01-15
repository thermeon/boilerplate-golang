// Code generated by go-swagger; DO NOT EDIT.

package selection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Getv1selectionlaneIDReader is a Reader for the Getv1selectionlaneID structure.
type Getv1selectionlaneIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Getv1selectionlaneIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetv1selectionlaneIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetv1selectionlaneIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetv1selectionlaneIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetv1selectionlaneIDOK creates a Getv1selectionlaneIDOK with default headers values
func NewGetv1selectionlaneIDOK() *Getv1selectionlaneIDOK {
	return &Getv1selectionlaneIDOK{}
}

/* Getv1selectionlaneIDOK describes a response with status code 200, with default header values.

Success
*/
type Getv1selectionlaneIDOK struct {
	Payload *models.SelectionResponse
}

func (o *Getv1selectionlaneIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/selection/{laneId}][%d] getv1selectionlaneIdOK  %+v", 200, o.Payload)
}
func (o *Getv1selectionlaneIDOK) GetPayload() *models.SelectionResponse {
	return o.Payload
}

func (o *Getv1selectionlaneIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SelectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetv1selectionlaneIDBadRequest creates a Getv1selectionlaneIDBadRequest with default headers values
func NewGetv1selectionlaneIDBadRequest() *Getv1selectionlaneIDBadRequest {
	return &Getv1selectionlaneIDBadRequest{}
}

/* Getv1selectionlaneIDBadRequest describes a response with status code 400, with default header values.

Returned if the 'form' parameter is empty or invalid.
*/
type Getv1selectionlaneIDBadRequest struct {
}

func (o *Getv1selectionlaneIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/selection/{laneId}][%d] getv1selectionlaneIdBadRequest ", 400)
}

func (o *Getv1selectionlaneIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetv1selectionlaneIDInternalServerError creates a Getv1selectionlaneIDInternalServerError with default headers values
func NewGetv1selectionlaneIDInternalServerError() *Getv1selectionlaneIDInternalServerError {
	return &Getv1selectionlaneIDInternalServerError{}
}

/* Getv1selectionlaneIDInternalServerError describes a response with status code 500, with default header values.

Returned if the cardholder doesn't respond within the specified amount of time.
*/
type Getv1selectionlaneIDInternalServerError struct {
}

func (o *Getv1selectionlaneIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/selection/{laneId}][%d] getv1selectionlaneIdInternalServerError ", 500)
}

func (o *Getv1selectionlaneIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}