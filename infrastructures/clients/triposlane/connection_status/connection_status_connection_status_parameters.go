// Code generated by go-swagger; DO NOT EDIT.

package connection_status

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

// NewConnectionStatusConnectionStatusParams creates a new ConnectionStatusConnectionStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConnectionStatusConnectionStatusParams() *ConnectionStatusConnectionStatusParams {
	return &ConnectionStatusConnectionStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConnectionStatusConnectionStatusParamsWithTimeout creates a new ConnectionStatusConnectionStatusParams object
// with the ability to set a timeout on a request.
func NewConnectionStatusConnectionStatusParamsWithTimeout(timeout time.Duration) *ConnectionStatusConnectionStatusParams {
	return &ConnectionStatusConnectionStatusParams{
		timeout: timeout,
	}
}

// NewConnectionStatusConnectionStatusParamsWithContext creates a new ConnectionStatusConnectionStatusParams object
// with the ability to set a context for a request.
func NewConnectionStatusConnectionStatusParamsWithContext(ctx context.Context) *ConnectionStatusConnectionStatusParams {
	return &ConnectionStatusConnectionStatusParams{
		Context: ctx,
	}
}

// NewConnectionStatusConnectionStatusParamsWithHTTPClient creates a new ConnectionStatusConnectionStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewConnectionStatusConnectionStatusParamsWithHTTPClient(client *http.Client) *ConnectionStatusConnectionStatusParams {
	return &ConnectionStatusConnectionStatusParams{
		HTTPClient: client,
	}
}

/* ConnectionStatusConnectionStatusParams contains all the parameters to send to the API endpoint
   for the connection status connection status operation.

   Typically these are written to a http.Request.
*/
type ConnectionStatusConnectionStatusParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	// LaneID.
	//
	// Format: int32
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

// WithDefaults hydrates default values in the connection status connection status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectionStatusConnectionStatusParams) WithDefaults() *ConnectionStatusConnectionStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the connection status connection status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectionStatusConnectionStatusParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := ConnectionStatusConnectionStatusParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithTimeout(timeout time.Duration) *ConnectionStatusConnectionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithContext(ctx context.Context) *ConnectionStatusConnectionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithHTTPClient(client *http.Client) *ConnectionStatusConnectionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithContentType(contentType string) *ConnectionStatusConnectionStatusParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithLaneID adds the laneID to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithLaneID(laneID int32) *ConnectionStatusConnectionStatusParams {
	o.SetLaneID(laneID)
	return o
}

// SetLaneID adds the laneId to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetLaneID(laneID int32) {
	o.LaneID = laneID
}

// WithTpApplicationID adds the tpApplicationID to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithTpApplicationID(tpApplicationID string) *ConnectionStatusConnectionStatusParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithTpApplicationName(tpApplicationName string) *ConnectionStatusConnectionStatusParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithTpApplicationVersion(tpApplicationVersion string) *ConnectionStatusConnectionStatusParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *ConnectionStatusConnectionStatusParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithTpExpressAccountID(tpExpressAccountID string) *ConnectionStatusConnectionStatusParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithTpExpressAccountToken(tpExpressAccountToken string) *ConnectionStatusConnectionStatusParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) WithTpRequestID(tpRequestID string) *ConnectionStatusConnectionStatusParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the connection status connection status params
func (o *ConnectionStatusConnectionStatusParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectionStatusConnectionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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