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

// PaymentAccountAuthorizationResponse PaymentAccountAuthorizationResponse
//
// swagger:model PaymentAccountAuthorizationResponse
type PaymentAccountAuthorizationResponse struct {

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

	// Approval number from the processor. Depending on card type and processor an approval number might not be returned.
	ApprovalNumber string `json:"approvalNumber,omitempty"`

	// The amount approved by the processor. This is the actual amount that will be charged or credited.
	ApprovedAmount float64 `json:"approvedAmount,omitempty"`

	// The AVS response code.
	AvsResponseCode string `json:"avsResponseCode,omitempty"`

	// The card logo. Possible values are: Visa, Mastercard, Discover, Amex, Diners Club, JCB, Carte Blanche, Other.
	CardLogo string `json:"cardLogo,omitempty"`

	// The convenience fee added to the transaction
	ConvenienceFeeAmount float64 `json:"convenienceFeeAmount,omitempty"`

	// Set to true if the host approved the transaction.
	IsApproved bool `json:"isApproved,omitempty"`

	// A boolean value indicating whether triPOS is disconnected from the host.
	IsOffline bool `json:"isOffline,omitempty"`

	// The Merchant used to process the transaction.
	MerchantID string `json:"merchantId,omitempty"`

	// The payment account ID.
	PaymentAccountID string `json:"paymentAccountId,omitempty"`

	// The payment account reference number.
	PaymentAccountReferenceNumber string `json:"paymentAccountReferenceNumber,omitempty"`

	// The reference number.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// The status code for the transaction.
	StatusCode string `json:"statusCode,omitempty"`

	// The original amount sent for the transaction.
	SubTotalAmount float64 `json:"subTotalAmount,omitempty"`

	// The ID of the terminal used during the transaction
	TerminalID string `json:"terminalId,omitempty"`

	// The tip amount added to the transaction.
	TipAmount float64 `json:"tipAmount,omitempty"`

	// The total amount of the transaction.
	TotalAmount float64 `json:"totalAmount,omitempty"`

	// Transaction date/time in ISO8601 format
	TransactionDateTime string `json:"transactionDateTime,omitempty"`

	// The transaction ID from the processor
	TransactionID string `json:"transactionId,omitempty"`
}

// Validate validates this payment account authorization response
func (m *PaymentAccountAuthorizationResponse) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAccountAuthorizationResponse) validateErrors(formats strfmt.Registry) error {
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

func (m *PaymentAccountAuthorizationResponse) validateLinks(formats strfmt.Registry) error {
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

func (m *PaymentAccountAuthorizationResponse) validateProcessor(formats strfmt.Registry) error {
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

func (m *PaymentAccountAuthorizationResponse) validateWarnings(formats strfmt.Registry) error {
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

// ContextValidate validate this payment account authorization response based on the context it is used
func (m *PaymentAccountAuthorizationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAccountAuthorizationResponse) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PaymentAccountAuthorizationResponse) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PaymentAccountAuthorizationResponse) contextValidateProcessor(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PaymentAccountAuthorizationResponse) contextValidateWarnings(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PaymentAccountAuthorizationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAccountAuthorizationResponse) UnmarshalBinary(b []byte) error {
	var res PaymentAccountAuthorizationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
