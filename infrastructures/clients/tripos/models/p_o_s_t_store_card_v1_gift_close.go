// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// POSTStoreCardV1GiftClose StoreCard
//
// swagger:model POST_StoreCard_v1_gift_close
type POSTStoreCardV1GiftClose struct {

	// The ID of a StoreCard.
	ID string `json:"Id,omitempty"`

	// The password of StoreCard.
	Password string `json:"Password,omitempty"`
}

// Validate validates this p o s t store card v1 gift close
func (m *POSTStoreCardV1GiftClose) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p o s t store card v1 gift close based on context it is used
func (m *POSTStoreCardV1GiftClose) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *POSTStoreCardV1GiftClose) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *POSTStoreCardV1GiftClose) UnmarshalBinary(b []byte) error {
	var res POSTStoreCardV1GiftClose
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
