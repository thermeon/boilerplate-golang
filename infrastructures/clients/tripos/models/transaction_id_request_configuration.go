// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransactionIDRequestConfiguration TransactionIdRequestConfiguration
//
// swagger:model TransactionIdRequestConfiguration
type TransactionIDRequestConfiguration struct {

	// The market code of the transaction.
	MarketCode string `json:"marketCode,omitempty"`
}

// Validate validates this transaction Id request configuration
func (m *TransactionIDRequestConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this transaction Id request configuration based on context it is used
func (m *TransactionIDRequestConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransactionIDRequestConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionIDRequestConfiguration) UnmarshalBinary(b []byte) error {
	var res TransactionIDRequestConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
