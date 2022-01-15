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

// PaymentAccountAuthorizationRequest PaymentAccountAuthorizationRequest
//
// swagger:model PaymentAccountAuthorizationRequest
type PaymentAccountAuthorizationRequest struct {

	// The cardholder address information for the transaction.
	Address *POSTAddressV1AuthorizationPaymentAccountID `json:"address,omitempty"`

	// An optional clerk number for reference.
	ClerkNumber string `json:"clerkNumber,omitempty"`

	// Any value included in this section will override the corresponding value set in the triPOS.config
	Configuration *POSTRequestConfigurationV1AuthorizationPaymentAccountID `json:"configuration,omitempty"`

	// The convenience fee amount of the transaction. This amount is added to the TotalAmount before the cardholder is charged
	ConvenienceFeeAmount float64 `json:"convenienceFeeAmount,omitempty"`

	// The healthcare section that contains all applicable healthcare-qualified amounts. <a href='https://triposcert.vantiv.com/api/help/kb/healthcare.html'>more&raquo;</a>
	Healthcare *POSTHealthcareV1AuthorizationPaymentAccountID `json:"healthcare,omitempty"`

	// Specifies which lane to use for the card sale.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// A user defined reference number.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// An optional shift id for reference.
	ShiftID string `json:"shiftId,omitempty"`

	// An optional ticket number.
	TicketNumber string `json:"ticketNumber,omitempty"`

	// The tip amount of the transaction. This amount is added to the TotalAmount before the cardholder is charged.
	TipAmount float64 `json:"tipAmount,omitempty"`

	// The total transaction amount. This is the amount of funds to move on the card
	// Required: true
	TransactionAmount *float64 `json:"transactionAmount"`
}

// Validate validates this payment account authorization request
func (m *PaymentAccountAuthorizationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthcare(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLaneID(formats); err != nil {
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

func (m *PaymentAccountAuthorizationRequest) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAccountAuthorizationRequest) validateConfiguration(formats strfmt.Registry) error {
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

func (m *PaymentAccountAuthorizationRequest) validateHealthcare(formats strfmt.Registry) error {
	if swag.IsZero(m.Healthcare) { // not required
		return nil
	}

	if m.Healthcare != nil {
		if err := m.Healthcare.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthcare")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("healthcare")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAccountAuthorizationRequest) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAccountAuthorizationRequest) validateTransactionAmount(formats strfmt.Registry) error {

	if err := validate.Required("transactionAmount", "body", m.TransactionAmount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this payment account authorization request based on the context it is used
func (m *PaymentAccountAuthorizationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthcare(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAccountAuthorizationRequest) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

	if m.Address != nil {
		if err := m.Address.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAccountAuthorizationRequest) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PaymentAccountAuthorizationRequest) contextValidateHealthcare(ctx context.Context, formats strfmt.Registry) error {

	if m.Healthcare != nil {
		if err := m.Healthcare.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthcare")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("healthcare")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAccountAuthorizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAccountAuthorizationRequest) UnmarshalBinary(b []byte) error {
	var res PaymentAccountAuthorizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
