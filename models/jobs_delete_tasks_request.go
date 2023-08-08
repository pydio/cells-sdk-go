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

// JobsDeleteTasksRequest jobs delete tasks request
//
// swagger:model jobsDeleteTasksRequest
type JobsDeleteTasksRequest struct {

	// Id of the job
	JobID string `json:"JobId,omitempty"`

	// If deleting by status, optionally keep only a number of tasks
	PruneLimit int32 `json:"PruneLimit,omitempty"`

	// If no TaskID and/or no JobID are passed, delete tasks by status
	Status []*JobsTaskStatus `json:"Status"`

	// Ids of tasks to delete
	TaskID []string `json:"TaskID"`
}

// Validate validates this jobs delete tasks request
func (m *JobsDeleteTasksRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsDeleteTasksRequest) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	for i := 0; i < len(m.Status); i++ {
		if swag.IsZero(m.Status[i]) { // not required
			continue
		}

		if m.Status[i] != nil {
			if err := m.Status[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this jobs delete tasks request based on the context it is used
func (m *JobsDeleteTasksRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsDeleteTasksRequest) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Status); i++ {

		if m.Status[i] != nil {

			if swag.IsZero(m.Status[i]) { // not required
				return nil
			}

			if err := m.Status[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Status" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Status" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsDeleteTasksRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsDeleteTasksRequest) UnmarshalBinary(b []byte) error {
	var res JobsDeleteTasksRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
