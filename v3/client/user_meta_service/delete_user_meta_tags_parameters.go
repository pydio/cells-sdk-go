// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

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

// NewDeleteUserMetaTagsParams creates a new DeleteUserMetaTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUserMetaTagsParams() *DeleteUserMetaTagsParams {
	return &DeleteUserMetaTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserMetaTagsParamsWithTimeout creates a new DeleteUserMetaTagsParams object
// with the ability to set a timeout on a request.
func NewDeleteUserMetaTagsParamsWithTimeout(timeout time.Duration) *DeleteUserMetaTagsParams {
	return &DeleteUserMetaTagsParams{
		timeout: timeout,
	}
}

// NewDeleteUserMetaTagsParamsWithContext creates a new DeleteUserMetaTagsParams object
// with the ability to set a context for a request.
func NewDeleteUserMetaTagsParamsWithContext(ctx context.Context) *DeleteUserMetaTagsParams {
	return &DeleteUserMetaTagsParams{
		Context: ctx,
	}
}

// NewDeleteUserMetaTagsParamsWithHTTPClient creates a new DeleteUserMetaTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUserMetaTagsParamsWithHTTPClient(client *http.Client) *DeleteUserMetaTagsParams {
	return &DeleteUserMetaTagsParams{
		HTTPClient: client,
	}
}

/* DeleteUserMetaTagsParams contains all the parameters to send to the API endpoint
   for the delete user meta tags operation.

   Typically these are written to a http.Request.
*/
type DeleteUserMetaTagsParams struct {

	// Namespace.
	Namespace string

	// Tags.
	Tags string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete user meta tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserMetaTagsParams) WithDefaults() *DeleteUserMetaTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete user meta tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUserMetaTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) WithTimeout(timeout time.Duration) *DeleteUserMetaTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) WithContext(ctx context.Context) *DeleteUserMetaTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) WithHTTPClient(client *http.Client) *DeleteUserMetaTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) WithNamespace(namespace string) *DeleteUserMetaTagsParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithTags adds the tags to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) WithTags(tags string) *DeleteUserMetaTagsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the delete user meta tags params
func (o *DeleteUserMetaTagsParams) SetTags(tags string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserMetaTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Namespace
	if err := r.SetPathParam("Namespace", o.Namespace); err != nil {
		return err
	}

	// path param Tags
	if err := r.SetPathParam("Tags", o.Tags); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
