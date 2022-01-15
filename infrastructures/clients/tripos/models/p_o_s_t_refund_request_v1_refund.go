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

// POSTRefundRequestV1Refund RefundRequest
//
// swagger:model POST_RefundRequest_v1_refund
type POSTRefundRequestV1Refund struct {

	// The cardholder address information for the transaction.
	Address *POSTAddressV1Refund `json:"address,omitempty"`

	// Defines whether the cardholder is present at the transaction. This value is optional, but recommended to be set. If this value is not set in the request, it will be automatically set based on the market code.
	CardHolderPresentCode string `json:"cardHolderPresentCode,omitempty"`

	// An optional clerk number for reference.
	ClerkNumber string `json:"clerkNumber,omitempty"`

	// The commercial card customer code for the transaction. This is for Level II.
	CommercialCardCustomerCode string `json:"commercialCardCustomerCode,omitempty"`

	// Any value included in this section will override the corresponding value set in the triPOS.config
	Configuration *POSTRefundRequestConfigurationV1Refund `json:"configuration,omitempty"`

	// The convenience fee amount.
	ConvenienceFeeAmount float64 `json:"convenienceFeeAmount,omitempty"`

	// Display the transaction amount during card read.
	DisplayTransactionAmount bool `json:"displayTransactionAmount,omitempty"`

	// To retrieve Token in the transaction. <a href='https://triposcert.vantiv.com/api/help/kb/GetToken.html'>more&raquo;</a>
	GetToken string `json:"getToken,omitempty"`

	// This field is for Valutec gift/loyalty card program. <br>'01' = Original Gift Card Program <br>'02' = Promotional Gift Card Program <br>'03' = Original Combo Card Program <br>'04' = Auto Rewards LPR Card Program <br>'05' = Original Loyalty Card Program
	GiftCardProgram string `json:"giftCardProgram,omitempty"`

	// This field is for Valutec gift/loyalty card type. The value 0 is for gift and 1 for loyalty
	GiftProgramType string `json:"giftProgramType,omitempty"`

	//  Invokes manual card entry.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`

	// Invokes prompt for cardholder to enter card security code for manual keyed card entry.
	IsCscSupported string `json:"isCscSupported,omitempty"`

	// Specifies which lane to use.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// Flag indicating desire to convert Credit card Refund transaction to a PINless Debit Card transaction. <a href='https://triposcert.vantiv.com/api/help/kb/PinlessDebitConversion.html'>more&raquo;</a>
	PinlessPosConversionIndicator string `json:"pinlessPosConversionIndicator,omitempty"`

	// Flag indicating desire to process refund as PreRead (i.e. - do ONLY Pre-Read operations, then expecting next endpoint call to be QuickChip) <a href='https://triposcert.vantiv.com/api\help\kb\cloud\QuickChipConfiguration.html'>more&raquo;</a>.
	PreRead bool `json:"preRead,omitempty"`

	// Flag indicating desire to process refund as QuickChip (i.e. - Pre-Read has been previously performed) <a href='https://triposcert.vantiv.com/api\help\kb\cloud\QuickChipConfiguration.html'>more&raquo;</a>.
	QuickChip bool `json:"quickChip,omitempty"`

	// A user defined reference number.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// The amount of sales tax for the transaction. This is for Level II, submit amount for tax or 0.00 for tax-exempt.
	SalesTaxAmount float64 `json:"salesTaxAmount,omitempty"`

	// An optional shift id for reference.
	ShiftID string `json:"shiftId,omitempty"`

	// Information necessary to process a StoreCard transaction.
	StoreCard *POSTStoreCardV1Refund `json:"storeCard,omitempty"`

	// An optional ticket number.
	TicketNumber string `json:"ticketNumber,omitempty"`

	// The total transaction amount. This is the amount of funds to move on the card
	// Required: true
	TransactionAmount *string `json:"transactionAmount"`
}

// Validate validates this p o s t refund request v1 refund
func (m *POSTRefundRequestV1Refund) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

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

func (m *POSTRefundRequestV1Refund) validateAddress(formats strfmt.Registry) error {
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

func (m *POSTRefundRequestV1Refund) validateConfiguration(formats strfmt.Registry) error {
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

func (m *POSTRefundRequestV1Refund) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *POSTRefundRequestV1Refund) validateStoreCard(formats strfmt.Registry) error {
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

func (m *POSTRefundRequestV1Refund) validateTransactionAmount(formats strfmt.Registry) error {

	if err := validate.Required("transactionAmount", "body", m.TransactionAmount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this p o s t refund request v1 refund based on the context it is used
func (m *POSTRefundRequestV1Refund) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

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

func (m *POSTRefundRequestV1Refund) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *POSTRefundRequestV1Refund) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

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

func (m *POSTRefundRequestV1Refund) contextValidateStoreCard(ctx context.Context, formats strfmt.Registry) error {

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
func (m *POSTRefundRequestV1Refund) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTRefundRequestV1Refund) UnmarshalBinary(b []byte) error {
	var res POSTRefundRequestV1Refund
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
