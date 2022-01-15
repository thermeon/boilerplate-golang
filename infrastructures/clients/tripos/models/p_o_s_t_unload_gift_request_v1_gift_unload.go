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

// POSTUnloadGiftRequestV1GiftUnload UnloadGiftRequest
//
// swagger:model POST_UnloadGiftRequest_v1_gift_unload
type POSTUnloadGiftRequestV1GiftUnload struct {

	// Defines whether the cardholder is present at the transaction. This value is optional, but recommended to be set. If this value is not set in the request, it will be automatically set based on the market code.
	CardHolderPresentCode string `json:"cardHolderPresentCode,omitempty"`

	// An optional clerk number for reference.
	ClerkNumber string `json:"clerkNumber,omitempty"`

	// Any value included in this section will override the corresponding value set in the triPOS.config
	Configuration *POSTRequestConfigurationV1GiftUnload `json:"configuration,omitempty"`

	//  Invokes manual card entry.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`

	// Invokes prompt for cardholder to enter card security code for manual keyed card entry.
	IsCscSupported string `json:"isCscSupported,omitempty"`

	// Specifies which lane to use.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// A user defined reference number.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// An optional shift id for reference.
	ShiftID string `json:"shiftId,omitempty"`

	// Information necessary to process a StoreCard transaction.
	StoreCard *POSTStoreCardV1GiftUnload `json:"storeCard,omitempty"`

	// An optional ticket number.
	TicketNumber string `json:"ticketNumber,omitempty"`

	// The total transaction amount. This is the amount of funds to move on the card
	// Required: true
	TransactionAmount *float64 `json:"transactionAmount"`
}

// Validate validates this p o s t unload gift request v1 gift unload
func (m *POSTUnloadGiftRequestV1GiftUnload) Validate(formats strfmt.Registry) error {
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

func (m *POSTUnloadGiftRequestV1GiftUnload) validateConfiguration(formats strfmt.Registry) error {
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

func (m *POSTUnloadGiftRequestV1GiftUnload) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *POSTUnloadGiftRequestV1GiftUnload) validateStoreCard(formats strfmt.Registry) error {
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

func (m *POSTUnloadGiftRequestV1GiftUnload) validateTransactionAmount(formats strfmt.Registry) error {

	if err := validate.Required("transactionAmount", "body", m.TransactionAmount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p o s t unload gift request v1 gift unload based on the context it is used
func (m *POSTUnloadGiftRequestV1GiftUnload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *POSTUnloadGiftRequestV1GiftUnload) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

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

func (m *POSTUnloadGiftRequestV1GiftUnload) contextValidateStoreCard(ctx context.Context, formats strfmt.Registry) error {

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
func (m *POSTUnloadGiftRequestV1GiftUnload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTUnloadGiftRequestV1GiftUnload) UnmarshalBinary(b []byte) error {
	var res POSTUnloadGiftRequestV1GiftUnload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
