// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IdmRoleSingleQuery RoleSingleQuery is the basic unit for building queries to Roles.
//
// swagger:model idmRoleSingleQuery
type IdmRoleSingleQuery struct {

	// Look for roles that have any value in the autoApplies field
	HasAutoApply bool `json:"HasAutoApply,omitempty"`

	// Look for roles associated with a Group
	IsGroupRole bool `json:"IsGroupRole,omitempty"`

	// Look up for roles associated with a Team
	IsTeam bool `json:"IsTeam,omitempty"`

	// Look for roles associated with a User
	IsUserRole bool `json:"IsUserRole,omitempty"`

	// Look for roles by label, eventually using "wildchar"
	Label string `json:"Label,omitempty"`

	// Look for roles by Uuid
	UUID []string `json:"Uuid"`

	// Internal - Negate the query
	Not bool `json:"not,omitempty"`
}

// Validate validates this idm role single query
func (m *IdmRoleSingleQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this idm role single query based on context it is used
func (m *IdmRoleSingleQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IdmRoleSingleQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmRoleSingleQuery) UnmarshalBinary(b []byte) error {
	var res IdmRoleSingleQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
