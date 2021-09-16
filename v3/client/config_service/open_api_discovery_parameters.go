// Code generated by go-swagger; DO NOT EDIT.

package config_service

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

// NewOpenAPIDiscoveryParams creates a new OpenAPIDiscoveryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpenAPIDiscoveryParams() *OpenAPIDiscoveryParams {
	return &OpenAPIDiscoveryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpenAPIDiscoveryParamsWithTimeout creates a new OpenAPIDiscoveryParams object
// with the ability to set a timeout on a request.
func NewOpenAPIDiscoveryParamsWithTimeout(timeout time.Duration) *OpenAPIDiscoveryParams {
	return &OpenAPIDiscoveryParams{
		timeout: timeout,
	}
}

// NewOpenAPIDiscoveryParamsWithContext creates a new OpenAPIDiscoveryParams object
// with the ability to set a context for a request.
func NewOpenAPIDiscoveryParamsWithContext(ctx context.Context) *OpenAPIDiscoveryParams {
	return &OpenAPIDiscoveryParams{
		Context: ctx,
	}
}

// NewOpenAPIDiscoveryParamsWithHTTPClient creates a new OpenAPIDiscoveryParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpenAPIDiscoveryParamsWithHTTPClient(client *http.Client) *OpenAPIDiscoveryParams {
	return &OpenAPIDiscoveryParams{
		HTTPClient: client,
	}
}

/* OpenAPIDiscoveryParams contains all the parameters to send to the API endpoint
   for the open Api discovery operation.

   Typically these are written to a http.Request.
*/
type OpenAPIDiscoveryParams struct {

	/* EndpointType.

	   Filter result to a specific endpoint type.
	*/
	EndpointType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the open Api discovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAPIDiscoveryParams) WithDefaults() *OpenAPIDiscoveryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the open Api discovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAPIDiscoveryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the open Api discovery params
func (o *OpenAPIDiscoveryParams) WithTimeout(timeout time.Duration) *OpenAPIDiscoveryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the open Api discovery params
func (o *OpenAPIDiscoveryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the open Api discovery params
func (o *OpenAPIDiscoveryParams) WithContext(ctx context.Context) *OpenAPIDiscoveryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the open Api discovery params
func (o *OpenAPIDiscoveryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the open Api discovery params
func (o *OpenAPIDiscoveryParams) WithHTTPClient(client *http.Client) *OpenAPIDiscoveryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the open Api discovery params
func (o *OpenAPIDiscoveryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndpointType adds the endpointType to the open Api discovery params
func (o *OpenAPIDiscoveryParams) WithEndpointType(endpointType *string) *OpenAPIDiscoveryParams {
	o.SetEndpointType(endpointType)
	return o
}

// SetEndpointType adds the endpointType to the open Api discovery params
func (o *OpenAPIDiscoveryParams) SetEndpointType(endpointType *string) {
	o.EndpointType = endpointType
}

// WriteToRequest writes these params to a swagger request
func (o *OpenAPIDiscoveryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndpointType != nil {

		// query param EndpointType
		var qrEndpointType string

		if o.EndpointType != nil {
			qrEndpointType = *o.EndpointType
		}
		qEndpointType := qrEndpointType
		if qEndpointType != "" {

			if err := r.SetQueryParam("EndpointType", qEndpointType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
