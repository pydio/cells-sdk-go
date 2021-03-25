// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JobsAction jobs action
// swagger:model jobsAction
type JobsAction struct {

	// Previous action output filter
	ActionOutputFilter *JobsActionOutputFilter `json:"ActionOutputFilter,omitempty"`

	// Other actions to perform after this one is finished,
	// using the Output of this action as Input for the next.
	// If there are many, it is considered they can be triggered
	// in parallel
	ChainedActions []*JobsAction `json:"ChainedActions"`

	// Metadata policy-based filter
	ContextMetaFilter *JobsContextMetaFilter `json:"ContextMetaFilter,omitempty"`

	// User-defined comment for this action
	Description string `json:"Description,omitempty"`

	// If any Filter is used, next actions can be triggered on Failure
	// This adds ability to create conditional Yes/No branches
	FailedFilterActions []*JobsAction `json:"FailedFilterActions"`

	// String Identifier for specific action
	ID string `json:"ID,omitempty"`

	// Idm objects filter
	IdmFilter *JobsIdmSelector `json:"IdmFilter,omitempty"`

	// Idm objects collector
	IdmSelector *JobsIdmSelector `json:"IdmSelector,omitempty"`

	// User-defined label for this action
	Label string `json:"Label,omitempty"`

	// Node Filter
	NodesFilter *JobsNodesSelector `json:"NodesFilter,omitempty"`

	// Nodes Selector
	NodesSelector *JobsNodesSelector `json:"NodesSelector,omitempty"`

	// Defined parameters for this action
	Parameters map[string]string `json:"Parameters,omitempty"`

	// User Filter (deprecated in favor of IdmSelector)
	UsersFilter *JobsUsersSelector `json:"UsersFilter,omitempty"`

	// Users Selector (deprecated in favor of IdmSelector)
	UsersSelector *JobsUsersSelector `json:"UsersSelector,omitempty"`
}

// Validate validates this jobs action
func (m *JobsAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionOutputFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChainedActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextMetaFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailedFilterActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdmFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdmSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodesFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodesSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsersFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsersSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsAction) validateActionOutputFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionOutputFilter) { // not required
		return nil
	}

	if m.ActionOutputFilter != nil {
		if err := m.ActionOutputFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActionOutputFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateChainedActions(formats strfmt.Registry) error {

	if swag.IsZero(m.ChainedActions) { // not required
		return nil
	}

	for i := 0; i < len(m.ChainedActions); i++ {
		if swag.IsZero(m.ChainedActions[i]) { // not required
			continue
		}

		if m.ChainedActions[i] != nil {
			if err := m.ChainedActions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChainedActions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsAction) validateContextMetaFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.ContextMetaFilter) { // not required
		return nil
	}

	if m.ContextMetaFilter != nil {
		if err := m.ContextMetaFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContextMetaFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateFailedFilterActions(formats strfmt.Registry) error {

	if swag.IsZero(m.FailedFilterActions) { // not required
		return nil
	}

	for i := 0; i < len(m.FailedFilterActions); i++ {
		if swag.IsZero(m.FailedFilterActions[i]) { // not required
			continue
		}

		if m.FailedFilterActions[i] != nil {
			if err := m.FailedFilterActions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FailedFilterActions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsAction) validateIdmFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.IdmFilter) { // not required
		return nil
	}

	if m.IdmFilter != nil {
		if err := m.IdmFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IdmFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateIdmSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.IdmSelector) { // not required
		return nil
	}

	if m.IdmSelector != nil {
		if err := m.IdmSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IdmSelector")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateNodesFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.NodesFilter) { // not required
		return nil
	}

	if m.NodesFilter != nil {
		if err := m.NodesFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodesFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateNodesSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.NodesSelector) { // not required
		return nil
	}

	if m.NodesSelector != nil {
		if err := m.NodesSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodesSelector")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateUsersFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.UsersFilter) { // not required
		return nil
	}

	if m.UsersFilter != nil {
		if err := m.UsersFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UsersFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateUsersSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.UsersSelector) { // not required
		return nil
	}

	if m.UsersSelector != nil {
		if err := m.UsersSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UsersSelector")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsAction) UnmarshalBinary(b []byte) error {
	var res JobsAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
