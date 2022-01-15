// Code generated by go-swagger; DO NOT EDIT.

package display

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

	"github.com/thermeon/boilerplate-golang/infrastructures/clients/tripos/models"
)

// NewPostv1displayParams creates a new Postv1displayParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostv1displayParams() *Postv1displayParams {
	return &Postv1displayParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostv1displayParamsWithTimeout creates a new Postv1displayParams object
// with the ability to set a timeout on a request.
func NewPostv1displayParamsWithTimeout(timeout time.Duration) *Postv1displayParams {
	return &Postv1displayParams{
		timeout: timeout,
	}
}

// NewPostv1displayParamsWithContext creates a new Postv1displayParams object
// with the ability to set a context for a request.
func NewPostv1displayParamsWithContext(ctx context.Context) *Postv1displayParams {
	return &Postv1displayParams{
		Context: ctx,
	}
}

// NewPostv1displayParamsWithHTTPClient creates a new Postv1displayParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostv1displayParamsWithHTTPClient(client *http.Client) *Postv1displayParams {
	return &Postv1displayParams{
		HTTPClient: client,
	}
}

/* Postv1displayParams contains all the parameters to send to the API endpoint
   for the postv1display operation.

   Typically these are written to a http.Request.
*/
type Postv1displayParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	// DisplayRequest.
	DisplayRequest *models.POSTDisplayRequestV1Display

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

// WithDefaults hydrates default values in the postv1display params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1displayParams) WithDefaults() *Postv1displayParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the postv1display params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1displayParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Postv1displayParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the postv1display params
func (o *Postv1displayParams) WithTimeout(timeout time.Duration) *Postv1displayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the postv1display params
func (o *Postv1displayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the postv1display params
func (o *Postv1displayParams) WithContext(ctx context.Context) *Postv1displayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the postv1display params
func (o *Postv1displayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the postv1display params
func (o *Postv1displayParams) WithHTTPClient(client *http.Client) *Postv1displayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the postv1display params
func (o *Postv1displayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the postv1display params
func (o *Postv1displayParams) WithContentType(contentType string) *Postv1displayParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the postv1display params
func (o *Postv1displayParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithDisplayRequest adds the displayRequest to the postv1display params
func (o *Postv1displayParams) WithDisplayRequest(displayRequest *models.POSTDisplayRequestV1Display) *Postv1displayParams {
	o.SetDisplayRequest(displayRequest)
	return o
}

// SetDisplayRequest adds the displayRequest to the postv1display params
func (o *Postv1displayParams) SetDisplayRequest(displayRequest *models.POSTDisplayRequestV1Display) {
	o.DisplayRequest = displayRequest
}

// WithTpApplicationID adds the tpApplicationID to the postv1display params
func (o *Postv1displayParams) WithTpApplicationID(tpApplicationID string) *Postv1displayParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the postv1display params
func (o *Postv1displayParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the postv1display params
func (o *Postv1displayParams) WithTpApplicationName(tpApplicationName string) *Postv1displayParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the postv1display params
func (o *Postv1displayParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the postv1display params
func (o *Postv1displayParams) WithTpApplicationVersion(tpApplicationVersion string) *Postv1displayParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the postv1display params
func (o *Postv1displayParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the postv1display params
func (o *Postv1displayParams) WithTpAuthorization(tpAuthorization string) *Postv1displayParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the postv1display params
func (o *Postv1displayParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the postv1display params
func (o *Postv1displayParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Postv1displayParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the postv1display params
func (o *Postv1displayParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the postv1display params
func (o *Postv1displayParams) WithTpExpressAccountID(tpExpressAccountID string) *Postv1displayParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the postv1display params
func (o *Postv1displayParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the postv1display params
func (o *Postv1displayParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Postv1displayParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the postv1display params
func (o *Postv1displayParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the postv1display params
func (o *Postv1displayParams) WithTpRequestID(tpRequestID string) *Postv1displayParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the postv1display params
func (o *Postv1displayParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the postv1display params
func (o *Postv1displayParams) WithTpReturnLogs(tpReturnLogs *bool) *Postv1displayParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the postv1display params
func (o *Postv1displayParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WriteToRequest writes these params to a swagger request
func (o *Postv1displayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}
	if o.DisplayRequest != nil {
		if err := r.SetBodyParam(o.DisplayRequest); err != nil {
			return err
		}
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