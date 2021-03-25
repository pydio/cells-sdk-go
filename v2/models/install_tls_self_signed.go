// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InstallTLSSelfSigned TLSSelfSigned generates a selfsigned certificate
// swagger:model installTLSSelfSigned
type InstallTLSSelfSigned struct {

	// hostnames
	Hostnames []string `json:"Hostnames"`
}

// Validate validates this install TLS self signed
func (m *InstallTLSSelfSigned) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstallTLSSelfSigned) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstallTLSSelfSigned) UnmarshalBinary(b []byte) error {
	var res InstallTLSSelfSigned
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
