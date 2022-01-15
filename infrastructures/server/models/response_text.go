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

// ResponseText Response code description
// Example: APPROVED
//
// swagger:model ResponseText
type ResponseText string

// Validate validates this response text
func (m ResponseText) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MaxLength("", "body", string(m), 64); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this response text based on context it is used
func (m ResponseText) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
