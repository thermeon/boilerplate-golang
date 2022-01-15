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

// POSTNullableBooleanV1GiftClose Nullable`Boolean`
//
// swagger:model POST_Nullable_Boolean_v1_gift_close
type POSTNullableBooleanV1GiftClose struct {

	// has value
	// Required: true
	HasValue *bool `json:"hasValue"`

	// value
	// Required: true
	Value *bool `json:"value"`
}

// Validate validates this p o s t nullable boolean v1 gift close
func (m *POSTNullableBooleanV1GiftClose) Validate(formats strfmt.Registry) error {
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

func (m *POSTNullableBooleanV1GiftClose) validateHasValue(formats strfmt.Registry) error {

	if err := validate.Required("hasValue", "body", m.HasValue); err != nil {
		return err
	}

	return nil
}

func (m *POSTNullableBooleanV1GiftClose) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p o s t nullable boolean v1 gift close based on context it is used
func (m *POSTNullableBooleanV1GiftClose) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTNullableBooleanV1GiftClose) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTNullableBooleanV1GiftClose) UnmarshalBinary(b []byte) error {
	var res POSTNullableBooleanV1GiftClose
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
