// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// POSTRequestConfigurationV1AuthorizationToken RequestConfiguration
//
// swagger:model POST_RequestConfiguration_v1_authorization_token
type POSTRequestConfigurationV1AuthorizationToken struct {

	// If set to true, partial approvals are allowed
	AllowPartialApprovals bool `json:"allowPartialApprovals,omitempty"`

	// If set to true, enables duplicate checking logic for the transaction at the host.
	CheckForDuplicateTransactions bool `json:"checkForDuplicateTransactions,omitempty"`

	// The market code of the transaction.
	MarketCode string `json:"marketCode,omitempty"`
}

// Validate validates this p o s t request configuration v1 authorization token
func (m *POSTRequestConfigurationV1AuthorizationToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p o s t request configuration v1 authorization token based on context it is used
func (m *POSTRequestConfigurationV1AuthorizationToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTRequestConfigurationV1AuthorizationToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTRequestConfigurationV1AuthorizationToken) UnmarshalBinary(b []byte) error {
	var res POSTRequestConfigurationV1AuthorizationToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}