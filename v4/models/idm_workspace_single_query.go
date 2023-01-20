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

// IdmWorkspaceSingleQuery idm workspace single query
//
// swagger:model idmWorkspaceSingleQuery
type IdmWorkspaceSingleQuery struct {

	// Lookup for an attribute, to be used in combination with AttributeValue
	AttributeName string `json:"AttributeName,omitempty"`

	// Value used for comparison
	AttributeValue string `json:"AttributeValue,omitempty"`

	// Lookup for the presence of a specific attribute, whatever its value
	HasAttribute string `json:"HasAttribute,omitempty"`

	// Compared to workspace update date (Golang duration with a leading comparator < or >)
	LastUpdated string `json:"LastUpdated,omitempty"`

	// Lookup by description
	Description string `json:"description,omitempty"`

	// Lookup by workspace Label
	Label string `json:"label,omitempty"`

	// Internal - Negate the query result
	Not bool `json:"not,omitempty"`

	// Restrict to a specific workspace type
	Scope *IdmWorkspaceScope `json:"scope,omitempty"`

	// Select workspace by slug
	Slug string `json:"slug,omitempty"`

	// Lookup by workspace Uuid
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

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this idm workspace single query based on the context it is used
func (m *IdmWorkspaceSingleQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmWorkspaceSingleQuery) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {
		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
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
