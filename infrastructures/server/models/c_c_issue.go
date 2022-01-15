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

// CCIssue Credit card issue number
// Example: 1
//
// swagger:model CCIssue
type CCIssue int64

// Validate validates this c c issue
func (m CCIssue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.MinimumInt("", "body", int64(m), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("", "body", int64(m), 99, false); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this c c issue based on context it is used
func (m CCIssue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}