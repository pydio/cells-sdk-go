// Code generated by go-swagger; DO NOT EDIT.

package config_service

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

// NewListVirtualNodesParams creates a new ListVirtualNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVirtualNodesParams() *ListVirtualNodesParams {
	return &ListVirtualNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVirtualNodesParamsWithTimeout creates a new ListVirtualNodesParams object
// with the ability to set a timeout on a request.
func NewListVirtualNodesParamsWithTimeout(timeout time.Duration) *ListVirtualNodesParams {
	return &ListVirtualNodesParams{
		timeout: timeout,
	}
}

// NewListVirtualNodesParamsWithContext creates a new ListVirtualNodesParams object
// with the ability to set a context for a request.
func NewListVirtualNodesParamsWithContext(ctx context.Context) *ListVirtualNodesParams {
	return &ListVirtualNodesParams{
		Context: ctx,
	}
}

// NewListVirtualNodesParamsWithHTTPClient creates a new ListVirtualNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVirtualNodesParamsWithHTTPClient(client *http.Client) *ListVirtualNodesParams {
	return &ListVirtualNodesParams{
		HTTPClient: client,
	}
}

/*
ListVirtualNodesParams contains all the parameters to send to the API endpoint

	for the list virtual nodes operation.

	Typically these are written to a http.Request.
*/
type ListVirtualNodesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list virtual nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVirtualNodesParams) WithDefaults() *ListVirtualNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list virtual nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVirtualNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list virtual nodes params
func (o *ListVirtualNodesParams) WithTimeout(timeout time.Duration) *ListVirtualNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list virtual nodes params
func (o *ListVirtualNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list virtual nodes params
func (o *ListVirtualNodesParams) WithContext(ctx context.Context) *ListVirtualNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list virtual nodes params
func (o *ListVirtualNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list virtual nodes params
func (o *ListVirtualNodesParams) WithHTTPClient(client *http.Client) *ListVirtualNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list virtual nodes params
func (o *ListVirtualNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListVirtualNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
