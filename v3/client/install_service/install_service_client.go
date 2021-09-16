// Code generated by go-swagger; DO NOT EDIT.

package install_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new install service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for install service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetAgreement(params *GetAgreementParams, opts ...ClientOption) (*GetAgreementOK, error)

	GetInstall(params *GetInstallParams, opts ...ClientOption) (*GetInstallOK, error)

	InstallEvents(params *InstallEventsParams, opts ...ClientOption) (*InstallEventsOK, error)

	PerformInstallCheck(params *PerformInstallCheckParams, opts ...ClientOption) (*PerformInstallCheckOK, error)

	PostInstall(params *PostInstallParams, opts ...ClientOption) (*PostInstallOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAgreement loads a textual agreement for using the software
*/
func (a *Client) GetAgreement(params *GetAgreementParams, opts ...ClientOption) (*GetAgreementOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAgreementParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAgreement",
		Method:             "GET",
		PathPattern:        "/install/agreement",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &GetAgreementReader{formats: a.formats},
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
	success, ok := result.(*GetAgreementOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAgreement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInstall loads default values for install form
*/
func (a *Client) GetInstall(params *GetInstallParams, opts ...ClientOption) (*GetInstallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetInstall",
		Method:             "GET",
		PathPattern:        "/install",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &GetInstallReader{formats: a.formats},
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
	success, ok := result.(*GetInstallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetInstall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InstallEvents install events API
*/
func (a *Client) InstallEvents(params *InstallEventsParams, opts ...ClientOption) (*InstallEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInstallEventsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "InstallEvents",
		Method:             "GET",
		PathPattern:        "/install/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &InstallEventsReader{formats: a.formats},
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
	success, ok := result.(*InstallEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InstallEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PerformInstallCheck performs a check during install like a valid d b connection
*/
func (a *Client) PerformInstallCheck(params *PerformInstallCheckParams, opts ...ClientOption) (*PerformInstallCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformInstallCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PerformInstallCheck",
		Method:             "POST",
		PathPattern:        "/install/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PerformInstallCheckReader{formats: a.formats},
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
	success, ok := result.(*PerformInstallCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PerformInstallCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostInstall posts values to be saved for install
*/
func (a *Client) PostInstall(params *PostInstallParams, opts ...ClientOption) (*PostInstallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostInstallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostInstall",
		Method:             "POST",
		PathPattern:        "/install",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PostInstallReader{formats: a.formats},
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
	success, ok := result.(*PostInstallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostInstall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
