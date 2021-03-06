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

// GETHostStatusRequestV1StatusHost HostStatusRequest
//
// swagger:model GET_HostStatusRequest_v1_status_host
type GETHostStatusRequestV1StatusHost struct {

	// The Express acceptorId. If all three Express credentials are specified in the request headers, triPOS will use these credentials instead of the credentials in the triPOS.config for that request only.
	// Required: true
	AcceptorID *string `json:"acceptorId"`

	// The Express accountId. If all three Express credentials are specified in the request headers, triPOS will use these credentials instead of the credentials in the triPOS.config for that request only.
	// Required: true
	AccountID *string `json:"accountId"`

	// The Express accountToken. If all three Express credentials are specified in the request headers, triPOS will use these credentials instead of the credentials in the triPOS.config for that request only.
	// Required: true
	AccountToken *string `json:"accountToken"`

	// The ID of the business application.
	// Required: true
	ApplicationID *string `json:"applicationId"`

	// The name of the business application.
	// Required: true
	ApplicationName *string `json:"applicationName"`

	// The version of the business application.
	// Required: true
	ApplicationVersion *string `json:"applicationVersion"`

	// The authorization header.
	// Required: true
	AuthorizationSignature *string `json:"authorizationSignature"`

	// ctaf object
	// Required: true
	CtafObject *bool `json:"ctafObject"`

	// lane lock
	// Required: true
	LaneLock *GETILaneLockV1StatusHost `json:"laneLock"`

	// A unique ID for this request. This value should be a UUID or GUID. <a href='https://triposcert.vantiv.com/api/help/kb/requestid.html'>more&raquo;</a>
	// Required: true
	RequestID *string `json:"requestId"`

	// request Id as Guid
	// Required: true
	RequestIDAsGUID GETGUIDV1StatusHost `json:"requestIdAsGuid"`

	// Set to true to have logs populated in the response.
	ReturnLogs bool `json:"returnLogs,omitempty"`
}

// Validate validates this g e t host status request v1 status host
func (m *GETHostStatusRequestV1StatusHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizationSignature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCtafObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLaneLock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestIDAsGUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateAcceptorID(formats strfmt.Registry) error {

	if err := validate.Required("acceptorId", "body", m.AcceptorID); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateAccountToken(formats strfmt.Registry) error {

	if err := validate.Required("accountToken", "body", m.AccountToken); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateApplicationID(formats strfmt.Registry) error {

	if err := validate.Required("applicationId", "body", m.ApplicationID); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateApplicationName(formats strfmt.Registry) error {

	if err := validate.Required("applicationName", "body", m.ApplicationName); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateApplicationVersion(formats strfmt.Registry) error {

	if err := validate.Required("applicationVersion", "body", m.ApplicationVersion); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateAuthorizationSignature(formats strfmt.Registry) error {

	if err := validate.Required("authorizationSignature", "body", m.AuthorizationSignature); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateCtafObject(formats strfmt.Registry) error {

	if err := validate.Required("ctafObject", "body", m.CtafObject); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateLaneLock(formats strfmt.Registry) error {

	if err := validate.Required("laneLock", "body", m.LaneLock); err != nil {
		return err
	}

	if m.LaneLock != nil {
		if err := m.LaneLock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("laneLock")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("laneLock")
			}
			return err
		}
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateRequestID(formats strfmt.Registry) error {

	if err := validate.Required("requestId", "body", m.RequestID); err != nil {
		return err
	}

	return nil
}

func (m *GETHostStatusRequestV1StatusHost) validateRequestIDAsGUID(formats strfmt.Registry) error {

	if m.RequestIDAsGUID == nil {
		return errors.Required("requestIdAsGuid", "body", nil)
	}

	return nil
}

// ContextValidate validate this g e t host status request v1 status host based on the context it is used
func (m *GETHostStatusRequestV1StatusHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLaneLock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GETHostStatusRequestV1StatusHost) contextValidateLaneLock(ctx context.Context, formats strfmt.Registry) error {

	if m.LaneLock != nil {
		if err := m.LaneLock.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("laneLock")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("laneLock")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GETHostStatusRequestV1StatusHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GETHostStatusRequestV1StatusHost) UnmarshalBinary(b []byte) error {
	var res GETHostStatusRequestV1StatusHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
