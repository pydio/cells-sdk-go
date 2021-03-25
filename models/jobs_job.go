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

// JobsJob jobs job
// swagger:model jobsJob
type JobsJob struct {

	// Chain of actions to perform
	Actions []*JobsAction `json:"Actions"`

	// Remove job automatically once it is finished (success only)
	AutoClean bool `json:"AutoClean,omitempty"`

	// Start task as soon as job is inserted
	AutoStart bool `json:"AutoStart,omitempty"`

	// Event Context Filter
	ContextMetaFilter *JobsContextMetaFilter `json:"ContextMetaFilter,omitempty"`

	// How the job will be triggered.
	// One of these must be set (not exclusive)
	// Listen to a given set of events
	EventNames []string `json:"EventNames"`

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

	// Filter out specific events
	NodeEventFilter *JobsNodesSelector `json:"NodeEventFilter,omitempty"`

	// Who created this Job
	Owner string `json:"Owner,omitempty"`

	// Job-level parameters that can be passed to underlying actions
	Parameters []*JobsJobParameter `json:"Parameters"`

	// Schedule a periodic repetition
	Schedule *JobsSchedule `json:"Schedule,omitempty"`

	// Filled with currently running tasks
	Tasks []*JobsTask `json:"Tasks"`

	// Do not send notification on task update
	TasksSilentUpdate bool `json:"TasksSilentUpdate,omitempty"`

	// Deprecated in favor of more generic IdmSelector
	UserEventFilter *JobsUsersSelector `json:"UserEventFilter,omitempty"`
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

	if err := m.validateIdmFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeEventFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
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
			}
			return err
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
