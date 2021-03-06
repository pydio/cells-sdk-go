// Code generated by go-swagger; DO NOT EDIT.

package user_service

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

// NewSearchUsersParams creates a new SearchUsersParams object
// with the default values initialized.
func NewSearchUsersParams() *SearchUsersParams {
	var ()
	return &SearchUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUsersParamsWithTimeout creates a new SearchUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchUsersParamsWithTimeout(timeout time.Duration) *SearchUsersParams {
	var ()
	return &SearchUsersParams{

		timeout: timeout,
	}
}

// NewSearchUsersParamsWithContext creates a new SearchUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchUsersParamsWithContext(ctx context.Context) *SearchUsersParams {
	var ()
	return &SearchUsersParams{

		Context: ctx,
	}
}

// NewSearchUsersParamsWithHTTPClient creates a new SearchUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchUsersParamsWithHTTPClient(client *http.Client) *SearchUsersParams {
	var ()
	return &SearchUsersParams{
		HTTPClient: client,
	}
}

/*SearchUsersParams contains all the parameters to send to the API endpoint
for the search users operation typically these are written to a http.Request
*/
type SearchUsersParams struct {

	/*Body*/
	Body *models.RestSearchUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search users params
func (o *SearchUsersParams) WithTimeout(timeout time.Duration) *SearchUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search users params
func (o *SearchUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search users params
func (o *SearchUsersParams) WithContext(ctx context.Context) *SearchUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search users params
func (o *SearchUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search users params
func (o *SearchUsersParams) WithHTTPClient(client *http.Client) *SearchUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search users params
func (o *SearchUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search users params
func (o *SearchUsersParams) WithBody(body *models.RestSearchUserRequest) *SearchUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search users params
func (o *SearchUsersParams) SetBody(body *models.RestSearchUserRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
