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

	"github.com/pydio/cells-sdk-go/v3/models"
)

// NewListSharedResourcesParams creates a new ListSharedResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListSharedResourcesParams() *ListSharedResourcesParams {
	return &ListSharedResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListSharedResourcesParamsWithTimeout creates a new ListSharedResourcesParams object
// with the ability to set a timeout on a request.
func NewListSharedResourcesParamsWithTimeout(timeout time.Duration) *ListSharedResourcesParams {
	return &ListSharedResourcesParams{
		timeout: timeout,
	}
}

// NewListSharedResourcesParamsWithContext creates a new ListSharedResourcesParams object
// with the ability to set a context for a request.
func NewListSharedResourcesParamsWithContext(ctx context.Context) *ListSharedResourcesParams {
	return &ListSharedResourcesParams{
		Context: ctx,
	}
}

// NewListSharedResourcesParamsWithHTTPClient creates a new ListSharedResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListSharedResourcesParamsWithHTTPClient(client *http.Client) *ListSharedResourcesParams {
	return &ListSharedResourcesParams{
		HTTPClient: client,
	}
}

/* ListSharedResourcesParams contains all the parameters to send to the API endpoint
   for the list shared resources operation.

   Typically these are written to a http.Request.
*/
type ListSharedResourcesParams struct {

	// Body.
	Body *models.RestListSharedResourcesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list shared resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSharedResourcesParams) WithDefaults() *ListSharedResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list shared resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListSharedResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list shared resources params
func (o *ListSharedResourcesParams) WithTimeout(timeout time.Duration) *ListSharedResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list shared resources params
func (o *ListSharedResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list shared resources params
func (o *ListSharedResourcesParams) WithContext(ctx context.Context) *ListSharedResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list shared resources params
func (o *ListSharedResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list shared resources params
func (o *ListSharedResourcesParams) WithHTTPClient(client *http.Client) *ListSharedResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list shared resources params
func (o *ListSharedResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list shared resources params
func (o *ListSharedResourcesParams) WithBody(body *models.RestListSharedResourcesRequest) *ListSharedResourcesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list shared resources params
func (o *ListSharedResourcesParams) SetBody(body *models.RestListSharedResourcesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListSharedResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
