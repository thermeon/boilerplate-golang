// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AuthorizationCompletionResponse AuthorizationCompletionResponse
//
// swagger:model AuthorizationCompletionResponse
type AuthorizationCompletionResponse struct {

	// A list of errors that occurred.
	Errors []*APIError `json:"_errors"`

	// Indicates if there are errors.
	HasErrors bool `json:"_hasErrors,omitempty"`

	// A list of resource links
	Links []*APILink `json:"_links"`

	// A list of log entries detailing what happened during the request. Ideally only used during development or troubleshooting as this can be quite verbose.
	Logs []string `json:"_logs"`

	// Response information from the processor.
	Processor *Processor `json:"_processor,omitempty"`

	// The type of object held in the result.
	Type string `json:"_type,omitempty"`

	// A list of warnings that occurred.
	Warnings []*APIWarning `json:"_warnings"`

	// The card account number.
	AccountNumber string `json:"accountNumber,omitempty"`

	// The account type selected and used in the transaction.
	AccountType string `json:"accountType,omitempty"`

	// Approval number from the processor. Depending on card type and processor an approval number might not be returned.
	ApprovalNumber string `json:"approvalNumber,omitempty"`

	// The amount approved by the processor. This is the actual amount that will be charged or credited.
	ApprovedAmount float64 `json:"approvedAmount,omitempty"`

	// The balance of the gift card.
	BalanceAmount float64 `json:"balanceAmount,omitempty"`

	// The balance currency code.
	BalanceCurrencyCode string `json:"balanceCurrencyCode,omitempty"`

	// The BIN entry that matched the account number.
	BinValue string `json:"binValue,omitempty"`

	// The cardholder name.
	CardHolderName string `json:"cardHolderName,omitempty"`

	// The card logo. Possible values are: Visa, Mastercard, Discover, Amex, Diners Club, JCB, Carte Blanche, Other.
	CardLogo string `json:"cardLogo,omitempty"`

	// The convenience fee added to the transaction
	ConvenienceFeeAmount float64 `json:"convenienceFeeAmount,omitempty"`

	// The country code used in the transaction.
	CountryCode string `json:"countryCode,omitempty"`

	// The currency code used in the transaction.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// The type of the EBT card
	EbtType string `json:"ebtType,omitempty"`

	// The fields used on the receipt for an EMV transaction. Null if the transaction was not EMV.
	Emv *Emv `json:"emv,omitempty"`

	// Description of how card was entered.
	EntryMode string `json:"entryMode,omitempty"`

	// The card's expiration month
	ExpirationMonth string `json:"expirationMonth,omitempty"`

	// The card's expiration year
	ExpirationYear string `json:"expirationYear,omitempty"`

	// Indicates whether the card used was a FSA card.<br />Note: Maybe = No BIN entry to determine if FSA.
	FsaCard string `json:"fsaCard,omitempty"`

	// Set to true if the host approved the transaction.
	IsApproved bool `json:"isApproved,omitempty"`

	// Indicates whether the EMV card was still inserted into the payment device when the transaction completed.
	IsCardInserted bool `json:"isCardInserted,omitempty"`

	// A boolean value indicating whether triPOS is disconnected from the host.
	IsOffline bool `json:"isOffline,omitempty"`

	// The language used in the transaction.
	Language string `json:"language,omitempty"`

	// The Merchant used to process the transaction.
	MerchantID string `json:"merchantId,omitempty"`

	// Label that shows the network where the transaction was routed for authorization
	NetworkLabel string `json:"networkLabel,omitempty"`

	// Non Financial data.
	NonFinancialData *NonFinancialCard `json:"nonFinancialData,omitempty"`

	// Description of payment type utilized.
	PaymentType string `json:"paymentType,omitempty"`

	// True if the PIN was verified, false if not verified or undetermined.
	PinVerified bool `json:"pinVerified,omitempty"`

	// The message unique to QuickChip pre-read functionality only.
	QuickChipMessage string `json:"quickChipMessage,omitempty"`

	// The signature data.
	Signature *Signature `json:"signature,omitempty"`

	// The status code for the transaction.
	StatusCode string `json:"statusCode,omitempty"`

	// The original amount sent for the transaction.
	SubTotalAmount float64 `json:"subTotalAmount,omitempty"`

	// The ID of the terminal used during the transaction
	TerminalID string `json:"terminalId,omitempty"`

	// The tip amount added to the transaction.
	TipAmount float64 `json:"tipAmount,omitempty"`

	// The Token ID.
	TokenID string `json:"tokenId,omitempty"`

	// The Token Provider.
	TokenProvider string `json:"tokenProvider,omitempty"`

	// The total amount of the transaction.
	TotalAmount float64 `json:"totalAmount,omitempty"`

	// Transaction date/time in ISO8601 format
	TransactionDateTime string `json:"transactionDateTime,omitempty"`

	// The transaction ID from the processor
	TransactionID string `json:"transactionId,omitempty"`
}

// Validate validates this authorization completion response
func (m *AuthorizationCompletionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonFinancialData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationCompletionResponse) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthorizationCompletionResponse) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for i := 0; i < len(m.Links); i++ {
		if swag.IsZero(m.Links[i]) { // not required
			continue
		}

		if m.Links[i] != nil {
			if err := m.Links[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthorizationCompletionResponse) validateProcessor(formats strfmt.Registry) error {
	if swag.IsZero(m.Processor) { // not required
		return nil
	}

	if m.Processor != nil {
		if err := m.Processor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_processor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_processor")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationCompletionResponse) validateWarnings(formats strfmt.Registry) error {
	if swag.IsZero(m.Warnings) { // not required
		return nil
	}

	for i := 0; i < len(m.Warnings); i++ {
		if swag.IsZero(m.Warnings[i]) { // not required
			continue
		}

		if m.Warnings[i] != nil {
			if err := m.Warnings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthorizationCompletionResponse) validateEmv(formats strfmt.Registry) error {
	if swag.IsZero(m.Emv) { // not required
		return nil
	}

	if m.Emv != nil {
		if err := m.Emv.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emv")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emv")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationCompletionResponse) validateNonFinancialData(formats strfmt.Registry) error {
	if swag.IsZero(m.NonFinancialData) { // not required
		return nil
	}

	if m.NonFinancialData != nil {
		if err := m.NonFinancialData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nonFinancialData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nonFinancialData")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationCompletionResponse) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(m.Signature) { // not required
		return nil
	}

	if m.Signature != nil {
		if err := m.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this authorization completion response based on the context it is used
func (m *AuthorizationCompletionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProcessor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWarnings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNonFinancialData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationCompletionResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthorizationCompletionResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Links); i++ {

		if m.Links[i] != nil {
			if err := m.Links[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_links" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_links" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthorizationCompletionResponse) contextValidateProcessor(ctx context.Context, formats strfmt.Registry) error {

	if m.Processor != nil {
		if err := m.Processor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_processor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("_processor")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationCompletionResponse) contextValidateWarnings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Warnings); i++ {

		if m.Warnings[i] != nil {
			if err := m.Warnings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("_warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("_warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthorizationCompletionResponse) contextValidateEmv(ctx context.Context, formats strfmt.Registry) error {

	if m.Emv != nil {
		if err := m.Emv.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emv")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emv")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationCompletionResponse) contextValidateNonFinancialData(ctx context.Context, formats strfmt.Registry) error {

	if m.NonFinancialData != nil {
		if err := m.NonFinancialData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nonFinancialData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nonFinancialData")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationCompletionResponse) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if m.Signature != nil {
		if err := m.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizationCompletionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizationCompletionResponse) UnmarshalBinary(b []byte) error {
	var res AuthorizationCompletionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}