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

// ReversalRequest ReversalRequest
//
// swagger:model ReversalRequest
type ReversalRequest struct {

	// Defines whether the cardholder is present at the transaction. This value is optional, but recommended to be set. If this value is not set in the request, it will be automatically set based on the market code.
	CardHolderPresentCode string `json:"cardHolderPresentCode,omitempty"`

	// An optional clerk number for reference.
	ClerkNumber string `json:"clerkNumber,omitempty"`

	// Any value included in this section will override the corresponding value set in the triPOS.config
	Configuration *POSTTransactionIDRequestConfigurationV1ReversalTransactionIDPaymentType `json:"configuration,omitempty"`

	// The convenience fee amount.
	ConvenienceFeeAmount float64 `json:"convenienceFeeAmount,omitempty"`

	// EBT card type of the original transaction. This is required when payment type is EBT.
	EbtType string `json:"ebtType,omitempty"`

	// To retrieve Token in the transaction.
	GetToken string `json:"getToken,omitempty"`

	// This field is for Valutec gift/loyalty card program. <br>'01' = Original Gift Card Program <br>'02' = Promotional Gift Card Program <br>'03' = Original Combo Card Program <br>'04' = Auto Rewards LPR Card Program <br>'05' = Original Loyalty Card Program
	GiftCardProgram string `json:"giftCardProgram,omitempty"`

	// This field is for Valutec gift/loyalty card type. The value 0 is for gift and 1 for loyalty
	GiftProgramType string `json:"giftProgramType,omitempty"`

	// Specifies which lane to use.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// The card brand specific transaction ID that should be stored by the integrator.
	NetworkTransactionID string `json:"networkTransactionID,omitempty"`

	// Intent of why the credentials are being stored.
	RecurringPaymentType string `json:"recurringPaymentType,omitempty"`

	// A user defined reference number.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// An optional shift id for reference.
	ShiftID string `json:"shiftId,omitempty"`

	// Information necessary to process a StoreCard transaction.
	StoreCard *POSTStoreCardV1ReversalTransactionIDPaymentType `json:"storeCard,omitempty"`

	// Initial vs Subsequent Transaction.
	SubmissionType string `json:"submissionType,omitempty"`

	// An optional ticket number.
	TicketNumber string `json:"ticketNumber,omitempty"`

	// The original transaction amount.
	// Required: true
	TransactionAmount *float64 `json:"transactionAmount"`

	// The type of reversal.
	Type string `json:"type,omitempty"`
}

// Validate validates this reversal request
func (m *ReversalRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLaneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoreCard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalRequest) validateConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalRequest) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *ReversalRequest) validateStoreCard(formats strfmt.Registry) error {
	if swag.IsZero(m.StoreCard) { // not required
		return nil
	}

	if m.StoreCard != nil {
		if err := m.StoreCard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storeCard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storeCard")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalRequest) validateTransactionAmount(formats strfmt.Registry) error {

	if err := validate.Required("transactionAmount", "body", m.TransactionAmount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this reversal request based on the context it is used
func (m *ReversalRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoreCard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalRequest) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.Configuration != nil {
		if err := m.Configuration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalRequest) contextValidateStoreCard(ctx context.Context, formats strfmt.Registry) error {

	if m.StoreCard != nil {
		if err := m.StoreCard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storeCard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storeCard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalRequest) UnmarshalBinary(b []byte) error {
	var res ReversalRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
