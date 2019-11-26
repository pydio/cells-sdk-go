// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IdmACL ACL are the basic flags that can be put anywhere in the tree to provide some specific rights to a given role.
// The context of how they apply can be fine-tuned by workspace.
// swagger:model idmACL
type IdmACL struct {

	// Action on which this ACL provides control
	Action *IdmACLAction `json:"Action,omitempty"`

	// Unique ID of this ACL
	ID string `json:"ID,omitempty"`

	// Associated Node
	NodeID string `json:"NodeID,omitempty"`

	// Associated Role
	RoleID string `json:"RoleID,omitempty"`

	// Associated Workspace
	WorkspaceID string `json:"WorkspaceID,omitempty"`
}

// Validate validates this idm ACL
func (m *IdmACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmACL) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Action")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdmACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmACL) UnmarshalBinary(b []byte) error {
	var res IdmACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
