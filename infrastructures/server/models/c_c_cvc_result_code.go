// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CCCvcResultCode Issuer CVC check response code:
//   * D - Transaction determined suspicious by issuing bank.
//   * I - Card verification number failed processor's data validation check.
//   * M - Card verification number matched.
//   * N - Card verification number not matched.
//   * P - Card verification number not processed by processor for unspecified reason.
//   * S - Card verification number is on the card but was not included in the request.
//   * U - Card verification is not supported by the issuing bank.
//   * X - Card verification is not supported by the card association.
//   * 1 - Card verification is not supported for this processor or card type.
//   * 2 - Unrecognized result code returned by processor for card verification response.
//   * 3 - No result code returned by processor.
//
//
// swagger:model CCCvcResultCode
type CCCvcResultCode string

func NewCCCvcResultCode(value CCCvcResultCode) *CCCvcResultCode {
	v := value
	return &v
}

const (

	// CCCvcResultCodeD captures enum value "D"
	CCCvcResultCodeD CCCvcResultCode = "D"

	// CCCvcResultCodeI captures enum value "I"
	CCCvcResultCodeI CCCvcResultCode = "I"

	// CCCvcResultCodeM captures enum value "M"
	CCCvcResultCodeM CCCvcResultCode = "M"

	// CCCvcResultCodeN captures enum value "N"
	CCCvcResultCodeN CCCvcResultCode = "N"

	// CCCvcResultCodeP captures enum value "P"
	CCCvcResultCodeP CCCvcResultCode = "P"

	// CCCvcResultCodeS captures enum value "S"
	CCCvcResultCodeS CCCvcResultCode = "S"

	// CCCvcResultCodeU captures enum value "U"
	CCCvcResultCodeU CCCvcResultCode = "U"

	// CCCvcResultCodeX captures enum value "X"
	CCCvcResultCodeX CCCvcResultCode = "X"

	// CCCvcResultCodeNr1 captures enum value "1"
	CCCvcResultCodeNr1 CCCvcResultCode = "1"

	// CCCvcResultCodeNr2 captures enum value "2"
	CCCvcResultCodeNr2 CCCvcResultCode = "2"

	// CCCvcResultCodeNr3 captures enum value "3"
	CCCvcResultCodeNr3 CCCvcResultCode = "3"
)

// for schema
var cCCvcResultCodeEnum []interface{}

func init() {
	var res []CCCvcResultCode
	if err := json.Unmarshal([]byte(`["D","I","M","N","P","S","U","X","1","2","3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cCCvcResultCodeEnum = append(cCCvcResultCodeEnum, v)
	}
}

func (m CCCvcResultCode) validateCCCvcResultCodeEnum(path, location string, value CCCvcResultCode) error {
	if err := validate.EnumCase(path, location, value, cCCvcResultCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this c c cvc result code
func (m CCCvcResultCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCCCvcResultCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this c c cvc result code based on context it is used
func (m CCCvcResultCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
