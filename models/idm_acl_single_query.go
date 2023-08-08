// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IdmACLSingleQuery idm ACL single query
//
// swagger:model idmACLSingleQuery
type IdmACLSingleQuery struct {

	// List of Actions to lookup
	Actions []*IdmACLAction `json:"Actions"`

	// Filter by one or more node UUID
	NodeIDs []string `json:"NodeIDs"`

	// Filter by one or more role UUID
	RoleIDs []string `json:"RoleIDs"`

	// Filter by one or more workspace UUID
	WorkspaceIDs []string `json:"WorkspaceIDs"`

	// Internal - Negate query
	Not bool `json:"not,omitempty"`
}

// Validate validates this idm ACL single query
func (m *IdmACLSingleQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmACLSingleQuery) validateActions(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this idm ACL single query based on the context it is used
func (m *IdmACLSingleQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmACLSingleQuery) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Actions); i++ {

		if m.Actions[i] != nil {

			if swag.IsZero(m.Actions[i]) { // not required
				return nil
			}

			if err := m.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdmACLSingleQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmACLSingleQuery) UnmarshalBinary(b []byte) error {
	var res IdmACLSingleQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}