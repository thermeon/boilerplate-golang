// Code generated by go-swagger; DO NOT EDIT.

package payment_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Postv1paymentAccountReader is a Reader for the Postv1paymentAccount structure.
type Postv1paymentAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1paymentAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1paymentAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostv1paymentAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1paymentAccountOK creates a Postv1paymentAccountOK with default headers values
func NewPostv1paymentAccountOK() *Postv1paymentAccountOK {
	return &Postv1paymentAccountOK{}
}

/* Postv1paymentAccountOK describes a response with status code 200, with default header values.

Success
*/
type Postv1paymentAccountOK struct {
	Payload *models.CreatePaymentAccountCreditResponse
}

func (o *Postv1paymentAccountOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/paymentAccount][%d] postv1paymentAccountOK  %+v", 200, o.Payload)
}
func (o *Postv1paymentAccountOK) GetPayload() *models.CreatePaymentAccountCreditResponse {
	return o.Payload
}

func (o *Postv1paymentAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatePaymentAccountCreditResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1paymentAccountUnauthorized creates a Postv1paymentAccountUnauthorized with default headers values
func NewPostv1paymentAccountUnauthorized() *Postv1paymentAccountUnauthorized {
	return &Postv1paymentAccountUnauthorized{}
}

/* Postv1paymentAccountUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Postv1paymentAccountUnauthorized struct {
}

func (o *Postv1paymentAccountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/paymentAccount][%d] postv1paymentAccountUnauthorized ", 401)
}

func (o *Postv1paymentAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}