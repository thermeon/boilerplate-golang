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

// UpdatePinPadProfileIdleResponse update pin pad profile idle response
//
// swagger:model UpdatePinPadProfileIdleResponse
type UpdatePinPadProfileIdleResponse struct {

	// acceptor Id
	AcceptorID string `json:"acceptorId,omitempty"`

	// account Id
	AccountID int32 `json:"accountId,omitempty"`

	// command Id
	// Example: 00000000-0000-0000-0000-000000000000
	// Format: uuid
	CommandID strfmt.UUID `json:"commandId,omitempty"`

	// exception type
	ExceptionType string `json:"exceptionType,omitempty"`

	// idle image file
	IdleImageFile string `json:"idleImageFile,omitempty"`

	// idle image Url
	IdleImageURL string `json:"idleImageUrl,omitempty"`

	// idle message
	IdleMessage string `json:"idleMessage,omitempty"`

	// lane Id
	LaneID int32 `json:"laneId,omitempty"`

	// raw exception
	RawException string `json:"rawException,omitempty"`

	// request Id
	// Example: 00000000-0000-0000-0000-000000000000
	// Format: uuid
	RequestID strfmt.UUID `json:"requestId,omitempty"`

	// response type
	// Enum: [0 1 2 3 4 5 6 7]
	ResponseType int32 `json:"responseType,omitempty"`
}

// Validate validates this update pin pad profile idle response
func (m *UpdatePinPadProfileIdleResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePinPadProfileIdleResponse) validateCommandID(formats strfmt.Registry) error {
	if swag.IsZero(m.CommandID) { // not required
		return nil
	}

	if err := validate.FormatOf("commandId", "body", "uuid", m.CommandID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdatePinPadProfileIdleResponse) validateRequestID(formats strfmt.Registry) error {
	if swag.IsZero(m.RequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("requestId", "body", "uuid", m.RequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

var updatePinPadProfileIdleResponseTypeResponseTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2,3,4,5,6,7]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updatePinPadProfileIdleResponseTypeResponseTypePropEnum = append(updatePinPadProfileIdleResponseTypeResponseTypePropEnum, v)
	}
}

// prop value enum
func (m *UpdatePinPadProfileIdleResponse) validateResponseTypeEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, updatePinPadProfileIdleResponseTypeResponseTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdatePinPadProfileIdleResponse) validateResponseType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResponseType) { // not required
		return nil
	}

	// value enum
	if err := m.validateResponseTypeEnum("responseType", "body", m.ResponseType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update pin pad profile idle response based on context it is used
func (m *UpdatePinPadProfileIdleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePinPadProfileIdleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePinPadProfileIdleResponse) UnmarshalBinary(b []byte) error {
	var res UpdatePinPadProfileIdleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
