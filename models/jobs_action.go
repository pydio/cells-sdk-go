// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobsAction jobs action
//
// swagger:model jobsAction
type JobsAction struct {

	// Previous action output filter
	ActionOutputFilter *JobsActionOutputFilter `json:"ActionOutputFilter,omitempty"`

	// Stop full chain now : do not carry on executing next actions
	BreakAfter bool `json:"BreakAfter,omitempty"`

	// Bypass this action (forward input to output and do nothing)
	Bypass bool `json:"Bypass,omitempty"`

	// Other actions to perform after this one is finished,
	// using the Output of this action as Input for the next.
	// If there are many, it is considered they can be triggered
	// in parallel
	ChainedActions []*JobsAction `json:"ChainedActions"`

	// Metadata policy-based filter
	ContextMetaFilter *JobsContextMetaFilter `json:"ContextMetaFilter,omitempty"`

	// DataSource objects filter
	DataSourceFilter *JobsDataSourceSelector `json:"DataSourceFilter,omitempty"`

	// DataSource objects collector
	DataSourceSelector *JobsDataSourceSelector `json:"DataSourceSelector,omitempty"`

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

	// CollectAction adds starts another chain after the whole ChainedAction/FailedFilterActions have been performed
	MergeAction *JobsAction `json:"MergeAction,omitempty"`

	// Node Filter
	NodesFilter *JobsNodesSelector `json:"NodesFilter,omitempty"`

	// Nodes Selector
	NodesSelector *JobsNodesSelector `json:"NodesSelector,omitempty"`

	// Defined parameters for this action
	Parameters map[string]string `json:"Parameters,omitempty"`

	// Optional timeout for this action
	Timeout string `json:"Timeout,omitempty"`

	// Filter on specific triggers
	TriggerFilter *JobsTriggerFilter `json:"TriggerFilter,omitempty"`

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

	if err := m.validateDataSourceFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSourceSelector(formats); err != nil {
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

	if err := m.validateMergeAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodesFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodesSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggerFilter(formats); err != nil {
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActionOutputFilter")
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChainedActions" + "." + strconv.Itoa(i))
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ContextMetaFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateDataSourceFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSourceFilter) { // not required
		return nil
	}

	if m.DataSourceFilter != nil {
		if err := m.DataSourceFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DataSourceFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DataSourceFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateDataSourceSelector(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSourceSelector) { // not required
		return nil
	}

	if m.DataSourceSelector != nil {
		if err := m.DataSourceSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DataSourceSelector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DataSourceSelector")
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FailedFilterActions" + "." + strconv.Itoa(i))
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("IdmFilter")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("IdmSelector")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateMergeAction(formats strfmt.Registry) error {
	if swag.IsZero(m.MergeAction) { // not required
		return nil
	}

	if m.MergeAction != nil {
		if err := m.MergeAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MergeAction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MergeAction")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NodesFilter")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NodesSelector")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) validateTriggerFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.TriggerFilter) { // not required
		return nil
	}

	if m.TriggerFilter != nil {
		if err := m.TriggerFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TriggerFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TriggerFilter")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UsersFilter")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UsersSelector")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jobs action based on the context it is used
func (m *JobsAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActionOutputFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChainedActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContextMetaFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFailedFilterActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdmFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdmSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMergeAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodesFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodesSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTriggerFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsersFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsersSelector(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsAction) contextValidateActionOutputFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.ActionOutputFilter != nil {

		if swag.IsZero(m.ActionOutputFilter) { // not required
			return nil
		}

		if err := m.ActionOutputFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ActionOutputFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ActionOutputFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateChainedActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChainedActions); i++ {

		if m.ChainedActions[i] != nil {

			if swag.IsZero(m.ChainedActions[i]) { // not required
				return nil
			}

			if err := m.ChainedActions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChainedActions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChainedActions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsAction) contextValidateContextMetaFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.ContextMetaFilter != nil {

		if swag.IsZero(m.ContextMetaFilter) { // not required
			return nil
		}

		if err := m.ContextMetaFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContextMetaFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ContextMetaFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateDataSourceFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.DataSourceFilter != nil {

		if swag.IsZero(m.DataSourceFilter) { // not required
			return nil
		}

		if err := m.DataSourceFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DataSourceFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DataSourceFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateDataSourceSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.DataSourceSelector != nil {

		if swag.IsZero(m.DataSourceSelector) { // not required
			return nil
		}

		if err := m.DataSourceSelector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DataSourceSelector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("DataSourceSelector")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateFailedFilterActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FailedFilterActions); i++ {

		if m.FailedFilterActions[i] != nil {

			if swag.IsZero(m.FailedFilterActions[i]) { // not required
				return nil
			}

			if err := m.FailedFilterActions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("FailedFilterActions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("FailedFilterActions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsAction) contextValidateIdmFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.IdmFilter != nil {

		if swag.IsZero(m.IdmFilter) { // not required
			return nil
		}

		if err := m.IdmFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IdmFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("IdmFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateIdmSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.IdmSelector != nil {

		if swag.IsZero(m.IdmSelector) { // not required
			return nil
		}

		if err := m.IdmSelector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("IdmSelector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("IdmSelector")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateMergeAction(ctx context.Context, formats strfmt.Registry) error {

	if m.MergeAction != nil {

		if swag.IsZero(m.MergeAction) { // not required
			return nil
		}

		if err := m.MergeAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MergeAction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("MergeAction")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateNodesFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.NodesFilter != nil {

		if swag.IsZero(m.NodesFilter) { // not required
			return nil
		}

		if err := m.NodesFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodesFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NodesFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateNodesSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.NodesSelector != nil {

		if swag.IsZero(m.NodesSelector) { // not required
			return nil
		}

		if err := m.NodesSelector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodesSelector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NodesSelector")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateTriggerFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.TriggerFilter != nil {

		if swag.IsZero(m.TriggerFilter) { // not required
			return nil
		}

		if err := m.TriggerFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("TriggerFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("TriggerFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateUsersFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.UsersFilter != nil {

		if swag.IsZero(m.UsersFilter) { // not required
			return nil
		}

		if err := m.UsersFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UsersFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UsersFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsAction) contextValidateUsersSelector(ctx context.Context, formats strfmt.Registry) error {

	if m.UsersSelector != nil {

		if swag.IsZero(m.UsersSelector) { // not required
			return nil
		}

		if err := m.UsersSelector.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UsersSelector")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UsersSelector")
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
