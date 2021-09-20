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

// NewStatAdminTreeParams creates a new StatAdminTreeParams object
// with the default values initialized.
func NewStatAdminTreeParams() *StatAdminTreeParams {
	var ()
	return &StatAdminTreeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatAdminTreeParamsWithTimeout creates a new StatAdminTreeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatAdminTreeParamsWithTimeout(timeout time.Duration) *StatAdminTreeParams {
	var ()
	return &StatAdminTreeParams{

		timeout: timeout,
	}
}

// NewStatAdminTreeParamsWithContext creates a new StatAdminTreeParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatAdminTreeParamsWithContext(ctx context.Context) *StatAdminTreeParams {
	var ()
	return &StatAdminTreeParams{

		Context: ctx,
	}
}

// NewStatAdminTreeParamsWithHTTPClient creates a new StatAdminTreeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatAdminTreeParamsWithHTTPClient(client *http.Client) *StatAdminTreeParams {
	var ()
	return &StatAdminTreeParams{
		HTTPClient: client,
	}
}

/*StatAdminTreeParams contains all the parameters to send to the API endpoint
for the stat admin tree operation typically these are written to a http.Request
*/
type StatAdminTreeParams struct {

	/*Body*/
	Body *models.TreeReadNodeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stat admin tree params
func (o *StatAdminTreeParams) WithTimeout(timeout time.Duration) *StatAdminTreeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stat admin tree params
func (o *StatAdminTreeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stat admin tree params
func (o *StatAdminTreeParams) WithContext(ctx context.Context) *StatAdminTreeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stat admin tree params
func (o *StatAdminTreeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stat admin tree params
func (o *StatAdminTreeParams) WithHTTPClient(client *http.Client) *StatAdminTreeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stat admin tree params
func (o *StatAdminTreeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the stat admin tree params
func (o *StatAdminTreeParams) WithBody(body *models.TreeReadNodeRequest) *StatAdminTreeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the stat admin tree params
func (o *StatAdminTreeParams) SetBody(body *models.TreeReadNodeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StatAdminTreeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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