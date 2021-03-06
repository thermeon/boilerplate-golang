// Code generated by go-swagger; DO NOT EDIT.

package connection_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new connection status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for connection status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConnectionStatusConnectionStatus(params *ConnectionStatusConnectionStatusParams, opts ...ClientOption) (*ConnectionStatusConnectionStatusOK, error)

	ConnectionStatusHistory(params *ConnectionStatusHistoryParams, opts ...ClientOption) (*ConnectionStatusHistoryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ConnectionStatusConnectionStatus Use this endpoint to see the connection status for a specific lane. Note that the triPOS Cloud environment obtains a snapshot of the PIN pad status every 75 seconds via GET/v1/lanes/{laneId}/connectionstatus endpoint. This endpoint does not connect to the PIN pad in real-time and any updates to the status of the PIN pad may take up to 75 seconds to be reflected in the status returned by the endpoint.
*/
func (a *Client) ConnectionStatusConnectionStatus(params *ConnectionStatusConnectionStatusParams, opts ...ClientOption) (*ConnectionStatusConnectionStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectionStatusConnectionStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConnectionStatus_ConnectionStatus",
		Method:             "GET",
		PathPattern:        "/v1/lanes/{laneId}/connectionstatus",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConnectionStatusConnectionStatusReader{formats: a.formats},
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
	success, ok := result.(*ConnectionStatusConnectionStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConnectionStatus_ConnectionStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ConnectionStatusHistory Use this endpoint to see the connection history for a specific lane.
*/
func (a *Client) ConnectionStatusHistory(params *ConnectionStatusHistoryParams, opts ...ClientOption) (*ConnectionStatusHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectionStatusHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConnectionStatus_History",
		Method:             "GET",
		PathPattern:        "/v1/lanes/{laneId}/connectionstatus/history",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConnectionStatusHistoryReader{formats: a.formats},
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
	success, ok := result.(*ConnectionStatusHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConnectionStatus_History: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
