// Code generated by go-swagger; DO NOT EDIT.

package token

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

// NewPostv1tokenpaymetrictransactionIDParams creates a new Postv1tokenpaymetrictransactionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostv1tokenpaymetrictransactionIDParams() *Postv1tokenpaymetrictransactionIDParams {
	return &Postv1tokenpaymetrictransactionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostv1tokenpaymetrictransactionIDParamsWithTimeout creates a new Postv1tokenpaymetrictransactionIDParams object
// with the ability to set a timeout on a request.
func NewPostv1tokenpaymetrictransactionIDParamsWithTimeout(timeout time.Duration) *Postv1tokenpaymetrictransactionIDParams {
	return &Postv1tokenpaymetrictransactionIDParams{
		timeout: timeout,
	}
}

// NewPostv1tokenpaymetrictransactionIDParamsWithContext creates a new Postv1tokenpaymetrictransactionIDParams object
// with the ability to set a context for a request.
func NewPostv1tokenpaymetrictransactionIDParamsWithContext(ctx context.Context) *Postv1tokenpaymetrictransactionIDParams {
	return &Postv1tokenpaymetrictransactionIDParams{
		Context: ctx,
	}
}

// NewPostv1tokenpaymetrictransactionIDParamsWithHTTPClient creates a new Postv1tokenpaymetrictransactionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostv1tokenpaymetrictransactionIDParamsWithHTTPClient(client *http.Client) *Postv1tokenpaymetrictransactionIDParams {
	return &Postv1tokenpaymetrictransactionIDParams{
		HTTPClient: client,
	}
}

/* Postv1tokenpaymetrictransactionIDParams contains all the parameters to send to the API endpoint
   for the postv1tokenpaymetrictransaction Id operation.

   Typically these are written to a http.Request.
*/
type Postv1tokenpaymetrictransactionIDParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	// CreatePaymetricTokenWithTransIDRequest.
	CreatePaymetricTokenWithTransIDRequest *models.POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID

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

	   The ID of a previous transaction.
	*/
	TransactionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the postv1tokenpaymetrictransaction Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1tokenpaymetrictransactionIDParams) WithDefaults() *Postv1tokenpaymetrictransactionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the postv1tokenpaymetrictransaction Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1tokenpaymetrictransactionIDParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Postv1tokenpaymetrictransactionIDParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTimeout(timeout time.Duration) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithContext(ctx context.Context) *Postv1tokenpaymetrictransactionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithHTTPClient(client *http.Client) *Postv1tokenpaymetrictransactionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithContentType(contentType string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithCreatePaymetricTokenWithTransIDRequest adds the createPaymetricTokenWithTransIDRequest to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithCreatePaymetricTokenWithTransIDRequest(createPaymetricTokenWithTransIDRequest *models.POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID) *Postv1tokenpaymetrictransactionIDParams {
	o.SetCreatePaymetricTokenWithTransIDRequest(createPaymetricTokenWithTransIDRequest)
	return o
}

// SetCreatePaymetricTokenWithTransIDRequest adds the createPaymetricTokenWithTransIdRequest to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetCreatePaymetricTokenWithTransIDRequest(createPaymetricTokenWithTransIDRequest *models.POSTCreatePaymetricTokenWithTransIDRequestV1TokenPaymetricTransactionID) {
	o.CreatePaymetricTokenWithTransIDRequest = createPaymetricTokenWithTransIDRequest
}

// WithTpApplicationID adds the tpApplicationID to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpApplicationID(tpApplicationID string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpApplicationName(tpApplicationName string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpApplicationVersion(tpApplicationVersion string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpAuthorization(tpAuthorization string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpExpressAccountID(tpExpressAccountID string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpRequestID(tpRequestID string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTpReturnLogs(tpReturnLogs *bool) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WithTransactionID adds the transactionID to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) WithTransactionID(transactionID string) *Postv1tokenpaymetrictransactionIDParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the postv1tokenpaymetrictransaction Id params
func (o *Postv1tokenpaymetrictransactionIDParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *Postv1tokenpaymetrictransactionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}
	if o.CreatePaymetricTokenWithTransIDRequest != nil {
		if err := r.SetBodyParam(o.CreatePaymetricTokenWithTransIDRequest); err != nil {
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
