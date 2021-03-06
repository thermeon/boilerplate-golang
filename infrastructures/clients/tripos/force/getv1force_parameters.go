// Code generated by go-swagger; DO NOT EDIT.

package force

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

// NewGetv1forceParams creates a new Getv1forceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetv1forceParams() *Getv1forceParams {
	return &Getv1forceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetv1forceParamsWithTimeout creates a new Getv1forceParams object
// with the ability to set a timeout on a request.
func NewGetv1forceParamsWithTimeout(timeout time.Duration) *Getv1forceParams {
	return &Getv1forceParams{
		timeout: timeout,
	}
}

// NewGetv1forceParamsWithContext creates a new Getv1forceParams object
// with the ability to set a context for a request.
func NewGetv1forceParamsWithContext(ctx context.Context) *Getv1forceParams {
	return &Getv1forceParams{
		Context: ctx,
	}
}

// NewGetv1forceParamsWithHTTPClient creates a new Getv1forceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetv1forceParamsWithHTTPClient(client *http.Client) *Getv1forceParams {
	return &Getv1forceParams{
		HTTPClient: client,
	}
}

/* Getv1forceParams contains all the parameters to send to the API endpoint
   for the getv1force operation.

   Typically these are written to a http.Request.
*/
type Getv1forceParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

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

// WithDefaults hydrates default values in the getv1force params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Getv1forceParams) WithDefaults() *Getv1forceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the getv1force params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Getv1forceParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Getv1forceParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the getv1force params
func (o *Getv1forceParams) WithTimeout(timeout time.Duration) *Getv1forceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getv1force params
func (o *Getv1forceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getv1force params
func (o *Getv1forceParams) WithContext(ctx context.Context) *Getv1forceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getv1force params
func (o *Getv1forceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getv1force params
func (o *Getv1forceParams) WithHTTPClient(client *http.Client) *Getv1forceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getv1force params
func (o *Getv1forceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the getv1force params
func (o *Getv1forceParams) WithContentType(contentType string) *Getv1forceParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the getv1force params
func (o *Getv1forceParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithTpApplicationID adds the tpApplicationID to the getv1force params
func (o *Getv1forceParams) WithTpApplicationID(tpApplicationID string) *Getv1forceParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the getv1force params
func (o *Getv1forceParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the getv1force params
func (o *Getv1forceParams) WithTpApplicationName(tpApplicationName string) *Getv1forceParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the getv1force params
func (o *Getv1forceParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the getv1force params
func (o *Getv1forceParams) WithTpApplicationVersion(tpApplicationVersion string) *Getv1forceParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the getv1force params
func (o *Getv1forceParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the getv1force params
func (o *Getv1forceParams) WithTpAuthorization(tpAuthorization string) *Getv1forceParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the getv1force params
func (o *Getv1forceParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the getv1force params
func (o *Getv1forceParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Getv1forceParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the getv1force params
func (o *Getv1forceParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the getv1force params
func (o *Getv1forceParams) WithTpExpressAccountID(tpExpressAccountID string) *Getv1forceParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the getv1force params
func (o *Getv1forceParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the getv1force params
func (o *Getv1forceParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Getv1forceParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the getv1force params
func (o *Getv1forceParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the getv1force params
func (o *Getv1forceParams) WithTpRequestID(tpRequestID string) *Getv1forceParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the getv1force params
func (o *Getv1forceParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the getv1force params
func (o *Getv1forceParams) WithTpReturnLogs(tpReturnLogs *bool) *Getv1forceParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the getv1force params
func (o *Getv1forceParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WriteToRequest writes these params to a swagger request
func (o *Getv1forceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
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
