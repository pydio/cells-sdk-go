// Code generated by go-swagger; DO NOT EDIT.

package tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewHeadNodeParams creates a new HeadNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHeadNodeParams() *HeadNodeParams {
	return &HeadNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHeadNodeParamsWithTimeout creates a new HeadNodeParams object
// with the ability to set a timeout on a request.
func NewHeadNodeParamsWithTimeout(timeout time.Duration) *HeadNodeParams {
	return &HeadNodeParams{
		timeout: timeout,
	}
}

// NewHeadNodeParamsWithContext creates a new HeadNodeParams object
// with the ability to set a context for a request.
func NewHeadNodeParamsWithContext(ctx context.Context) *HeadNodeParams {
	return &HeadNodeParams{
		Context: ctx,
	}
}

// NewHeadNodeParamsWithHTTPClient creates a new HeadNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewHeadNodeParamsWithHTTPClient(client *http.Client) *HeadNodeParams {
	return &HeadNodeParams{
		HTTPClient: client,
	}
}

/*
HeadNodeParams contains all the parameters to send to the API endpoint

	for the head node operation.

	Typically these are written to a http.Request.
*/
type HeadNodeParams struct {

	/* Node.

	   The node to state
	*/
	Node string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the head node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadNodeParams) WithDefaults() *HeadNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the head node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HeadNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the head node params
func (o *HeadNodeParams) WithTimeout(timeout time.Duration) *HeadNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the head node params
func (o *HeadNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the head node params
func (o *HeadNodeParams) WithContext(ctx context.Context) *HeadNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the head node params
func (o *HeadNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the head node params
func (o *HeadNodeParams) WithHTTPClient(client *http.Client) *HeadNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the head node params
func (o *HeadNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNode adds the node to the head node params
func (o *HeadNodeParams) WithNode(node string) *HeadNodeParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the head node params
func (o *HeadNodeParams) SetNode(node string) {
	o.Node = node
}

// WriteToRequest writes these params to a swagger request
func (o *HeadNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Node
	if err := r.SetPathParam("Node", o.Node); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
