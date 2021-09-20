// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

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

// NewUserDeleteTasksParams creates a new UserDeleteTasksParams object
// with the default values initialized.
func NewUserDeleteTasksParams() *UserDeleteTasksParams {
	var ()
	return &UserDeleteTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserDeleteTasksParamsWithTimeout creates a new UserDeleteTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserDeleteTasksParamsWithTimeout(timeout time.Duration) *UserDeleteTasksParams {
	var ()
	return &UserDeleteTasksParams{

		timeout: timeout,
	}
}

// NewUserDeleteTasksParamsWithContext creates a new UserDeleteTasksParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserDeleteTasksParamsWithContext(ctx context.Context) *UserDeleteTasksParams {
	var ()
	return &UserDeleteTasksParams{

		Context: ctx,
	}
}

// NewUserDeleteTasksParamsWithHTTPClient creates a new UserDeleteTasksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserDeleteTasksParamsWithHTTPClient(client *http.Client) *UserDeleteTasksParams {
	var ()
	return &UserDeleteTasksParams{
		HTTPClient: client,
	}
}

/*UserDeleteTasksParams contains all the parameters to send to the API endpoint
for the user delete tasks operation typically these are written to a http.Request
*/
type UserDeleteTasksParams struct {

	/*Body*/
	Body *models.JobsDeleteTasksRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user delete tasks params
func (o *UserDeleteTasksParams) WithTimeout(timeout time.Duration) *UserDeleteTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user delete tasks params
func (o *UserDeleteTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user delete tasks params
func (o *UserDeleteTasksParams) WithContext(ctx context.Context) *UserDeleteTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user delete tasks params
func (o *UserDeleteTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user delete tasks params
func (o *UserDeleteTasksParams) WithHTTPClient(client *http.Client) *UserDeleteTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user delete tasks params
func (o *UserDeleteTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the user delete tasks params
func (o *UserDeleteTasksParams) WithBody(body *models.JobsDeleteTasksRequest) *UserDeleteTasksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user delete tasks params
func (o *UserDeleteTasksParams) SetBody(body *models.JobsDeleteTasksRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UserDeleteTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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