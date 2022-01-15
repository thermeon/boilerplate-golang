// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// POSTHealthcareV1Authorization Healthcare
//
// swagger:model POST_Healthcare_v1_authorization
type POSTHealthcareV1Authorization struct {

	// The total amount of healthcare-qualified goods that fall into the category of 'clinic'.
	Clinic float64 `json:"clinic,omitempty"`

	// The total amount of healthcare-qualified goods that fall into the category of 'dental'.
	Dental float64 `json:"dental,omitempty"`

	// The total amount of healthcare-qualified goods that fall into the category of 'prescription'.
	Prescription float64 `json:"prescription,omitempty"`

	// The total amount of healthcare-qualified goods. If any healthcare values are included, this value is required to be present.
	Total float64 `json:"total,omitempty"`

	// The total amount of healthcare-qualified goods that fall into the category of 'vision'.
	Vision float64 `json:"vision,omitempty"`
}

// Validate validates this p o s t healthcare v1 authorization
func (m *POSTHealthcareV1Authorization) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p o s t healthcare v1 authorization based on context it is used
func (m *POSTHealthcareV1Authorization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTHealthcareV1Authorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTHealthcareV1Authorization) UnmarshalBinary(b []byte) error {
	var res POSTHealthcareV1Authorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
