// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NonFinancialCardRequest NonFinancialCardRequest
//
// swagger:model NonFinancialCardRequest
type NonFinancialCardRequest struct {

	// The flag that invokes manual entry directly.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`
}

// Validate validates this non financial card request
func (m *NonFinancialCardRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this non financial card request based on context it is used
func (m *NonFinancialCardRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NonFinancialCardRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NonFinancialCardRequest) UnmarshalBinary(b []byte) error {
	var res NonFinancialCardRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
