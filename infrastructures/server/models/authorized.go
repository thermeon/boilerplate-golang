// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// Authorized Authorization from the bank approved
// Example: true
//
// swagger:model Authorized
type Authorized bool

// Validate validates this authorized
func (m Authorized) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this authorized based on context it is used
func (m Authorized) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
