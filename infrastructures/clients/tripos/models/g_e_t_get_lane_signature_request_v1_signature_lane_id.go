// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GETGetLaneSignatureRequestV1SignatureLaneID GetLaneSignatureRequest
//
// swagger:model GET_GetLaneSignatureRequest_v1_signature_laneId
type GETGetLaneSignatureRequestV1SignatureLaneID struct {

	// goto idle
	GotoIdle bool `json:"gotoIdle,omitempty"`
}

// Validate validates this g e t get lane signature request v1 signature lane Id
func (m *GETGetLaneSignatureRequestV1SignatureLaneID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g e t get lane signature request v1 signature lane Id based on context it is used
func (m *GETGetLaneSignatureRequestV1SignatureLaneID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GETGetLaneSignatureRequestV1SignatureLaneID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GETGetLaneSignatureRequestV1SignatureLaneID) UnmarshalBinary(b []byte) error {
	var res GETGetLaneSignatureRequestV1SignatureLaneID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}