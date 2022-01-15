// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// DeviceDescription EFTPOS device description
// Example: Main device for location AB123
//
// swagger:model DeviceDescription
type DeviceDescription string

// Validate validates this device description
func (m DeviceDescription) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this device description based on context it is used
func (m DeviceDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
