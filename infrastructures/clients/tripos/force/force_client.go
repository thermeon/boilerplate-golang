// Code generated by go-swagger; DO NOT EDIT.

package force

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new force API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for force API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Getv1force(params *Getv1forceParams, opts ...ClientOption) (*Getv1forceOK, error)

	Postv1forcecredit(params *Postv1forcecreditParams, opts ...ClientOption) (*Postv1forcecreditOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Getv1force returns a the list of force endpoints
*/
func (a *Client) Getv1force(params *Getv1forceParams, opts ...ClientOption) (*Getv1forceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1forceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1force",
		Method:             "GET",
		PathPattern:        "/api/v1/force",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1forceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Getv1forceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1force: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Postv1forcecredit creates a new credit card force based on the passed in amounts

  This endpoint supports QuickChip/PreRead functionality. See <a href='https://triposcert.vantiv.com/api/help/kb/cloud/QuickChipConfiguration.html'> QuickChip/PreRead documentation</a> for more information.
*/
func (a *Client) Postv1forcecredit(params *Postv1forcecreditParams, opts ...ClientOption) (*Postv1forcecreditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostv1forcecreditParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postv1forcecredit",
		Method:             "POST",
		PathPattern:        "/api/v1/force/credit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Postv1forcecreditReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Postv1forcecreditOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postv1forcecredit: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
