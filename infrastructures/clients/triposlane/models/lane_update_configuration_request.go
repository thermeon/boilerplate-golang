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

// LaneUpdateConfigurationRequest lane update configuration request
//
// swagger:model LaneUpdateConfigurationRequest
type LaneUpdateConfigurationRequest struct {

	// A Boolean value (true|false) that enables Contactless EMV, if supported.
	// This feature is supported on Ingenico devices with RBA version equal or greater than 23.0.44.
	ContactlessEmvEnabled string `json:"contactlessEmvEnabled,omitempty"`

	// A Boolean value (true|false) that enables Contactless MSD, if supported.
	// This feature is supported on Ingenico devices with RBA version less than 23.0.44.
	ContactlessMsdEnabled string `json:"contactlessMsdEnabled,omitempty"`

	// A string value (English|French) that sets default lanugage on the device.
	// This feature is supported on Canadian Ingenico Tetra Lane 3000 and Link 2500 devices.
	Language string `json:"language,omitempty"`

	// A positive integer between 30 and 600 indicating the number of seconds the pre-read.
	// data will be stored for the Quick-Chip feature, if supported and enabled.
	// Maximum: 600
	// Minimum: 30
	QuickChipDataLifetime int32 `json:"quickChipDataLifetime,omitempty"`

	// A Boolean value (true|false) that enables Quick-Chip feature, if supported.
	// This feature is supported on Ingenico devices with RBA version equal or greater than 23.0.44.
	QuickChipEnabled string `json:"quickChipEnabled,omitempty"`
}

// Validate validates this lane update configuration request
func (m *LaneUpdateConfigurationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuickChipDataLifetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LaneUpdateConfigurationRequest) validateQuickChipDataLifetime(formats strfmt.Registry) error {
	if swag.IsZero(m.QuickChipDataLifetime) { // not required
		return nil
	}

	if err := validate.MinimumInt("quickChipDataLifetime", "body", int64(m.QuickChipDataLifetime), 30, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("quickChipDataLifetime", "body", int64(m.QuickChipDataLifetime), 600, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this lane update configuration request based on context it is used
func (m *LaneUpdateConfigurationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LaneUpdateConfigurationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LaneUpdateConfigurationRequest) UnmarshalBinary(b []byte) error {
	var res LaneUpdateConfigurationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
