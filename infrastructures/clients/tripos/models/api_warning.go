// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIWarning ApiWarning
//
// swagger:model ApiWarning
type APIWarning struct {

	// A warning message targeted at the developer of the integrated business application.
	DeveloperMessage string `json:"developerMessage,omitempty"`

	// A warning message targeted at the end user of the integrated business application.
	UserMessage string `json:"userMessage,omitempty"`
}

// Validate validates this Api warning
func (m *APIWarning) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api warning based on context it is used
func (m *APIWarning) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIWarning) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIWarning) UnmarshalBinary(b []byte) error {
	var res APIWarning
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
