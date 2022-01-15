// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Amount A monetary amount
// Example: 12.34
//
// swagger:model Amount
type Amount float32

// Validate validates this amount
func (m Amount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Minimum("", "body", float64(m), 0, true); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this amount based on context it is used
func (m Amount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
