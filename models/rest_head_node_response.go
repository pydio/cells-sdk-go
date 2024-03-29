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

// RestHeadNodeResponse rest head node response
//
// swagger:model restHeadNodeResponse
type RestHeadNodeResponse struct {

	// node
	Node *TreeNode `json:"Node,omitempty"`
}

// Validate validates this rest head node response
func (m *RestHeadNodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestHeadNodeResponse) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rest head node response based on the context it is used
func (m *RestHeadNodeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestHeadNodeResponse) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {

		if swag.IsZero(m.Node) { // not required
			return nil
		}

		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Node")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestHeadNodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestHeadNodeResponse) UnmarshalBinary(b []byte) error {
	var res RestHeadNodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
