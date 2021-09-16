// Code generated by go-swagger; DO NOT EDIT.

package install_service

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

// NewPostInstallParams creates a new PostInstallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostInstallParams() *PostInstallParams {
	return &PostInstallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostInstallParamsWithTimeout creates a new PostInstallParams object
// with the ability to set a timeout on a request.
func NewPostInstallParamsWithTimeout(timeout time.Duration) *PostInstallParams {
	return &PostInstallParams{
		timeout: timeout,
	}
}

// NewPostInstallParamsWithContext creates a new PostInstallParams object
// with the ability to set a context for a request.
func NewPostInstallParamsWithContext(ctx context.Context) *PostInstallParams {
	return &PostInstallParams{
		Context: ctx,
	}
}

// NewPostInstallParamsWithHTTPClient creates a new PostInstallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostInstallParamsWithHTTPClient(client *http.Client) *PostInstallParams {
	return &PostInstallParams{
		HTTPClient: client,
	}
}

/* PostInstallParams contains all the parameters to send to the API endpoint
   for the post install operation.

   Typically these are written to a http.Request.
*/
type PostInstallParams struct {

	// Body.
	Body *models.InstallInstallRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInstallParams) WithDefaults() *PostInstallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post install params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostInstallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post install params
func (o *PostInstallParams) WithTimeout(timeout time.Duration) *PostInstallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post install params
func (o *PostInstallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post install params
func (o *PostInstallParams) WithContext(ctx context.Context) *PostInstallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post install params
func (o *PostInstallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post install params
func (o *PostInstallParams) WithHTTPClient(client *http.Client) *PostInstallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post install params
func (o *PostInstallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post install params
func (o *PostInstallParams) WithBody(body *models.InstallInstallRequest) *PostInstallParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post install params
func (o *PostInstallParams) SetBody(body *models.InstallInstallRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostInstallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
