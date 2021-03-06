// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/thermeon/boilerplate-golang/infrastructures/server/models"
)

// NewSaleTransactionParams creates a new SaleTransactionParams object
//
// There are no default values defined in the spec.
func NewSaleTransactionParams() SaleTransactionParams {

	return SaleTransactionParams{}
}

// SaleTransactionParams contains all the bound params for the sale transaction operation
// typically these are obtained from a http.Request
//
// swagger:parameters SaleTransaction
type SaleTransactionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	RequestCreate *models.Create
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSaleTransactionParams() beforehand.
func (o *SaleTransactionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Create
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("requestCreate", "body", ""))
			} else {
				res = append(res, errors.NewParseError("requestCreate", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.RequestCreate = &body
			}
		}
	} else {
		res = append(res, errors.Required("requestCreate", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
