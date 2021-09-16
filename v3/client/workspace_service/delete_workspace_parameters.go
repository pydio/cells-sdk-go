// Code generated by go-swagger; DO NOT EDIT.

package workspace_service

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

// NewDeleteWorkspaceParams creates a new DeleteWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWorkspaceParams() *DeleteWorkspaceParams {
	return &DeleteWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkspaceParamsWithTimeout creates a new DeleteWorkspaceParams object
// with the ability to set a timeout on a request.
func NewDeleteWorkspaceParamsWithTimeout(timeout time.Duration) *DeleteWorkspaceParams {
	return &DeleteWorkspaceParams{
		timeout: timeout,
	}
}

// NewDeleteWorkspaceParamsWithContext creates a new DeleteWorkspaceParams object
// with the ability to set a context for a request.
func NewDeleteWorkspaceParamsWithContext(ctx context.Context) *DeleteWorkspaceParams {
	return &DeleteWorkspaceParams{
		Context: ctx,
	}
}

// NewDeleteWorkspaceParamsWithHTTPClient creates a new DeleteWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWorkspaceParamsWithHTTPClient(client *http.Client) *DeleteWorkspaceParams {
	return &DeleteWorkspaceParams{
		HTTPClient: client,
	}
}

/* DeleteWorkspaceParams contains all the parameters to send to the API endpoint
   for the delete workspace operation.

   Typically these are written to a http.Request.
*/
type DeleteWorkspaceParams struct {

	// Slug.
	Slug string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkspaceParams) WithDefaults() *DeleteWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete workspace params
func (o *DeleteWorkspaceParams) WithTimeout(timeout time.Duration) *DeleteWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workspace params
func (o *DeleteWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workspace params
func (o *DeleteWorkspaceParams) WithContext(ctx context.Context) *DeleteWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workspace params
func (o *DeleteWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workspace params
func (o *DeleteWorkspaceParams) WithHTTPClient(client *http.Client) *DeleteWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workspace params
func (o *DeleteWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSlug adds the slug to the delete workspace params
func (o *DeleteWorkspaceParams) WithSlug(slug string) *DeleteWorkspaceParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete workspace params
func (o *DeleteWorkspaceParams) SetSlug(slug string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Slug
	if err := r.SetPathParam("Slug", o.Slug); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
