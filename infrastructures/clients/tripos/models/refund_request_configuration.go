// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RefundRequestConfiguration RefundRequestConfiguration
//
// swagger:model RefundRequestConfiguration
type RefundRequestConfiguration struct {

	// If set to true, partial approvals are allowed
	AllowPartialApprovals bool `json:"allowPartialApprovals,omitempty"`

	// If set to true, enables duplicate checking logic for the transaction at the host.
	CheckForDuplicateTransactions bool `json:"checkForDuplicateTransactions,omitempty"`

	// Use this parameter to allow Vantiv & Valutec gift cards to be used.
	IsGiftSupported string `json:"isGiftSupported,omitempty"`

	// Use this parameter to allow manual entry of card account numbers on the PIN pad.
	IsManualEntryAllowed string `json:"isManualEntryAllowed,omitempty"`

	// The market code of the transaction.
	MarketCode string `json:"marketCode,omitempty"`

	// Specifies how the signature prompt should be handled for the request. If a value is not provided, Always will be used.  See <a href='https://triposcert.vantiv.com/api/help/kb/signaturePrompt.html'>Signature Prompt</a>.
	PromptForSignature string `json:"promptForSignature,omitempty"`

	// Override provisional amount used for QC transactions, which is $1.00 by default
	ProvisionalAmount string `json:"provisionalAmount,omitempty"`

	// Specifies the user input timeout.
	UserInputTimeout string `json:"userInputTimeout,omitempty"`
}

// Validate validates this refund request configuration
func (m *RefundRequestConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this refund request configuration based on context it is used
func (m *RefundRequestConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RefundRequestConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RefundRequestConfiguration) UnmarshalBinary(b []byte) error {
	var res RefundRequestConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}