// Code generated by go-swagger; DO NOT EDIT.

package activity_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new activity service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for activity service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SearchSubscriptions(params *SearchSubscriptionsParams, opts ...ClientOption) (*SearchSubscriptionsOK, error)

	Stream(params *StreamParams, opts ...ClientOption) (*StreamOK, error)

	Subscribe(params *SubscribeParams, opts ...ClientOption) (*SubscribeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SearchSubscriptions loads subscriptions to other users nodes feeds
*/
func (a *Client) SearchSubscriptions(params *SearchSubscriptionsParams, opts ...ClientOption) (*SearchSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchSubscriptionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchSubscriptions",
		Method:             "POST",
		PathPattern:        "/activity/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SearchSubscriptionsReader{formats: a.formats},
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
	success, ok := result.(*SearchSubscriptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SearchSubscriptions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Stream loads the the feeds of the currently logged user
*/
func (a *Client) Stream(params *StreamParams, opts ...ClientOption) (*StreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Stream",
		Method:             "POST",
		PathPattern:        "/activity/stream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &StreamReader{formats: a.formats},
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
	success, ok := result.(*StreamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Stream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Subscribe manages subscriptions to other users nodes feeds
*/
func (a *Client) Subscribe(params *SubscribeParams, opts ...ClientOption) (*SubscribeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscribeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Subscribe",
		Method:             "POST",
		PathPattern:        "/activity/subscribe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SubscribeReader{formats: a.formats},
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
	success, ok := result.(*SubscribeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Subscribe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}