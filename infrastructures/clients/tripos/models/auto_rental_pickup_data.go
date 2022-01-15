// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AutoRentalPickupData AutoRentalPickupData
//
// swagger:model AutoRentalPickupData
type AutoRentalPickupData struct {

	// Auto rental pickup city.
	City string `json:"city,omitempty"`

	// Country code where vehicle was picked up. For US, use 840.
	CountryCode string `json:"countryCode,omitempty"`

	// return date in yyyy-MM-ddTHH:mm:ss format.
	Date string `json:"date,omitempty"`

	// Name of business where vehicle was picked up.
	Location string `json:"location,omitempty"`

	// Abbreviation of state, province, or other country subdivision where vehicle was picked up.
	State string `json:"state,omitempty"`
}

// Validate validates this auto rental pickup data
func (m *AutoRentalPickupData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this auto rental pickup data based on context it is used
func (m *AutoRentalPickupData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AutoRentalPickupData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoRentalPickupData) UnmarshalBinary(b []byte) error {
	var res AutoRentalPickupData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}