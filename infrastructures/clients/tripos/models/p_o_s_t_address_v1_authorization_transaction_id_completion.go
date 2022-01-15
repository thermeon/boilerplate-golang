// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// POSTAddressV1AuthorizationTransactionIDCompletion Address
//
// swagger:model POST_Address_v1_authorization_transactionId_completion
type POSTAddressV1AuthorizationTransactionIDCompletion struct {

	// The street address used for billing purposes.
	BillingAddress1 string `json:"billingAddress1,omitempty"`

	// The street address used for billing purposes.
	BillingAddress2 string `json:"billingAddress2,omitempty"`

	// The name of the city used for billing purposes.
	BillingCity string `json:"billingCity,omitempty"`

	// The e-mail address used for billing purposes.
	BillingEmail string `json:"billingEmail,omitempty"`

	// The name used for billing purposes.
	BillingName string `json:"billingName,omitempty"`

	// The phone number used for billing purposes. The recommended format is (800)555-1212.
	BillingPhone string `json:"billingPhone,omitempty"`

	// The postal code used for billing purposes.
	BillingPostalCode string `json:"billingPostalCode,omitempty"`

	// The name of the state used for billing purposes. This value may be any 2 character state code or the full state name.
	BillingState string `json:"billingState,omitempty"`

	// The street address used for shipping purposes.
	ShippingAddress1 string `json:"shippingAddress1,omitempty"`

	// The street address used for shipping purposes.
	ShippingAddress2 string `json:"shippingAddress2,omitempty"`

	// The name of the city used for shipping purposes.
	ShippingCity string `json:"shippingCity,omitempty"`

	// The e-mail address used for shipping purposes.
	ShippingEmail string `json:"shippingEmail,omitempty"`

	// The name used for shipping purposes.
	ShippingName string `json:"shippingName,omitempty"`

	// The phone number used for shipping purposes. The recommended format is (800)555-1212
	ShippingPhone string `json:"shippingPhone,omitempty"`

	// The postal code used for shipping purposes.
	ShippingPostalCode string `json:"shippingPostalCode,omitempty"`

	// The name of the state used for shipping purposes. This value may be any 2 character state code or the full state name.
	ShippingState string `json:"shippingState,omitempty"`
}

// Validate validates this p o s t address v1 authorization transaction Id completion
func (m *POSTAddressV1AuthorizationTransactionIDCompletion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p o s t address v1 authorization transaction Id completion based on context it is used
func (m *POSTAddressV1AuthorizationTransactionIDCompletion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTAddressV1AuthorizationTransactionIDCompletion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTAddressV1AuthorizationTransactionIDCompletion) UnmarshalBinary(b []byte) error {
	var res POSTAddressV1AuthorizationTransactionIDCompletion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
