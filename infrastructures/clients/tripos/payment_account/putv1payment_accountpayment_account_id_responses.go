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

// Putv1paymentAccountpaymentAccountIDReader is a Reader for the Putv1paymentAccountpaymentAccountID structure.
type Putv1paymentAccountpaymentAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Putv1paymentAccountpaymentAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutv1paymentAccountpaymentAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutv1paymentAccountpaymentAccountIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutv1paymentAccountpaymentAccountIDOK creates a Putv1paymentAccountpaymentAccountIDOK with default headers values
func NewPutv1paymentAccountpaymentAccountIDOK() *Putv1paymentAccountpaymentAccountIDOK {
	return &Putv1paymentAccountpaymentAccountIDOK{}
}

/* Putv1paymentAccountpaymentAccountIDOK describes a response with status code 200, with default header values.

Success
*/
type Putv1paymentAccountpaymentAccountIDOK struct {
	Payload *models.UpdatePaymentAccountCreditResponse
}

func (o *Putv1paymentAccountpaymentAccountIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/paymentAccount/{paymentAccountId}][%d] putv1paymentAccountpaymentAccountIdOK  %+v", 200, o.Payload)
}
func (o *Putv1paymentAccountpaymentAccountIDOK) GetPayload() *models.UpdatePaymentAccountCreditResponse {
	return o.Payload
}

func (o *Putv1paymentAccountpaymentAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdatePaymentAccountCreditResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutv1paymentAccountpaymentAccountIDUnauthorized creates a Putv1paymentAccountpaymentAccountIDUnauthorized with default headers values
func NewPutv1paymentAccountpaymentAccountIDUnauthorized() *Putv1paymentAccountpaymentAccountIDUnauthorized {
	return &Putv1paymentAccountpaymentAccountIDUnauthorized{}
}

/* Putv1paymentAccountpaymentAccountIDUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Putv1paymentAccountpaymentAccountIDUnauthorized struct {
}

func (o *Putv1paymentAccountpaymentAccountIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/paymentAccount/{paymentAccountId}][%d] putv1paymentAccountpaymentAccountIdUnauthorized ", 401)
}

func (o *Putv1paymentAccountpaymentAccountIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
