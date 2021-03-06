// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetLaneSignatureRequest GetLaneSignatureRequest
//
// swagger:model GetLaneSignatureRequest
type GetLaneSignatureRequest struct {

	// goto idle
	GotoIdle bool `json:"gotoIdle,omitempty"`
}

// Validate validates this get lane signature request
func (m *GetLaneSignatureRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get lane signature request based on context it is used
func (m *GetLaneSignatureRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetLaneSignatureRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLaneSignatureRequest) UnmarshalBinary(b []byte) error {
	var res GetLaneSignatureRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
