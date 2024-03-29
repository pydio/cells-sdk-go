// Code generated by go-swagger; DO NOT EDIT.

package token_service

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

// NewGenerateDocumentAccessTokenParams creates a new GenerateDocumentAccessTokenParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGenerateDocumentAccessTokenParams() *GenerateDocumentAccessTokenParams {
	return &GenerateDocumentAccessTokenParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGenerateDocumentAccessTokenParamsWithTimeout creates a new GenerateDocumentAccessTokenParams object
// with the ability to set a timeout on a request.
func NewGenerateDocumentAccessTokenParamsWithTimeout(timeout time.Duration) *GenerateDocumentAccessTokenParams {
	return &GenerateDocumentAccessTokenParams{
		timeout: timeout,
	}
}

// NewGenerateDocumentAccessTokenParamsWithContext creates a new GenerateDocumentAccessTokenParams object
// with the ability to set a context for a request.
func NewGenerateDocumentAccessTokenParamsWithContext(ctx context.Context) *GenerateDocumentAccessTokenParams {
	return &GenerateDocumentAccessTokenParams{
		Context: ctx,
	}
}

// NewGenerateDocumentAccessTokenParamsWithHTTPClient creates a new GenerateDocumentAccessTokenParams object
// with the ability to set a custom HTTPClient for a request.
func NewGenerateDocumentAccessTokenParamsWithHTTPClient(client *http.Client) *GenerateDocumentAccessTokenParams {
	return &GenerateDocumentAccessTokenParams{
		HTTPClient: client,
	}
}

/*
GenerateDocumentAccessTokenParams contains all the parameters to send to the API endpoint

	for the generate document access token operation.

	Typically these are written to a http.Request.
*/
type GenerateDocumentAccessTokenParams struct {

	// Body.
	Body *models.RestDocumentAccessTokenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the generate document access token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateDocumentAccessTokenParams) WithDefaults() *GenerateDocumentAccessTokenParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the generate document access token params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GenerateDocumentAccessTokenParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the generate document access token params
func (o *GenerateDocumentAccessTokenParams) WithTimeout(timeout time.Duration) *GenerateDocumentAccessTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the generate document access token params
func (o *GenerateDocumentAccessTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the generate document access token params
func (o *GenerateDocumentAccessTokenParams) WithContext(ctx context.Context) *GenerateDocumentAccessTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the generate document access token params
func (o *GenerateDocumentAccessTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the generate document access token params
func (o *GenerateDocumentAccessTokenParams) WithHTTPClient(client *http.Client) *GenerateDocumentAccessTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the generate document access token params
func (o *GenerateDocumentAccessTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the generate document access token params
func (o *GenerateDocumentAccessTokenParams) WithBody(body *models.RestDocumentAccessTokenRequest) *GenerateDocumentAccessTokenParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the generate document access token params
func (o *GenerateDocumentAccessTokenParams) SetBody(body *models.RestDocumentAccessTokenRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateDocumentAccessTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
