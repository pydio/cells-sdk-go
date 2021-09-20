// Code generated by go-swagger; DO NOT EDIT.

package activity_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/v2/models"
)

// NewSubscribeParams creates a new SubscribeParams object
// with the default values initialized.
func NewSubscribeParams() *SubscribeParams {
	var ()
	return &SubscribeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscribeParamsWithTimeout creates a new SubscribeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscribeParamsWithTimeout(timeout time.Duration) *SubscribeParams {
	var ()
	return &SubscribeParams{

		timeout: timeout,
	}
}

// NewSubscribeParamsWithContext creates a new SubscribeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscribeParamsWithContext(ctx context.Context) *SubscribeParams {
	var ()
	return &SubscribeParams{

		Context: ctx,
	}
}

// NewSubscribeParamsWithHTTPClient creates a new SubscribeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscribeParamsWithHTTPClient(client *http.Client) *SubscribeParams {
	var ()
	return &SubscribeParams{
		HTTPClient: client,
	}
}

/*SubscribeParams contains all the parameters to send to the API endpoint
for the subscribe operation typically these are written to a http.Request
*/
type SubscribeParams struct {

	/*Body*/
	Body *models.ActivitySubscription

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscribe params
func (o *SubscribeParams) WithTimeout(timeout time.Duration) *SubscribeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscribe params
func (o *SubscribeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscribe params
func (o *SubscribeParams) WithContext(ctx context.Context) *SubscribeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscribe params
func (o *SubscribeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscribe params
func (o *SubscribeParams) WithHTTPClient(client *http.Client) *SubscribeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscribe params
func (o *SubscribeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the subscribe params
func (o *SubscribeParams) WithBody(body *models.ActivitySubscription) *SubscribeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the subscribe params
func (o *SubscribeParams) SetBody(body *models.ActivitySubscription) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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