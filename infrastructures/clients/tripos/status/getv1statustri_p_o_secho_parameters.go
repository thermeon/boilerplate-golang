// Code generated by go-swagger; DO NOT EDIT.

package status

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

// NewGetv1statustriPOSechoParams creates a new Getv1statustriPOSechoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetv1statustriPOSechoParams() *Getv1statustriPOSechoParams {
	return &Getv1statustriPOSechoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetv1statustriPOSechoParamsWithTimeout creates a new Getv1statustriPOSechoParams object
// with the ability to set a timeout on a request.
func NewGetv1statustriPOSechoParamsWithTimeout(timeout time.Duration) *Getv1statustriPOSechoParams {
	return &Getv1statustriPOSechoParams{
		timeout: timeout,
	}
}

// NewGetv1statustriPOSechoParamsWithContext creates a new Getv1statustriPOSechoParams object
// with the ability to set a context for a request.
func NewGetv1statustriPOSechoParamsWithContext(ctx context.Context) *Getv1statustriPOSechoParams {
	return &Getv1statustriPOSechoParams{
		Context: ctx,
	}
}

// NewGetv1statustriPOSechoParamsWithHTTPClient creates a new Getv1statustriPOSechoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetv1statustriPOSechoParamsWithHTTPClient(client *http.Client) *Getv1statustriPOSechoParams {
	return &Getv1statustriPOSechoParams{
		HTTPClient: client,
	}
}

/* Getv1statustriPOSechoParams contains all the parameters to send to the API endpoint
   for the getv1statustri p o secho operation.

   Typically these are written to a http.Request.
*/
type Getv1statustriPOSechoParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	/* Echo.

	   Text value that will be echoed back in the response.

	   Format: int32
	*/
	Echo int32

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

	/* TpAuthorization.

	   The authorization header.
	*/
	TpAuthorization string

	/* TpExpressAcceptorID.

	   The Express acceptorId. If all three Express credentials are specified in the request headers, triPOS will use these credentials instead of the credentials in the triPOS.config for that request only.
	*/
	TpExpressAcceptorID string

	/* TpExpressAccountID.

	   The Express accountId. If all three Express credentials are specified in the request headers, triPOS will use these credentials instead of the credentials in the triPOS.config for that request only.
	*/
	TpExpressAccountID string

	/* TpExpressAccountToken.

	   The Express accountToken. If all three Express credentials are specified in the request headers, triPOS will use these credentials instead of the credentials in the triPOS.config for that request only.
	*/
	TpExpressAccountToken string

	/* TpRequestID.

	   A unique ID for this request. This value should be a UUID or GUID. <a href='https://triposcert.vantiv.com/api/help/kb/requestid.html'>more&raquo;</a>
	*/
	TpRequestID string

	/* TpReturnLogs.

	   Set to true to have logs populated in the response.
	*/
	TpReturnLogs *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the getv1statustri p o secho params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Getv1statustriPOSechoParams) WithDefaults() *Getv1statustriPOSechoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the getv1statustri p o secho params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Getv1statustriPOSechoParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Getv1statustriPOSechoParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTimeout(timeout time.Duration) *Getv1statustriPOSechoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithContext(ctx context.Context) *Getv1statustriPOSechoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithHTTPClient(client *http.Client) *Getv1statustriPOSechoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithContentType(contentType string) *Getv1statustriPOSechoParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithEcho adds the echo to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithEcho(echo int32) *Getv1statustriPOSechoParams {
	o.SetEcho(echo)
	return o
}

// SetEcho adds the echo to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetEcho(echo int32) {
	o.Echo = echo
}

// WithTpApplicationID adds the tpApplicationID to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpApplicationID(tpApplicationID string) *Getv1statustriPOSechoParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpApplicationName(tpApplicationName string) *Getv1statustriPOSechoParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpApplicationVersion(tpApplicationVersion string) *Getv1statustriPOSechoParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpAuthorization(tpAuthorization string) *Getv1statustriPOSechoParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Getv1statustriPOSechoParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpExpressAccountID(tpExpressAccountID string) *Getv1statustriPOSechoParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Getv1statustriPOSechoParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpRequestID(tpRequestID string) *Getv1statustriPOSechoParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) WithTpReturnLogs(tpReturnLogs *bool) *Getv1statustriPOSechoParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the getv1statustri p o secho params
func (o *Getv1statustriPOSechoParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WriteToRequest writes these params to a swagger request
func (o *Getv1statustriPOSechoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	// path param echo
	if err := r.SetPathParam("echo", swag.FormatInt32(o.Echo)); err != nil {
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

	// header param tp-authorization
	if err := r.SetHeaderParam("tp-authorization", o.TpAuthorization); err != nil {
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

	if o.TpReturnLogs != nil {

		// header param tp-return-logs
		if err := r.SetHeaderParam("tp-return-logs", swag.FormatBool(*o.TpReturnLogs)); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
