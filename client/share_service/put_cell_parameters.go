// Code generated by go-swagger; DO NOT EDIT.

package share_service

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

// NewPutCellParams creates a new PutCellParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutCellParams() *PutCellParams {
	return &PutCellParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutCellParamsWithTimeout creates a new PutCellParams object
// with the ability to set a timeout on a request.
func NewPutCellParamsWithTimeout(timeout time.Duration) *PutCellParams {
	return &PutCellParams{
		timeout: timeout,
	}
}

// NewPutCellParamsWithContext creates a new PutCellParams object
// with the ability to set a context for a request.
func NewPutCellParamsWithContext(ctx context.Context) *PutCellParams {
	return &PutCellParams{
		Context: ctx,
	}
}

// NewPutCellParamsWithHTTPClient creates a new PutCellParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutCellParamsWithHTTPClient(client *http.Client) *PutCellParams {
	return &PutCellParams{
		HTTPClient: client,
	}
}

/*
PutCellParams contains all the parameters to send to the API endpoint

	for the put cell operation.

	Typically these are written to a http.Request.
*/
type PutCellParams struct {

	// Body.
	Body *models.RestPutCellRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put cell params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCellParams) WithDefaults() *PutCellParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put cell params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutCellParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put cell params
func (o *PutCellParams) WithTimeout(timeout time.Duration) *PutCellParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cell params
func (o *PutCellParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cell params
func (o *PutCellParams) WithContext(ctx context.Context) *PutCellParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cell params
func (o *PutCellParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cell params
func (o *PutCellParams) WithHTTPClient(client *http.Client) *PutCellParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cell params
func (o *PutCellParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put cell params
func (o *PutCellParams) WithBody(body *models.RestPutCellRequest) *PutCellParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put cell params
func (o *PutCellParams) SetBody(body *models.RestPutCellRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutCellParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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