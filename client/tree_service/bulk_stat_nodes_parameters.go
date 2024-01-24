// Code generated by go-swagger; DO NOT EDIT.

package tree_service

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

	"github.com/pydio/cells-sdk-go/v5/models"
)

// NewBulkStatNodesParams creates a new BulkStatNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBulkStatNodesParams() *BulkStatNodesParams {
	return &BulkStatNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBulkStatNodesParamsWithTimeout creates a new BulkStatNodesParams object
// with the ability to set a timeout on a request.
func NewBulkStatNodesParamsWithTimeout(timeout time.Duration) *BulkStatNodesParams {
	return &BulkStatNodesParams{
		timeout: timeout,
	}
}

// NewBulkStatNodesParamsWithContext creates a new BulkStatNodesParams object
// with the ability to set a context for a request.
func NewBulkStatNodesParamsWithContext(ctx context.Context) *BulkStatNodesParams {
	return &BulkStatNodesParams{
		Context: ctx,
	}
}

// NewBulkStatNodesParamsWithHTTPClient creates a new BulkStatNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewBulkStatNodesParamsWithHTTPClient(client *http.Client) *BulkStatNodesParams {
	return &BulkStatNodesParams{
		HTTPClient: client,
	}
}

/*
BulkStatNodesParams contains all the parameters to send to the API endpoint

	for the bulk stat nodes operation.

	Typically these are written to a http.Request.
*/
type BulkStatNodesParams struct {

	// Body.
	Body *models.RestGetBulkMetaRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the bulk stat nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkStatNodesParams) WithDefaults() *BulkStatNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the bulk stat nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BulkStatNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the bulk stat nodes params
func (o *BulkStatNodesParams) WithTimeout(timeout time.Duration) *BulkStatNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bulk stat nodes params
func (o *BulkStatNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bulk stat nodes params
func (o *BulkStatNodesParams) WithContext(ctx context.Context) *BulkStatNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bulk stat nodes params
func (o *BulkStatNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bulk stat nodes params
func (o *BulkStatNodesParams) WithHTTPClient(client *http.Client) *BulkStatNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bulk stat nodes params
func (o *BulkStatNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the bulk stat nodes params
func (o *BulkStatNodesParams) WithBody(body *models.RestGetBulkMetaRequest) *BulkStatNodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bulk stat nodes params
func (o *BulkStatNodesParams) SetBody(body *models.RestGetBulkMetaRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BulkStatNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
