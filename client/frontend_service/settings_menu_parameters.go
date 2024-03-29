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

// NewSettingsMenuParams creates a new SettingsMenuParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSettingsMenuParams() *SettingsMenuParams {
	return &SettingsMenuParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSettingsMenuParamsWithTimeout creates a new SettingsMenuParams object
// with the ability to set a timeout on a request.
func NewSettingsMenuParamsWithTimeout(timeout time.Duration) *SettingsMenuParams {
	return &SettingsMenuParams{
		timeout: timeout,
	}
}

// NewSettingsMenuParamsWithContext creates a new SettingsMenuParams object
// with the ability to set a context for a request.
func NewSettingsMenuParamsWithContext(ctx context.Context) *SettingsMenuParams {
	return &SettingsMenuParams{
		Context: ctx,
	}
}

// NewSettingsMenuParamsWithHTTPClient creates a new SettingsMenuParams object
// with the ability to set a custom HTTPClient for a request.
func NewSettingsMenuParamsWithHTTPClient(client *http.Client) *SettingsMenuParams {
	return &SettingsMenuParams{
		HTTPClient: client,
	}
}

/*
SettingsMenuParams contains all the parameters to send to the API endpoint

	for the settings menu operation.

	Typically these are written to a http.Request.
*/
type SettingsMenuParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the settings menu params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsMenuParams) WithDefaults() *SettingsMenuParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the settings menu params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SettingsMenuParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the settings menu params
func (o *SettingsMenuParams) WithTimeout(timeout time.Duration) *SettingsMenuParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the settings menu params
func (o *SettingsMenuParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the settings menu params
func (o *SettingsMenuParams) WithContext(ctx context.Context) *SettingsMenuParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the settings menu params
func (o *SettingsMenuParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the settings menu params
func (o *SettingsMenuParams) WithHTTPClient(client *http.Client) *SettingsMenuParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the settings menu params
func (o *SettingsMenuParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SettingsMenuParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
