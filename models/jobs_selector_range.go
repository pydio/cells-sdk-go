// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobsSelectorRange jobs selector range
//
// swagger:model jobsSelectorRange
type JobsSelectorRange struct {

	// Limit number of results - use string to allow templating
	Limit string `json:"Limit,omitempty"`

	// Start offset - use string to allow templating
	Offset string `json:"Offset,omitempty"`

	// OrderBy a given field of the object
	OrderBy string `json:"OrderBy,omitempty"`

	// Order direction (asc/desc)
	OrderDir string `json:"OrderDir,omitempty"`
}

// Validate validates this jobs selector range
func (m *JobsSelectorRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this jobs selector range based on context it is used
func (m *JobsSelectorRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobsSelectorRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsSelectorRange) UnmarshalBinary(b []byte) error {
	var res JobsSelectorRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}