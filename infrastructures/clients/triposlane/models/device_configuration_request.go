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

// DeviceConfigurationRequest device configuration request
//
// swagger:model DeviceConfigurationRequest
type DeviceConfigurationRequest struct {

	// is display custom aid screen
	IsDisplayCustomAidScreen bool `json:"isDisplayCustomAidScreen,omitempty"`

	// reboot time
	// Required: true
	RebootTime *string `json:"rebootTime"`
}

// Validate validates this device configuration request
func (m *DeviceConfigurationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRebootTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceConfigurationRequest) validateRebootTime(formats strfmt.Registry) error {

	if err := validate.Required("rebootTime", "body", m.RebootTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this device configuration request based on context it is used
func (m *DeviceConfigurationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceConfigurationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceConfigurationRequest) UnmarshalBinary(b []byte) error {
	var res DeviceConfigurationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}