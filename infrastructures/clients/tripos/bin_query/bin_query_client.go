// Code generated by go-swagger; DO NOT EDIT.

package bin_query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bin query API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bin query API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Getv1binQuerylaneID(params *Getv1binQuerylaneIDParams, opts ...ClientOption) (*Getv1binQuerylaneIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Getv1binQuerylaneID determines if a card falls into a specific b i n range such as prepaid healthcare or debit

  Use the GET /binQuery/{laneId} endpoint to determine if a card falls into a specific BIN range such as prepaid, healthcare, or debit. The BIN query endpoint always makes an online call to the host in order to ensure the most up-to-date BIN information. If the card swiped or inserted is not found in the online BIN table, this endpoint returns a 404 Not Found HTTP response.
*/
func (a *Client) Getv1binQuerylaneID(params *Getv1binQuerylaneIDParams, opts ...ClientOption) (*Getv1binQuerylaneIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1binQuerylaneIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1binQuerylaneId",
		Method:             "GET",
		PathPattern:        "/api/v1/binQuery/{laneId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1binQuerylaneIDReader{formats: a.formats},
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
	success, ok := result.(*Getv1binQuerylaneIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1binQuerylaneId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
