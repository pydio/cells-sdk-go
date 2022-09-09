// Code generated by go-swagger; DO NOT EDIT.

package role_service

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

// NewGetRoleParams creates a new GetRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoleParams() *GetRoleParams {
	return &GetRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleParamsWithTimeout creates a new GetRoleParams object
// with the ability to set a timeout on a request.
func NewGetRoleParamsWithTimeout(timeout time.Duration) *GetRoleParams {
	return &GetRoleParams{
		timeout: timeout,
	}
}

// NewGetRoleParamsWithContext creates a new GetRoleParams object
// with the ability to set a context for a request.
func NewGetRoleParamsWithContext(ctx context.Context) *GetRoleParams {
	return &GetRoleParams{
		Context: ctx,
	}
}

// NewGetRoleParamsWithHTTPClient creates a new GetRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoleParamsWithHTTPClient(client *http.Client) *GetRoleParams {
	return &GetRoleParams{
		HTTPClient: client,
	}
}

/*
GetRoleParams contains all the parameters to send to the API endpoint

	for the get role operation.

	Typically these are written to a http.Request.
*/
type GetRoleParams struct {

	/* AutoApplies.

	   List of profiles (standard, shared, admin) on which the role will be automatically applied.
	*/
	AutoApplies []string

	/* ForceOverride.

	   Is used in a stack of roles, this one will always be applied last.
	*/
	ForceOverride *bool

	/* GroupRole.

	   Whether this role is attached to a Group object.
	*/
	GroupRole *bool

	/* IsTeam.

	   Whether this role represents a user team or not.
	*/
	IsTeam *bool

	/* Label.

	   Label of this role.
	*/
	Label *string

	/* LastUpdated.

	   Last modification date of the role.

	   Format: int32
	*/
	LastUpdated *int32

	/* PoliciesContextEditable.

	   Whether the policies resolve into an editable state.
	*/
	PoliciesContextEditable *bool

	/* UserRole.

	   Whether this role is attached to a User object.
	*/
	UserRole *bool

	/* UUID.

	   Unique identifier of this role
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoleParams) WithDefaults() *GetRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get role params
func (o *GetRoleParams) WithTimeout(timeout time.Duration) *GetRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role params
func (o *GetRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role params
func (o *GetRoleParams) WithContext(ctx context.Context) *GetRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role params
func (o *GetRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role params
func (o *GetRoleParams) WithHTTPClient(client *http.Client) *GetRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role params
func (o *GetRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAutoApplies adds the autoApplies to the get role params
func (o *GetRoleParams) WithAutoApplies(autoApplies []string) *GetRoleParams {
	o.SetAutoApplies(autoApplies)
	return o
}

// SetAutoApplies adds the autoApplies to the get role params
func (o *GetRoleParams) SetAutoApplies(autoApplies []string) {
	o.AutoApplies = autoApplies
}

// WithForceOverride adds the forceOverride to the get role params
func (o *GetRoleParams) WithForceOverride(forceOverride *bool) *GetRoleParams {
	o.SetForceOverride(forceOverride)
	return o
}

// SetForceOverride adds the forceOverride to the get role params
func (o *GetRoleParams) SetForceOverride(forceOverride *bool) {
	o.ForceOverride = forceOverride
}

// WithGroupRole adds the groupRole to the get role params
func (o *GetRoleParams) WithGroupRole(groupRole *bool) *GetRoleParams {
	o.SetGroupRole(groupRole)
	return o
}

// SetGroupRole adds the groupRole to the get role params
func (o *GetRoleParams) SetGroupRole(groupRole *bool) {
	o.GroupRole = groupRole
}

// WithIsTeam adds the isTeam to the get role params
func (o *GetRoleParams) WithIsTeam(isTeam *bool) *GetRoleParams {
	o.SetIsTeam(isTeam)
	return o
}

// SetIsTeam adds the isTeam to the get role params
func (o *GetRoleParams) SetIsTeam(isTeam *bool) {
	o.IsTeam = isTeam
}

// WithLabel adds the label to the get role params
func (o *GetRoleParams) WithLabel(label *string) *GetRoleParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the get role params
func (o *GetRoleParams) SetLabel(label *string) {
	o.Label = label
}

// WithLastUpdated adds the lastUpdated to the get role params
func (o *GetRoleParams) WithLastUpdated(lastUpdated *int32) *GetRoleParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the get role params
func (o *GetRoleParams) SetLastUpdated(lastUpdated *int32) {
	o.LastUpdated = lastUpdated
}

// WithPoliciesContextEditable adds the policiesContextEditable to the get role params
func (o *GetRoleParams) WithPoliciesContextEditable(policiesContextEditable *bool) *GetRoleParams {
	o.SetPoliciesContextEditable(policiesContextEditable)
	return o
}

// SetPoliciesContextEditable adds the policiesContextEditable to the get role params
func (o *GetRoleParams) SetPoliciesContextEditable(policiesContextEditable *bool) {
	o.PoliciesContextEditable = policiesContextEditable
}

// WithUserRole adds the userRole to the get role params
func (o *GetRoleParams) WithUserRole(userRole *bool) *GetRoleParams {
	o.SetUserRole(userRole)
	return o
}

// SetUserRole adds the userRole to the get role params
func (o *GetRoleParams) SetUserRole(userRole *bool) {
	o.UserRole = userRole
}

// WithUUID adds the uuid to the get role params
func (o *GetRoleParams) WithUUID(uuid string) *GetRoleParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get role params
func (o *GetRoleParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AutoApplies != nil {

		// binding items for AutoApplies
		joinedAutoApplies := o.bindParamAutoApplies(reg)

		// query array param AutoApplies
		if err := r.SetQueryParam("AutoApplies", joinedAutoApplies...); err != nil {
			return err
		}
	}

	if o.ForceOverride != nil {

		// query param ForceOverride
		var qrForceOverride bool

		if o.ForceOverride != nil {
			qrForceOverride = *o.ForceOverride
		}
		qForceOverride := swag.FormatBool(qrForceOverride)
		if qForceOverride != "" {

			if err := r.SetQueryParam("ForceOverride", qForceOverride); err != nil {
				return err
			}
		}
	}

	if o.GroupRole != nil {

		// query param GroupRole
		var qrGroupRole bool

		if o.GroupRole != nil {
			qrGroupRole = *o.GroupRole
		}
		qGroupRole := swag.FormatBool(qrGroupRole)
		if qGroupRole != "" {

			if err := r.SetQueryParam("GroupRole", qGroupRole); err != nil {
				return err
			}
		}
	}

	if o.IsTeam != nil {

		// query param IsTeam
		var qrIsTeam bool

		if o.IsTeam != nil {
			qrIsTeam = *o.IsTeam
		}
		qIsTeam := swag.FormatBool(qrIsTeam)
		if qIsTeam != "" {

			if err := r.SetQueryParam("IsTeam", qIsTeam); err != nil {
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

	if o.UserRole != nil {

		// query param UserRole
		var qrUserRole bool

		if o.UserRole != nil {
			qrUserRole = *o.UserRole
		}
		qUserRole := swag.FormatBool(qrUserRole)
		if qUserRole != "" {

			if err := r.SetQueryParam("UserRole", qUserRole); err != nil {
				return err
			}
		}
	}

	// path param Uuid
	if err := r.SetPathParam("Uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetRole binds the parameter AutoApplies
func (o *GetRoleParams) bindParamAutoApplies(formats strfmt.Registry) []string {
	autoAppliesIR := o.AutoApplies

	var autoAppliesIC []string
	for _, autoAppliesIIR := range autoAppliesIR { // explode []string

		autoAppliesIIV := autoAppliesIIR // string as string
		autoAppliesIC = append(autoAppliesIC, autoAppliesIIV)
	}

	// items.CollectionFormat: "multi"
	autoAppliesIS := swag.JoinByFormat(autoAppliesIC, "multi")

	return autoAppliesIS
}
