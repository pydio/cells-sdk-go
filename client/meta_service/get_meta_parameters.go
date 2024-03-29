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

// NewGetMetaParams creates a new GetMetaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMetaParams() *GetMetaParams {
	return &GetMetaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMetaParamsWithTimeout creates a new GetMetaParams object
// with the ability to set a timeout on a request.
func NewGetMetaParamsWithTimeout(timeout time.Duration) *GetMetaParams {
	return &GetMetaParams{
		timeout: timeout,
	}
}

// NewGetMetaParamsWithContext creates a new GetMetaParams object
// with the ability to set a context for a request.
func NewGetMetaParamsWithContext(ctx context.Context) *GetMetaParams {
	return &GetMetaParams{
		Context: ctx,
	}
}

// NewGetMetaParamsWithHTTPClient creates a new GetMetaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMetaParamsWithHTTPClient(client *http.Client) *GetMetaParams {
	return &GetMetaParams{
		HTTPClient: client,
	}
}

/*
GetMetaParams contains all the parameters to send to the API endpoint

	for the get meta operation.

	Typically these are written to a http.Request.
*/
type GetMetaParams struct {

	/* NodePath.

	   Path to the requested node
	*/
	NodePath string

	// Body.
	Body GetMetaBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMetaParams) WithDefaults() *GetMetaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMetaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get meta params
func (o *GetMetaParams) WithTimeout(timeout time.Duration) *GetMetaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get meta params
func (o *GetMetaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get meta params
func (o *GetMetaParams) WithContext(ctx context.Context) *GetMetaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get meta params
func (o *GetMetaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get meta params
func (o *GetMetaParams) WithHTTPClient(client *http.Client) *GetMetaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get meta params
func (o *GetMetaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodePath adds the nodePath to the get meta params
func (o *GetMetaParams) WithNodePath(nodePath string) *GetMetaParams {
	o.SetNodePath(nodePath)
	return o
}

// SetNodePath adds the nodePath to the get meta params
func (o *GetMetaParams) SetNodePath(nodePath string) {
	o.NodePath = nodePath
}

// WithBody adds the body to the get meta params
func (o *GetMetaParams) WithBody(body GetMetaBody) *GetMetaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get meta params
func (o *GetMetaParams) SetBody(body GetMetaBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetMetaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
