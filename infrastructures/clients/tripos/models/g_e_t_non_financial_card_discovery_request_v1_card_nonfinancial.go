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

// GETNonFinancialCardDiscoveryRequestV1CardNonfinancial NonFinancialCardDiscoveryRequest
//
// swagger:model GET_NonFinancialCardDiscoveryRequest_v1_card_nonfinancial
type GETNonFinancialCardDiscoveryRequestV1CardNonfinancial struct {

	// Specifies which lane to use.
	// Required: true
	LaneID *int32 `json:"laneId"`
}

// Validate validates this g e t non financial card discovery request v1 card nonfinancial
func (m *GETNonFinancialCardDiscoveryRequestV1CardNonfinancial) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLaneID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GETNonFinancialCardDiscoveryRequestV1CardNonfinancial) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this g e t non financial card discovery request v1 card nonfinancial based on context it is used
func (m *GETNonFinancialCardDiscoveryRequestV1CardNonfinancial) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GETNonFinancialCardDiscoveryRequestV1CardNonfinancial) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GETNonFinancialCardDiscoveryRequestV1CardNonfinancial) UnmarshalBinary(b []byte) error {
	var res GETNonFinancialCardDiscoveryRequestV1CardNonfinancial
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
