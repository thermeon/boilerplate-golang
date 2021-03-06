// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinkDto link dto
//
// swagger:model LinkDto
type LinkDto struct {

	// href
	Href string `json:"href,omitempty"`

	// method
	Method string `json:"method,omitempty"`

	// rel
	Rel string `json:"rel,omitempty"`
}

// Validate validates this link dto
func (m *LinkDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this link dto based on context it is used
func (m *LinkDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinkDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkDto) UnmarshalBinary(b []byte) error {
	var res LinkDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
