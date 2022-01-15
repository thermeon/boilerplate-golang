// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIError ApiError
//
// swagger:model ApiError
type APIError struct {

	// An error message targeted at the developer of the integrated business application.
	DeveloperMessage string `json:"developerMessage,omitempty"`

	// Code associated with the error if it exists.
	ErrorType string `json:"errorType,omitempty"`

	// The body of the exception message.
	ExceptionMessage string `json:"exceptionMessage,omitempty"`

	// The full name of the exception.
	ExceptionTypeFullName string `json:"exceptionTypeFullName,omitempty"`

	// The short name of the exception.
	ExceptionTypeShortName string `json:"exceptionTypeShortName,omitempty"`

	// An error message targeted at the end user of the integrated business application.
	UserMessage string `json:"userMessage,omitempty"`
}

// Validate validates this Api error
func (m *APIError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api error based on context it is used
func (m *APIError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIError) UnmarshalBinary(b []byte) error {
	var res APIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
