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

// Postv1saleReader is a Reader for the Postv1sale structure.
type Postv1saleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Postv1saleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostv1saleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostv1saleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostv1saleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostv1saleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostv1saleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostv1saleOK creates a Postv1saleOK with default headers values
func NewPostv1saleOK() *Postv1saleOK {
	return &Postv1saleOK{}
}

/* Postv1saleOK describes a response with status code 200, with default header values.

Success
*/
type Postv1saleOK struct {
	Payload *models.SaleResponse
}

func (o *Postv1saleOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/sale][%d] postv1saleOK  %+v", 200, o.Payload)
}
func (o *Postv1saleOK) GetPayload() *models.SaleResponse {
	return o.Payload
}

func (o *Postv1saleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SaleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostv1saleBadRequest creates a Postv1saleBadRequest with default headers values
func NewPostv1saleBadRequest() *Postv1saleBadRequest {
	return &Postv1saleBadRequest{}
}

/* Postv1saleBadRequest describes a response with status code 400, with default header values.

1) Returned if action is forward and if the requestIdToForward is missing or is not a valid GUID. 2) Returned if client sends in all 5 of the healthcare values as positive values. 3) Returned if any healthcare field has a negative value. 4) Returned if healthcare Total field is less than sum of the other healthcare fields. 5) Returned if action is store and the customer swipes a healthcare card. 6) Returned if MarketCode is not HotelLodging but the request includes lodging parameters on a valid lodging request. 7) Returned if the lodging checkInDate or checkOutDate is not in the ISO 8601 format of YYYY-MM-DD on a valid lodging request. 8) Returned if the lodging roomAmount is less than 0.00 on a valid lodging request. 9) Returned if the lodging duration is less than 0 on a valid lodging request. 10) Returned if the cashbackOptions maximumAmount is either not provided, or is not an integer greater than zero. 11) Returned if the cashbackOptions amountIncrement is not an integer greater than zero. If it is not provided, it will default to 1.
*/
type Postv1saleBadRequest struct {
}

func (o *Postv1saleBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/sale][%d] postv1saleBadRequest ", 400)
}

func (o *Postv1saleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostv1saleUnauthorized creates a Postv1saleUnauthorized with default headers values
func NewPostv1saleUnauthorized() *Postv1saleUnauthorized {
	return &Postv1saleUnauthorized{}
}

/* Postv1saleUnauthorized describes a response with status code 401, with default header values.

Returned if the AcceptorId, AccountId, AccountToken and TerminalId (config file only) are not provided in the HTTP request header or in the config file.
*/
type Postv1saleUnauthorized struct {
}

func (o *Postv1saleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/sale][%d] postv1saleUnauthorized ", 401)
}

func (o *Postv1saleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostv1saleNotFound creates a Postv1saleNotFound with default headers values
func NewPostv1saleNotFound() *Postv1saleNotFound {
	return &Postv1saleNotFound{}
}

/* Postv1saleNotFound describes a response with status code 404, with default header values.

Returned if action is store or forward and if the given requestIdToForward cannot be found in the database.
*/
type Postv1saleNotFound struct {
}

func (o *Postv1saleNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/sale][%d] postv1saleNotFound ", 404)
}

func (o *Postv1saleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostv1saleInternalServerError creates a Postv1saleInternalServerError with default headers values
func NewPostv1saleInternalServerError() *Postv1saleInternalServerError {
	return &Postv1saleInternalServerError{}
}

/* Postv1saleInternalServerError describes a response with status code 500, with default header values.

1) Returned if action is store or forward and if an error occurs while saving the sale to the database. 2) Returned if action is store or forward and if the stored message cannot be deserialized into the original stored object. 3) Returned if action is store or forward and if the stored message content is blank. 4) Returned if action is store or forward and if triPOS is unable to pull the stored record from the database.
*/
type Postv1saleInternalServerError struct {
}

func (o *Postv1saleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/sale][%d] postv1saleInternalServerError ", 500)
}

func (o *Postv1saleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
