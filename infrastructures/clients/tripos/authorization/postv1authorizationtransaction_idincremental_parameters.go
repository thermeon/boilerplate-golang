// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewPostv1authorizationtransactionIdincrementalParams creates a new Postv1authorizationtransactionIdincrementalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostv1authorizationtransactionIdincrementalParams() *Postv1authorizationtransactionIdincrementalParams {
	return &Postv1authorizationtransactionIdincrementalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostv1authorizationtransactionIdincrementalParamsWithTimeout creates a new Postv1authorizationtransactionIdincrementalParams object
// with the ability to set a timeout on a request.
func NewPostv1authorizationtransactionIdincrementalParamsWithTimeout(timeout time.Duration) *Postv1authorizationtransactionIdincrementalParams {
	return &Postv1authorizationtransactionIdincrementalParams{
		timeout: timeout,
	}
}

// NewPostv1authorizationtransactionIdincrementalParamsWithContext creates a new Postv1authorizationtransactionIdincrementalParams object
// with the ability to set a context for a request.
func NewPostv1authorizationtransactionIdincrementalParamsWithContext(ctx context.Context) *Postv1authorizationtransactionIdincrementalParams {
	return &Postv1authorizationtransactionIdincrementalParams{
		Context: ctx,
	}
}

// NewPostv1authorizationtransactionIdincrementalParamsWithHTTPClient creates a new Postv1authorizationtransactionIdincrementalParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostv1authorizationtransactionIdincrementalParamsWithHTTPClient(client *http.Client) *Postv1authorizationtransactionIdincrementalParams {
	return &Postv1authorizationtransactionIdincrementalParams{
		HTTPClient: client,
	}
}

/* Postv1authorizationtransactionIdincrementalParams contains all the parameters to send to the API endpoint
   for the postv1authorizationtransaction idincremental operation.

   Typically these are written to a http.Request.
*/
type Postv1authorizationtransactionIdincrementalParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	// IncrementalAuthorizationRequest.
	IncrementalAuthorizationRequest *models.POSTIncrementalAuthorizationRequestV1AuthorizationTransactionIDIncremental

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

	/* TransactionID.

	   The ID of a previous authorization transaction.
	*/
	TransactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the postv1authorizationtransaction idincremental params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1authorizationtransactionIdincrementalParams) WithDefaults() *Postv1authorizationtransactionIdincrementalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the postv1authorizationtransaction idincremental params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1authorizationtransactionIdincrementalParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Postv1authorizationtransactionIdincrementalParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTimeout(timeout time.Duration) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithContext(ctx context.Context) *Postv1authorizationtransactionIdincrementalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithHTTPClient(client *http.Client) *Postv1authorizationtransactionIdincrementalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithContentType(contentType string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithIncrementalAuthorizationRequest adds the incrementalAuthorizationRequest to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithIncrementalAuthorizationRequest(incrementalAuthorizationRequest *models.POSTIncrementalAuthorizationRequestV1AuthorizationTransactionIDIncremental) *Postv1authorizationtransactionIdincrementalParams {
	o.SetIncrementalAuthorizationRequest(incrementalAuthorizationRequest)
	return o
}

// SetIncrementalAuthorizationRequest adds the incrementalAuthorizationRequest to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetIncrementalAuthorizationRequest(incrementalAuthorizationRequest *models.POSTIncrementalAuthorizationRequestV1AuthorizationTransactionIDIncremental) {
	o.IncrementalAuthorizationRequest = incrementalAuthorizationRequest
}

// WithTpApplicationID adds the tpApplicationID to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpApplicationID(tpApplicationID string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpApplicationName(tpApplicationName string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpApplicationVersion(tpApplicationVersion string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpAuthorization(tpAuthorization string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpExpressAccountID(tpExpressAccountID string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpRequestID(tpRequestID string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTpReturnLogs(tpReturnLogs *bool) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WithTransactionID adds the transactionID to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) WithTransactionID(transactionID string) *Postv1authorizationtransactionIdincrementalParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the postv1authorizationtransaction idincremental params
func (o *Postv1authorizationtransactionIdincrementalParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *Postv1authorizationtransactionIdincrementalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}
	if o.IncrementalAuthorizationRequest != nil {
		if err := r.SetBodyParam(o.IncrementalAuthorizationRequest); err != nil {
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

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
