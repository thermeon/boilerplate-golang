// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Emv Emv
//
// swagger:model Emv
type Emv struct {

	// The Application Identifier also known as the AID. Identifies the application as described in ISO/IEC 7816-5. Printed receipts are required to contain the AID as hexadecimal characters.
	ApplicationIdentifier string `json:"applicationIdentifier,omitempty"`

	// Mnemonic associated with the AID according to ISO/IEC 7816-5. If the Application Preferred Name is not available or the Issuer code table index is not supported, then the Application Label should be used on the receipt instead of the Application Preferred Name.
	ApplicationLabel string `json:"applicationLabel,omitempty"`

	// Preferred mnemonic associated with the AID. When the Application Preferred Name is present and the Issuer code table index is supported, then this data element is mandatory on the receipt.
	ApplicationPreferredName string `json:"applicationPreferredName,omitempty"`

	// The application transaction counter.
	ApplicationTransactionCounter string `json:"applicationTransactionCounter,omitempty"`

	// The EMV cryptogram type and value. It is a preferred best practice to include this data element on the receipt, but is not mandatory. This field contains cryptogram type followed by the cryptogram value.
	Cryptogram string `json:"cryptogram,omitempty"`

	// Indicates the code table according to ISO/IEC 8859 for displaying the Application Preferred Name.
	IssuerCodeTableIndex string `json:"issuerCodeTableIndex,omitempty"`

	// pin bypassed
	PinBypassed bool `json:"pinBypassed,omitempty"`

	// A name value collection of additional EMV tags that are required to appear on the receipt.
	Tags []*Tag `json:"tags"`
}

// Validate validates this emv
func (m *Emv) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Emv) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this emv based on the context it is used
func (m *Emv) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Emv) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Emv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Emv) UnmarshalBinary(b []byte) error {
	var res Emv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
