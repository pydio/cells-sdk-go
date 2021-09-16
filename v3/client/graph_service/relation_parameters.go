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
)

// NewRelationParams creates a new RelationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRelationParams() *RelationParams {
	return &RelationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRelationParamsWithTimeout creates a new RelationParams object
// with the ability to set a timeout on a request.
func NewRelationParamsWithTimeout(timeout time.Duration) *RelationParams {
	return &RelationParams{
		timeout: timeout,
	}
}

// NewRelationParamsWithContext creates a new RelationParams object
// with the ability to set a context for a request.
func NewRelationParamsWithContext(ctx context.Context) *RelationParams {
	return &RelationParams{
		Context: ctx,
	}
}

// NewRelationParamsWithHTTPClient creates a new RelationParams object
// with the ability to set a custom HTTPClient for a request.
func NewRelationParamsWithHTTPClient(client *http.Client) *RelationParams {
	return &RelationParams{
		HTTPClient: client,
	}
}

/* RelationParams contains all the parameters to send to the API endpoint
   for the relation operation.

   Typically these are written to a http.Request.
*/
type RelationParams struct {

	// UserID.
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the relation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RelationParams) WithDefaults() *RelationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the relation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RelationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the relation params
func (o *RelationParams) WithTimeout(timeout time.Duration) *RelationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the relation params
func (o *RelationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the relation params
func (o *RelationParams) WithContext(ctx context.Context) *RelationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the relation params
func (o *RelationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the relation params
func (o *RelationParams) WithHTTPClient(client *http.Client) *RelationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the relation params
func (o *RelationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the relation params
func (o *RelationParams) WithUserID(userID string) *RelationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the relation params
func (o *RelationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *RelationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param UserId
	if err := r.SetPathParam("UserId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
