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

// POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID CreatePaymetricTokenWithTransIdRequest
//
// swagger:model POST_CreatePaymetricTokenWithTransIdRequest_v1_token_paymetric_transactionId
type POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID struct {

	// The lane ID.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// The Vault ID with Paymetric.
	// Required: true
	VaultID *string `json:"vaultId"`
}

// Validate validates this p o s t create paymetric token with trans Id request v1 token paymetric transaction Id
func (m *POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLaneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID) validateVaultID(formats strfmt.Registry) error {

	if err := validate.Required("vaultId", "body", m.VaultID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p o s t create paymetric token with trans Id request v1 token paymetric transaction Id based on context it is used
func (m *POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID) UnmarshalBinary(b []byte) error {
	var res POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
