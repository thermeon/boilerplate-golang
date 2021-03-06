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

// POSTTokenSaleRequestV1SaleToken TokenSaleRequest
//
// swagger:model POST_TokenSaleRequest_v1_sale_token
type POSTTokenSaleRequestV1SaleToken struct {

	// The cardholder address information for the transaction.
	Address *POSTAddressV1SaleToken `json:"address,omitempty"`

	// The auto rental parameters to be passed in for an auto rental transaction. <a href='https://triposcert.vantiv.com/api/help/kb/autorental.html'>more&raquo;</a>
	AutoRental *POSTAutoRentalV1SaleToken `json:"autoRental,omitempty"`

	// The card logo of the card that the token represents.
	// Required: true
	CardLogo *string `json:"cardLogo"`

	// An optional clerk number for reference.
	ClerkNumber string `json:"clerkNumber,omitempty"`

	// Any value included in this section will override the corresponding value set in the triPOS.config
	Configuration *POSTRequestConfigurationV1SaleToken `json:"configuration,omitempty"`

	// The convenience fee amount of the transaction. This amount is added to the TotalAmount before the cardholder is charged
	ConvenienceFeeAmount float64 `json:"convenienceFeeAmount,omitempty"`

	// The expiration month of the card the token represents. Although this is not required, it is recommended for fraud protection.
	ExpirationMonth string `json:"expirationMonth,omitempty"`

	// The expiration year of the card the token represents. Although this is not required, it is recommended for fraud protection.
	ExpirationYear string `json:"expirationYear,omitempty"`

	// The healthcare section that contains all applicable healthcare-qualified amounts. <a href='https://triposcert.vantiv.com/api/help/kb/healthcare.html'>more&raquo;</a>
	Healthcare *POSTHealthcareV1SaleToken `json:"healthcare,omitempty"`

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

	// The token ID used for the sale.
	// Required: true
	TokenID *string `json:"tokenId"`

	// The type of token ID used for the sale.
	// Required: true
	TokenProvider *string `json:"tokenProvider"`

	// The total transaction amount. This is the amount of funds to move on the card
	// Required: true
	TransactionAmount *float64 `json:"transactionAmount"`

	// The Vault ID used to create the token.
	// Required: true
	VaultID *string `json:"vaultId"`
}

// Validate validates this p o s t token sale request v1 sale token
func (m *POSTTokenSaleRequestV1SaleToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoRental(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCardLogo(formats); err != nil {
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

	if err := m.validateTokenID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokenProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *POSTTokenSaleRequestV1SaleToken) validateAddress(formats strfmt.Registry) error {
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

func (m *POSTTokenSaleRequestV1SaleToken) validateAutoRental(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoRental) { // not required
		return nil
	}

	if m.AutoRental != nil {
		if err := m.AutoRental.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoRental")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoRental")
			}
			return err
		}
	}

	return nil
}

func (m *POSTTokenSaleRequestV1SaleToken) validateCardLogo(formats strfmt.Registry) error {

	if err := validate.Required("cardLogo", "body", m.CardLogo); err != nil {
		return err
	}

	return nil
}

func (m *POSTTokenSaleRequestV1SaleToken) validateConfiguration(formats strfmt.Registry) error {
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

func (m *POSTTokenSaleRequestV1SaleToken) validateHealthcare(formats strfmt.Registry) error {
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

func (m *POSTTokenSaleRequestV1SaleToken) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *POSTTokenSaleRequestV1SaleToken) validateTokenID(formats strfmt.Registry) error {

	if err := validate.Required("tokenId", "body", m.TokenID); err != nil {
		return err
	}

	return nil
}

func (m *POSTTokenSaleRequestV1SaleToken) validateTokenProvider(formats strfmt.Registry) error {

	if err := validate.Required("tokenProvider", "body", m.TokenProvider); err != nil {
		return err
	}

	return nil
}

func (m *POSTTokenSaleRequestV1SaleToken) validateTransactionAmount(formats strfmt.Registry) error {

	if err := validate.Required("transactionAmount", "body", m.TransactionAmount); err != nil {
		return err
	}

	return nil
}

func (m *POSTTokenSaleRequestV1SaleToken) validateVaultID(formats strfmt.Registry) error {

	if err := validate.Required("vaultId", "body", m.VaultID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p o s t token sale request v1 sale token based on the context it is used
func (m *POSTTokenSaleRequestV1SaleToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutoRental(ctx, formats); err != nil {
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

func (m *POSTTokenSaleRequestV1SaleToken) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *POSTTokenSaleRequestV1SaleToken) contextValidateAutoRental(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoRental != nil {
		if err := m.AutoRental.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoRental")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoRental")
			}
			return err
		}
	}

	return nil
}

func (m *POSTTokenSaleRequestV1SaleToken) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

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

func (m *POSTTokenSaleRequestV1SaleToken) contextValidateHealthcare(ctx context.Context, formats strfmt.Registry) error {

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
func (m *POSTTokenSaleRequestV1SaleToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTTokenSaleRequestV1SaleToken) UnmarshalBinary(b []byte) error {
	var res POSTTokenSaleRequestV1SaleToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
