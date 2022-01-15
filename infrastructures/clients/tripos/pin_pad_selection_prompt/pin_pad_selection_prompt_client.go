// Code generated by go-swagger; DO NOT EDIT.

package pin_pad_selection_prompt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pin pad selection prompt API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pin pad selection prompt API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Getv1pinPadSelectionPrompt(params *Getv1pinPadSelectionPromptParams, opts ...ClientOption) (*Getv1pinPadSelectionPromptOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Getv1pinPadSelectionPrompt deprecateds displays a custom p i n pad prompt with custom selection options
*/
func (a *Client) Getv1pinPadSelectionPrompt(params *Getv1pinPadSelectionPromptParams, opts ...ClientOption) (*Getv1pinPadSelectionPromptOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetv1pinPadSelectionPromptParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getv1pinPadSelectionPrompt",
		Method:             "GET",
		PathPattern:        "/api/v1/pinPadSelectionPrompt",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Getv1pinPadSelectionPromptReader{formats: a.formats},
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
	success, ok := result.(*Getv1pinPadSelectionPromptOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getv1pinPadSelectionPrompt: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
