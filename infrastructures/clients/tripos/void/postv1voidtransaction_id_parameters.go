// Code generated by go-swagger; DO NOT EDIT.

package void

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

// NewPostv1voidtransactionIDParams creates a new Postv1voidtransactionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostv1voidtransactionIDParams() *Postv1voidtransactionIDParams {
	return &Postv1voidtransactionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostv1voidtransactionIDParamsWithTimeout creates a new Postv1voidtransactionIDParams object
// with the ability to set a timeout on a request.
func NewPostv1voidtransactionIDParamsWithTimeout(timeout time.Duration) *Postv1voidtransactionIDParams {
	return &Postv1voidtransactionIDParams{
		timeout: timeout,
	}
}

// NewPostv1voidtransactionIDParamsWithContext creates a new Postv1voidtransactionIDParams object
// with the ability to set a context for a request.
func NewPostv1voidtransactionIDParamsWithContext(ctx context.Context) *Postv1voidtransactionIDParams {
	return &Postv1voidtransactionIDParams{
		Context: ctx,
	}
}

// NewPostv1voidtransactionIDParamsWithHTTPClient creates a new Postv1voidtransactionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostv1voidtransactionIDParamsWithHTTPClient(client *http.Client) *Postv1voidtransactionIDParams {
	return &Postv1voidtransactionIDParams{
		HTTPClient: client,
	}
}

/* Postv1voidtransactionIDParams contains all the parameters to send to the API endpoint
   for the postv1voidtransaction Id operation.

   Typically these are written to a http.Request.
*/
type Postv1voidtransactionIDParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	// VoidRequest.
	VoidRequest *models.POSTVoidRequestV1VoidTransactionID

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

// WithDefaults hydrates default values in the postv1voidtransaction Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1voidtransactionIDParams) WithDefaults() *Postv1voidtransactionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the postv1voidtransaction Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1voidtransactionIDParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Postv1voidtransactionIDParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTimeout(timeout time.Duration) *Postv1voidtransactionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithContext(ctx context.Context) *Postv1voidtransactionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithHTTPClient(client *http.Client) *Postv1voidtransactionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithContentType(contentType string) *Postv1voidtransactionIDParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithVoidRequest adds the voidRequest to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithVoidRequest(voidRequest *models.POSTVoidRequestV1VoidTransactionID) *Postv1voidtransactionIDParams {
	o.SetVoidRequest(voidRequest)
	return o
}

// SetVoidRequest adds the voidRequest to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetVoidRequest(voidRequest *models.POSTVoidRequestV1VoidTransactionID) {
	o.VoidRequest = voidRequest
}

// WithTpApplicationID adds the tpApplicationID to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpApplicationID(tpApplicationID string) *Postv1voidtransactionIDParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpApplicationName(tpApplicationName string) *Postv1voidtransactionIDParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpApplicationVersion(tpApplicationVersion string) *Postv1voidtransactionIDParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpAuthorization(tpAuthorization string) *Postv1voidtransactionIDParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Postv1voidtransactionIDParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpExpressAccountID(tpExpressAccountID string) *Postv1voidtransactionIDParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Postv1voidtransactionIDParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpRequestID(tpRequestID string) *Postv1voidtransactionIDParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTpReturnLogs(tpReturnLogs *bool) *Postv1voidtransactionIDParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WithTransactionID adds the transactionID to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) WithTransactionID(transactionID string) *Postv1voidtransactionIDParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the postv1voidtransaction Id params
func (o *Postv1voidtransactionIDParams) SetTransactionID(transactionID string) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *Postv1voidtransactionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}
	if o.VoidRequest != nil {
		if err := r.SetBodyParam(o.VoidRequest); err != nil {
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
