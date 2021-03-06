// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FinancialCardRequest FinancialCardRequest
//
// swagger:model FinancialCardRequest
type FinancialCardRequest struct {

	// The flag that invokes manual entry directly.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`
}

// Validate validates this financial card request
func (m *FinancialCardRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this financial card request based on context it is used
func (m *FinancialCardRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FinancialCardRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FinancialCardRequest) UnmarshalBinary(b []byte) error {
	var res FinancialCardRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
