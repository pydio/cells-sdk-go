// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RestCellACL Group collected acls by subjects
// swagger:model restCellAcl
type RestCellACL struct {

	// List of Acl Actions and their effect
	Actions []*IdmACLAction `json:"Actions"`

	// Associated Group
	Group *IdmUser `json:"Group,omitempty"`

	// Flag for detecting if it's a user role or not
	IsUserRole bool `json:"IsUserRole,omitempty"`

	// Associated Role
	Role *IdmRole `json:"Role,omitempty"`

	// Associated Role ID
	RoleID string `json:"RoleId,omitempty"`

	// Associated User
	User *IdmUser `json:"User,omitempty"`
}

// Validate validates this rest cell Acl
func (m *RestCellACL) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestCellACL) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestCellACL) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Group")
			}
			return err
		}
	}

	return nil
}

func (m *RestCellACL) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Role")
			}
			return err
		}
	}

	return nil
}

func (m *RestCellACL) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("User")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestCellACL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestCellACL) UnmarshalBinary(b []byte) error {
	var res RestCellACL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
