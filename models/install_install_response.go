// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstallInstallResponse install install response
//
// swagger:model installInstallResponse
type InstallInstallResponse struct {

	// success
	Success bool `json:"success,omitempty"`
}

// Validate validates this install install response
func (m *InstallInstallResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this install install response based on context it is used
func (m *InstallInstallResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstallInstallResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstallInstallResponse) UnmarshalBinary(b []byte) error {
	var res InstallInstallResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
