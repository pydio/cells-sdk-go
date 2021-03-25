// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EncryptionAdminCreateKeyRequest encryption admin create key request
// swagger:model encryptionAdminCreateKeyRequest
type EncryptionAdminCreateKeyRequest struct {

	// Create a key with this ID
	KeyID string `json:"KeyID,omitempty"`

	// Provide label for the newly created key
	Label string `json:"Label,omitempty"`
}

// Validate validates this encryption admin create key request
func (m *EncryptionAdminCreateKeyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionAdminCreateKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionAdminCreateKeyRequest) UnmarshalBinary(b []byte) error {
	var res EncryptionAdminCreateKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
