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

// JobsDataSourceSelector Selector/Filter for DataSource objects
//
// swagger:model jobsDataSourceSelector
type JobsDataSourceSelector struct {

	// Select all
	All bool `json:"All,omitempty"`

	// Collect results
	Collect bool `json:"Collect,omitempty"`

	// Selector additional description
	Description string `json:"Description,omitempty"`

	// Selector custom label
	Label string `json:"Label,omitempty"`

	// Composition of DataSourceSingleQueries
	Query *ServiceQuery `json:"Query,omitempty"`

	// Optional Timeout for this selector
	Timeout string `json:"Timeout,omitempty"`

	// Selector type, either DataSource or Object service
	Type *JobsDataSourceSelectorType `json:"Type,omitempty"`
}

// Validate validates this jobs data source selector
func (m *JobsDataSourceSelector) Validate(formats strfmt.Registry) error {
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

func (m *JobsDataSourceSelector) validateQuery(formats strfmt.Registry) error {
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

func (m *JobsDataSourceSelector) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this jobs data source selector based on the context it is used
func (m *JobsDataSourceSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *JobsDataSourceSelector) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.Query != nil {
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

func (m *JobsDataSourceSelector) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
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
func (m *JobsDataSourceSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsDataSourceSelector) UnmarshalBinary(b []byte) error {
	var res JobsDataSourceSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
