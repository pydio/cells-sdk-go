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

// JobsTriggerFilter Filter for events, can be applied on action branches
//
// swagger:model jobsTriggerFilter
type JobsTriggerFilter struct {

	// Filter additional description
	Description string `json:"Description,omitempty"`

	// Filter custom label
	Label string `json:"Label,omitempty"`

	// Filter type
	Query *ServiceQuery `json:"Query,omitempty"`
}

// Validate validates this jobs trigger filter
func (m *JobsTriggerFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsTriggerFilter) validateQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Query")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jobs trigger filter based on the context it is used
func (m *JobsTriggerFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsTriggerFilter) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.Query != nil {

		if swag.IsZero(m.Query) { // not required
			return nil
		}

		if err := m.Query.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsTriggerFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsTriggerFilter) UnmarshalBinary(b []byte) error {
	var res JobsTriggerFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
