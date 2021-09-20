// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TreeQuery Search Queries
// swagger:model treeQuery
type TreeQuery struct {

	// Search in content
	Content string `json:"Content,omitempty"`

	// Pass a duration with > or < to compute MinDate / MaxDate
	DurationDate string `json:"DurationDate,omitempty"`

	// Search files by extension
	Extension string `json:"Extension,omitempty"`

	// Search in filename
	FileName string `json:"FileName,omitempty"`

	// Search in either filename or content (but at least one of them)
	FileNameOrContent string `json:"FileNameOrContent,omitempty"`

	// Free Query String (for metadata)
	FreeString string `json:"FreeString,omitempty"`

	// Search geographically
	GeoQuery *TreeGeoQuery `json:"GeoQuery,omitempty"`

	// max date
	MaxDate string `json:"MaxDate,omitempty"`

	// max size
	MaxSize string `json:"MaxSize,omitempty"`

	// Range for date
	MinDate string `json:"MinDate,omitempty"`

	// Range for size
	MinSize string `json:"MinSize,omitempty"`

	// Negate this query
	Not bool `json:"Not,omitempty"`

	// Limit to a given level of the tree - can be used in filters
	PathDepth int32 `json:"PathDepth,omitempty"`

	// Limit to a given subtree
	PathPrefix []string `json:"PathPrefix"`

	// Preset list of nodes by Path
	Paths []string `json:"Paths"`

	// Limit to a given node type
	Type TreeNodeType `json:"Type,omitempty"`

	// Preset list of Node by UUIDs
	Uuids []string `json:"UUIDs"`
}

// Validate validates this tree query
func (m *TreeQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeoQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeQuery) validateGeoQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.GeoQuery) { // not required
		return nil
	}

	if m.GeoQuery != nil {
		if err := m.GeoQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GeoQuery")
			}
			return err
		}
	}

	return nil
}

func (m *TreeQuery) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreeQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeQuery) UnmarshalBinary(b []byte) error {
	var res TreeQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}