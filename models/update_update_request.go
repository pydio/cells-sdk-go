// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UpdateUpdateRequest update update request
// swagger:model updateUpdateRequest
type UpdateUpdateRequest struct {

	// Channel name
	Channel string `json:"Channel,omitempty"`

	// Current version of the application
	CurrentVersion string `json:"CurrentVersion,omitempty"`

	// Current GOARCH
	GOARCH string `json:"GOARCH,omitempty"`

	// Current GOOS
	GOOS string `json:"GOOS,omitempty"`

	// For enterprise version, info about the current license
	LicenseInfo map[string]string `json:"LicenseInfo,omitempty"`

	// Name of the currently running application
	PackageName string `json:"PackageName,omitempty"`

	// Not Used : specific service to get updates for
	ServiceName string `json:"ServiceName,omitempty"`
}

// Validate validates this update update request
func (m *UpdateUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateUpdateRequest) UnmarshalBinary(b []byte) error {
	var res UpdateUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
