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

// TreeChangeLog tree change log
//
// swagger:model treeChangeLog
type TreeChangeLog struct {

	// Arbitrary additional data
	// Format: byte
	Data strfmt.Base64 `json:"Data,omitempty"`

	// Human-readable description of what happened
	Description string `json:"Description,omitempty"`

	// Event that triggered this change
	Event *TreeNodeChangeEvent `json:"Event,omitempty"`

	// Actual location of the stored version
	Location *TreeNode `json:"Location,omitempty"`

	// Unix Timestamp
	MTime string `json:"MTime,omitempty"`

	// Who performed this action
	OwnerUUID string `json:"OwnerUuid,omitempty"`

	// Content Size at that moment
	Size string `json:"Size,omitempty"`

	// Unique commit ID
	UUID string `json:"Uuid,omitempty"`
}

// Validate validates this tree change log
func (m *TreeChangeLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeChangeLog) validateEvent(formats strfmt.Registry) error {
	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {
		if err := m.Event.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Event")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Event")
			}
			return err
		}
	}

	return nil
}

func (m *TreeChangeLog) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tree change log based on the context it is used
func (m *TreeChangeLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeChangeLog) contextValidateEvent(ctx context.Context, formats strfmt.Registry) error {

	if m.Event != nil {
		if err := m.Event.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Event")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Event")
			}
			return err
		}
	}

	return nil
}

func (m *TreeChangeLog) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreeChangeLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeChangeLog) UnmarshalBinary(b []byte) error {
	var res TreeChangeLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
