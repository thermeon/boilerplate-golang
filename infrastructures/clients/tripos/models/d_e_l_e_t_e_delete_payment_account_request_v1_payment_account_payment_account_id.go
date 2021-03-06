// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DELETEDeletePaymentAccountRequestV1PaymentAccountPaymentAccountID DeletePaymentAccountRequest
//
// swagger:model DELETE_DeletePaymentAccountRequest_v1_paymentAccount_paymentAccountId
type DELETEDeletePaymentAccountRequestV1PaymentAccountPaymentAccountID struct {

	// Defines whether the cardholder is present at the transaction. This value is optional, but recommended to be set. If this value is not set in the request, it will be automatically set based on the market code.
	CardHolderPresentCode string `json:"cardHolderPresentCode,omitempty"`

	// An optional clerk number for reference.
	ClerkNumber string `json:"clerkNumber,omitempty"`

	// Specifies which lane to use.
	LaneID int32 `json:"laneId,omitempty"`

	// A user defined reference number.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// An optional shift id for reference.
	ShiftID string `json:"shiftId,omitempty"`

	// An optional ticket number.
	TicketNumber string `json:"ticketNumber,omitempty"`
}

// Validate validates this d e l e t e delete payment account request v1 payment account payment account Id
func (m *DELETEDeletePaymentAccountRequestV1PaymentAccountPaymentAccountID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this d e l e t e delete payment account request v1 payment account payment account Id based on context it is used
func (m *DELETEDeletePaymentAccountRequestV1PaymentAccountPaymentAccountID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DELETEDeletePaymentAccountRequestV1PaymentAccountPaymentAccountID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DELETEDeletePaymentAccountRequestV1PaymentAccountPaymentAccountID) UnmarshalBinary(b []byte) error {
	var res DELETEDeletePaymentAccountRequestV1PaymentAccountPaymentAccountID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
