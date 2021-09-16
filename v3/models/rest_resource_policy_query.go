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

// RestResourcePolicyQuery Generic Query for limiting results based on resource permissions
//
// swagger:model restResourcePolicyQuery
type RestResourcePolicyQuery struct {

	// The type can be CONTEXT, ANY, NODE or USER. This restricts the may filter out the result set based on their policies
	Type *ResourcePolicyQueryQueryType `json:"Type,omitempty"`

	// Limit to one given user ID
	UserID string `json:"UserId,omitempty"`
}

// Validate validates this rest resource policy query
func (m *RestResourcePolicyQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestResourcePolicyQuery) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rest resource policy query based on the context it is used
func (m *RestResourcePolicyQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestResourcePolicyQuery) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestResourcePolicyQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestResourcePolicyQuery) UnmarshalBinary(b []byte) error {
	var res RestResourcePolicyQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
