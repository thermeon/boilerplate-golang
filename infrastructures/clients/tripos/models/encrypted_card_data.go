// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EncryptedCardData EncryptedCardData
//
// swagger:model EncryptedCardData
type EncryptedCardData struct {

	// Encrypted card data key serial number.
	CardDataKeySerialNumber string `json:"cardDataKeySerialNumber,omitempty"`

	// The masked card number.
	CardNumberMasked string `json:"cardNumberMasked,omitempty"`

	// Card verification value found on the card.
	Cvv string `json:"cvv,omitempty"`

	// The format of the encrypted manual keyed card data.
	EncryptedFormat string `json:"encryptedFormat,omitempty"`

	// Encrypted manually keyed card data.
	EncryptedManualKeyedCardData string `json:"encryptedManualKeyedCardData,omitempty"`

	// Encrypted swiped track 1 data.
	EncryptedTrack1Data string `json:"encryptedTrack1Data,omitempty"`

	// Encrypted swiped track 2 data.
	EncryptedTrack2Data string `json:"encryptedTrack2Data,omitempty"`

	// PIN key serial number.
	KeySerialNumber string `json:"keySerialNumber,omitempty"`

	// Raw output of certain devices directly supported by Express.
	MagneprintData string `json:"magneprintData,omitempty"`

	// Encrypted PIN block information sent from the PIN encryption device.
	PinBlock string `json:"pinBlock,omitempty"`

	// Clear text swiped track 1 data.
	Track1Data string `json:"track1Data,omitempty"`

	// Clear text swiped track 2 data.
	Track2Data string `json:"track2Data,omitempty"`

	// Clear text swiped track 3 data.
	Track3Data string `json:"track3Data,omitempty"`
}

// Validate validates this encrypted card data
func (m *EncryptedCardData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this encrypted card data based on context it is used
func (m *EncryptedCardData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncryptedCardData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptedCardData) UnmarshalBinary(b []byte) error {
	var res EncryptedCardData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
