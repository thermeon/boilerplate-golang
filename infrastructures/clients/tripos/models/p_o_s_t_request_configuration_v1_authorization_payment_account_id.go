// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// POSTRequestConfigurationV1AuthorizationPaymentAccountID RequestConfiguration
//
// swagger:model POST_RequestConfiguration_v1_authorization_paymentAccountId
type POSTRequestConfigurationV1AuthorizationPaymentAccountID struct {

	// If set to true, partial approvals are allowed
	AllowPartialApprovals bool `json:"allowPartialApprovals,omitempty"`

	// If set to true, enables duplicate checking logic for the transaction at the host.
	CheckForDuplicateTransactions bool `json:"checkForDuplicateTransactions,omitempty"`

	// The market code of the transaction.
	MarketCode string `json:"marketCode,omitempty"`
}

// Validate validates this p o s t request configuration v1 authorization payment account Id
func (m *POSTRequestConfigurationV1AuthorizationPaymentAccountID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p o s t request configuration v1 authorization payment account Id based on context it is used
func (m *POSTRequestConfigurationV1AuthorizationPaymentAccountID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTRequestConfigurationV1AuthorizationPaymentAccountID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTRequestConfigurationV1AuthorizationPaymentAccountID) UnmarshalBinary(b []byte) error {
	var res POSTRequestConfigurationV1AuthorizationPaymentAccountID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
