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

// POSTCreateOmniTokenRequestV1TokenOmni CreateOmniTokenRequest
//
// swagger:model POST_CreateOmniTokenRequest_v1_token_omni
type POSTCreateOmniTokenRequestV1TokenOmni struct {

	// Invokes manual card entry.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`

	// Invokes prompt for cardholder to enter card security code for manual keyed card entry.
	IsCscSupported string `json:"isCscSupported,omitempty"`

	// Specifies which lane to use.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// The Vault ID with Omni.
	VaultID string `json:"vaultId,omitempty"`
}

// Validate validates this p o s t create omni token request v1 token omni
func (m *POSTCreateOmniTokenRequestV1TokenOmni) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLaneID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *POSTCreateOmniTokenRequestV1TokenOmni) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p o s t create omni token request v1 token omni based on context it is used
func (m *POSTCreateOmniTokenRequestV1TokenOmni) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTCreateOmniTokenRequestV1TokenOmni) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTCreateOmniTokenRequestV1TokenOmni) UnmarshalBinary(b []byte) error {
	var res POSTCreateOmniTokenRequestV1TokenOmni
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
