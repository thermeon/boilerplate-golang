// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewRetrieveTransactionHistoryParams creates a new RetrieveTransactionHistoryParams object
//
// There are no default values defined in the spec.
func NewRetrieveTransactionHistoryParams() RetrieveTransactionHistoryParams {

	return RetrieveTransactionHistoryParams{}
}

// RetrieveTransactionHistoryParams contains all the bound params for the retrieve transaction history operation
// typically these are obtained from a http.Request
//
// swagger:parameters RetrieveTransactionHistory
type RetrieveTransactionHistoryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Transaction ID
	  Required: true
	  In: path
	*/
	TransactionID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewRetrieveTransactionHistoryParams() beforehand.
func (o *RetrieveTransactionHistoryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rTransactionID, rhkTransactionID, _ := route.Params.GetOK("transaction_id")
	if err := o.bindTransactionID(rTransactionID, rhkTransactionID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindTransactionID binds and validates parameter TransactionID from path.
func (o *RetrieveTransactionHistoryParams) bindTransactionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.TransactionID = raw

	return nil
}