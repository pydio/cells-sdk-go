// Code generated by go-swagger; DO NOT EDIT.

package share_service

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

// NewPutShareLinkParams creates a new PutShareLinkParams object
// with the default values initialized.
func NewPutShareLinkParams() *PutShareLinkParams {
	var ()
	return &PutShareLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutShareLinkParamsWithTimeout creates a new PutShareLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutShareLinkParamsWithTimeout(timeout time.Duration) *PutShareLinkParams {
	var ()
	return &PutShareLinkParams{

		timeout: timeout,
	}
}

// NewPutShareLinkParamsWithContext creates a new PutShareLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutShareLinkParamsWithContext(ctx context.Context) *PutShareLinkParams {
	var ()
	return &PutShareLinkParams{

		Context: ctx,
	}
}

// NewPutShareLinkParamsWithHTTPClient creates a new PutShareLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutShareLinkParamsWithHTTPClient(client *http.Client) *PutShareLinkParams {
	var ()
	return &PutShareLinkParams{
		HTTPClient: client,
	}
}

/*PutShareLinkParams contains all the parameters to send to the API endpoint
for the put share link operation typically these are written to a http.Request
*/
type PutShareLinkParams struct {

	/*Body*/
	Body *models.RestPutShareLinkRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put share link params
func (o *PutShareLinkParams) WithTimeout(timeout time.Duration) *PutShareLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put share link params
func (o *PutShareLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put share link params
func (o *PutShareLinkParams) WithContext(ctx context.Context) *PutShareLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put share link params
func (o *PutShareLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put share link params
func (o *PutShareLinkParams) WithHTTPClient(client *http.Client) *PutShareLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put share link params
func (o *PutShareLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put share link params
func (o *PutShareLinkParams) WithBody(body *models.RestPutShareLinkRequest) *PutShareLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put share link params
func (o *PutShareLinkParams) SetBody(body *models.RestPutShareLinkRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutShareLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
