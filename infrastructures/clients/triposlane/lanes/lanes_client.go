// Code generated by go-swagger; DO NOT EDIT.

package lanes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new lanes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for lanes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	LanesActivationCode(params *LanesActivationCodeParams, opts ...ClientOption) (*LanesActivationCodeOK, error)

	LanesDelete(params *LanesDeleteParams, opts ...ClientOption) (*LanesDeleteOK, error)

	LanesGet(params *LanesGetParams, opts ...ClientOption) (*LanesGetOK, error)

	LanesGetList(params *LanesGetListParams, opts ...ClientOption) (*LanesGetListOK, error)

	LanesPatch(params *LanesPatchParams, opts ...ClientOption) (*LanesPatchOK, error)

	LanesPost(params *LanesPostParams, opts ...ClientOption) (*LanesPostOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  LanesActivationCode generates a new activation code

  Use this endpoint to generate a new activation code for your PIN pad. This allows for additional merchants to be associated with the PIN pad.
*/
func (a *Client) LanesActivationCode(params *LanesActivationCodeParams, opts ...ClientOption) (*LanesActivationCodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLanesActivationCodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Lanes_ActivationCode",
		Method:             "POST",
		PathPattern:        "/v1/lanes/{laneId}/activationCode",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LanesActivationCodeReader{formats: a.formats},
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
	success, ok := result.(*LanesActivationCodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Lanes_ActivationCode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LanesDelete Use this endpoint to delete a lane from your account. If currently connected, the connection to triPOS will be closed and a new activation code will be displayed.
*/
func (a *Client) LanesDelete(params *LanesDeleteParams, opts ...ClientOption) (*LanesDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLanesDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Lanes_Delete",
		Method:             "DELETE",
		PathPattern:        "/v1/lanes/{laneId}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LanesDeleteReader{formats: a.formats},
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
	success, ok := result.(*LanesDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Lanes_Delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LanesGet retrieves specific lane information

  Use this endpoint to see information for a specific lane.
*/
func (a *Client) LanesGet(params *LanesGetParams, opts ...ClientOption) (*LanesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLanesGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Lanes_Get",
		Method:             "GET",
		PathPattern:        "/v1/lanes/{laneId}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LanesGetReader{formats: a.formats},
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
	success, ok := result.(*LanesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Lanes_Get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LanesGetList returns a list of lane information

  Use this endpoint to see all the lanes currently associated with your account.
*/
func (a *Client) LanesGetList(params *LanesGetListParams, opts ...ClientOption) (*LanesGetListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLanesGetListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Lanes_GetList",
		Method:             "GET",
		PathPattern:        "/v1/lanes",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LanesGetListReader{formats: a.formats},
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
	success, ok := result.(*LanesGetListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Lanes_GetList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LanesPatch updates lane configuration

  Use this endpoint to update configuration for a specific lane.
<a href="https://triposcert.vantiv.com/api/help/kb/cloud/LaneRequest-Patch.html">more??</a>
*/
func (a *Client) LanesPatch(params *LanesPatchParams, opts ...ClientOption) (*LanesPatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLanesPatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Lanes_Patch",
		Method:             "PATCH",
		PathPattern:        "/v1/lanes/{laneId}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LanesPatchReader{formats: a.formats},
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
	success, ok := result.(*LanesPatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Lanes_Patch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LanesPost creates lane

  Use this endpoint to add a new lane to your account. An activation code should be displayed on the PIN pad you are trying to associate.
The intention is that the end user initiates the pairing by taking the activation code and entering it into the point of sale application.
<a href="https://triposcert.vantiv.com/api/help/kb/cloud/LaneRequest-Post.html">more??</a>
*/
func (a *Client) LanesPost(params *LanesPostParams, opts ...ClientOption) (*LanesPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLanesPostParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Lanes_Post",
		Method:             "POST",
		PathPattern:        "/v1/lanes",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "text/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LanesPostReader{formats: a.formats},
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
	success, ok := result.(*LanesPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Lanes_Post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
