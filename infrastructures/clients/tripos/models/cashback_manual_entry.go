// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CashbackManualEntry CashbackManualEntry
//
// swagger:model CashbackManualEntry
type CashbackManualEntry struct {

	// The PIN pad will only accept a cashback amount in the provided increment (i.e. Set to 5 to only allow cashback amounts in increments of $5.00).
	AmountIncrement int32 `json:"amountIncrement,omitempty"`

	// The PIN pad will only accept a cashback amount that does not exceed the provided maximum.
	MaximumAmount int32 `json:"maximumAmount,omitempty"`
}

// Validate validates this cashback manual entry
func (m *CashbackManualEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cashback manual entry based on context it is used
func (m *CashbackManualEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CashbackManualEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CashbackManualEntry) UnmarshalBinary(b []byte) error {
	var res CashbackManualEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
