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

// RestNodesCollection rest nodes collection
// swagger:model restNodesCollection
type RestNodesCollection struct {

	// children
	Children []*TreeNode `json:"Children"`

	// parent
	Parent *TreeNode `json:"Parent,omitempty"`
}

// Validate validates this rest nodes collection
func (m *RestNodesCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestNodesCollection) validateChildren(formats strfmt.Registry) error {

	if swag.IsZero(m.Children) { // not required
		return nil
	}

	for i := 0; i < len(m.Children); i++ {
		if swag.IsZero(m.Children[i]) { // not required
			continue
		}

		if m.Children[i] != nil {
			if err := m.Children[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Children" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestNodesCollection) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Parent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestNodesCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestNodesCollection) UnmarshalBinary(b []byte) error {
	var res RestNodesCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
