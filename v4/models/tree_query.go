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

// TreeQuery Search Queries
//
// swagger:model treeQuery
type TreeQuery struct {

	// Search in textual content (if search engine has this feature enabled)
	Content string `json:"Content,omitempty"`

	// Compute MinDate/MaxDate with a Golang duration with a leading comparator (> or <)
	// Duration may contain "s" second, "m" minute, "d" day.
	// Example: ">10m" for files modified before 10minutes ago
	DurationDate string `json:"DurationDate,omitempty"`

	// Look for a specific ETag value, may only be useful to lookup for files with __temporary__ ETag
	ETag string `json:"ETag,omitempty"`

	// Search files by their extension, use pipe symbol | if you wish to allow many extensions.
	// Example png|pdf|jpg
	Extension string `json:"Extension,omitempty"`

	// Lookup by file basename
	FileName string `json:"FileName,omitempty"`

	// Search in either filename or content (if search engine has this feature enabled)
	FileNameOrContent string `json:"FileNameOrContent,omitempty"`

	// Bleve-like search query to search for a specific metadata value.
	// When querying nodes, this will redirect this query to the Search Engine. When filtering an input, this will load an in-memory bleve engine to evaluate the node.
	//
	// Bleve query string format is a space separated list of `[+-]key:value`, where node meta keys must be prepended with "Meta."
	// For Example, for tags: `+Meta.usermeta-tags:myvalue`
	FreeString string `json:"FreeString,omitempty"`

	// Search geographically
	GeoQuery *TreeGeoQuery `json:"GeoQuery,omitempty"`

	// Range for modification date - node was modified before this date
	MaxDate string `json:"MaxDate,omitempty"`

	// Range for file size - size is smaller than
	MaxSize string `json:"MaxSize,omitempty"`

	// Range for modification date - node was modified after this date
	MinDate string `json:"MinDate,omitempty"`

	// Range for file size - size bigger than
	MinSize string `json:"MinSize,omitempty"`

	// Negate this query
	Not bool `json:"Not,omitempty"`

	// Restrict recursive listing to a given level of the tree starting from root.
	// Special value "-1" should list only one level in the folder defined by PathPrefix
	PathDepth int32 `json:"PathDepth,omitempty"`

	// Recursive listing of nodes below a given path. Combine with the PathDepth parameter to limit request results
	PathPrefix []string `json:"PathPrefix"`

	// List of nodes paths, exactly matching
	Paths []string `json:"Paths"`

	// Limit to a given node type (file or folder)
	Type *TreeNodeType `json:"Type,omitempty"`

	// Preset list of specific node defined by their UUIDs
	UUIDs []string `json:"UUIDs"`
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GeoQuery")
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

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tree query based on the context it is used
func (m *TreeQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGeoQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeQuery) contextValidateGeoQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.GeoQuery != nil {
		if err := m.GeoQuery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("GeoQuery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("GeoQuery")
			}
			return err
		}
	}

	return nil
}

func (m *TreeQuery) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Type")
			}
			return err
		}
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
