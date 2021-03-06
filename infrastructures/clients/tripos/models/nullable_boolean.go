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

// NullableBoolean Nullable`Boolean`
//
// swagger:model Nullable`Boolean`
type NullableBoolean struct {

	// has value
	// Required: true
	HasValue *bool `json:"hasValue"`

	// value
	// Required: true
	Value *bool `json:"value"`
}

// Validate validates this nullable boolean
func (m *NullableBoolean) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NullableBoolean) validateHasValue(formats strfmt.Registry) error {

	if err := validate.Required("hasValue", "body", m.HasValue); err != nil {
		return err
	}

	return nil
}

func (m *NullableBoolean) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nullable boolean based on context it is used
func (m *NullableBoolean) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NullableBoolean) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NullableBoolean) UnmarshalBinary(b []byte) error {
	var res NullableBoolean
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
