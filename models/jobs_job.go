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

// JobsJob jobs job
//
// swagger:model jobsJob
type JobsJob struct {

	// Chain of actions to perform
	Actions []*JobsAction `json:"Actions"`

	// Remove job automatically once it is finished (success only)
	AutoClean bool `json:"AutoClean,omitempty"`

	// Start task as soon as server is started
	AutoRestart bool `json:"AutoRestart,omitempty"`

	// Start task as soon as job is inserted
	AutoStart bool `json:"AutoStart,omitempty"`

	// Event Context Filter
	ContextMetaFilter *JobsContextMetaFilter `json:"ContextMetaFilter,omitempty"`

	// Timestamp for creation time
	CreatedAt int32 `json:"CreatedAt,omitempty"`

	// Job created by application or by administrator
	Custom bool `json:"Custom,omitempty"`

	// DataSource objects filter
	DataSourceFilter *JobsDataSourceSelector `json:"DataSourceFilter,omitempty"`

	// How the job will be triggered.
	// One of these must be set (not exclusive)
	// Listen to a given set of events
	EventNames []string `json:"EventNames"`

	// Expose this job through one or more-userspace APIs
	Hooks []*JobsJobHook `json:"Hooks"`

	// Unique ID for this Job
	ID string `json:"ID,omitempty"`

	// Idm objects filter
	IdmFilter *JobsIdmSelector `json:"IdmFilter,omitempty"`

	// Admin can temporarily disable this job
	Inactive bool `json:"Inactive,omitempty"`

	// Human-readable Label
	Label string `json:"Label,omitempty"`

	// Optional list of languages detected in the context at launch time
	Languages []string `json:"Languages"`

	// Task properties
	MaxConcurrency int32 `json:"MaxConcurrency,omitempty"`

	// Collect chain of actions into a merged output
	MergeAction *JobsAction `json:"MergeAction,omitempty"`

	// Timestamp for modification time
	ModifiedAt int32 `json:"ModifiedAt,omitempty"`

	// Filter out specific events
	NodeEventFilter *JobsNodesSelector `json:"NodeEventFilter,omitempty"`

	// Who created this Job
	Owner string `json:"Owner,omitempty"`

	// Job-level parameters that can be passed to underlying actions
	Parameters []*JobsJobParameter `json:"Parameters"`

	// Additional dependencies that may be required when running the job
	ResourcesDependencies []*ProtobufAny `json:"ResourcesDependencies"`

	// Schedule a periodic repetition
	Schedule *JobsSchedule `json:"Schedule,omitempty"`

	// Filled with currently running tasks
	Tasks []*JobsTask `json:"Tasks"`

	// Do not send notification on task update
	TasksSilentUpdate bool `json:"TasksSilentUpdate,omitempty"`

	// Optional Timeout any running job
	Timeout string `json:"Timeout,omitempty"`

	// Deprecated in favor of more generic IdmSelector
	UserEventFilter *JobsUsersSelector `json:"UserEventFilter,omitempty"`

	// Additional Versioning Metadata
	VersionMeta map[string]string `json:"VersionMeta,omitempty"`
}

// Validate validates this jobs job
func (m *JobsJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextMetaFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataSourceFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHooks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdmFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMergeAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeEventFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourcesDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserEventFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsJob) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) validateContextMetaFilter(formats strfmt.Registry) error {
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

func (m *JobsJob) validateDataSourceFilter(formats strfmt.Registry) error {
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

func (m *JobsJob) validateHooks(formats strfmt.Registry) error {
	if swag.IsZero(m.Hooks) { // not required
		return nil
	}

	for i := 0; i < len(m.Hooks); i++ {
		if swag.IsZero(m.Hooks[i]) { // not required
			continue
		}

		if m.Hooks[i] != nil {
			if err := m.Hooks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Hooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Hooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) validateIdmFilter(formats strfmt.Registry) error {
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

func (m *JobsJob) validateMergeAction(formats strfmt.Registry) error {
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

func (m *JobsJob) validateNodeEventFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeEventFilter) { // not required
		return nil
	}

	if m.NodeEventFilter != nil {
		if err := m.NodeEventFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodeEventFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NodeEventFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsJob) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) validateResourcesDependencies(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourcesDependencies) { // not required
		return nil
	}

	for i := 0; i < len(m.ResourcesDependencies); i++ {
		if swag.IsZero(m.ResourcesDependencies[i]) { // not required
			continue
		}

		if m.ResourcesDependencies[i] != nil {
			if err := m.ResourcesDependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ResourcesDependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ResourcesDependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Schedule")
			}
			return err
		}
	}

	return nil
}

func (m *JobsJob) validateTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Tasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) validateUserEventFilter(formats strfmt.Registry) error {
	if swag.IsZero(m.UserEventFilter) { // not required
		return nil
	}

	if m.UserEventFilter != nil {
		if err := m.UserEventFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UserEventFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UserEventFilter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jobs job based on the context it is used
func (m *JobsJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContextMetaFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDataSourceFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHooks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdmFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMergeAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeEventFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourcesDependencies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserEventFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsJob) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Actions); i++ {

		if m.Actions[i] != nil {

			if swag.IsZero(m.Actions[i]) { // not required
				return nil
			}

			if err := m.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) contextValidateContextMetaFilter(ctx context.Context, formats strfmt.Registry) error {

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

func (m *JobsJob) contextValidateDataSourceFilter(ctx context.Context, formats strfmt.Registry) error {

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

func (m *JobsJob) contextValidateHooks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Hooks); i++ {

		if m.Hooks[i] != nil {

			if swag.IsZero(m.Hooks[i]) { // not required
				return nil
			}

			if err := m.Hooks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Hooks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Hooks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) contextValidateIdmFilter(ctx context.Context, formats strfmt.Registry) error {

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

func (m *JobsJob) contextValidateMergeAction(ctx context.Context, formats strfmt.Registry) error {

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

func (m *JobsJob) contextValidateNodeEventFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeEventFilter != nil {

		if swag.IsZero(m.NodeEventFilter) { // not required
			return nil
		}

		if err := m.NodeEventFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NodeEventFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("NodeEventFilter")
			}
			return err
		}
	}

	return nil
}

func (m *JobsJob) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {

			if swag.IsZero(m.Parameters[i]) { // not required
				return nil
			}

			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) contextValidateResourcesDependencies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ResourcesDependencies); i++ {

		if m.ResourcesDependencies[i] != nil {

			if swag.IsZero(m.ResourcesDependencies[i]) { // not required
				return nil
			}

			if err := m.ResourcesDependencies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ResourcesDependencies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ResourcesDependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if swag.IsZero(m.Schedule) { // not required
			return nil
		}

		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Schedule")
			}
			return err
		}
	}

	return nil
}

func (m *JobsJob) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tasks); i++ {

		if m.Tasks[i] != nil {

			if swag.IsZero(m.Tasks[i]) { // not required
				return nil
			}

			if err := m.Tasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Tasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JobsJob) contextValidateUserEventFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.UserEventFilter != nil {

		if swag.IsZero(m.UserEventFilter) { // not required
			return nil
		}

		if err := m.UserEventFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("UserEventFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("UserEventFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsJob) UnmarshalBinary(b []byte) error {
	var res JobsJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
