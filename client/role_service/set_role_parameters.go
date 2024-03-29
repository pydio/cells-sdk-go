// Code generated by go-swagger; DO NOT EDIT.

package role_service

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

// NewSetRoleParams creates a new SetRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetRoleParams() *SetRoleParams {
	return &SetRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetRoleParamsWithTimeout creates a new SetRoleParams object
// with the ability to set a timeout on a request.
func NewSetRoleParamsWithTimeout(timeout time.Duration) *SetRoleParams {
	return &SetRoleParams{
		timeout: timeout,
	}
}

// NewSetRoleParamsWithContext creates a new SetRoleParams object
// with the ability to set a context for a request.
func NewSetRoleParamsWithContext(ctx context.Context) *SetRoleParams {
	return &SetRoleParams{
		Context: ctx,
	}
}

// NewSetRoleParamsWithHTTPClient creates a new SetRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetRoleParamsWithHTTPClient(client *http.Client) *SetRoleParams {
	return &SetRoleParams{
		HTTPClient: client,
	}
}

/*
SetRoleParams contains all the parameters to send to the API endpoint

	for the set role operation.

	Typically these are written to a http.Request.
*/
type SetRoleParams struct {

	/* UUID.

	   Unique identifier of this role
	*/
	UUID string

	// Body.
	Body SetRoleBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetRoleParams) WithDefaults() *SetRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set role params
func (o *SetRoleParams) WithTimeout(timeout time.Duration) *SetRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set role params
func (o *SetRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set role params
func (o *SetRoleParams) WithContext(ctx context.Context) *SetRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set role params
func (o *SetRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set role params
func (o *SetRoleParams) WithHTTPClient(client *http.Client) *SetRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set role params
func (o *SetRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the set role params
func (o *SetRoleParams) WithUUID(uuid string) *SetRoleParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the set role params
func (o *SetRoleParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WithBody adds the body to the set role params
func (o *SetRoleParams) WithBody(body SetRoleBody) *SetRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set role params
func (o *SetRoleParams) SetBody(body SetRoleBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Uuid
	if err := r.SetPathParam("Uuid", o.UUID); err != nil {
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
