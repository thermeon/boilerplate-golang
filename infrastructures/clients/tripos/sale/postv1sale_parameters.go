// Code generated by go-swagger; DO NOT EDIT.

package sale

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

// NewPostv1saleParams creates a new Postv1saleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostv1saleParams() *Postv1saleParams {
	return &Postv1saleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostv1saleParamsWithTimeout creates a new Postv1saleParams object
// with the ability to set a timeout on a request.
func NewPostv1saleParamsWithTimeout(timeout time.Duration) *Postv1saleParams {
	return &Postv1saleParams{
		timeout: timeout,
	}
}

// NewPostv1saleParamsWithContext creates a new Postv1saleParams object
// with the ability to set a context for a request.
func NewPostv1saleParamsWithContext(ctx context.Context) *Postv1saleParams {
	return &Postv1saleParams{
		Context: ctx,
	}
}

// NewPostv1saleParamsWithHTTPClient creates a new Postv1saleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostv1saleParamsWithHTTPClient(client *http.Client) *Postv1saleParams {
	return &Postv1saleParams{
		HTTPClient: client,
	}
}

/* Postv1saleParams contains all the parameters to send to the API endpoint
   for the postv1sale operation.

   Typically these are written to a http.Request.
*/
type Postv1saleParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	// SaleRequest.
	SaleRequest *models.POSTSaleRequestV1Sale

	/* Action.

	   The action to perform with the sale. This can either be 'store' or 'forward'. More information on <a href='https://triposcert.vantiv.com/api/help/kb/storeForward.html'>Store and Forward</a> here.
	*/
	Action *string

	/* RequestIDToForward.

	   The request ID of the stored transaction to forward.
	*/
	RequestIDToForward *string

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

// WithDefaults hydrates default values in the postv1sale params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1saleParams) WithDefaults() *Postv1saleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the postv1sale params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Postv1saleParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Postv1saleParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the postv1sale params
func (o *Postv1saleParams) WithTimeout(timeout time.Duration) *Postv1saleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the postv1sale params
func (o *Postv1saleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the postv1sale params
func (o *Postv1saleParams) WithContext(ctx context.Context) *Postv1saleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the postv1sale params
func (o *Postv1saleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the postv1sale params
func (o *Postv1saleParams) WithHTTPClient(client *http.Client) *Postv1saleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the postv1sale params
func (o *Postv1saleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the postv1sale params
func (o *Postv1saleParams) WithContentType(contentType string) *Postv1saleParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the postv1sale params
func (o *Postv1saleParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithSaleRequest adds the saleRequest to the postv1sale params
func (o *Postv1saleParams) WithSaleRequest(saleRequest *models.POSTSaleRequestV1Sale) *Postv1saleParams {
	o.SetSaleRequest(saleRequest)
	return o
}

// SetSaleRequest adds the saleRequest to the postv1sale params
func (o *Postv1saleParams) SetSaleRequest(saleRequest *models.POSTSaleRequestV1Sale) {
	o.SaleRequest = saleRequest
}

// WithAction adds the action to the postv1sale params
func (o *Postv1saleParams) WithAction(action *string) *Postv1saleParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the postv1sale params
func (o *Postv1saleParams) SetAction(action *string) {
	o.Action = action
}

// WithRequestIDToForward adds the requestIDToForward to the postv1sale params
func (o *Postv1saleParams) WithRequestIDToForward(requestIDToForward *string) *Postv1saleParams {
	o.SetRequestIDToForward(requestIDToForward)
	return o
}

// SetRequestIDToForward adds the requestIdToForward to the postv1sale params
func (o *Postv1saleParams) SetRequestIDToForward(requestIDToForward *string) {
	o.RequestIDToForward = requestIDToForward
}

// WithTpApplicationID adds the tpApplicationID to the postv1sale params
func (o *Postv1saleParams) WithTpApplicationID(tpApplicationID string) *Postv1saleParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the postv1sale params
func (o *Postv1saleParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the postv1sale params
func (o *Postv1saleParams) WithTpApplicationName(tpApplicationName string) *Postv1saleParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the postv1sale params
func (o *Postv1saleParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the postv1sale params
func (o *Postv1saleParams) WithTpApplicationVersion(tpApplicationVersion string) *Postv1saleParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the postv1sale params
func (o *Postv1saleParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the postv1sale params
func (o *Postv1saleParams) WithTpAuthorization(tpAuthorization string) *Postv1saleParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the postv1sale params
func (o *Postv1saleParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the postv1sale params
func (o *Postv1saleParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Postv1saleParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the postv1sale params
func (o *Postv1saleParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the postv1sale params
func (o *Postv1saleParams) WithTpExpressAccountID(tpExpressAccountID string) *Postv1saleParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the postv1sale params
func (o *Postv1saleParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the postv1sale params
func (o *Postv1saleParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Postv1saleParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the postv1sale params
func (o *Postv1saleParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the postv1sale params
func (o *Postv1saleParams) WithTpRequestID(tpRequestID string) *Postv1saleParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the postv1sale params
func (o *Postv1saleParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the postv1sale params
func (o *Postv1saleParams) WithTpReturnLogs(tpReturnLogs *bool) *Postv1saleParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the postv1sale params
func (o *Postv1saleParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WriteToRequest writes these params to a swagger request
func (o *Postv1saleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}
	if o.SaleRequest != nil {
		if err := r.SetBodyParam(o.SaleRequest); err != nil {
			return err
		}
	}

	if o.Action != nil {

		// query param action
		var qrAction string

		if o.Action != nil {
			qrAction = *o.Action
		}
		qAction := qrAction
		if qAction != "" {

			if err := r.SetQueryParam("action", qAction); err != nil {
				return err
			}
		}
	}

	if o.RequestIDToForward != nil {

		// query param requestIdToForward
		var qrRequestIDToForward string

		if o.RequestIDToForward != nil {
			qrRequestIDToForward = *o.RequestIDToForward
		}
		qRequestIDToForward := qrRequestIDToForward
		if qRequestIDToForward != "" {

			if err := r.SetQueryParam("requestIdToForward", qRequestIDToForward); err != nil {
				return err
			}
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
