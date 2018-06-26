// Code generated by go-swagger; DO NOT EDIT.

package workspace_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new workspace service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workspace service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteWorkspace deletes an existing workspace
*/
func (a *Client) DeleteWorkspace(params *DeleteWorkspaceParams) (*DeleteWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteWorkspace",
		Method:             "DELETE",
		PathPattern:        "/workspace/{Slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &DeleteWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWorkspaceOK), nil

}

/*
PutWorkspace creates or update a workspace
*/
func (a *Client) PutWorkspace(params *PutWorkspaceParams) (*PutWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutWorkspace",
		Method:             "PUT",
		PathPattern:        "/workspace/{Slug}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutWorkspaceOK), nil

}

/*
SearchWorkspaces searches workspaces on certain keys
*/
func (a *Client) SearchWorkspaces(params *SearchWorkspacesParams) (*SearchWorkspacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchWorkspacesParams()
	}

	op := runtime.ClientOperation{
		ID:                 "SearchWorkspaces",
		Method:             "POST",
		PathPattern:        "/workspace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SearchWorkspacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}

	restr, err := json.Marshal(op)
	if err != nil {
		fmt.Println("Cannot marshall request", err.Error())
	}
	fmt.Println("###### Marshalled request", string(restr))

	result, err := a.transport.Submit(&op)
	if err != nil {
		return nil, err
	}
	return result.(*SearchWorkspacesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
