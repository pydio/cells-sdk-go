// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// JobsCommand jobs command
//
// swagger:model jobsCommand
type JobsCommand string

func NewJobsCommand(value JobsCommand) *JobsCommand {
	v := value
	return &v
}

const (

	// JobsCommandNone captures enum value "None"
	JobsCommandNone JobsCommand = "None"

	// JobsCommandPause captures enum value "Pause"
	JobsCommandPause JobsCommand = "Pause"

	// JobsCommandResume captures enum value "Resume"
	JobsCommandResume JobsCommand = "Resume"

	// JobsCommandStop captures enum value "Stop"
	JobsCommandStop JobsCommand = "Stop"

	// JobsCommandDelete captures enum value "Delete"
	JobsCommandDelete JobsCommand = "Delete"

	// JobsCommandRunOnce captures enum value "RunOnce"
	JobsCommandRunOnce JobsCommand = "RunOnce"

	// JobsCommandInactive captures enum value "Inactive"
	JobsCommandInactive JobsCommand = "Inactive"

	// JobsCommandActive captures enum value "Active"
	JobsCommandActive JobsCommand = "Active"
)

// for schema
var jobsCommandEnum []interface{}

func init() {
	var res []JobsCommand
	if err := json.Unmarshal([]byte(`["None","Pause","Resume","Stop","Delete","RunOnce","Inactive","Active"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobsCommandEnum = append(jobsCommandEnum, v)
	}
}

func (m JobsCommand) validateJobsCommandEnum(path, location string, value JobsCommand) error {
	if err := validate.EnumCase(path, location, value, jobsCommandEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this jobs command
func (m JobsCommand) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateJobsCommandEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this jobs command based on context it is used
func (m JobsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
