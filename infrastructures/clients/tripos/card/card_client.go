// Code generated by go-swagger; DO NOT EDIT.

package card

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new card API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for card API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Getv1card(params *Getv1cardParams, opts ...ClientOption) (*Getv1cardOK, error)

	Getv1cardfinancial(params *Getv1cardfinancialParams, opts ...ClientOption) (*Getv1cardfinancialOK, error)

	Getv1cardfinanciallaneID(params *Getv1cardfinanciallaneIDParams, opts ...ClientOption) (*Getv1cardfinanciallaneIDOK, error)

	Getv1cardlaneID(params *Getv1cardlaneIDParams, opts ...ClientOption) (*Getv1cardlaneIDOK, error)

	Getv1cardnonfinancial(params *Getv1cardnonfinancialParams, opts ...ClientOption) (*Getv1cardnonfinancialOK, error)

	Getv1cardnonfinanciallaneID(params *Getv1cardnonfinanciallaneIDParams, opts ...ClientOption) (*Getv1cardnonfinanciallaneIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Getv1card returns a list of the tri p o s card endpoints
*/
func (a *Client) Getv1card(params *Getv1cardParams, opts ...ClientOption) (*Getv1cardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1cardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1card",
		Method:             "GET",
		PathPattern:        "/api/v1/card",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1cardReader{formats: a.formats},
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
	success, ok := result.(*Getv1cardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1card: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Getv1cardfinancial returns a list of the tri p o s financial card endpoints
*/
func (a *Client) Getv1cardfinancial(params *Getv1cardfinancialParams, opts ...ClientOption) (*Getv1cardfinancialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1cardfinancialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1cardfinancial",
		Method:             "GET",
		PathPattern:        "/api/v1/card/financial",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1cardfinancialReader{formats: a.formats},
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
	success, ok := result.(*Getv1cardfinancialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1cardfinancial: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Getv1cardfinanciallaneID gets financial card info
*/
func (a *Client) Getv1cardfinanciallaneID(params *Getv1cardfinanciallaneIDParams, opts ...ClientOption) (*Getv1cardfinanciallaneIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1cardfinanciallaneIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1cardfinanciallaneId",
		Method:             "GET",
		PathPattern:        "/api/v1/card/financial/{laneId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1cardfinanciallaneIDReader{formats: a.formats},
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
	success, ok := result.(*Getv1cardfinanciallaneIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1cardfinanciallaneId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Getv1cardlaneID gets card info
*/
func (a *Client) Getv1cardlaneID(params *Getv1cardlaneIDParams, opts ...ClientOption) (*Getv1cardlaneIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1cardlaneIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1cardlaneId",
		Method:             "GET",
		PathPattern:        "/api/v1/card/{laneId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1cardlaneIDReader{formats: a.formats},
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
	success, ok := result.(*Getv1cardlaneIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1cardlaneId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Getv1cardnonfinancial returns a list of the tri p o s non financial card endpoints
*/
func (a *Client) Getv1cardnonfinancial(params *Getv1cardnonfinancialParams, opts ...ClientOption) (*Getv1cardnonfinancialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1cardnonfinancialParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1cardnonfinancial",
		Method:             "GET",
		PathPattern:        "/api/v1/card/nonfinancial",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1cardnonfinancialReader{formats: a.formats},
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
	success, ok := result.(*Getv1cardnonfinancialOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1cardnonfinancial: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Getv1cardnonfinanciallaneID gets non financial card info
*/
func (a *Client) Getv1cardnonfinanciallaneID(params *Getv1cardnonfinanciallaneIDParams, opts ...ClientOption) (*Getv1cardnonfinanciallaneIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1cardnonfinanciallaneIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1cardnonfinanciallaneId",
		Method:             "GET",
		PathPattern:        "/api/v1/card/nonfinancial/{laneId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1cardnonfinanciallaneIDReader{formats: a.formats},
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
	success, ok := result.(*Getv1cardnonfinanciallaneIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1cardnonfinanciallaneId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}