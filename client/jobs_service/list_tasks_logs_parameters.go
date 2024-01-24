// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

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

// NewListTasksLogsParams creates a new ListTasksLogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTasksLogsParams() *ListTasksLogsParams {
	return &ListTasksLogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTasksLogsParamsWithTimeout creates a new ListTasksLogsParams object
// with the ability to set a timeout on a request.
func NewListTasksLogsParamsWithTimeout(timeout time.Duration) *ListTasksLogsParams {
	return &ListTasksLogsParams{
		timeout: timeout,
	}
}

// NewListTasksLogsParamsWithContext creates a new ListTasksLogsParams object
// with the ability to set a context for a request.
func NewListTasksLogsParamsWithContext(ctx context.Context) *ListTasksLogsParams {
	return &ListTasksLogsParams{
		Context: ctx,
	}
}

// NewListTasksLogsParamsWithHTTPClient creates a new ListTasksLogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTasksLogsParamsWithHTTPClient(client *http.Client) *ListTasksLogsParams {
	return &ListTasksLogsParams{
		HTTPClient: client,
	}
}

/*
ListTasksLogsParams contains all the parameters to send to the API endpoint

	for the list tasks logs operation.

	Typically these are written to a http.Request.
*/
type ListTasksLogsParams struct {

	/* Body.

	   ListLogRequest launches a parameterised query in the log repository and streams the results.
	*/
	Body *models.LogListLogRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list tasks logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTasksLogsParams) WithDefaults() *ListTasksLogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list tasks logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTasksLogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list tasks logs params
func (o *ListTasksLogsParams) WithTimeout(timeout time.Duration) *ListTasksLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tasks logs params
func (o *ListTasksLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tasks logs params
func (o *ListTasksLogsParams) WithContext(ctx context.Context) *ListTasksLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tasks logs params
func (o *ListTasksLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tasks logs params
func (o *ListTasksLogsParams) WithHTTPClient(client *http.Client) *ListTasksLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tasks logs params
func (o *ListTasksLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list tasks logs params
func (o *ListTasksLogsParams) WithBody(body *models.LogListLogRequest) *ListTasksLogsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list tasks logs params
func (o *ListTasksLogsParams) SetBody(body *models.LogListLogRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListTasksLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
