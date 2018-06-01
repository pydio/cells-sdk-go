// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new frontend service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for frontend service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FrontBootConf adds some data to the initial set of parameters loaded by the frontend
*/
func (a *Client) FrontBootConf(params *FrontBootConfParams) (*FrontBootConfOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontBootConfParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontBootConf",
		Method:             "GET",
		PathPattern:        "/frontend/bootconf",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontBootConfReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontBootConfOK), nil

}

/*
FrontLog sends a log from front php to back
*/
func (a *Client) FrontLog(params *FrontLogParams) (*FrontLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontLog",
		Method:             "PUT",
		PathPattern:        "/frontend/frontlogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontLogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontLogOK), nil

}

/*
SettingsMenu sends a tree of nodes to be used a menu in the settings panel
*/
func (a *Client) SettingsMenu(params *SettingsMenuParams) (*SettingsMenuOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingsMenuParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SettingsMenu",
		Method:             "GET",
		PathPattern:        "/frontend/settings-menu",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SettingsMenuReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SettingsMenuOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}