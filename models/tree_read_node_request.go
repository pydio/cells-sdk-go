// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TreeReadNodeRequest Request / Responses Messages
// swagger:model treeReadNodeRequest
type TreeReadNodeRequest struct {

	// Input node
	Node *TreeNode `json:"Node,omitempty"`

	// Used internally for the router ReadNode request, stat the datasource instead of index
	ObjectStats bool `json:"ObjectStats,omitempty"`

	// Gather commit information
	WithCommits bool `json:"WithCommits,omitempty"`

	// Get extended stats - For folders, computes ChildrenCount
	WithExtendedStats bool `json:"WithExtendedStats,omitempty"`
}

// Validate validates this tree read node request
func (m *TreeReadNodeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeReadNodeRequest) validateNode(formats strfmt.Registry) error {

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
func (m *TreeReadNodeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeReadNodeRequest) UnmarshalBinary(b []byte) error {
	var res TreeReadNodeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
