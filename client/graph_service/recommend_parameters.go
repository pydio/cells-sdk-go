// Code generated by go-swagger; DO NOT EDIT.

package graph_service

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

	"github.com/pydio/cells-sdk-go/v5/models"
)

// NewRecommendParams creates a new RecommendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRecommendParams() *RecommendParams {
	return &RecommendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRecommendParamsWithTimeout creates a new RecommendParams object
// with the ability to set a timeout on a request.
func NewRecommendParamsWithTimeout(timeout time.Duration) *RecommendParams {
	return &RecommendParams{
		timeout: timeout,
	}
}

// NewRecommendParamsWithContext creates a new RecommendParams object
// with the ability to set a context for a request.
func NewRecommendParamsWithContext(ctx context.Context) *RecommendParams {
	return &RecommendParams{
		Context: ctx,
	}
}

// NewRecommendParamsWithHTTPClient creates a new RecommendParams object
// with the ability to set a custom HTTPClient for a request.
func NewRecommendParamsWithHTTPClient(client *http.Client) *RecommendParams {
	return &RecommendParams{
		HTTPClient: client,
	}
}

/*
RecommendParams contains all the parameters to send to the API endpoint

	for the recommend operation.

	Typically these are written to a http.Request.
*/
type RecommendParams struct {

	// Body.
	Body *models.RestRecommendRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the recommend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecommendParams) WithDefaults() *RecommendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the recommend params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RecommendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the recommend params
func (o *RecommendParams) WithTimeout(timeout time.Duration) *RecommendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the recommend params
func (o *RecommendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the recommend params
func (o *RecommendParams) WithContext(ctx context.Context) *RecommendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the recommend params
func (o *RecommendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the recommend params
func (o *RecommendParams) WithHTTPClient(client *http.Client) *RecommendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the recommend params
func (o *RecommendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the recommend params
func (o *RecommendParams) WithBody(body *models.RestRecommendRequest) *RecommendParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the recommend params
func (o *RecommendParams) SetBody(body *models.RestRecommendRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RecommendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
