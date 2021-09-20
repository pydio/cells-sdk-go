// Code generated by go-swagger; DO NOT EDIT.

package mailer_service

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

// NewSendParams creates a new SendParams object
// with the default values initialized.
func NewSendParams() *SendParams {
	var ()
	return &SendParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendParamsWithTimeout creates a new SendParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendParamsWithTimeout(timeout time.Duration) *SendParams {
	var ()
	return &SendParams{

		timeout: timeout,
	}
}

// NewSendParamsWithContext creates a new SendParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendParamsWithContext(ctx context.Context) *SendParams {
	var ()
	return &SendParams{

		Context: ctx,
	}
}

// NewSendParamsWithHTTPClient creates a new SendParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendParamsWithHTTPClient(client *http.Client) *SendParams {
	var ()
	return &SendParams{
		HTTPClient: client,
	}
}

/*SendParams contains all the parameters to send to the API endpoint
for the send operation typically these are written to a http.Request
*/
type SendParams struct {

	/*Body*/
	Body *models.MailerMail

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send params
func (o *SendParams) WithTimeout(timeout time.Duration) *SendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send params
func (o *SendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send params
func (o *SendParams) WithContext(ctx context.Context) *SendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send params
func (o *SendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send params
func (o *SendParams) WithHTTPClient(client *http.Client) *SendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send params
func (o *SendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send params
func (o *SendParams) WithBody(body *models.MailerMail) *SendParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send params
func (o *SendParams) SetBody(body *models.MailerMail) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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