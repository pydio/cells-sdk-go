// Code generated by go-swagger; DO NOT EDIT.

package update_service

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

// NewUpdateRequiredParams creates a new UpdateRequiredParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRequiredParams() *UpdateRequiredParams {
	return &UpdateRequiredParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRequiredParamsWithTimeout creates a new UpdateRequiredParams object
// with the ability to set a timeout on a request.
func NewUpdateRequiredParamsWithTimeout(timeout time.Duration) *UpdateRequiredParams {
	return &UpdateRequiredParams{
		timeout: timeout,
	}
}

// NewUpdateRequiredParamsWithContext creates a new UpdateRequiredParams object
// with the ability to set a context for a request.
func NewUpdateRequiredParamsWithContext(ctx context.Context) *UpdateRequiredParams {
	return &UpdateRequiredParams{
		Context: ctx,
	}
}

// NewUpdateRequiredParamsWithHTTPClient creates a new UpdateRequiredParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRequiredParamsWithHTTPClient(client *http.Client) *UpdateRequiredParams {
	return &UpdateRequiredParams{
		HTTPClient: client,
	}
}

/*
UpdateRequiredParams contains all the parameters to send to the API endpoint

	for the update required operation.

	Typically these are written to a http.Request.
*/
type UpdateRequiredParams struct {

	// Body.
	Body *models.UpdateUpdateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update required params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRequiredParams) WithDefaults() *UpdateRequiredParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update required params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRequiredParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update required params
func (o *UpdateRequiredParams) WithTimeout(timeout time.Duration) *UpdateRequiredParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update required params
func (o *UpdateRequiredParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update required params
func (o *UpdateRequiredParams) WithContext(ctx context.Context) *UpdateRequiredParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update required params
func (o *UpdateRequiredParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update required params
func (o *UpdateRequiredParams) WithHTTPClient(client *http.Client) *UpdateRequiredParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update required params
func (o *UpdateRequiredParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update required params
func (o *UpdateRequiredParams) WithBody(body *models.UpdateUpdateRequest) *UpdateRequiredParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update required params
func (o *UpdateRequiredParams) SetBody(body *models.UpdateUpdateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRequiredParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
