// Code generated by go-swagger; DO NOT EDIT.

package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewProfilesIdleProfileParams creates a new ProfilesIdleProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProfilesIdleProfileParams() *ProfilesIdleProfileParams {
	return &ProfilesIdleProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProfilesIdleProfileParamsWithTimeout creates a new ProfilesIdleProfileParams object
// with the ability to set a timeout on a request.
func NewProfilesIdleProfileParamsWithTimeout(timeout time.Duration) *ProfilesIdleProfileParams {
	return &ProfilesIdleProfileParams{
		timeout: timeout,
	}
}

// NewProfilesIdleProfileParamsWithContext creates a new ProfilesIdleProfileParams object
// with the ability to set a context for a request.
func NewProfilesIdleProfileParamsWithContext(ctx context.Context) *ProfilesIdleProfileParams {
	return &ProfilesIdleProfileParams{
		Context: ctx,
	}
}

// NewProfilesIdleProfileParamsWithHTTPClient creates a new ProfilesIdleProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewProfilesIdleProfileParamsWithHTTPClient(client *http.Client) *ProfilesIdleProfileParams {
	return &ProfilesIdleProfileParams{
		HTTPClient: client,
	}
}

/* ProfilesIdleProfileParams contains all the parameters to send to the API endpoint
   for the profiles idle profile operation.

   Typically these are written to a http.Request.
*/
type ProfilesIdleProfileParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	/* LaneID.

	   The lane ID.

	   Format: int32
	*/
	LaneID int32

	/* TpApplicationID.

	   The ID of the business application.
	*/
	TpApplicationID string

	/* TpApplicationName.

	   The name of the business application.
	*/
	TpApplicationName string

	/* TpApplicationVersion.

	   The version of the business application.
	*/
	TpApplicationVersion string

	/* TpExpressAcceptorID.

	   The Express acceptorId.
	*/
	TpExpressAcceptorID string

	/* TpExpressAccountID.

	   The Express accountId.
	*/
	TpExpressAccountID string

	/* TpExpressAccountToken.

	   The Express accountToken.
	*/
	TpExpressAccountToken string

	/* TpRequestID.

	   A unique ID for this request. This value should be a UUID or GUID.
	*/
	TpRequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the profiles idle profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProfilesIdleProfileParams) WithDefaults() *ProfilesIdleProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the profiles idle profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProfilesIdleProfileParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := ProfilesIdleProfileParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithTimeout(timeout time.Duration) *ProfilesIdleProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithContext(ctx context.Context) *ProfilesIdleProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithHTTPClient(client *http.Client) *ProfilesIdleProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithContentType(contentType string) *ProfilesIdleProfileParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithLaneID adds the laneID to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithLaneID(laneID int32) *ProfilesIdleProfileParams {
	o.SetLaneID(laneID)
	return o
}

// SetLaneID adds the laneId to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetLaneID(laneID int32) {
	o.LaneID = laneID
}

// WithTpApplicationID adds the tpApplicationID to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithTpApplicationID(tpApplicationID string) *ProfilesIdleProfileParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithTpApplicationName(tpApplicationName string) *ProfilesIdleProfileParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithTpApplicationVersion(tpApplicationVersion string) *ProfilesIdleProfileParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *ProfilesIdleProfileParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithTpExpressAccountID(tpExpressAccountID string) *ProfilesIdleProfileParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithTpExpressAccountToken(tpExpressAccountToken string) *ProfilesIdleProfileParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the profiles idle profile params
func (o *ProfilesIdleProfileParams) WithTpRequestID(tpRequestID string) *ProfilesIdleProfileParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the profiles idle profile params
func (o *ProfilesIdleProfileParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *ProfilesIdleProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// path param laneId
	if err := r.SetPathParam("laneId", swag.FormatInt32(o.LaneID)); err != nil {
		return err
	}

	// header param tp-application-id
	if err := r.SetHeaderParam("tp-application-id", o.TpApplicationID); err != nil {
		return err
	}

	// header param tp-application-name
	if err := r.SetHeaderParam("tp-application-name", o.TpApplicationName); err != nil {
		return err
	}

	// header param tp-application-version
	if err := r.SetHeaderParam("tp-application-version", o.TpApplicationVersion); err != nil {
		return err
	}

	// header param tp-express-acceptor-id
	if err := r.SetHeaderParam("tp-express-acceptor-id", o.TpExpressAcceptorID); err != nil {
		return err
	}

	// header param tp-express-account-id
	if err := r.SetHeaderParam("tp-express-account-id", o.TpExpressAccountID); err != nil {
		return err
	}

	// header param tp-express-account-token
	if err := r.SetHeaderParam("tp-express-account-token", o.TpExpressAccountToken); err != nil {
		return err
	}

	// header param tp-request-id
	if err := r.SetHeaderParam("tp-request-id", o.TpRequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}