// Code generated by go-swagger; DO NOT EDIT.

package meta_service

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

// NewSetMetaParams creates a new SetMetaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetMetaParams() *SetMetaParams {
	return &SetMetaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetMetaParamsWithTimeout creates a new SetMetaParams object
// with the ability to set a timeout on a request.
func NewSetMetaParamsWithTimeout(timeout time.Duration) *SetMetaParams {
	return &SetMetaParams{
		timeout: timeout,
	}
}

// NewSetMetaParamsWithContext creates a new SetMetaParams object
// with the ability to set a context for a request.
func NewSetMetaParamsWithContext(ctx context.Context) *SetMetaParams {
	return &SetMetaParams{
		Context: ctx,
	}
}

// NewSetMetaParamsWithHTTPClient creates a new SetMetaParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetMetaParamsWithHTTPClient(client *http.Client) *SetMetaParams {
	return &SetMetaParams{
		HTTPClient: client,
	}
}

/*
SetMetaParams contains all the parameters to send to the API endpoint

	for the set meta operation.

	Typically these are written to a http.Request.
*/
type SetMetaParams struct {

	// NodePath.
	NodePath string

	// Body.
	Body SetMetaBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetMetaParams) WithDefaults() *SetMetaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetMetaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set meta params
func (o *SetMetaParams) WithTimeout(timeout time.Duration) *SetMetaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set meta params
func (o *SetMetaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set meta params
func (o *SetMetaParams) WithContext(ctx context.Context) *SetMetaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set meta params
func (o *SetMetaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set meta params
func (o *SetMetaParams) WithHTTPClient(client *http.Client) *SetMetaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set meta params
func (o *SetMetaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodePath adds the nodePath to the set meta params
func (o *SetMetaParams) WithNodePath(nodePath string) *SetMetaParams {
	o.SetNodePath(nodePath)
	return o
}

// SetNodePath adds the nodePath to the set meta params
func (o *SetMetaParams) SetNodePath(nodePath string) {
	o.NodePath = nodePath
}

// WithBody adds the body to the set meta params
func (o *SetMetaParams) WithBody(body SetMetaBody) *SetMetaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set meta params
func (o *SetMetaParams) SetBody(body SetMetaBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetMetaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param NodePath
	if err := r.SetPathParam("NodePath", o.NodePath); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
