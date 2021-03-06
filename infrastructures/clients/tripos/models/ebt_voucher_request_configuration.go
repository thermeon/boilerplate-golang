// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EbtVoucherRequestConfiguration EbtVoucherRequestConfiguration
//
// swagger:model EbtVoucherRequestConfiguration
type EbtVoucherRequestConfiguration struct {

	// If set to true, enables duplicate checking logic for the transaction at the host.
	CheckForDuplicateTransactions bool `json:"checkForDuplicateTransactions,omitempty"`

	// The market code of the transaction.
	MarketCode string `json:"marketCode,omitempty"`
}

// Validate validates this ebt voucher request configuration
func (m *EbtVoucherRequestConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ebt voucher request configuration based on context it is used
func (m *EbtVoucherRequestConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EbtVoucherRequestConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EbtVoucherRequestConfiguration) UnmarshalBinary(b []byte) error {
	var res EbtVoucherRequestConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
