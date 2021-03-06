// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

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

	models "github.com/pydio/cells-sdk-go/models"
)

// NewFrontSessionParams creates a new FrontSessionParams object
// with the default values initialized.
func NewFrontSessionParams() *FrontSessionParams {
	var ()
	return &FrontSessionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFrontSessionParamsWithTimeout creates a new FrontSessionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFrontSessionParamsWithTimeout(timeout time.Duration) *FrontSessionParams {
	var ()
	return &FrontSessionParams{

		timeout: timeout,
	}
}

// NewFrontSessionParamsWithContext creates a new FrontSessionParams object
// with the default values initialized, and the ability to set a context for a request
func NewFrontSessionParamsWithContext(ctx context.Context) *FrontSessionParams {
	var ()
	return &FrontSessionParams{

		Context: ctx,
	}
}

// NewFrontSessionParamsWithHTTPClient creates a new FrontSessionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFrontSessionParamsWithHTTPClient(client *http.Client) *FrontSessionParams {
	var ()
	return &FrontSessionParams{
		HTTPClient: client,
	}
}

/*FrontSessionParams contains all the parameters to send to the API endpoint
for the front session operation typically these are written to a http.Request
*/
type FrontSessionParams struct {

	/*Body*/
	Body *models.RestFrontSessionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the front session params
func (o *FrontSessionParams) WithTimeout(timeout time.Duration) *FrontSessionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the front session params
func (o *FrontSessionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the front session params
func (o *FrontSessionParams) WithContext(ctx context.Context) *FrontSessionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the front session params
func (o *FrontSessionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the front session params
func (o *FrontSessionParams) WithHTTPClient(client *http.Client) *FrontSessionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the front session params
func (o *FrontSessionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the front session params
func (o *FrontSessionParams) WithBody(body *models.RestFrontSessionRequest) *FrontSessionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the front session params
func (o *FrontSessionParams) SetBody(body *models.RestFrontSessionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FrontSessionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
