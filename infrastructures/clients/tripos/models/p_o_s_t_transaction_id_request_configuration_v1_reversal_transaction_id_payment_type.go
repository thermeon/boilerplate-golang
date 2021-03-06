// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// POSTTransactionIDRequestConfigurationV1ReversalTransactionIDPaymentType TransactionIdRequestConfiguration
//
// swagger:model POST_TransactionIdRequestConfiguration_v1_reversal_transactionId_paymentType
type POSTTransactionIDRequestConfigurationV1ReversalTransactionIDPaymentType struct {

	// The market code of the transaction.
	MarketCode string `json:"marketCode,omitempty"`
}

// Validate validates this p o s t transaction Id request configuration v1 reversal transaction Id payment type
func (m *POSTTransactionIDRequestConfigurationV1ReversalTransactionIDPaymentType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p o s t transaction Id request configuration v1 reversal transaction Id payment type based on context it is used
func (m *POSTTransactionIDRequestConfigurationV1ReversalTransactionIDPaymentType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTTransactionIDRequestConfigurationV1ReversalTransactionIDPaymentType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTTransactionIDRequestConfigurationV1ReversalTransactionIDPaymentType) UnmarshalBinary(b []byte) error {
	var res POSTTransactionIDRequestConfigurationV1ReversalTransactionIDPaymentType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
