// Code generated by go-swagger; DO NOT EDIT.

package card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Getv1cardnonfinanciallaneIDReader is a Reader for the Getv1cardnonfinanciallaneID structure.
type Getv1cardnonfinanciallaneIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Getv1cardnonfinanciallaneIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetv1cardnonfinanciallaneIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetv1cardnonfinanciallaneIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetv1cardnonfinanciallaneIDOK creates a Getv1cardnonfinanciallaneIDOK with default headers values
func NewGetv1cardnonfinanciallaneIDOK() *Getv1cardnonfinanciallaneIDOK {
	return &Getv1cardnonfinanciallaneIDOK{}
}

/* Getv1cardnonfinanciallaneIDOK describes a response with status code 200, with default header values.

Success
*/
type Getv1cardnonfinanciallaneIDOK struct {
	Payload *models.NonFinancialCardResponse
}

func (o *Getv1cardnonfinanciallaneIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/card/nonfinancial/{laneId}][%d] getv1cardnonfinanciallaneIdOK  %+v", 200, o.Payload)
}
func (o *Getv1cardnonfinanciallaneIDOK) GetPayload() *models.NonFinancialCardResponse {
	return o.Payload
}

func (o *Getv1cardnonfinanciallaneIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NonFinancialCardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetv1cardnonfinanciallaneIDNotFound creates a Getv1cardnonfinanciallaneIDNotFound with default headers values
func NewGetv1cardnonfinanciallaneIDNotFound() *Getv1cardnonfinanciallaneIDNotFound {
	return &Getv1cardnonfinanciallaneIDNotFound{}
}

/* Getv1cardnonfinanciallaneIDNotFound describes a response with status code 404, with default header values.

If no 'non-financial' data can be found when the card is swiped, tripos returns an HTTP status code of 404.
*/
type Getv1cardnonfinanciallaneIDNotFound struct {
}

func (o *Getv1cardnonfinanciallaneIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/card/nonfinancial/{laneId}][%d] getv1cardnonfinanciallaneIdNotFound ", 404)
}

func (o *Getv1cardnonfinanciallaneIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
