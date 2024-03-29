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

	"github.com/pydio/cells-sdk-go/v4/models"
)

// NewImportEncryptionKeyParams creates a new ImportEncryptionKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImportEncryptionKeyParams() *ImportEncryptionKeyParams {
	return &ImportEncryptionKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImportEncryptionKeyParamsWithTimeout creates a new ImportEncryptionKeyParams object
// with the ability to set a timeout on a request.
func NewImportEncryptionKeyParamsWithTimeout(timeout time.Duration) *ImportEncryptionKeyParams {
	return &ImportEncryptionKeyParams{
		timeout: timeout,
	}
}

// NewImportEncryptionKeyParamsWithContext creates a new ImportEncryptionKeyParams object
// with the ability to set a context for a request.
func NewImportEncryptionKeyParamsWithContext(ctx context.Context) *ImportEncryptionKeyParams {
	return &ImportEncryptionKeyParams{
		Context: ctx,
	}
}

// NewImportEncryptionKeyParamsWithHTTPClient creates a new ImportEncryptionKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewImportEncryptionKeyParamsWithHTTPClient(client *http.Client) *ImportEncryptionKeyParams {
	return &ImportEncryptionKeyParams{
		HTTPClient: client,
	}
}

/*
ImportEncryptionKeyParams contains all the parameters to send to the API endpoint

	for the import encryption key operation.

	Typically these are written to a http.Request.
*/
type ImportEncryptionKeyParams struct {

	// Body.
	Body *models.EncryptionAdminImportKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the import encryption key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportEncryptionKeyParams) WithDefaults() *ImportEncryptionKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the import encryption key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImportEncryptionKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the import encryption key params
func (o *ImportEncryptionKeyParams) WithTimeout(timeout time.Duration) *ImportEncryptionKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the import encryption key params
func (o *ImportEncryptionKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the import encryption key params
func (o *ImportEncryptionKeyParams) WithContext(ctx context.Context) *ImportEncryptionKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the import encryption key params
func (o *ImportEncryptionKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the import encryption key params
func (o *ImportEncryptionKeyParams) WithHTTPClient(client *http.Client) *ImportEncryptionKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the import encryption key params
func (o *ImportEncryptionKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the import encryption key params
func (o *ImportEncryptionKeyParams) WithBody(body *models.EncryptionAdminImportKeyRequest) *ImportEncryptionKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the import encryption key params
func (o *ImportEncryptionKeyParams) SetBody(body *models.EncryptionAdminImportKeyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ImportEncryptionKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
