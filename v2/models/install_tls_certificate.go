// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InstallTLSCertificate TLSCertificate is a TLSConfig where user passes
// swagger:model installTLSCertificate
type InstallTLSCertificate struct {

	// cells root c a
	CellsRootCA string `json:"CellsRootCA,omitempty"`

	// cert file
	CertFile string `json:"CertFile,omitempty"`

	// key file
	KeyFile string `json:"KeyFile,omitempty"`
}

// Validate validates this install TLS certificate
func (m *InstallTLSCertificate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstallTLSCertificate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstallTLSCertificate) UnmarshalBinary(b []byte) error {
	var res InstallTLSCertificate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
