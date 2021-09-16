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

// NewListPeersAddressesParams creates a new ListPeersAddressesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPeersAddressesParams() *ListPeersAddressesParams {
	return &ListPeersAddressesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPeersAddressesParamsWithTimeout creates a new ListPeersAddressesParams object
// with the ability to set a timeout on a request.
func NewListPeersAddressesParamsWithTimeout(timeout time.Duration) *ListPeersAddressesParams {
	return &ListPeersAddressesParams{
		timeout: timeout,
	}
}

// NewListPeersAddressesParamsWithContext creates a new ListPeersAddressesParams object
// with the ability to set a context for a request.
func NewListPeersAddressesParamsWithContext(ctx context.Context) *ListPeersAddressesParams {
	return &ListPeersAddressesParams{
		Context: ctx,
	}
}

// NewListPeersAddressesParamsWithHTTPClient creates a new ListPeersAddressesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPeersAddressesParamsWithHTTPClient(client *http.Client) *ListPeersAddressesParams {
	return &ListPeersAddressesParams{
		HTTPClient: client,
	}
}

/* ListPeersAddressesParams contains all the parameters to send to the API endpoint
   for the list peers addresses operation.

   Typically these are written to a http.Request.
*/
type ListPeersAddressesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list peers addresses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPeersAddressesParams) WithDefaults() *ListPeersAddressesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list peers addresses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPeersAddressesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list peers addresses params
func (o *ListPeersAddressesParams) WithTimeout(timeout time.Duration) *ListPeersAddressesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list peers addresses params
func (o *ListPeersAddressesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list peers addresses params
func (o *ListPeersAddressesParams) WithContext(ctx context.Context) *ListPeersAddressesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list peers addresses params
func (o *ListPeersAddressesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list peers addresses params
func (o *ListPeersAddressesParams) WithHTTPClient(client *http.Client) *ListPeersAddressesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list peers addresses params
func (o *ListPeersAddressesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListPeersAddressesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
