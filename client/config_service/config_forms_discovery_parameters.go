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

// NewConfigFormsDiscoveryParams creates a new ConfigFormsDiscoveryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfigFormsDiscoveryParams() *ConfigFormsDiscoveryParams {
	return &ConfigFormsDiscoveryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfigFormsDiscoveryParamsWithTimeout creates a new ConfigFormsDiscoveryParams object
// with the ability to set a timeout on a request.
func NewConfigFormsDiscoveryParamsWithTimeout(timeout time.Duration) *ConfigFormsDiscoveryParams {
	return &ConfigFormsDiscoveryParams{
		timeout: timeout,
	}
}

// NewConfigFormsDiscoveryParamsWithContext creates a new ConfigFormsDiscoveryParams object
// with the ability to set a context for a request.
func NewConfigFormsDiscoveryParamsWithContext(ctx context.Context) *ConfigFormsDiscoveryParams {
	return &ConfigFormsDiscoveryParams{
		Context: ctx,
	}
}

// NewConfigFormsDiscoveryParamsWithHTTPClient creates a new ConfigFormsDiscoveryParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfigFormsDiscoveryParamsWithHTTPClient(client *http.Client) *ConfigFormsDiscoveryParams {
	return &ConfigFormsDiscoveryParams{
		HTTPClient: client,
	}
}

/*
ConfigFormsDiscoveryParams contains all the parameters to send to the API endpoint

	for the config forms discovery operation.

	Typically these are written to a http.Request.
*/
type ConfigFormsDiscoveryParams struct {

	/* ServiceName.

	   Retrieve a configuration form for a given service
	*/
	ServiceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the config forms discovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigFormsDiscoveryParams) WithDefaults() *ConfigFormsDiscoveryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the config forms discovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigFormsDiscoveryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the config forms discovery params
func (o *ConfigFormsDiscoveryParams) WithTimeout(timeout time.Duration) *ConfigFormsDiscoveryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the config forms discovery params
func (o *ConfigFormsDiscoveryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the config forms discovery params
func (o *ConfigFormsDiscoveryParams) WithContext(ctx context.Context) *ConfigFormsDiscoveryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the config forms discovery params
func (o *ConfigFormsDiscoveryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the config forms discovery params
func (o *ConfigFormsDiscoveryParams) WithHTTPClient(client *http.Client) *ConfigFormsDiscoveryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the config forms discovery params
func (o *ConfigFormsDiscoveryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceName adds the serviceName to the config forms discovery params
func (o *ConfigFormsDiscoveryParams) WithServiceName(serviceName string) *ConfigFormsDiscoveryParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the config forms discovery params
func (o *ConfigFormsDiscoveryParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigFormsDiscoveryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ServiceName
	if err := r.SetPathParam("ServiceName", o.ServiceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
