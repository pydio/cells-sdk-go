// Code generated by go-swagger; DO NOT EDIT.

package workspace_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteWorkspaceParams creates a new DeleteWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteWorkspaceParams() *DeleteWorkspaceParams {
	return &DeleteWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkspaceParamsWithTimeout creates a new DeleteWorkspaceParams object
// with the ability to set a timeout on a request.
func NewDeleteWorkspaceParamsWithTimeout(timeout time.Duration) *DeleteWorkspaceParams {
	return &DeleteWorkspaceParams{
		timeout: timeout,
	}
}

// NewDeleteWorkspaceParamsWithContext creates a new DeleteWorkspaceParams object
// with the ability to set a context for a request.
func NewDeleteWorkspaceParamsWithContext(ctx context.Context) *DeleteWorkspaceParams {
	return &DeleteWorkspaceParams{
		Context: ctx,
	}
}

// NewDeleteWorkspaceParamsWithHTTPClient creates a new DeleteWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteWorkspaceParamsWithHTTPClient(client *http.Client) *DeleteWorkspaceParams {
	return &DeleteWorkspaceParams{
		HTTPClient: client,
	}
}

/*
DeleteWorkspaceParams contains all the parameters to send to the API endpoint

	for the delete workspace operation.

	Typically these are written to a http.Request.
*/
type DeleteWorkspaceParams struct {

	/* Attributes.

	   JSON-encoded list of attributes.
	*/
	Attributes *string

	/* Description.

	   Description of the workspace (max length 1000).
	*/
	Description *string

	/* Label.

	   Label of the workspace (max length 500).
	*/
	Label *string

	/* LastUpdated.

	   Last modification time.

	   Format: int32
	*/
	LastUpdated *int32

	/* PoliciesContextEditable.

	   Context-resolved to quickly check if workspace is editable or not.
	*/
	PoliciesContextEditable *bool

	/* RootUUIDs.

	   Quick list of the RootNodes uuids.
	*/
	RootUUIDs []string

	/* Scope.

	   Scope can be ADMIN, ROOM (=CELL) or LINK.

	   Default: "ANY"
	*/
	Scope *string

	/* Slug.

	   Slug is an url-compatible form of the workspace label, or can be freely modified (max length 500)
	*/
	Slug string

	/* UUID.

	   Unique identifier of the workspace.
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkspaceParams) WithDefaults() *DeleteWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteWorkspaceParams) SetDefaults() {
	var (
		scopeDefault = string("ANY")
	)

	val := DeleteWorkspaceParams{
		Scope: &scopeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete workspace params
func (o *DeleteWorkspaceParams) WithTimeout(timeout time.Duration) *DeleteWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workspace params
func (o *DeleteWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workspace params
func (o *DeleteWorkspaceParams) WithContext(ctx context.Context) *DeleteWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workspace params
func (o *DeleteWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workspace params
func (o *DeleteWorkspaceParams) WithHTTPClient(client *http.Client) *DeleteWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workspace params
func (o *DeleteWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the delete workspace params
func (o *DeleteWorkspaceParams) WithAttributes(attributes *string) *DeleteWorkspaceParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the delete workspace params
func (o *DeleteWorkspaceParams) SetAttributes(attributes *string) {
	o.Attributes = attributes
}

// WithDescription adds the description to the delete workspace params
func (o *DeleteWorkspaceParams) WithDescription(description *string) *DeleteWorkspaceParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the delete workspace params
func (o *DeleteWorkspaceParams) SetDescription(description *string) {
	o.Description = description
}

// WithLabel adds the label to the delete workspace params
func (o *DeleteWorkspaceParams) WithLabel(label *string) *DeleteWorkspaceParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the delete workspace params
func (o *DeleteWorkspaceParams) SetLabel(label *string) {
	o.Label = label
}

// WithLastUpdated adds the lastUpdated to the delete workspace params
func (o *DeleteWorkspaceParams) WithLastUpdated(lastUpdated *int32) *DeleteWorkspaceParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the delete workspace params
func (o *DeleteWorkspaceParams) SetLastUpdated(lastUpdated *int32) {
	o.LastUpdated = lastUpdated
}

// WithPoliciesContextEditable adds the policiesContextEditable to the delete workspace params
func (o *DeleteWorkspaceParams) WithPoliciesContextEditable(policiesContextEditable *bool) *DeleteWorkspaceParams {
	o.SetPoliciesContextEditable(policiesContextEditable)
	return o
}

// SetPoliciesContextEditable adds the policiesContextEditable to the delete workspace params
func (o *DeleteWorkspaceParams) SetPoliciesContextEditable(policiesContextEditable *bool) {
	o.PoliciesContextEditable = policiesContextEditable
}

// WithRootUUIDs adds the rootUUIDs to the delete workspace params
func (o *DeleteWorkspaceParams) WithRootUUIDs(rootUUIDs []string) *DeleteWorkspaceParams {
	o.SetRootUUIDs(rootUUIDs)
	return o
}

// SetRootUUIDs adds the rootUUiDs to the delete workspace params
func (o *DeleteWorkspaceParams) SetRootUUIDs(rootUUIDs []string) {
	o.RootUUIDs = rootUUIDs
}

// WithScope adds the scope to the delete workspace params
func (o *DeleteWorkspaceParams) WithScope(scope *string) *DeleteWorkspaceParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the delete workspace params
func (o *DeleteWorkspaceParams) SetScope(scope *string) {
	o.Scope = scope
}

// WithSlug adds the slug to the delete workspace params
func (o *DeleteWorkspaceParams) WithSlug(slug string) *DeleteWorkspaceParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the delete workspace params
func (o *DeleteWorkspaceParams) SetSlug(slug string) {
	o.Slug = slug
}

// WithUUID adds the uuid to the delete workspace params
func (o *DeleteWorkspaceParams) WithUUID(uuid *string) *DeleteWorkspaceParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete workspace params
func (o *DeleteWorkspaceParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Attributes != nil {

		// query param Attributes
		var qrAttributes string

		if o.Attributes != nil {
			qrAttributes = *o.Attributes
		}
		qAttributes := qrAttributes
		if qAttributes != "" {

			if err := r.SetQueryParam("Attributes", qAttributes); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param Description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("Description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.Label != nil {

		// query param Label
		var qrLabel string

		if o.Label != nil {
			qrLabel = *o.Label
		}
		qLabel := qrLabel
		if qLabel != "" {

			if err := r.SetQueryParam("Label", qLabel); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param LastUpdated
		var qrLastUpdated int32

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := swag.FormatInt32(qrLastUpdated)
		if qLastUpdated != "" {

			if err := r.SetQueryParam("LastUpdated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.PoliciesContextEditable != nil {

		// query param PoliciesContextEditable
		var qrPoliciesContextEditable bool

		if o.PoliciesContextEditable != nil {
			qrPoliciesContextEditable = *o.PoliciesContextEditable
		}
		qPoliciesContextEditable := swag.FormatBool(qrPoliciesContextEditable)
		if qPoliciesContextEditable != "" {

			if err := r.SetQueryParam("PoliciesContextEditable", qPoliciesContextEditable); err != nil {
				return err
			}
		}
	}

	if o.RootUUIDs != nil {

		// binding items for RootUUIDs
		joinedRootUUIDs := o.bindParamRootUUIDs(reg)

		// query array param RootUUIDs
		if err := r.SetQueryParam("RootUUIDs", joinedRootUUIDs...); err != nil {
			return err
		}
	}

	if o.Scope != nil {

		// query param Scope
		var qrScope string

		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("Scope", qScope); err != nil {
				return err
			}
		}
	}

	// path param Slug
	if err := r.SetPathParam("Slug", o.Slug); err != nil {
		return err
	}

	if o.UUID != nil {

		// query param UUID
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("UUID", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteWorkspace binds the parameter RootUUIDs
func (o *DeleteWorkspaceParams) bindParamRootUUIDs(formats strfmt.Registry) []string {
	rootUUIDsIR := o.RootUUIDs

	var rootUUIDsIC []string
	for _, rootUUIDsIIR := range rootUUIDsIR { // explode []string

		rootUUIDsIIV := rootUUIDsIIR // string as string
		rootUUIDsIC = append(rootUUIDsIC, rootUUIDsIIV)
	}

	// items.CollectionFormat: "multi"
	rootUUIDsIS := swag.JoinByFormat(rootUUIDsIC, "multi")

	return rootUUIDsIS
}
