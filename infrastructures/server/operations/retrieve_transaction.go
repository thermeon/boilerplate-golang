// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// RetrieveTransactionHandlerFunc turns a function with the right signature into a retrieve transaction handler
type RetrieveTransactionHandlerFunc func(RetrieveTransactionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn RetrieveTransactionHandlerFunc) Handle(params RetrieveTransactionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// RetrieveTransactionHandler interface for that can handle valid retrieve transaction params
type RetrieveTransactionHandler interface {
	Handle(RetrieveTransactionParams, interface{}) middleware.Responder
}

// NewRetrieveTransaction creates a new http.Handler for the retrieve transaction operation
func NewRetrieveTransaction(ctx *middleware.Context, handler RetrieveTransactionHandler) *RetrieveTransaction {
	return &RetrieveTransaction{Context: ctx, Handler: handler}
}

/* RetrieveTransaction swagger:route GET /transaction/{transaction_id} retrieveTransaction

Returns the transaction matching the transaction_id

*/
type RetrieveTransaction struct {
	Context *middleware.Context
	Handler RetrieveTransactionHandler
}

func (o *RetrieveTransaction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewRetrieveTransactionParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
