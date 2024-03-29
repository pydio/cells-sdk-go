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
)

// NewGetCellParams creates a new GetCellParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCellParams() *GetCellParams {
	return &GetCellParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCellParamsWithTimeout creates a new GetCellParams object
// with the ability to set a timeout on a request.
func NewGetCellParamsWithTimeout(timeout time.Duration) *GetCellParams {
	return &GetCellParams{
		timeout: timeout,
	}
}

// NewGetCellParamsWithContext creates a new GetCellParams object
// with the ability to set a context for a request.
func NewGetCellParamsWithContext(ctx context.Context) *GetCellParams {
	return &GetCellParams{
		Context: ctx,
	}
}

// NewGetCellParamsWithHTTPClient creates a new GetCellParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCellParamsWithHTTPClient(client *http.Client) *GetCellParams {
	return &GetCellParams{
		HTTPClient: client,
	}
}

/*
GetCellParams contains all the parameters to send to the API endpoint

	for the get cell operation.

	Typically these are written to a http.Request.
*/
type GetCellParams struct {

	/* UUID.

	   Cell Uuid
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cell params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCellParams) WithDefaults() *GetCellParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cell params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCellParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cell params
func (o *GetCellParams) WithTimeout(timeout time.Duration) *GetCellParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cell params
func (o *GetCellParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cell params
func (o *GetCellParams) WithContext(ctx context.Context) *GetCellParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cell params
func (o *GetCellParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cell params
func (o *GetCellParams) WithHTTPClient(client *http.Client) *GetCellParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cell params
func (o *GetCellParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the get cell params
func (o *GetCellParams) WithUUID(uuid string) *GetCellParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get cell params
func (o *GetCellParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetCellParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Uuid
	if err := r.SetPathParam("Uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
