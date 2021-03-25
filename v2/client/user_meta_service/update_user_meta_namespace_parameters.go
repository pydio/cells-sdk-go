// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

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

// NewUpdateUserMetaNamespaceParams creates a new UpdateUserMetaNamespaceParams object
// with the default values initialized.
func NewUpdateUserMetaNamespaceParams() *UpdateUserMetaNamespaceParams {
	var ()
	return &UpdateUserMetaNamespaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserMetaNamespaceParamsWithTimeout creates a new UpdateUserMetaNamespaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserMetaNamespaceParamsWithTimeout(timeout time.Duration) *UpdateUserMetaNamespaceParams {
	var ()
	return &UpdateUserMetaNamespaceParams{

		timeout: timeout,
	}
}

// NewUpdateUserMetaNamespaceParamsWithContext creates a new UpdateUserMetaNamespaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserMetaNamespaceParamsWithContext(ctx context.Context) *UpdateUserMetaNamespaceParams {
	var ()
	return &UpdateUserMetaNamespaceParams{

		Context: ctx,
	}
}

// NewUpdateUserMetaNamespaceParamsWithHTTPClient creates a new UpdateUserMetaNamespaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserMetaNamespaceParamsWithHTTPClient(client *http.Client) *UpdateUserMetaNamespaceParams {
	var ()
	return &UpdateUserMetaNamespaceParams{
		HTTPClient: client,
	}
}

/*UpdateUserMetaNamespaceParams contains all the parameters to send to the API endpoint
for the update user meta namespace operation typically these are written to a http.Request
*/
type UpdateUserMetaNamespaceParams struct {

	/*Body*/
	Body *models.IdmUpdateUserMetaNamespaceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user meta namespace params
func (o *UpdateUserMetaNamespaceParams) WithTimeout(timeout time.Duration) *UpdateUserMetaNamespaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user meta namespace params
func (o *UpdateUserMetaNamespaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user meta namespace params
func (o *UpdateUserMetaNamespaceParams) WithContext(ctx context.Context) *UpdateUserMetaNamespaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user meta namespace params
func (o *UpdateUserMetaNamespaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user meta namespace params
func (o *UpdateUserMetaNamespaceParams) WithHTTPClient(client *http.Client) *UpdateUserMetaNamespaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user meta namespace params
func (o *UpdateUserMetaNamespaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update user meta namespace params
func (o *UpdateUserMetaNamespaceParams) WithBody(body *models.IdmUpdateUserMetaNamespaceRequest) *UpdateUserMetaNamespaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update user meta namespace params
func (o *UpdateUserMetaNamespaceParams) SetBody(body *models.IdmUpdateUserMetaNamespaceRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserMetaNamespaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
