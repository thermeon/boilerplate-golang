// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LodgingIncrementalAuthorization LodgingIncrementalAuthorization
//
// swagger:model LodgingIncrementalAuthorization
type LodgingIncrementalAuthorization struct {

	// The number of hotel stay days. Set to 1 for no show.
	Duration int32 `json:"duration,omitempty"`
}

// Validate validates this lodging incremental authorization
func (m *LodgingIncrementalAuthorization) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lodging incremental authorization based on context it is used
func (m *LodgingIncrementalAuthorization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LodgingIncrementalAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LodgingIncrementalAuthorization) UnmarshalBinary(b []byte) error {
	var res LodgingIncrementalAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
