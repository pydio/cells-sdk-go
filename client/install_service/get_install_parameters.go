// Code generated by go-swagger; DO NOT EDIT.

package install_service

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

// NewGetInstallParams creates a new GetInstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInstallParams() *GetInstallParams {
	return &GetInstallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstallParamsWithTimeout creates a new GetInstallParams object
// with the ability to set a timeout on a request.
func NewGetInstallParamsWithTimeout(timeout time.Duration) *GetInstallParams {
	return &GetInstallParams{
		timeout: timeout,
	}
}

// NewGetInstallParamsWithContext creates a new GetInstallParams object
// with the ability to set a context for a request.
func NewGetInstallParamsWithContext(ctx context.Context) *GetInstallParams {
	return &GetInstallParams{
		Context: ctx,
	}
}

// NewGetInstallParamsWithHTTPClient creates a new GetInstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInstallParamsWithHTTPClient(client *http.Client) *GetInstallParams {
	return &GetInstallParams{
		HTTPClient: client,
	}
}

/*
GetInstallParams contains all the parameters to send to the API endpoint

	for the get install operation.

	Typically these are written to a http.Request.
*/
type GetInstallParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallParams) WithDefaults() *GetInstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInstallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get install params
func (o *GetInstallParams) WithTimeout(timeout time.Duration) *GetInstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get install params
func (o *GetInstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get install params
func (o *GetInstallParams) WithContext(ctx context.Context) *GetInstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get install params
func (o *GetInstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get install params
func (o *GetInstallParams) WithHTTPClient(client *http.Client) *GetInstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get install params
func (o *GetInstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
