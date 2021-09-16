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

	"github.com/pydio/cells-sdk-go/v3/models"
)

// NewDeleteEncryptionKeyParams creates a new DeleteEncryptionKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteEncryptionKeyParams() *DeleteEncryptionKeyParams {
	return &DeleteEncryptionKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEncryptionKeyParamsWithTimeout creates a new DeleteEncryptionKeyParams object
// with the ability to set a timeout on a request.
func NewDeleteEncryptionKeyParamsWithTimeout(timeout time.Duration) *DeleteEncryptionKeyParams {
	return &DeleteEncryptionKeyParams{
		timeout: timeout,
	}
}

// NewDeleteEncryptionKeyParamsWithContext creates a new DeleteEncryptionKeyParams object
// with the ability to set a context for a request.
func NewDeleteEncryptionKeyParamsWithContext(ctx context.Context) *DeleteEncryptionKeyParams {
	return &DeleteEncryptionKeyParams{
		Context: ctx,
	}
}

// NewDeleteEncryptionKeyParamsWithHTTPClient creates a new DeleteEncryptionKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteEncryptionKeyParamsWithHTTPClient(client *http.Client) *DeleteEncryptionKeyParams {
	return &DeleteEncryptionKeyParams{
		HTTPClient: client,
	}
}

/* DeleteEncryptionKeyParams contains all the parameters to send to the API endpoint
   for the delete encryption key operation.

   Typically these are written to a http.Request.
*/
type DeleteEncryptionKeyParams struct {

	// Body.
	Body *models.EncryptionAdminDeleteKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete encryption key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEncryptionKeyParams) WithDefaults() *DeleteEncryptionKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete encryption key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteEncryptionKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete encryption key params
func (o *DeleteEncryptionKeyParams) WithTimeout(timeout time.Duration) *DeleteEncryptionKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete encryption key params
func (o *DeleteEncryptionKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete encryption key params
func (o *DeleteEncryptionKeyParams) WithContext(ctx context.Context) *DeleteEncryptionKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete encryption key params
func (o *DeleteEncryptionKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete encryption key params
func (o *DeleteEncryptionKeyParams) WithHTTPClient(client *http.Client) *DeleteEncryptionKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete encryption key params
func (o *DeleteEncryptionKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete encryption key params
func (o *DeleteEncryptionKeyParams) WithBody(body *models.EncryptionAdminDeleteKeyRequest) *DeleteEncryptionKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete encryption key params
func (o *DeleteEncryptionKeyParams) SetBody(body *models.EncryptionAdminDeleteKeyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEncryptionKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
