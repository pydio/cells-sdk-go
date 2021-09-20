// Code generated by go-swagger; DO NOT EDIT.

package admin_tree_service

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

// NewListAdminTreeParams creates a new ListAdminTreeParams object
// with the default values initialized.
func NewListAdminTreeParams() *ListAdminTreeParams {
	var ()
	return &ListAdminTreeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAdminTreeParamsWithTimeout creates a new ListAdminTreeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAdminTreeParamsWithTimeout(timeout time.Duration) *ListAdminTreeParams {
	var ()
	return &ListAdminTreeParams{

		timeout: timeout,
	}
}

// NewListAdminTreeParamsWithContext creates a new ListAdminTreeParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAdminTreeParamsWithContext(ctx context.Context) *ListAdminTreeParams {
	var ()
	return &ListAdminTreeParams{

		Context: ctx,
	}
}

// NewListAdminTreeParamsWithHTTPClient creates a new ListAdminTreeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAdminTreeParamsWithHTTPClient(client *http.Client) *ListAdminTreeParams {
	var ()
	return &ListAdminTreeParams{
		HTTPClient: client,
	}
}

/*ListAdminTreeParams contains all the parameters to send to the API endpoint
for the list admin tree operation typically these are written to a http.Request
*/
type ListAdminTreeParams struct {

	/*Body*/
	Body *models.TreeListNodesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list admin tree params
func (o *ListAdminTreeParams) WithTimeout(timeout time.Duration) *ListAdminTreeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list admin tree params
func (o *ListAdminTreeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list admin tree params
func (o *ListAdminTreeParams) WithContext(ctx context.Context) *ListAdminTreeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list admin tree params
func (o *ListAdminTreeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list admin tree params
func (o *ListAdminTreeParams) WithHTTPClient(client *http.Client) *ListAdminTreeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list admin tree params
func (o *ListAdminTreeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list admin tree params
func (o *ListAdminTreeParams) WithBody(body *models.TreeListNodesRequest) *ListAdminTreeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list admin tree params
func (o *ListAdminTreeParams) SetBody(body *models.TreeListNodesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListAdminTreeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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