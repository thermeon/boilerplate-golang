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

// AuthorizationRequest AuthorizationRequest
//
// swagger:model AuthorizationRequest
type AuthorizationRequest struct {

	// The cardholder address information for the transaction.
	Address *POSTAddressV1Authorization `json:"address,omitempty"`

	// The auto rental parameters to be passed in for an auto rental transaction. <a href='https://triposcert.vantiv.com/api/help/kb/autorental.html'>more&raquo;</a>
	AutoRental *POSTAutoRentalBasicV1Authorization `json:"autoRental,omitempty"`

	// Defines whether the cardholder is present at the transaction. This value is optional, but recommended to be set. If this value is not set in the request, it will be automatically set based on the market code.
	CardHolderPresentCode string `json:"cardHolderPresentCode,omitempty"`

	// An optional clerk number for reference.
	ClerkNumber string `json:"clerkNumber,omitempty"`

	// The commercial card customer code for the transaction. This is for Level II.
	CommercialCardCustomerCode string `json:"commercialCardCustomerCode,omitempty"`

	// Any value included in this section will override the corresponding value set in the triPOS.config
	Configuration *POSTAuthorizationRequestConfigurationV1Authorization `json:"configuration,omitempty"`

	// The convenience fee amount of the transaction. This amount is added to the TotalAmount before the cardholder is charged
	ConvenienceFeeAmount float64 `json:"convenienceFeeAmount,omitempty"`

	// Display the transaction amount during card read.
	DisplayTransactionAmount bool `json:"displayTransactionAmount,omitempty"`

	// To retrieve Token in the transaction. <a href='https://triposcert.vantiv.com/api/help/kb/GetToken.html'>more&raquo;</a>
	GetToken string `json:"getToken,omitempty"`

	// The healthcare section that contains all applicable healthcare-qualified amounts. <a href='https://triposcert.vantiv.com/api/help/kb/healthcare.html'>more&raquo;</a>
	Healthcare *POSTHealthcareV1Authorization `json:"healthcare,omitempty"`

	//  Invokes manual card entry.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`

	// Invokes prompt for cardholder to enter card security code for manual keyed card entry.
	IsCscSupported string `json:"isCscSupported,omitempty"`

	// Specifies which lane to use.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// The lodging parameters to be passed in for a lodging transaction. <a href='https://triposcert.vantiv.com/api/help/kb/lodging.html'>more&raquo;</a>
	Lodging *POSTLodgingAuthorizationV1Authorization `json:"lodging,omitempty"`

	// The card brand specific transaction ID that should be stored by the integrator.
	NetworkTransactionID string `json:"networkTransactionID,omitempty"`

	// Non Financial Card read Expected. Only supported on unattended devices.
	NonFinancialExpected bool `json:"nonFinancialExpected,omitempty"`

	// Flag indicating desire to process authorization as PreRead (i.e. - do ONLY Pre-Read operations, then expecting next endpoint call to be QuickChip) <a href='https://triposcert.vantiv.com/api\help\kb\cloud\QuickChipConfiguration.html'>more&raquo;</a>.
	PreRead bool `json:"preRead,omitempty"`

	// Flag indicating desire to process authorization as QuickChip (i.e. - Pre-Read has been previously performed) <a href='https://triposcert.vantiv.com/api\help\kb\cloud\QuickChipConfiguration.html'>more&raquo;</a>.
	QuickChip bool `json:"quickChip,omitempty"`

	// Recurring PaymentType,Intent of why the credentials are being stored.
	RecurringPaymentType string `json:"recurringPaymentType,omitempty"`

	// A user defined reference number.
	ReferenceNumber string `json:"referenceNumber,omitempty"`

	// The amount of sales tax for the transaction. This is for Level II, submit amount for tax or 0.00 for tax-exempt.
	SalesTaxAmount float64 `json:"salesTaxAmount,omitempty"`

	// An optional shift id for reference.
	ShiftID string `json:"shiftId,omitempty"`

	// SubmissionType,Initial vs Subsequent Transaction.
	SubmissionType string `json:"submissionType,omitempty"`

	// An optional ticket number.
	TicketNumber string `json:"ticketNumber,omitempty"`

	// The tip amount of the transaction. This amount is added to the TotalAmount before the cardholder is charged.
	TipAmount float64 `json:"tipAmount,omitempty"`

	// The total transaction amount. This is the amount of funds to move on the card
	// Required: true
	TransactionAmount *string `json:"transactionAmount"`
}

// Validate validates this authorization request
func (m *AuthorizationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoRental(formats); err != nil {
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

	if err := m.validateLodging(formats); err != nil {
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

func (m *AuthorizationRequest) validateAddress(formats strfmt.Registry) error {
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

func (m *AuthorizationRequest) validateAutoRental(formats strfmt.Registry) error {
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

func (m *AuthorizationRequest) validateConfiguration(formats strfmt.Registry) error {
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

func (m *AuthorizationRequest) validateHealthcare(formats strfmt.Registry) error {
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

func (m *AuthorizationRequest) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *AuthorizationRequest) validateLodging(formats strfmt.Registry) error {
	if swag.IsZero(m.Lodging) { // not required
		return nil
	}

	if m.Lodging != nil {
		if err := m.Lodging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lodging")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lodging")
			}
			return err
		}
	}

	return nil
}

func (m *AuthorizationRequest) validateTransactionAmount(formats strfmt.Registry) error {

	if err := validate.Required("transactionAmount", "body", m.TransactionAmount); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this authorization request based on the context it is used
func (m *AuthorizationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateLodging(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizationRequest) contextValidateAddress(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AuthorizationRequest) contextValidateAutoRental(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AuthorizationRequest) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AuthorizationRequest) contextValidateHealthcare(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AuthorizationRequest) contextValidateLodging(ctx context.Context, formats strfmt.Registry) error {

	if m.Lodging != nil {
		if err := m.Lodging.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lodging")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lodging")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizationRequest) UnmarshalBinary(b []byte) error {
	var res AuthorizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}