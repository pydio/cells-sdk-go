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

// NewPutShareLinkParams creates a new PutShareLinkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutShareLinkParams() *PutShareLinkParams {
	return &PutShareLinkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutShareLinkParamsWithTimeout creates a new PutShareLinkParams object
// with the ability to set a timeout on a request.
func NewPutShareLinkParamsWithTimeout(timeout time.Duration) *PutShareLinkParams {
	return &PutShareLinkParams{
		timeout: timeout,
	}
}

// NewPutShareLinkParamsWithContext creates a new PutShareLinkParams object
// with the ability to set a context for a request.
func NewPutShareLinkParamsWithContext(ctx context.Context) *PutShareLinkParams {
	return &PutShareLinkParams{
		Context: ctx,
	}
}

// NewPutShareLinkParamsWithHTTPClient creates a new PutShareLinkParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutShareLinkParamsWithHTTPClient(client *http.Client) *PutShareLinkParams {
	return &PutShareLinkParams{
		HTTPClient: client,
	}
}

/*
PutShareLinkParams contains all the parameters to send to the API endpoint

	for the put share link operation.

	Typically these are written to a http.Request.
*/
type PutShareLinkParams struct {

	// Body.
	Body *models.RestPutShareLinkRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put share link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutShareLinkParams) WithDefaults() *PutShareLinkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put share link params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutShareLinkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put share link params
func (o *PutShareLinkParams) WithTimeout(timeout time.Duration) *PutShareLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put share link params
func (o *PutShareLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put share link params
func (o *PutShareLinkParams) WithContext(ctx context.Context) *PutShareLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put share link params
func (o *PutShareLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put share link params
func (o *PutShareLinkParams) WithHTTPClient(client *http.Client) *PutShareLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put share link params
func (o *PutShareLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put share link params
func (o *PutShareLinkParams) WithBody(body *models.RestPutShareLinkRequest) *PutShareLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put share link params
func (o *PutShareLinkParams) SetBody(body *models.RestPutShareLinkRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutShareLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
