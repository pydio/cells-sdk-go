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

// NewListSitesParams creates a new ListSitesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSitesParams() *ListSitesParams {
	return &ListSitesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSitesParamsWithTimeout creates a new ListSitesParams object
// with the ability to set a timeout on a request.
func NewListSitesParamsWithTimeout(timeout time.Duration) *ListSitesParams {
	return &ListSitesParams{
		timeout: timeout,
	}
}

// NewListSitesParamsWithContext creates a new ListSitesParams object
// with the ability to set a context for a request.
func NewListSitesParamsWithContext(ctx context.Context) *ListSitesParams {
	return &ListSitesParams{
		Context: ctx,
	}
}

// NewListSitesParamsWithHTTPClient creates a new ListSitesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSitesParamsWithHTTPClient(client *http.Client) *ListSitesParams {
	return &ListSitesParams{
		HTTPClient: client,
	}
}

/*
ListSitesParams contains all the parameters to send to the API endpoint

	for the list sites operation.

	Typically these are written to a http.Request.
*/
type ListSitesParams struct {

	// Filter.
	Filter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list sites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSitesParams) WithDefaults() *ListSitesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list sites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSitesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list sites params
func (o *ListSitesParams) WithTimeout(timeout time.Duration) *ListSitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sites params
func (o *ListSitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sites params
func (o *ListSitesParams) WithContext(ctx context.Context) *ListSitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sites params
func (o *ListSitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sites params
func (o *ListSitesParams) WithHTTPClient(client *http.Client) *ListSitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sites params
func (o *ListSitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilter adds the filter to the list sites params
func (o *ListSitesParams) WithFilter(filter string) *ListSitesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the list sites params
func (o *ListSitesParams) SetFilter(filter string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *ListSitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Filter
	if err := r.SetPathParam("Filter", o.Filter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}