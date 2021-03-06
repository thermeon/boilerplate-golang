// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TipOptions TipOptions
//
// swagger:model TipOptions
type TipOptions struct {

	// Other option to show to select tip.
	OtherOption string `json:"otherOption,omitempty"`

	// CSV tip amount selections.
	TipSelections string `json:"tipSelections,omitempty"`

	// Type of tip entry: Amount selection, Percentage selection, or Prompt for custom value entry. Note: The Prompt value overrides the otherOption and tipSelections, skipping directly to Tip Amount entry screen.
	Type string `json:"type,omitempty"`
}

// Validate validates this tip options
func (m *TipOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tip options based on context it is used
func (m *TipOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TipOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TipOptions) UnmarshalBinary(b []byte) error {
	var res TipOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
