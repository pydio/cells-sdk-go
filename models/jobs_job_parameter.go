// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobsJobParameter jobs job parameter
//
// swagger:model jobsJobParameter
type JobsJobParameter struct {

	// Additional description
	Description string `json:"Description,omitempty"`

	// Additional data used by GUI elements
	JSONChoices string `json:"JsonChoices,omitempty"`

	// If mandatory, job cannot start without a value
	Mandatory bool `json:"Mandatory,omitempty"`

	// Parameter name
	Name string `json:"Name,omitempty"`

	// Parameter type used in GUI forms
	Type string `json:"Type,omitempty"`

	// Value saved for this parameter
	Value string `json:"Value,omitempty"`
}

// Validate validates this jobs job parameter
func (m *JobsJobParameter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this jobs job parameter based on context it is used
func (m *JobsJobParameter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobsJobParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsJobParameter) UnmarshalBinary(b []byte) error {
	var res JobsJobParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
