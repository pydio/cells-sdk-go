// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JobsActionOutputFilter ActionOutputFilter can be used to filter last message output
// swagger:model jobsActionOutputFilter
type JobsActionOutputFilter struct {

	// Query built from ActionOutputSingleQuery
	Query *ServiceQuery `json:"Query,omitempty"`
}

// Validate validates this jobs action output filter
func (m *JobsActionOutputFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsActionOutputFilter) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsActionOutputFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsActionOutputFilter) UnmarshalBinary(b []byte) error {
	var res JobsActionOutputFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
