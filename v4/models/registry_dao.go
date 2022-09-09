// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RegistryDao registry dao
//
// swagger:model registryDao
type RegistryDao struct {

	// driver
	Driver string `json:"driver,omitempty"`

	// dsn
	Dsn string `json:"dsn,omitempty"`
}

// Validate validates this registry dao
func (m *RegistryDao) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this registry dao based on context it is used
func (m *RegistryDao) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegistryDao) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistryDao) UnmarshalBinary(b []byte) error {
	var res RegistryDao
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
