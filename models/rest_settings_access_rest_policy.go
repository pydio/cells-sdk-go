// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestSettingsAccessRestPolicy rest settings access rest policy
// swagger:model restSettingsAccessRestPolicy
type RestSettingsAccessRestPolicy struct {

	// action
	Action string `json:"Action,omitempty"`

	// resource
	Resource string `json:"Resource,omitempty"`
}

// Validate validates this rest settings access rest policy
func (m *RestSettingsAccessRestPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestSettingsAccessRestPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestSettingsAccessRestPolicy) UnmarshalBinary(b []byte) error {
	var res RestSettingsAccessRestPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}