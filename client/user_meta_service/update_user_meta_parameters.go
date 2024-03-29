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

	"github.com/pydio/cells-sdk-go/v4/models"
)

// NewUpdateUserMetaParams creates a new UpdateUserMetaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserMetaParams() *UpdateUserMetaParams {
	return &UpdateUserMetaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserMetaParamsWithTimeout creates a new UpdateUserMetaParams object
// with the ability to set a timeout on a request.
func NewUpdateUserMetaParamsWithTimeout(timeout time.Duration) *UpdateUserMetaParams {
	return &UpdateUserMetaParams{
		timeout: timeout,
	}
}

// NewUpdateUserMetaParamsWithContext creates a new UpdateUserMetaParams object
// with the ability to set a context for a request.
func NewUpdateUserMetaParamsWithContext(ctx context.Context) *UpdateUserMetaParams {
	return &UpdateUserMetaParams{
		Context: ctx,
	}
}

// NewUpdateUserMetaParamsWithHTTPClient creates a new UpdateUserMetaParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserMetaParamsWithHTTPClient(client *http.Client) *UpdateUserMetaParams {
	return &UpdateUserMetaParams{
		HTTPClient: client,
	}
}

/*
UpdateUserMetaParams contains all the parameters to send to the API endpoint

	for the update user meta operation.

	Typically these are written to a http.Request.
*/
type UpdateUserMetaParams struct {

	// Body.
	Body *models.IdmUpdateUserMetaRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserMetaParams) WithDefaults() *UpdateUserMetaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserMetaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user meta params
func (o *UpdateUserMetaParams) WithTimeout(timeout time.Duration) *UpdateUserMetaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user meta params
func (o *UpdateUserMetaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user meta params
func (o *UpdateUserMetaParams) WithContext(ctx context.Context) *UpdateUserMetaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user meta params
func (o *UpdateUserMetaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user meta params
func (o *UpdateUserMetaParams) WithHTTPClient(client *http.Client) *UpdateUserMetaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user meta params
func (o *UpdateUserMetaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user meta params
func (o *UpdateUserMetaParams) WithBody(body *models.IdmUpdateUserMetaRequest) *UpdateUserMetaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user meta params
func (o *UpdateUserMetaParams) SetBody(body *models.IdmUpdateUserMetaRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserMetaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
