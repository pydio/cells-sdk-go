// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

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

// NewFrontPluginsParams creates a new FrontPluginsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFrontPluginsParams() *FrontPluginsParams {
	return &FrontPluginsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFrontPluginsParamsWithTimeout creates a new FrontPluginsParams object
// with the ability to set a timeout on a request.
func NewFrontPluginsParamsWithTimeout(timeout time.Duration) *FrontPluginsParams {
	return &FrontPluginsParams{
		timeout: timeout,
	}
}

// NewFrontPluginsParamsWithContext creates a new FrontPluginsParams object
// with the ability to set a context for a request.
func NewFrontPluginsParamsWithContext(ctx context.Context) *FrontPluginsParams {
	return &FrontPluginsParams{
		Context: ctx,
	}
}

// NewFrontPluginsParamsWithHTTPClient creates a new FrontPluginsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFrontPluginsParamsWithHTTPClient(client *http.Client) *FrontPluginsParams {
	return &FrontPluginsParams{
		HTTPClient: client,
	}
}

/*
FrontPluginsParams contains all the parameters to send to the API endpoint

	for the front plugins operation.

	Typically these are written to a http.Request.
*/
type FrontPluginsParams struct {

	// Lang.
	Lang string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the front plugins params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FrontPluginsParams) WithDefaults() *FrontPluginsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the front plugins params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FrontPluginsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the front plugins params
func (o *FrontPluginsParams) WithTimeout(timeout time.Duration) *FrontPluginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the front plugins params
func (o *FrontPluginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the front plugins params
func (o *FrontPluginsParams) WithContext(ctx context.Context) *FrontPluginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the front plugins params
func (o *FrontPluginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the front plugins params
func (o *FrontPluginsParams) WithHTTPClient(client *http.Client) *FrontPluginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the front plugins params
func (o *FrontPluginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLang adds the lang to the front plugins params
func (o *FrontPluginsParams) WithLang(lang string) *FrontPluginsParams {
	o.SetLang(lang)
	return o
}

// SetLang adds the lang to the front plugins params
func (o *FrontPluginsParams) SetLang(lang string) {
	o.Lang = lang
}

// WriteToRequest writes these params to a swagger request
func (o *FrontPluginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Lang
	if err := r.SetPathParam("Lang", o.Lang); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
