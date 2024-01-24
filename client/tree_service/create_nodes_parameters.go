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

// NewCreateNodesParams creates a new CreateNodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNodesParams() *CreateNodesParams {
	return &CreateNodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNodesParamsWithTimeout creates a new CreateNodesParams object
// with the ability to set a timeout on a request.
func NewCreateNodesParamsWithTimeout(timeout time.Duration) *CreateNodesParams {
	return &CreateNodesParams{
		timeout: timeout,
	}
}

// NewCreateNodesParamsWithContext creates a new CreateNodesParams object
// with the ability to set a context for a request.
func NewCreateNodesParamsWithContext(ctx context.Context) *CreateNodesParams {
	return &CreateNodesParams{
		Context: ctx,
	}
}

// NewCreateNodesParamsWithHTTPClient creates a new CreateNodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNodesParamsWithHTTPClient(client *http.Client) *CreateNodesParams {
	return &CreateNodesParams{
		HTTPClient: client,
	}
}

/*
CreateNodesParams contains all the parameters to send to the API endpoint

	for the create nodes operation.

	Typically these are written to a http.Request.
*/
type CreateNodesParams struct {

	// Body.
	Body *models.RestCreateNodesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNodesParams) WithDefaults() *CreateNodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create nodes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create nodes params
func (o *CreateNodesParams) WithTimeout(timeout time.Duration) *CreateNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create nodes params
func (o *CreateNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create nodes params
func (o *CreateNodesParams) WithContext(ctx context.Context) *CreateNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create nodes params
func (o *CreateNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create nodes params
func (o *CreateNodesParams) WithHTTPClient(client *http.Client) *CreateNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create nodes params
func (o *CreateNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create nodes params
func (o *CreateNodesParams) WithBody(body *models.RestCreateNodesRequest) *CreateNodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create nodes params
func (o *CreateNodesParams) SetBody(body *models.RestCreateNodesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
