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

// NewDeleteACLParams creates a new DeleteACLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteACLParams() *DeleteACLParams {
	return &DeleteACLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteACLParamsWithTimeout creates a new DeleteACLParams object
// with the ability to set a timeout on a request.
func NewDeleteACLParamsWithTimeout(timeout time.Duration) *DeleteACLParams {
	return &DeleteACLParams{
		timeout: timeout,
	}
}

// NewDeleteACLParamsWithContext creates a new DeleteACLParams object
// with the ability to set a context for a request.
func NewDeleteACLParamsWithContext(ctx context.Context) *DeleteACLParams {
	return &DeleteACLParams{
		Context: ctx,
	}
}

// NewDeleteACLParamsWithHTTPClient creates a new DeleteACLParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteACLParamsWithHTTPClient(client *http.Client) *DeleteACLParams {
	return &DeleteACLParams{
		HTTPClient: client,
	}
}

/* DeleteACLParams contains all the parameters to send to the API endpoint
   for the delete Acl operation.

   Typically these are written to a http.Request.
*/
type DeleteACLParams struct {

	// Body.
	Body *models.IdmACL

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteACLParams) WithDefaults() *DeleteACLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete Acl params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteACLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete Acl params
func (o *DeleteACLParams) WithTimeout(timeout time.Duration) *DeleteACLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Acl params
func (o *DeleteACLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Acl params
func (o *DeleteACLParams) WithContext(ctx context.Context) *DeleteACLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Acl params
func (o *DeleteACLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Acl params
func (o *DeleteACLParams) WithHTTPClient(client *http.Client) *DeleteACLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Acl params
func (o *DeleteACLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete Acl params
func (o *DeleteACLParams) WithBody(body *models.IdmACL) *DeleteACLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete Acl params
func (o *DeleteACLParams) SetBody(body *models.IdmACL) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteACLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
