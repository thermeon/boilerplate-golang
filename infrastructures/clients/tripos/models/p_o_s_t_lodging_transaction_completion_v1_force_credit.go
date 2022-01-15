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

// POSTLodgingTransactionCompletionV1ForceCredit LodgingTransactionCompletion
//
// swagger:model POST_LodgingTransactionCompletion_v1_force_credit
type POSTLodgingTransactionCompletionV1ForceCredit struct {

	// The lodging agreement number.
	AgreementNumber string `json:"agreementNumber,omitempty"`

	// The lodging charge type. For an empty or invalid value, this parameter defaults to Default.
	ChargeType string `json:"chargeType,omitempty"`

	// The check-in date for the lodging stay. Must be in the ISO 8601 format of YYYY-MM-DD.
	CheckInDate string `json:"checkInDate,omitempty"`

	// The check-out date for the lodging stay. Must be in the ISO 8601 format of YYYY-MM-DD.
	CheckOutDate string `json:"checkOutDate,omitempty"`

	// The lodging customer name.
	CustomerName string `json:"customerName,omitempty"`

	// The number of hotel stay days. Set to 1 for no show.
	Duration int32 `json:"duration,omitempty"`

	// Pass in up to 6 values that describe the extra charges. For an empty or invalid value, this parameter defaults to NotUsed.
	ExtraChargesDetail []string `json:"extraChargesDetail"`

	// Set to true to indicate no show. Otherwise, value defaults to false.
	NoShow bool `json:"noShow,omitempty"`

	// The lodging prestigious property code. For an empty or invalid value, this parameter defaults to NonParticipant.
	PrestigiousPropertyCode string `json:"prestigiousPropertyCode,omitempty"`

	// The nightly rate for one room.
	RoomAmount float64 `json:"roomAmount,omitempty"`

	// The lodging special program code. For an empty or invalid value, this parameter defaults to Default.
	SpecialProgramCode string `json:"specialProgramCode,omitempty"`
}

// Validate validates this p o s t lodging transaction completion v1 force credit
func (m *POSTLodgingTransactionCompletionV1ForceCredit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtraChargesDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pOSTLodgingTransactionCompletionV1ForceCreditExtraChargesDetailItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotUsed","Reserved","Restaurant","GiftShop","MiniBar","Telephone","Other","Laundry"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pOSTLodgingTransactionCompletionV1ForceCreditExtraChargesDetailItemsEnum = append(pOSTLodgingTransactionCompletionV1ForceCreditExtraChargesDetailItemsEnum, v)
	}
}

func (m *POSTLodgingTransactionCompletionV1ForceCredit) validateExtraChargesDetailItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pOSTLodgingTransactionCompletionV1ForceCreditExtraChargesDetailItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *POSTLodgingTransactionCompletionV1ForceCredit) validateExtraChargesDetail(formats strfmt.Registry) error {
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

// ContextValidate validates this p o s t lodging transaction completion v1 force credit based on context it is used
func (m *POSTLodgingTransactionCompletionV1ForceCredit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTLodgingTransactionCompletionV1ForceCredit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTLodgingTransactionCompletionV1ForceCredit) UnmarshalBinary(b []byte) error {
	var res POSTLodgingTransactionCompletionV1ForceCredit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}