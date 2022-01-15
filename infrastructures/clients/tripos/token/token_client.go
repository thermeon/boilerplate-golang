// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new token API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for token API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Postv1tokenomni(params *Postv1tokenomniParams, opts ...ClientOption) (*Postv1tokenomniOK, error)

	Postv1tokenomnitransactionID(params *Postv1tokenomnitransactionIDParams, opts ...ClientOption) (*Postv1tokenomnitransactionIDOK, error)

	Postv1tokenpaymetric(params *Postv1tokenpaymetricParams, opts ...ClientOption) (*Postv1tokenpaymetricOK, error)

	Postv1tokenpaymetrictransactionID(params *Postv1tokenpaymetrictransactionIDParams, opts ...ClientOption) (*Postv1tokenpaymetrictransactionIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Postv1tokenomni creates a omni token based on a card swipe

  This endpoint is for swipe capable omni token creation.
*/
func (a *Client) Postv1tokenomni(params *Postv1tokenomniParams, opts ...ClientOption) (*Postv1tokenomniOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostv1tokenomniParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postv1tokenomni",
		Method:             "POST",
		PathPattern:        "/api/v1/token/omni",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Postv1tokenomniReader{formats: a.formats},
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
	success, ok := result.(*Postv1tokenomniOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postv1tokenomni: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Postv1tokenomnitransactionID creates a omni token based on a previous transaction ID
*/
func (a *Client) Postv1tokenomnitransactionID(params *Postv1tokenomnitransactionIDParams, opts ...ClientOption) (*Postv1tokenomnitransactionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostv1tokenomnitransactionIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postv1tokenomnitransactionId",
		Method:             "POST",
		PathPattern:        "/api/v1/token/omni/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Postv1tokenomnitransactionIDReader{formats: a.formats},
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
	success, ok := result.(*Postv1tokenomnitransactionIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postv1tokenomnitransactionId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Postv1tokenpaymetric creates a paymetric token based on a card swipe

  This endpoint is for swipe capable paymetric token creation.
*/
func (a *Client) Postv1tokenpaymetric(params *Postv1tokenpaymetricParams, opts ...ClientOption) (*Postv1tokenpaymetricOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostv1tokenpaymetricParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postv1tokenpaymetric",
		Method:             "POST",
		PathPattern:        "/api/v1/token/paymetric",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Postv1tokenpaymetricReader{formats: a.formats},
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
	success, ok := result.(*Postv1tokenpaymetricOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postv1tokenpaymetric: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Postv1tokenpaymetrictransactionID creates a paymetric token based on a previous transaction ID
*/
func (a *Client) Postv1tokenpaymetrictransactionID(params *Postv1tokenpaymetrictransactionIDParams, opts ...ClientOption) (*Postv1tokenpaymetrictransactionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostv1tokenpaymetrictransactionIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "postv1tokenpaymetrictransactionId",
		Method:             "POST",
		PathPattern:        "/api/v1/token/paymetric/{transactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Postv1tokenpaymetrictransactionIDReader{formats: a.formats},
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
	success, ok := result.(*Postv1tokenpaymetrictransactionIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postv1tokenpaymetrictransactionId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
