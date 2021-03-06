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

// POSTNullableDecimalV1AuthorizationPaymentAccountID Nullable`Decimal`
//
// swagger:model POST_Nullable_Decimal_v1_authorization_paymentAccountId
type POSTNullableDecimalV1AuthorizationPaymentAccountID struct {

	// has value
	// Required: true
	HasValue *bool `json:"hasValue"`

	// value
	// Required: true
	Value *float64 `json:"value"`
}

// Validate validates this p o s t nullable decimal v1 authorization payment account Id
func (m *POSTNullableDecimalV1AuthorizationPaymentAccountID) Validate(formats strfmt.Registry) error {
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

func (m *POSTNullableDecimalV1AuthorizationPaymentAccountID) validateHasValue(formats strfmt.Registry) error {

	if err := validate.Required("hasValue", "body", m.HasValue); err != nil {
		return err
	}

	return nil
}

func (m *POSTNullableDecimalV1AuthorizationPaymentAccountID) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p o s t nullable decimal v1 authorization payment account Id based on context it is used
func (m *POSTNullableDecimalV1AuthorizationPaymentAccountID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTNullableDecimalV1AuthorizationPaymentAccountID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTNullableDecimalV1AuthorizationPaymentAccountID) UnmarshalBinary(b []byte) error {
	var res POSTNullableDecimalV1AuthorizationPaymentAccountID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
