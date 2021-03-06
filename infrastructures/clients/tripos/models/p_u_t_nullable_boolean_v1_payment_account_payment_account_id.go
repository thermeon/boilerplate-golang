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

// PUTNullableBooleanV1PaymentAccountPaymentAccountID Nullable`Boolean`
//
// swagger:model PUT_Nullable_Boolean_v1_paymentAccount_paymentAccountId
type PUTNullableBooleanV1PaymentAccountPaymentAccountID struct {

	// has value
	// Required: true
	HasValue *bool `json:"hasValue"`

	// value
	// Required: true
	Value *bool `json:"value"`
}

// Validate validates this p u t nullable boolean v1 payment account payment account Id
func (m *PUTNullableBooleanV1PaymentAccountPaymentAccountID) Validate(formats strfmt.Registry) error {
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

func (m *PUTNullableBooleanV1PaymentAccountPaymentAccountID) validateHasValue(formats strfmt.Registry) error {

	if err := validate.Required("hasValue", "body", m.HasValue); err != nil {
		return err
	}

	return nil
}

func (m *PUTNullableBooleanV1PaymentAccountPaymentAccountID) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p u t nullable boolean v1 payment account payment account Id based on context it is used
func (m *PUTNullableBooleanV1PaymentAccountPaymentAccountID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PUTNullableBooleanV1PaymentAccountPaymentAccountID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PUTNullableBooleanV1PaymentAccountPaymentAccountID) UnmarshalBinary(b []byte) error {
	var res PUTNullableBooleanV1PaymentAccountPaymentAccountID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
