// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Signature Signature
//
// swagger:model Signature
type Signature struct {

	// The byte array of the signature in the format specified by Format.
	Data []strfmt.Base64 `json:"data"`

	// The format of the signature.
	Format string `json:"format,omitempty"`

	// Indicates why a signature is or is not present.
	StatusCode string `json:"statusCode,omitempty"`
}

// Validate validates this signature
func (m *Signature) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this signature based on context it is used
func (m *Signature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Signature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Signature) UnmarshalBinary(b []byte) error {
	var res Signature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
