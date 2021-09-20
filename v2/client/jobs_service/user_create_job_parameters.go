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

// NewUserCreateJobParams creates a new UserCreateJobParams object
// with the default values initialized.
func NewUserCreateJobParams() *UserCreateJobParams {
	var ()
	return &UserCreateJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserCreateJobParamsWithTimeout creates a new UserCreateJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserCreateJobParamsWithTimeout(timeout time.Duration) *UserCreateJobParams {
	var ()
	return &UserCreateJobParams{

		timeout: timeout,
	}
}

// NewUserCreateJobParamsWithContext creates a new UserCreateJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserCreateJobParamsWithContext(ctx context.Context) *UserCreateJobParams {
	var ()
	return &UserCreateJobParams{

		Context: ctx,
	}
}

// NewUserCreateJobParamsWithHTTPClient creates a new UserCreateJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserCreateJobParamsWithHTTPClient(client *http.Client) *UserCreateJobParams {
	var ()
	return &UserCreateJobParams{
		HTTPClient: client,
	}
}

/*UserCreateJobParams contains all the parameters to send to the API endpoint
for the user create job operation typically these are written to a http.Request
*/
type UserCreateJobParams struct {

	/*JobName*/
	JobName string
	/*Body*/
	Body *models.RestUserJobRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user create job params
func (o *UserCreateJobParams) WithTimeout(timeout time.Duration) *UserCreateJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user create job params
func (o *UserCreateJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user create job params
func (o *UserCreateJobParams) WithContext(ctx context.Context) *UserCreateJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user create job params
func (o *UserCreateJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user create job params
func (o *UserCreateJobParams) WithHTTPClient(client *http.Client) *UserCreateJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user create job params
func (o *UserCreateJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobName adds the jobName to the user create job params
func (o *UserCreateJobParams) WithJobName(jobName string) *UserCreateJobParams {
	o.SetJobName(jobName)
	return o
}

// SetJobName adds the jobName to the user create job params
func (o *UserCreateJobParams) SetJobName(jobName string) {
	o.JobName = jobName
}

// WithBody adds the body to the user create job params
func (o *UserCreateJobParams) WithBody(body *models.RestUserJobRequest) *UserCreateJobParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the user create job params
func (o *UserCreateJobParams) SetBody(body *models.RestUserJobRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UserCreateJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param JobName
	if err := r.SetPathParam("JobName", o.JobName); err != nil {
		return err
	}

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