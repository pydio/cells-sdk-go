// Code generated by go-swagger; DO NOT EDIT.

package admin_tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new admin tree service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admin tree service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListAdminTree(params *ListAdminTreeParams, opts ...ClientOption) (*ListAdminTreeOK, error)

	StatAdminTree(params *StatAdminTreeParams, opts ...ClientOption) (*StatAdminTreeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListAdminTree lists files and folders starting at the root first level lists the datasources
*/
func (a *Client) ListAdminTree(params *ListAdminTreeParams, opts ...ClientOption) (*ListAdminTreeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAdminTreeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListAdminTree",
		Method:             "POST",
		PathPattern:        "/tree/admin/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListAdminTreeReader{formats: a.formats},
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
	success, ok := result.(*ListAdminTreeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListAdminTree: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StatAdminTree reads a node information inside the admin tree
*/
func (a *Client) StatAdminTree(params *StatAdminTreeParams, opts ...ClientOption) (*StatAdminTreeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatAdminTreeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StatAdminTree",
		Method:             "POST",
		PathPattern:        "/tree/admin/stat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &StatAdminTreeReader{formats: a.formats},
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
	success, ok := result.(*StatAdminTreeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for StatAdminTree: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
