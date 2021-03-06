// Code generated by go-swagger; DO NOT EDIT.

package void

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Postv1voidtransactionIDReader is a Reader for the Postv1voidtransactionID structure.
type Postv1voidtransactionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1voidtransactionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1voidtransactionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostv1voidtransactionIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1voidtransactionIDOK creates a Postv1voidtransactionIDOK with default headers values
func NewPostv1voidtransactionIDOK() *Postv1voidtransactionIDOK {
	return &Postv1voidtransactionIDOK{}
}

/* Postv1voidtransactionIDOK describes a response with status code 200, with default header values.

Success
*/
type Postv1voidtransactionIDOK struct {
	Payload *models.VoidResponse
}

func (o *Postv1voidtransactionIDOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/void/{transactionId}][%d] postv1voidtransactionIdOK  %+v", 200, o.Payload)
}
func (o *Postv1voidtransactionIDOK) GetPayload() *models.VoidResponse {
	return o.Payload
}

func (o *Postv1voidtransactionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoidResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1voidtransactionIDUnauthorized creates a Postv1voidtransactionIDUnauthorized with default headers values
func NewPostv1voidtransactionIDUnauthorized() *Postv1voidtransactionIDUnauthorized {
	return &Postv1voidtransactionIDUnauthorized{}
}

/* Postv1voidtransactionIDUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Postv1voidtransactionIDUnauthorized struct {
}

func (o *Postv1voidtransactionIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/void/{transactionId}][%d] postv1voidtransactionIdUnauthorized ", 401)
}

func (o *Postv1voidtransactionIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
