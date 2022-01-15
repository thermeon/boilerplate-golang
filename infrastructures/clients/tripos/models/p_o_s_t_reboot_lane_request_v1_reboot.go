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

// POSTRebootLaneRequestV1Reboot RebootLaneRequest
//
// swagger:model POST_RebootLaneRequest_v1_reboot
type POSTRebootLaneRequestV1Reboot struct {

	// The desired lane ID. The lane ID should be a maximum of six digits long.
	// Required: true
	LaneID *int32 `json:"laneId"`
}

// Validate validates this p o s t reboot lane request v1 reboot
func (m *POSTRebootLaneRequestV1Reboot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLaneID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *POSTRebootLaneRequestV1Reboot) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this p o s t reboot lane request v1 reboot based on context it is used
func (m *POSTRebootLaneRequestV1Reboot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTRebootLaneRequestV1Reboot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTRebootLaneRequestV1Reboot) UnmarshalBinary(b []byte) error {
	var res POSTRebootLaneRequestV1Reboot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
