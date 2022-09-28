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

	"github.com/pydio/cells-sdk-go/v4/models"
)

// NewListProcessesParams creates a new ListProcessesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProcessesParams() *ListProcessesParams {
	return &ListProcessesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProcessesParamsWithTimeout creates a new ListProcessesParams object
// with the ability to set a timeout on a request.
func NewListProcessesParamsWithTimeout(timeout time.Duration) *ListProcessesParams {
	return &ListProcessesParams{
		timeout: timeout,
	}
}

// NewListProcessesParamsWithContext creates a new ListProcessesParams object
// with the ability to set a context for a request.
func NewListProcessesParamsWithContext(ctx context.Context) *ListProcessesParams {
	return &ListProcessesParams{
		Context: ctx,
	}
}

// NewListProcessesParamsWithHTTPClient creates a new ListProcessesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProcessesParamsWithHTTPClient(client *http.Client) *ListProcessesParams {
	return &ListProcessesParams{
		HTTPClient: client,
	}
}

/*
ListProcessesParams contains all the parameters to send to the API endpoint

	for the list processes operation.

	Typically these are written to a http.Request.
*/
type ListProcessesParams struct {

	// Body.
	Body *models.RestListProcessesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list processes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProcessesParams) WithDefaults() *ListProcessesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list processes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProcessesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list processes params
func (o *ListProcessesParams) WithTimeout(timeout time.Duration) *ListProcessesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list processes params
func (o *ListProcessesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list processes params
func (o *ListProcessesParams) WithContext(ctx context.Context) *ListProcessesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list processes params
func (o *ListProcessesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list processes params
func (o *ListProcessesParams) WithHTTPClient(client *http.Client) *ListProcessesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list processes params
func (o *ListProcessesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list processes params
func (o *ListProcessesParams) WithBody(body *models.RestListProcessesRequest) *ListProcessesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list processes params
func (o *ListProcessesParams) SetBody(body *models.RestListProcessesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListProcessesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}