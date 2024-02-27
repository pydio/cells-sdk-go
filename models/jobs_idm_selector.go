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

// JobsIdmSelector Generic container for select/filter idm objects
//
// swagger:model jobsIdmSelector
type JobsIdmSelector struct {

	// Load all objects
	All bool `json:"All,omitempty"`

	// Clear previous selection
	ClearInput bool `json:"ClearInput,omitempty"`

	// Pass a slice of objects to one action, or trigger all actions in parallel
	Collect bool `json:"Collect,omitempty"`

	// Selector additional description
	Description string `json:"Description,omitempty"`

	// Ignore query and just fan out input.[Type]
	FanOutInput bool `json:"FanOutInput,omitempty"`

	// Selector custom label
	Label string `json:"Label,omitempty"`

	// Serialized search query
	Query *ServiceQuery `json:"Query,omitempty"`

	// Handle ranges
	Range *JobsSelectorRange `json:"Range,omitempty"`

	// Optional Timeout for this selector
	Timeout string `json:"Timeout,omitempty"`

	// Type of objects to look for
	Type *JobsIdmSelectorType `json:"Type,omitempty"`
}

// Validate validates this jobs idm selector
func (m *JobsIdmSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
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

func (m *JobsIdmSelector) validateQuery(formats strfmt.Registry) error {
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

func (m *JobsIdmSelector) validateRange(formats strfmt.Registry) error {
	if swag.IsZero(m.Range) { // not required
		return nil
	}

	if m.Range != nil {
		if err := m.Range.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Range")
			}
			return err
		}
	}

	return nil
}

func (m *JobsIdmSelector) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this jobs idm selector based on the context it is used
func (m *JobsIdmSelector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRange(ctx, formats); err != nil {
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

func (m *JobsIdmSelector) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

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

func (m *JobsIdmSelector) contextValidateRange(ctx context.Context, formats strfmt.Registry) error {

	if m.Range != nil {

		if swag.IsZero(m.Range) { // not required
			return nil
		}

		if err := m.Range.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Range")
			}
			return err
		}
	}

	return nil
}

func (m *JobsIdmSelector) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *JobsIdmSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsIdmSelector) UnmarshalBinary(b []byte) error {
	var res JobsIdmSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
