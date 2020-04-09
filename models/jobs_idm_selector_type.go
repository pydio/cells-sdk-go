// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// JobsIdmSelectorType Possible values for IdmSelector.Type
// swagger:model jobsIdmSelectorType
type JobsIdmSelectorType string

const (

	// JobsIdmSelectorTypeUser captures enum value "User"
	JobsIdmSelectorTypeUser JobsIdmSelectorType = "User"

	// JobsIdmSelectorTypeRole captures enum value "Role"
	JobsIdmSelectorTypeRole JobsIdmSelectorType = "Role"

	// JobsIdmSelectorTypeWorkspace captures enum value "Workspace"
	JobsIdmSelectorTypeWorkspace JobsIdmSelectorType = "Workspace"

	// JobsIdmSelectorTypeACL captures enum value "Acl"
	JobsIdmSelectorTypeACL JobsIdmSelectorType = "Acl"
)

// for schema
var jobsIdmSelectorTypeEnum []interface{}

func init() {
	var res []JobsIdmSelectorType
	if err := json.Unmarshal([]byte(`["User","Role","Workspace","Acl"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobsIdmSelectorTypeEnum = append(jobsIdmSelectorTypeEnum, v)
	}
}

func (m JobsIdmSelectorType) validateJobsIdmSelectorTypeEnum(path, location string, value JobsIdmSelectorType) error {
	if err := validate.Enum(path, location, value, jobsIdmSelectorTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this jobs idm selector type
func (m JobsIdmSelectorType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateJobsIdmSelectorTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
