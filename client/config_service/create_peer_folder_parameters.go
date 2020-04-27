// Code generated by go-swagger; DO NOT EDIT.

package config_service

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

// NewCreatePeerFolderParams creates a new CreatePeerFolderParams object
// with the default values initialized.
func NewCreatePeerFolderParams() *CreatePeerFolderParams {
	var ()
	return &CreatePeerFolderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePeerFolderParamsWithTimeout creates a new CreatePeerFolderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePeerFolderParamsWithTimeout(timeout time.Duration) *CreatePeerFolderParams {
	var ()
	return &CreatePeerFolderParams{

		timeout: timeout,
	}
}

// NewCreatePeerFolderParamsWithContext creates a new CreatePeerFolderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePeerFolderParamsWithContext(ctx context.Context) *CreatePeerFolderParams {
	var ()
	return &CreatePeerFolderParams{

		Context: ctx,
	}
}

// NewCreatePeerFolderParamsWithHTTPClient creates a new CreatePeerFolderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePeerFolderParamsWithHTTPClient(client *http.Client) *CreatePeerFolderParams {
	var ()
	return &CreatePeerFolderParams{
		HTTPClient: client,
	}
}

/*CreatePeerFolderParams contains all the parameters to send to the API endpoint
for the create peer folder operation typically these are written to a http.Request
*/
type CreatePeerFolderParams struct {

	/*PeerAddress*/
	PeerAddress string
	/*Body*/
	Body *models.RestCreatePeerFolderRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create peer folder params
func (o *CreatePeerFolderParams) WithTimeout(timeout time.Duration) *CreatePeerFolderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create peer folder params
func (o *CreatePeerFolderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create peer folder params
func (o *CreatePeerFolderParams) WithContext(ctx context.Context) *CreatePeerFolderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create peer folder params
func (o *CreatePeerFolderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create peer folder params
func (o *CreatePeerFolderParams) WithHTTPClient(client *http.Client) *CreatePeerFolderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create peer folder params
func (o *CreatePeerFolderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeerAddress adds the peerAddress to the create peer folder params
func (o *CreatePeerFolderParams) WithPeerAddress(peerAddress string) *CreatePeerFolderParams {
	o.SetPeerAddress(peerAddress)
	return o
}

// SetPeerAddress adds the peerAddress to the create peer folder params
func (o *CreatePeerFolderParams) SetPeerAddress(peerAddress string) {
	o.PeerAddress = peerAddress
}

// WithBody adds the body to the create peer folder params
func (o *CreatePeerFolderParams) WithBody(body *models.RestCreatePeerFolderRequest) *CreatePeerFolderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create peer folder params
func (o *CreatePeerFolderParams) SetBody(body *models.RestCreatePeerFolderRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePeerFolderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param PeerAddress
	if err := r.SetPathParam("PeerAddress", o.PeerAddress); err != nil {
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
