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

// Getv1cardfinancialReader is a Reader for the Getv1cardfinancial structure.
type Getv1cardfinancialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Getv1cardfinancialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetv1cardfinancialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetv1cardfinancialOK creates a Getv1cardfinancialOK with default headers values
func NewGetv1cardfinancialOK() *Getv1cardfinancialOK {
	return &Getv1cardfinancialOK{}
}

/* Getv1cardfinancialOK describes a response with status code 200, with default header values.

Success
*/
type Getv1cardfinancialOK struct {
	Payload *models.FinancialCardResponse
}

func (o *Getv1cardfinancialOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/card/financial][%d] getv1cardfinancialOK  %+v", 200, o.Payload)
}
func (o *Getv1cardfinancialOK) GetPayload() *models.FinancialCardResponse {
	return o.Payload
}

func (o *Getv1cardfinancialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FinancialCardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}