// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobsListJobsRequest jobs list jobs request
//
// swagger:model jobsListJobsRequest
type JobsListJobsRequest struct {

	// Filter with only event-based jobs
	EventsOnly bool `json:"EventsOnly,omitempty"`

	// Load jobs by their ID
	JobIDs []string `json:"JobIDs"`

	// Load tasks that correspond to the given TaskStatus
	LoadTasks *JobsTaskStatus `json:"LoadTasks,omitempty"`

	// Restrict to a specific owner (current user by default)
	Owner string `json:"Owner,omitempty"`

	// Lmit the number of results
	TasksLimit int32 `json:"TasksLimit,omitempty"`

	// Start listing at a given position
	TasksOffset int32 `json:"TasksOffset,omitempty"`

	// Filter with only timer-based jobs
	TimersOnly bool `json:"TimersOnly,omitempty"`
}

// Validate validates this jobs list jobs request
func (m *JobsListJobsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsListJobsRequest) validateLoadTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.LoadTasks) { // not required
		return nil
	}

	if m.LoadTasks != nil {
		if err := m.LoadTasks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LoadTasks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LoadTasks")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jobs list jobs request based on the context it is used
func (m *JobsListJobsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoadTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsListJobsRequest) contextValidateLoadTasks(ctx context.Context, formats strfmt.Registry) error {

	if m.LoadTasks != nil {
		if err := m.LoadTasks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("LoadTasks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("LoadTasks")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsListJobsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsListJobsRequest) UnmarshalBinary(b []byte) error {
	var res JobsListJobsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
