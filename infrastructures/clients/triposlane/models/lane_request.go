// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LaneRequest lane request
//
// swagger:model LaneRequest
type LaneRequest struct {

	// activation code
	// Required: true
	ActivationCode *string `json:"activationCode"`

	// A Boolean value (true|false) that enables Contactless EMV, if supported.
	// This feature is supported on Ingenico devices with RBA version equal or greater than 23.0.44.
	ContactlessEmvEnabled string `json:"contactlessEmvEnabled,omitempty"`

	// A Boolean value (true|false) that enables Contactless MSD, if supported.
	// This feature is supported on Ingenico devices with RBA version less than 23.0.44.
	ContactlessMsdEnabled string `json:"contactlessMsdEnabled,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// lane Id
	// Required: true
	// Maximum: 999999
	// Minimum: 1
	LaneID *int32 `json:"laneId"`

	// market code
	// Read Only: true
	// Enum: [0 1 2 3 4 5 6 7 8]
	MarketCode int32 `json:"marketCode,omitempty"`

	// A string value (Vantiv|VantivCanada) that enables settings on device for USA|CAN.
	// This feature is supported on Ingenico Tetra Lane 3000 and Link 2500 devices.
	Processor string `json:"processor,omitempty"`

	// A positive integer between 30 and 600 indicating the number of seconds the pre-read.
	// data will be stored for the Quick-Chip feature, if supported and enabled.
	// Maximum: 600
	// Minimum: 30
	QuickChipDataLifetime int32 `json:"quickChipDataLifetime,omitempty"`

	// A Boolean value (true|false) that enables Quick-Chip feature, if supported.
	// This feature is supported on Ingenico devices with RBA version equal or greater than 23.0.44.
	QuickChipEnabled string `json:"quickChipEnabled,omitempty"`

	// terminal Id
	// Required: true
	TerminalID *string `json:"terminalId"`
}

// Validate validates this lane request
func (m *LaneRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivationCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLaneID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuickChipDataLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminalID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LaneRequest) validateActivationCode(formats strfmt.Registry) error {

	if err := validate.Required("activationCode", "body", m.ActivationCode); err != nil {
		return err
	}

	return nil
}

func (m *LaneRequest) validateLaneID(formats strfmt.Registry) error {

	if err := validate.Required("laneId", "body", m.LaneID); err != nil {
		return err
	}

	if err := validate.MinimumInt("laneId", "body", int64(*m.LaneID), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("laneId", "body", int64(*m.LaneID), 999999, false); err != nil {
		return err
	}

	return nil
}

var laneRequestTypeMarketCodePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6,7,8]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		laneRequestTypeMarketCodePropEnum = append(laneRequestTypeMarketCodePropEnum, v)
	}
}

// prop value enum
func (m *LaneRequest) validateMarketCodeEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, laneRequestTypeMarketCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LaneRequest) validateMarketCode(formats strfmt.Registry) error {
	if swag.IsZero(m.MarketCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateMarketCodeEnum("marketCode", "body", m.MarketCode); err != nil {
		return err
	}

	return nil
}

func (m *LaneRequest) validateQuickChipDataLifetime(formats strfmt.Registry) error {
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

func (m *LaneRequest) validateTerminalID(formats strfmt.Registry) error {

	if err := validate.Required("terminalId", "body", m.TerminalID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this lane request based on the context it is used
func (m *LaneRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMarketCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LaneRequest) contextValidateMarketCode(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "marketCode", "body", int32(m.MarketCode)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LaneRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LaneRequest) UnmarshalBinary(b []byte) error {
	var res LaneRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
