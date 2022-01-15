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

// ScrollingDisplayRequest ScrollingDisplayRequest
//
// swagger:model ScrollingDisplayRequest
type ScrollingDisplayRequest struct {

	// The lane ID.
	// Required: true
	LaneID *int32 `json:"laneId"`

	// The new text line to be added in the scrolling area of the PIN pad screen. It typically includes a description and the price of an individual item (i.e. Cookies... $3.99).
	// Required: true
	LineItem *string `json:"lineItem"`

	// The new subtotal of all the scanned items including the new line item.
	// Required: true
	Subtotal *string `json:"subtotal"`

	// The tax value of the new subtotal.
	// Required: true
	Tax *string `json:"tax"`

	// The value of the new subtotal and tax.
	// Required: true
	Total *string `json:"total"`
}

// Validate validates this scrolling display request
func (m *ScrollingDisplayRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLaneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScrollingDisplayRequest) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

func (m *ScrollingDisplayRequest) validateLineItem(formats strfmt.Registry) error {

	if err := validate.Required("lineItem", "body", m.LineItem); err != nil {
		return err
	}

	return nil
}

func (m *ScrollingDisplayRequest) validateSubtotal(formats strfmt.Registry) error {

	if err := validate.Required("subtotal", "body", m.Subtotal); err != nil {
		return err
	}

	return nil
}

func (m *ScrollingDisplayRequest) validateTax(formats strfmt.Registry) error {

	if err := validate.Required("tax", "body", m.Tax); err != nil {
		return err
	}

	return nil
}

func (m *ScrollingDisplayRequest) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scrolling display request based on context it is used
func (m *ScrollingDisplayRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScrollingDisplayRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScrollingDisplayRequest) UnmarshalBinary(b []byte) error {
	var res ScrollingDisplayRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}