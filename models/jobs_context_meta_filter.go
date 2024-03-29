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

// JobsContextMetaFilter PolicyContextFilter can be used to filter request metadata
//
// swagger:model jobsContextMetaFilter
type JobsContextMetaFilter struct {

	// Selector additional description
	Description string `json:"Description,omitempty"`

	// Selector custom label
	Label string `json:"Label,omitempty"`

	// Can be built with ContextMetaSingleQuery
	Query *ServiceQuery `json:"Query,omitempty"`

	// Type of context filter
	Type *JobsContextMetaFilterType `json:"Type,omitempty"`
}

// Validate validates this jobs context meta filter
func (m *JobsContextMetaFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsContextMetaFilter) validateQuery(formats strfmt.Registry) error {
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

func (m *JobsContextMetaFilter) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jobs context meta filter based on the context it is used
func (m *JobsContextMetaFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsContextMetaFilter) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

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

func (m *JobsContextMetaFilter) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {

		if swag.IsZero(m.Type) { // not required
			return nil
		}

		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsContextMetaFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsContextMetaFilter) UnmarshalBinary(b []byte) error {
	var res JobsContextMetaFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
