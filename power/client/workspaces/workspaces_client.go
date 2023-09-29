// Code generated by go-swagger; DO NOT EDIT.

package workspaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new workspaces API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workspaces API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	V1WorkspacesGet(params *V1WorkspacesGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1WorkspacesGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
V1WorkspacesGet gets a workspace s information and capabilities
*/
func (a *Client) V1WorkspacesGet(params *V1WorkspacesGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*V1WorkspacesGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewV1WorkspacesGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "v1.workspaces.get",
		Method:             "GET",
		PathPattern:        "/v1/workspaces/{workspace_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &V1WorkspacesGetReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*V1WorkspacesGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for v1.workspaces.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
