// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CCExpire Credit card expiry date, full-date notation as per RFC 3339
// Example: 2017-07-21
//
// swagger:model CCExpire
type CCExpire strfmt.Date

// UnmarshalJSON sets a CCExpire value from JSON input
func (m *CCExpire) UnmarshalJSON(b []byte) error {
	return ((*strfmt.Date)(m)).UnmarshalJSON(b)
}

// MarshalJSON retrieves a CCExpire value as JSON output
func (m CCExpire) MarshalJSON() ([]byte, error) {
	return (strfmt.Date(m)).MarshalJSON()
}

// Validate validates this c c expire
func (m CCExpire) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.FormatOf("", "body", "date", strfmt.Date(m).String(), formats); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this c c expire based on context it is used
func (m CCExpire) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CCExpire) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CCExpire) UnmarshalBinary(b []byte) error {
	var res CCExpire
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
