// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GETBinQueryRequestV1BinQueryLaneID BinQueryRequest
//
// swagger:model GET_BinQueryRequest_v1_binQuery_laneId
type GETBinQueryRequestV1BinQueryLaneID struct {

	// The flag that invokes manual entry directly.
	InvokeManualEntry bool `json:"invokeManualEntry,omitempty"`
}

// Validate validates this g e t bin query request v1 bin query lane Id
func (m *GETBinQueryRequestV1BinQueryLaneID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this g e t bin query request v1 bin query lane Id based on context it is used
func (m *GETBinQueryRequestV1BinQueryLaneID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GETBinQueryRequestV1BinQueryLaneID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GETBinQueryRequestV1BinQueryLaneID) UnmarshalBinary(b []byte) error {
	var res GETBinQueryRequestV1BinQueryLaneID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
