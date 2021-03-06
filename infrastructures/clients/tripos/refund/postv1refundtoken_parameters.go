// Code generated by go-swagger; DO NOT EDIT.

package refund

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

// NewPostv1refundtokenParams creates a new Postv1refundtokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostv1refundtokenParams() *Postv1refundtokenParams {
	return &Postv1refundtokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostv1refundtokenParamsWithTimeout creates a new Postv1refundtokenParams object
// with the ability to set a timeout on a request.
func NewPostv1refundtokenParamsWithTimeout(timeout time.Duration) *Postv1refundtokenParams {
	return &Postv1refundtokenParams{
		timeout: timeout,
	}
}

// NewPostv1refundtokenParamsWithContext creates a new Postv1refundtokenParams object
// with the ability to set a context for a request.
func NewPostv1refundtokenParamsWithContext(ctx context.Context) *Postv1refundtokenParams {
	return &Postv1refundtokenParams{
		Context: ctx,
	}
}

// NewPostv1refundtokenParamsWithHTTPClient creates a new Postv1refundtokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostv1refundtokenParamsWithHTTPClient(client *http.Client) *Postv1refundtokenParams {
	return &Postv1refundtokenParams{
		HTTPClient: client,
	}
}

/* Postv1refundtokenParams contains all the parameters to send to the API endpoint
   for the postv1refundtoken operation.

   Typically these are written to a http.Request.
*/
type Postv1refundtokenParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	// TokenRefundRequest.
	TokenRefundRequest *models.POSTTokenRefundRequestV1RefundToken

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

// WithDefaults hydrates default values in the postv1refundtoken params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1refundtokenParams) WithDefaults() *Postv1refundtokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the postv1refundtoken params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1refundtokenParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Postv1refundtokenParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTimeout(timeout time.Duration) *Postv1refundtokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithContext(ctx context.Context) *Postv1refundtokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithHTTPClient(client *http.Client) *Postv1refundtokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithContentType(contentType string) *Postv1refundtokenParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithTokenRefundRequest adds the tokenRefundRequest to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTokenRefundRequest(tokenRefundRequest *models.POSTTokenRefundRequestV1RefundToken) *Postv1refundtokenParams {
	o.SetTokenRefundRequest(tokenRefundRequest)
	return o
}

// SetTokenRefundRequest adds the tokenRefundRequest to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTokenRefundRequest(tokenRefundRequest *models.POSTTokenRefundRequestV1RefundToken) {
	o.TokenRefundRequest = tokenRefundRequest
}

// WithTpApplicationID adds the tpApplicationID to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpApplicationID(tpApplicationID string) *Postv1refundtokenParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpApplicationName(tpApplicationName string) *Postv1refundtokenParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpApplicationVersion(tpApplicationVersion string) *Postv1refundtokenParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpAuthorization(tpAuthorization string) *Postv1refundtokenParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Postv1refundtokenParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpExpressAccountID(tpExpressAccountID string) *Postv1refundtokenParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Postv1refundtokenParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpRequestID(tpRequestID string) *Postv1refundtokenParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the postv1refundtoken params
func (o *Postv1refundtokenParams) WithTpReturnLogs(tpReturnLogs *bool) *Postv1refundtokenParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the postv1refundtoken params
func (o *Postv1refundtokenParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WriteToRequest writes these params to a swagger request
func (o *Postv1refundtokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}
	if o.TokenRefundRequest != nil {
		if err := r.SetBodyParam(o.TokenRefundRequest); err != nil {
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
