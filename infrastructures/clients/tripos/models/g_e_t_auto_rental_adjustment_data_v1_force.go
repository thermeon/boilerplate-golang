// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GETAutoRentalAdjustmentDataV1Force AutoRentalAdjustmentData
//
// swagger:model GET_AutoRentalAdjustmentData_v1_force
type GETAutoRentalAdjustmentDataV1Force struct {

	// Adjustment amount.
	Amount float64 `json:"amount,omitempty"`

	// Indicates if adjustment for mileage, fuel, vehicle damage, etc. was made to a rental agreement.
	Code string `json:"code,omitempty"`

	// Pass in up to 6 values that describe the extra charges. For an empty or invalid value, this parameter defaults to NoExtraCharge.
	ExtraChargesDetail []string `json:"extraChargesDetail"`
}

// Validate validates this g e t auto rental adjustment data v1 force
func (m *GETAutoRentalAdjustmentDataV1Force) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtraChargesDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var gETAutoRentalAdjustmentDataV1ForceExtraChargesDetailItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NoExtraCharge","Gasoline","ExtraMileage","LateReturn","OneWayServiceFee","ParkingOrMovingViolation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gETAutoRentalAdjustmentDataV1ForceExtraChargesDetailItemsEnum = append(gETAutoRentalAdjustmentDataV1ForceExtraChargesDetailItemsEnum, v)
	}
}

func (m *GETAutoRentalAdjustmentDataV1Force) validateExtraChargesDetailItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gETAutoRentalAdjustmentDataV1ForceExtraChargesDetailItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GETAutoRentalAdjustmentDataV1Force) validateExtraChargesDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtraChargesDetail) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtraChargesDetail); i++ {

		// value enum
		if err := m.validateExtraChargesDetailItemsEnum("extraChargesDetail"+"."+strconv.Itoa(i), "body", m.ExtraChargesDetail[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this g e t auto rental adjustment data v1 force based on context it is used
func (m *GETAutoRentalAdjustmentDataV1Force) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GETAutoRentalAdjustmentDataV1Force) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GETAutoRentalAdjustmentDataV1Force) UnmarshalBinary(b []byte) error {
	var res GETAutoRentalAdjustmentDataV1Force
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
