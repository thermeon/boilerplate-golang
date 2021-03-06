// Code generated by go-swagger; DO NOT EDIT.

package pinpad

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

// NewGetv1pinpadcardfinanciallaneIDParams creates a new Getv1pinpadcardfinanciallaneIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetv1pinpadcardfinanciallaneIDParams() *Getv1pinpadcardfinanciallaneIDParams {
	return &Getv1pinpadcardfinanciallaneIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetv1pinpadcardfinanciallaneIDParamsWithTimeout creates a new Getv1pinpadcardfinanciallaneIDParams object
// with the ability to set a timeout on a request.
func NewGetv1pinpadcardfinanciallaneIDParamsWithTimeout(timeout time.Duration) *Getv1pinpadcardfinanciallaneIDParams {
	return &Getv1pinpadcardfinanciallaneIDParams{
		timeout: timeout,
	}
}

// NewGetv1pinpadcardfinanciallaneIDParamsWithContext creates a new Getv1pinpadcardfinanciallaneIDParams object
// with the ability to set a context for a request.
func NewGetv1pinpadcardfinanciallaneIDParamsWithContext(ctx context.Context) *Getv1pinpadcardfinanciallaneIDParams {
	return &Getv1pinpadcardfinanciallaneIDParams{
		Context: ctx,
	}
}

// NewGetv1pinpadcardfinanciallaneIDParamsWithHTTPClient creates a new Getv1pinpadcardfinanciallaneIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetv1pinpadcardfinanciallaneIDParamsWithHTTPClient(client *http.Client) *Getv1pinpadcardfinanciallaneIDParams {
	return &Getv1pinpadcardfinanciallaneIDParams{
		HTTPClient: client,
	}
}

/* Getv1pinpadcardfinanciallaneIDParams contains all the parameters to send to the API endpoint
   for the getv1pinpadcardfinanciallane Id operation.

   Typically these are written to a http.Request.
*/
type Getv1pinpadcardfinanciallaneIDParams struct {

	/* ContentType.

	   Content type for request.

	   Default: "application/json"
	*/
	ContentType string

	/* InvokeManualEntry.

	   The flag that invokes manual entry directly.
	*/
	InvokeManualEntry *bool

	/* IsCscSupported.

	   Invokes prompt for cardholder to enter card security code for manual keyed card entry.
	*/
	IsCscSupported *string

	/* IsEncryptedDataNeeded.

	   The flag that indicates if encrypted card data needs to be obtained.
	*/
	IsEncryptedDataNeeded *bool

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

// WithDefaults hydrates default values in the getv1pinpadcardfinanciallane Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Getv1pinpadcardfinanciallaneIDParams) WithDefaults() *Getv1pinpadcardfinanciallaneIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the getv1pinpadcardfinanciallane Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *Getv1pinpadcardfinanciallaneIDParams) SetDefaults() {
	var (
		contentTypeDefault = string("application/json")
	)

	val := Getv1pinpadcardfinanciallaneIDParams{
		ContentType: contentTypeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTimeout(timeout time.Duration) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithContext(ctx context.Context) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithHTTPClient(client *http.Client) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentType adds the contentType to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithContentType(contentType string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetContentType(contentType string) {
	o.ContentType = contentType
}

// WithInvokeManualEntry adds the invokeManualEntry to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithInvokeManualEntry(invokeManualEntry *bool) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetInvokeManualEntry(invokeManualEntry)
	return o
}

// SetInvokeManualEntry adds the invokeManualEntry to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetInvokeManualEntry(invokeManualEntry *bool) {
	o.InvokeManualEntry = invokeManualEntry
}

// WithIsCscSupported adds the isCscSupported to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithIsCscSupported(isCscSupported *string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetIsCscSupported(isCscSupported)
	return o
}

// SetIsCscSupported adds the isCscSupported to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetIsCscSupported(isCscSupported *string) {
	o.IsCscSupported = isCscSupported
}

// WithIsEncryptedDataNeeded adds the isEncryptedDataNeeded to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithIsEncryptedDataNeeded(isEncryptedDataNeeded *bool) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetIsEncryptedDataNeeded(isEncryptedDataNeeded)
	return o
}

// SetIsEncryptedDataNeeded adds the isEncryptedDataNeeded to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetIsEncryptedDataNeeded(isEncryptedDataNeeded *bool) {
	o.IsEncryptedDataNeeded = isEncryptedDataNeeded
}

// WithLaneID adds the laneID to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithLaneID(laneID int32) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetLaneID(laneID)
	return o
}

// SetLaneID adds the laneId to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetLaneID(laneID int32) {
	o.LaneID = laneID
}

// WithTpApplicationID adds the tpApplicationID to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpApplicationID(tpApplicationID string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpApplicationID(tpApplicationID)
	return o
}

// SetTpApplicationID adds the tpApplicationId to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpApplicationID(tpApplicationID string) {
	o.TpApplicationID = tpApplicationID
}

// WithTpApplicationName adds the tpApplicationName to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpApplicationName(tpApplicationName string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpApplicationName(tpApplicationName)
	return o
}

// SetTpApplicationName adds the tpApplicationName to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpApplicationName(tpApplicationName string) {
	o.TpApplicationName = tpApplicationName
}

// WithTpApplicationVersion adds the tpApplicationVersion to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpApplicationVersion(tpApplicationVersion string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpApplicationVersion(tpApplicationVersion)
	return o
}

// SetTpApplicationVersion adds the tpApplicationVersion to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpApplicationVersion(tpApplicationVersion string) {
	o.TpApplicationVersion = tpApplicationVersion
}

// WithTpAuthorization adds the tpAuthorization to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpAuthorization(tpAuthorization string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpAuthorization(tpAuthorization)
	return o
}

// SetTpAuthorization adds the tpAuthorization to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpAuthorization(tpAuthorization string) {
	o.TpAuthorization = tpAuthorization
}

// WithTpExpressAcceptorID adds the tpExpressAcceptorID to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpExpressAcceptorID(tpExpressAcceptorID string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpExpressAcceptorID(tpExpressAcceptorID)
	return o
}

// SetTpExpressAcceptorID adds the tpExpressAcceptorId to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpExpressAcceptorID(tpExpressAcceptorID string) {
	o.TpExpressAcceptorID = tpExpressAcceptorID
}

// WithTpExpressAccountID adds the tpExpressAccountID to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpExpressAccountID(tpExpressAccountID string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpExpressAccountID(tpExpressAccountID)
	return o
}

// SetTpExpressAccountID adds the tpExpressAccountId to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpExpressAccountID(tpExpressAccountID string) {
	o.TpExpressAccountID = tpExpressAccountID
}

// WithTpExpressAccountToken adds the tpExpressAccountToken to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpExpressAccountToken(tpExpressAccountToken string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpExpressAccountToken(tpExpressAccountToken)
	return o
}

// SetTpExpressAccountToken adds the tpExpressAccountToken to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpExpressAccountToken(tpExpressAccountToken string) {
	o.TpExpressAccountToken = tpExpressAccountToken
}

// WithTpRequestID adds the tpRequestID to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpRequestID(tpRequestID string) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpRequestID(tpRequestID)
	return o
}

// SetTpRequestID adds the tpRequestId to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpRequestID(tpRequestID string) {
	o.TpRequestID = tpRequestID
}

// WithTpReturnLogs adds the tpReturnLogs to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) WithTpReturnLogs(tpReturnLogs *bool) *Getv1pinpadcardfinanciallaneIDParams {
	o.SetTpReturnLogs(tpReturnLogs)
	return o
}

// SetTpReturnLogs adds the tpReturnLogs to the getv1pinpadcardfinanciallane Id params
func (o *Getv1pinpadcardfinanciallaneIDParams) SetTpReturnLogs(tpReturnLogs *bool) {
	o.TpReturnLogs = tpReturnLogs
}

// WriteToRequest writes these params to a swagger request
func (o *Getv1pinpadcardfinanciallaneIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Content-Type
	if err := r.SetHeaderParam("Content-Type", o.ContentType); err != nil {
		return err
	}

	if o.InvokeManualEntry != nil {

		// query param invokeManualEntry
		var qrInvokeManualEntry bool

		if o.InvokeManualEntry != nil {
			qrInvokeManualEntry = *o.InvokeManualEntry
		}
		qInvokeManualEntry := swag.FormatBool(qrInvokeManualEntry)
		if qInvokeManualEntry != "" {

			if err := r.SetQueryParam("invokeManualEntry", qInvokeManualEntry); err != nil {
				return err
			}
		}
	}

	if o.IsCscSupported != nil {

		// query param isCscSupported
		var qrIsCscSupported string

		if o.IsCscSupported != nil {
			qrIsCscSupported = *o.IsCscSupported
		}
		qIsCscSupported := qrIsCscSupported
		if qIsCscSupported != "" {

			if err := r.SetQueryParam("isCscSupported", qIsCscSupported); err != nil {
				return err
			}
		}
	}

	if o.IsEncryptedDataNeeded != nil {

		// query param isEncryptedDataNeeded
		var qrIsEncryptedDataNeeded bool

		if o.IsEncryptedDataNeeded != nil {
			qrIsEncryptedDataNeeded = *o.IsEncryptedDataNeeded
		}
		qIsEncryptedDataNeeded := swag.FormatBool(qrIsEncryptedDataNeeded)
		if qIsEncryptedDataNeeded != "" {

			if err := r.SetQueryParam("isEncryptedDataNeeded", qIsEncryptedDataNeeded); err != nil {
				return err
			}
		}
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
