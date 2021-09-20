// Code generated by go-swagger; DO NOT EDIT.

package search_service

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

// NewNodesParams creates a new NodesParams object
// with the default values initialized.
func NewNodesParams() *NodesParams {
	var ()
	return &NodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesParamsWithTimeout creates a new NodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesParamsWithTimeout(timeout time.Duration) *NodesParams {
	var ()
	return &NodesParams{

		timeout: timeout,
	}
}

// NewNodesParamsWithContext creates a new NodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodesParamsWithContext(ctx context.Context) *NodesParams {
	var ()
	return &NodesParams{

		Context: ctx,
	}
}

// NewNodesParamsWithHTTPClient creates a new NodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesParamsWithHTTPClient(client *http.Client) *NodesParams {
	var ()
	return &NodesParams{
		HTTPClient: client,
	}
}

/*NodesParams contains all the parameters to send to the API endpoint
for the nodes operation typically these are written to a http.Request
*/
type NodesParams struct {

	/*Body*/
	Body *models.TreeSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes params
func (o *NodesParams) WithTimeout(timeout time.Duration) *NodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes params
func (o *NodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes params
func (o *NodesParams) WithContext(ctx context.Context) *NodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes params
func (o *NodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes params
func (o *NodesParams) WithHTTPClient(client *http.Client) *NodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes params
func (o *NodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the nodes params
func (o *NodesParams) WithBody(body *models.TreeSearchRequest) *NodesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the nodes params
func (o *NodesParams) SetBody(body *models.TreeSearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *NodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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