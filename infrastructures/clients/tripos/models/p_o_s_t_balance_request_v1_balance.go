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

// POSTBalanceRequestV1Balance BalanceRequest
//
// swagger:model POST_BalanceRequest_v1_balance
type POSTBalanceRequestV1Balance struct {

	// Defines whether the cardholder is present at the transaction. This value is optional, but recommended to be set. If this value is not set in the request, it will be automatically set based on the market code.
	CardHolderPresentCode string `json:"cardHolderPresentCode,omitempty"`

	// An optional clerk number for reference.
	ClerkNumber string `json:"clerkNumber,omitempty"`

	// Any value included in this section will override the corresponding value set in the triPOS.config
	Configuration *POSTRequestConfigurationV1Balance `json:"configuration,omitempty"`

	// EbtType may be FoodStamp or CashBenefit
	EbtType string `json:"ebtType,omitempty"`

	// This field is for Valutec gift/loyalty card program. <br>'01' = Original Gift Card Program <br>'02' = Promotional Gift Card Program <br>'03' = Original Combo Card Program <br>'04' = Auto Rewards LPR Card Program <br>'05' = Original Loyalty Card Program
	GiftCardProgram string `json:"giftCardProgram,omitempty"`

	// This field is for Valutec gift/loyalty card type. The value 0 is for gift and 1 for loyalty
	GiftProgramType string `json:"giftProgramType,omitempty"`

	// Invokes manual card entry.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`

	// Invokes prompt for cardholder to enter card security code for manual keyed card entry.
	IsCscSupported string `json:"isCscSupported,omitempty"`

	// Allows to use gift card functionality
	IsGiftSupported string `json:"isGiftSupported,omitempty"`

	// Specifies which lane to use.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// A user defined reference number.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// An optional shift id for reference.
	ShiftID string `json:"shiftId,omitempty"`

	// Information necessary to process a StoreCard transaction.
	StoreCard *POSTStoreCardV1Balance `json:"storeCard,omitempty"`

	// An optional ticket number.
	TicketNumber string `json:"ticketNumber,omitempty"`
}

// Validate validates this p o s t balance request v1 balance
func (m *POSTBalanceRequestV1Balance) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *POSTBalanceRequestV1Balance) validateConfiguration(formats strfmt.Registry) error {
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

func (m *POSTBalanceRequestV1Balance) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *POSTBalanceRequestV1Balance) validateStoreCard(formats strfmt.Registry) error {
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

// ContextValidate validate this p o s t balance request v1 balance based on the context it is used
func (m *POSTBalanceRequestV1Balance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *POSTBalanceRequestV1Balance) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

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

func (m *POSTBalanceRequestV1Balance) contextValidateStoreCard(ctx context.Context, formats strfmt.Registry) error {

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
func (m *POSTBalanceRequestV1Balance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTBalanceRequestV1Balance) UnmarshalBinary(b []byte) error {
	var res POSTBalanceRequestV1Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}