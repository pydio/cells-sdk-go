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

// TreeNodeChangeEvent tree node change event
//
// swagger:model treeNodeChangeEvent
type TreeNodeChangeEvent struct {

	// metadata
	Metadata map[string]string `json:"Metadata,omitempty"`

	// optimistic
	Optimistic bool `json:"Optimistic,omitempty"`

	// silent
	Silent bool `json:"Silent,omitempty"`

	// source
	Source *TreeNode `json:"Source,omitempty"`

	// target
	Target *TreeNode `json:"Target,omitempty"`

	// type
	Type *TreeNodeChangeEventEventType `json:"Type,omitempty"`
}

// Validate validates this tree node change event
func (m *TreeNodeChangeEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
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

func (m *TreeNodeChangeEvent) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Source")
			}
			return err
		}
	}

	return nil
}

func (m *TreeNodeChangeEvent) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Target")
			}
			return err
		}
	}

	return nil
}

func (m *TreeNodeChangeEvent) validateType(formats strfmt.Registry) error {
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

// ContextValidate validate this tree node change event based on the context it is used
func (m *TreeNodeChangeEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
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

func (m *TreeNodeChangeEvent) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if swag.IsZero(m.Source) { // not required
			return nil
		}

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Source")
			}
			return err
		}
	}

	return nil
}

func (m *TreeNodeChangeEvent) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.Target != nil {

		if swag.IsZero(m.Target) { // not required
			return nil
		}

		if err := m.Target.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Target")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Target")
			}
			return err
		}
	}

	return nil
}

func (m *TreeNodeChangeEvent) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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
func (m *TreeNodeChangeEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeNodeChangeEvent) UnmarshalBinary(b []byte) error {
	var res TreeNodeChangeEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
