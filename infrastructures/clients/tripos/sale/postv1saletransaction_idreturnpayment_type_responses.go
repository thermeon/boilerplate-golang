// Code generated by go-swagger; DO NOT EDIT.

package sale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Postv1saletransactionIdreturnpaymentTypeReader is a Reader for the Postv1saletransactionIdreturnpaymentType structure.
type Postv1saletransactionIdreturnpaymentTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1saletransactionIdreturnpaymentTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1saletransactionIdreturnpaymentTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostv1saletransactionIdreturnpaymentTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1saletransactionIdreturnpaymentTypeOK creates a Postv1saletransactionIdreturnpaymentTypeOK with default headers values
func NewPostv1saletransactionIdreturnpaymentTypeOK() *Postv1saletransactionIdreturnpaymentTypeOK {
	return &Postv1saletransactionIdreturnpaymentTypeOK{}
}

/* Postv1saletransactionIdreturnpaymentTypeOK describes a response with status code 200, with default header values.

Success
*/
type Postv1saletransactionIdreturnpaymentTypeOK struct {
	Payload *models.ReturnResponse
}

func (o *Postv1saletransactionIdreturnpaymentTypeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/sale/{transactionId}/return/{paymentType}][%d] postv1saletransactionIdreturnpaymentTypeOK  %+v", 200, o.Payload)
}
func (o *Postv1saletransactionIdreturnpaymentTypeOK) GetPayload() *models.ReturnResponse {
	return o.Payload
}

func (o *Postv1saletransactionIdreturnpaymentTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReturnResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1saletransactionIdreturnpaymentTypeUnauthorized creates a Postv1saletransactionIdreturnpaymentTypeUnauthorized with default headers values
func NewPostv1saletransactionIdreturnpaymentTypeUnauthorized() *Postv1saletransactionIdreturnpaymentTypeUnauthorized {
	return &Postv1saletransactionIdreturnpaymentTypeUnauthorized{}
}

/* Postv1saletransactionIdreturnpaymentTypeUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Postv1saletransactionIdreturnpaymentTypeUnauthorized struct {
}

func (o *Postv1saletransactionIdreturnpaymentTypeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/sale/{transactionId}/return/{paymentType}][%d] postv1saletransactionIdreturnpaymentTypeUnauthorized ", 401)
}

func (o *Postv1saletransactionIdreturnpaymentTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
