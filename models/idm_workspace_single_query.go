// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IdmWorkspaceSingleQuery idm workspace single query
// swagger:model idmWorkspaceSingleQuery
type IdmWorkspaceSingleQuery struct {

	// attribute name
	AttributeName string `json:"AttributeName,omitempty"`

	// attribute value
	AttributeValue string `json:"AttributeValue,omitempty"`

	// has attribute
	HasAttribute string `json:"HasAttribute,omitempty"`

	// last updated
	LastUpdated string `json:"LastUpdated,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// not
	Not bool `json:"not,omitempty"`

	// scope
	Scope IdmWorkspaceScope `json:"scope,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this idm workspace single query
func (m *IdmWorkspaceSingleQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmWorkspaceSingleQuery) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := m.Scope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdmWorkspaceSingleQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmWorkspaceSingleQuery) UnmarshalBinary(b []byte) error {
	var res IdmWorkspaceSingleQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
