// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GETFinancialCardRequestV1PinpadCardFinancialLaneID FinancialCardRequest
//
// swagger:model GET_FinancialCardRequest_v1_pinpad_card_financial_laneId
type GETFinancialCardRequestV1PinpadCardFinancialLaneID struct {

	// The flag that invokes manual entry directly.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`
}

// Validate validates this g e t financial card request v1 pinpad card financial lane Id
func (m *GETFinancialCardRequestV1PinpadCardFinancialLaneID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g e t financial card request v1 pinpad card financial lane Id based on context it is used
func (m *GETFinancialCardRequestV1PinpadCardFinancialLaneID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GETFinancialCardRequestV1PinpadCardFinancialLaneID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GETFinancialCardRequestV1PinpadCardFinancialLaneID) UnmarshalBinary(b []byte) error {
	var res GETFinancialCardRequestV1PinpadCardFinancialLaneID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
