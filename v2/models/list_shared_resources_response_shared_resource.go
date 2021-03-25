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

// ListSharedResourcesResponseSharedResource Container for ShareLink or Cell
// swagger:model ListSharedResourcesResponseSharedResource
type ListSharedResourcesResponseSharedResource struct {

	// cells
	Cells []*RestCell `json:"Cells"`

	// link
	Link *RestShareLink `json:"Link,omitempty"`

	// node
	Node *TreeNode `json:"Node,omitempty"`
}

// Validate validates this list shared resources response shared resource
func (m *ListSharedResourcesResponseSharedResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCells(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListSharedResourcesResponseSharedResource) validateCells(formats strfmt.Registry) error {

	if swag.IsZero(m.Cells) { // not required
		return nil
	}

	for i := 0; i < len(m.Cells); i++ {
		if swag.IsZero(m.Cells[i]) { // not required
			continue
		}

		if m.Cells[i] != nil {
			if err := m.Cells[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Cells" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListSharedResourcesResponseSharedResource) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if m.Link != nil {
		if err := m.Link.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Link")
			}
			return err
		}
	}

	return nil
}

func (m *ListSharedResourcesResponseSharedResource) validateNode(formats strfmt.Registry) error {

	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListSharedResourcesResponseSharedResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListSharedResourcesResponseSharedResource) UnmarshalBinary(b []byte) error {
	var res ListSharedResourcesResponseSharedResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
