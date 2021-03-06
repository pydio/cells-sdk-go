// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EncryptionExport encryption export
// swagger:model encryptionExport
type EncryptionExport struct {

	// Name of exporter
	By string `json:"By,omitempty"`

	// Date of export
	Date int32 `json:"Date,omitempty"`
}

// Validate validates this encryption export
func (m *EncryptionExport) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionExport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionExport) UnmarshalBinary(b []byte) error {
	var res EncryptionExport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
