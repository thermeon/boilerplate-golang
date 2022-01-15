// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// POSTAuthorizationCompletionRequestConfigurationV1AuthorizationTransactionIDCompletion AuthorizationCompletionRequestConfiguration
//
// swagger:model POST_AuthorizationCompletionRequestConfiguration_v1_authorization_transactionId_completion
type POSTAuthorizationCompletionRequestConfigurationV1AuthorizationTransactionIDCompletion struct {

	// If set to true, partial approvals are allowed
	AllowPartialApprovals bool `json:"allowPartialApprovals,omitempty"`

	// If set to true, enables duplicate checking logic for the transaction at the host.
	CheckForDuplicateTransactions bool `json:"checkForDuplicateTransactions,omitempty"`

	// Use this parameter to allow manual entry of card account numbers on the PIN pad.
	IsManualEntryAllowed string `json:"isManualEntryAllowed,omitempty"`

	// The market code of the transaction.
	MarketCode string `json:"marketCode,omitempty"`

	// Specifies how the signature prompt should be handled for the request. If a value is not provided, Never will be used. See <a href='https://triposcert.vantiv.com/api/help/kb/signaturePrompt.html'>Signature Prompt</a>.
	PromptForSignature string `json:"promptForSignature,omitempty"`

	// Override provisional amount used for QC transactions, which is $1.00 by default
	ProvisionalAmount string `json:"provisionalAmount,omitempty"`

	// Specifies the user input timeout.
	UserInputTimeout string `json:"userInputTimeout,omitempty"`
}

// Validate validates this p o s t authorization completion request configuration v1 authorization transaction Id completion
func (m *POSTAuthorizationCompletionRequestConfigurationV1AuthorizationTransactionIDCompletion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p o s t authorization completion request configuration v1 authorization transaction Id completion based on context it is used
func (m *POSTAuthorizationCompletionRequestConfigurationV1AuthorizationTransactionIDCompletion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTAuthorizationCompletionRequestConfigurationV1AuthorizationTransactionIDCompletion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTAuthorizationCompletionRequestConfigurationV1AuthorizationTransactionIDCompletion) UnmarshalBinary(b []byte) error {
	var res POSTAuthorizationCompletionRequestConfigurationV1AuthorizationTransactionIDCompletion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}