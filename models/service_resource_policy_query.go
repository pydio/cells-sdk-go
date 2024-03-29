// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceResourcePolicyQuery service resource policy query
//
// swagger:model serviceResourcePolicyQuery
type ServiceResourcePolicyQuery struct {

	// any
	Any bool `json:"Any,omitempty"`

	// empty
	Empty bool `json:"Empty,omitempty"`

	// subjects
	Subjects []string `json:"Subjects"`
}

// Validate validates this service resource policy query
func (m *ServiceResourcePolicyQuery) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service resource policy query based on context it is used
func (m *ServiceResourcePolicyQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceResourcePolicyQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceResourcePolicyQuery) UnmarshalBinary(b []byte) error {
	var res ServiceResourcePolicyQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
