// Code generated by go-swagger; DO NOT EDIT.

package acl_service

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

	"github.com/pydio/cells-sdk-go/v3/models"
)

// NewPutACLParams creates a new PutACLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutACLParams() *PutACLParams {
	return &PutACLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutACLParamsWithTimeout creates a new PutACLParams object
// with the ability to set a timeout on a request.
func NewPutACLParamsWithTimeout(timeout time.Duration) *PutACLParams {
	return &PutACLParams{
		timeout: timeout,
	}
}

// NewPutACLParamsWithContext creates a new PutACLParams object
// with the ability to set a context for a request.
func NewPutACLParamsWithContext(ctx context.Context) *PutACLParams {
	return &PutACLParams{
		Context: ctx,
	}
}

// NewPutACLParamsWithHTTPClient creates a new PutACLParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutACLParamsWithHTTPClient(client *http.Client) *PutACLParams {
	return &PutACLParams{
		HTTPClient: client,
	}
}

/* PutACLParams contains all the parameters to send to the API endpoint
   for the put Acl operation.

   Typically these are written to a http.Request.
*/
type PutACLParams struct {

	// Body.
	Body *models.IdmACL

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutACLParams) WithDefaults() *PutACLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutACLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put Acl params
func (o *PutACLParams) WithTimeout(timeout time.Duration) *PutACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put Acl params
func (o *PutACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put Acl params
func (o *PutACLParams) WithContext(ctx context.Context) *PutACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put Acl params
func (o *PutACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put Acl params
func (o *PutACLParams) WithHTTPClient(client *http.Client) *PutACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put Acl params
func (o *PutACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put Acl params
func (o *PutACLParams) WithBody(body *models.IdmACL) *PutACLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put Acl params
func (o *PutACLParams) SetBody(body *models.IdmACL) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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