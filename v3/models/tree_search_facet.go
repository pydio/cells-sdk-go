// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TreeSearchFacet tree search facet
//
// swagger:model treeSearchFacet
type TreeSearchFacet struct {

	// Segment results count
	Count int32 `json:"Count,omitempty"`

	// end
	End int32 `json:"End,omitempty"`

	// Facet field name
	FieldName string `json:"FieldName,omitempty"`

	// Segment Label
	Label string `json:"Label,omitempty"`

	// max
	Max string `json:"Max,omitempty"`

	// For NumericRange facets, min/max values
	Min string `json:"Min,omitempty"`

	// For DateRange facets, start/end values
	Start int32 `json:"Start,omitempty"`

	// For string facets, term value
	Term string `json:"Term,omitempty"`
}

// Validate validates this tree search facet
func (m *TreeSearchFacet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tree search facet based on context it is used
func (m *TreeSearchFacet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TreeSearchFacet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeSearchFacet) UnmarshalBinary(b []byte) error {
	var res TreeSearchFacet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
