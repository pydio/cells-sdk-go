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

	"github.com/pydio/cells-sdk-go/v5/models"
)

// NewSearchUserMetaParams creates a new SearchUserMetaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchUserMetaParams() *SearchUserMetaParams {
	return &SearchUserMetaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUserMetaParamsWithTimeout creates a new SearchUserMetaParams object
// with the ability to set a timeout on a request.
func NewSearchUserMetaParamsWithTimeout(timeout time.Duration) *SearchUserMetaParams {
	return &SearchUserMetaParams{
		timeout: timeout,
	}
}

// NewSearchUserMetaParamsWithContext creates a new SearchUserMetaParams object
// with the ability to set a context for a request.
func NewSearchUserMetaParamsWithContext(ctx context.Context) *SearchUserMetaParams {
	return &SearchUserMetaParams{
		Context: ctx,
	}
}

// NewSearchUserMetaParamsWithHTTPClient creates a new SearchUserMetaParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchUserMetaParamsWithHTTPClient(client *http.Client) *SearchUserMetaParams {
	return &SearchUserMetaParams{
		HTTPClient: client,
	}
}

/*
SearchUserMetaParams contains all the parameters to send to the API endpoint

	for the search user meta operation.

	Typically these are written to a http.Request.
*/
type SearchUserMetaParams struct {

	// Body.
	Body *models.IdmSearchUserMetaRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search user meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUserMetaParams) WithDefaults() *SearchUserMetaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search user meta params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchUserMetaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search user meta params
func (o *SearchUserMetaParams) WithTimeout(timeout time.Duration) *SearchUserMetaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search user meta params
func (o *SearchUserMetaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search user meta params
func (o *SearchUserMetaParams) WithContext(ctx context.Context) *SearchUserMetaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search user meta params
func (o *SearchUserMetaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search user meta params
func (o *SearchUserMetaParams) WithHTTPClient(client *http.Client) *SearchUserMetaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search user meta params
func (o *SearchUserMetaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search user meta params
func (o *SearchUserMetaParams) WithBody(body *models.IdmSearchUserMetaRequest) *SearchUserMetaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search user meta params
func (o *SearchUserMetaParams) SetBody(body *models.IdmSearchUserMetaRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUserMetaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
