// Code generated by go-swagger; DO NOT EDIT.

package refund

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// Postv1refundtokenReader is a Reader for the Postv1refundtoken structure.
type Postv1refundtokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1refundtokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1refundtokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostv1refundtokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1refundtokenOK creates a Postv1refundtokenOK with default headers values
func NewPostv1refundtokenOK() *Postv1refundtokenOK {
	return &Postv1refundtokenOK{}
}

/* Postv1refundtokenOK describes a response with status code 200, with default header values.

Success
*/
type Postv1refundtokenOK struct {
	Payload *models.TokenRefundResponse
}

func (o *Postv1refundtokenOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/refund/token][%d] postv1refundtokenOK  %+v", 200, o.Payload)
}
func (o *Postv1refundtokenOK) GetPayload() *models.TokenRefundResponse {
	return o.Payload
}

func (o *Postv1refundtokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenRefundResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1refundtokenUnauthorized creates a Postv1refundtokenUnauthorized with default headers values
func NewPostv1refundtokenUnauthorized() *Postv1refundtokenUnauthorized {
	return &Postv1refundtokenUnauthorized{}
}

/* Postv1refundtokenUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Postv1refundtokenUnauthorized struct {
}

func (o *Postv1refundtokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/refund/token][%d] postv1refundtokenUnauthorized ", 401)
}

func (o *Postv1refundtokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
