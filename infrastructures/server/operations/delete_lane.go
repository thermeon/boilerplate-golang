// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteLaneHandlerFunc turns a function with the right signature into a delete lane handler
type DeleteLaneHandlerFunc func(DeleteLaneParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteLaneHandlerFunc) Handle(params DeleteLaneParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteLaneHandler interface for that can handle valid delete lane params
type DeleteLaneHandler interface {
	Handle(DeleteLaneParams, interface{}) middleware.Responder
}

// NewDeleteLane creates a new http.Handler for the delete lane operation
func NewDeleteLane(ctx *middleware.Context, handler DeleteLaneHandler) *DeleteLane {
	return &DeleteLane{Context: ctx, Handler: handler}
}

/* DeleteLane swagger:route DELETE /lane/{processor}/{lane_id} deleteLane

Delete a configured lane

*/
type DeleteLane struct {
	Context *middleware.Context
	Handler DeleteLaneHandler
}

func (o *DeleteLane) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteLaneParams()
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
