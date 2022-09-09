// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EncryptionAdminDeleteKeyRequest encryption admin delete key request
//
// swagger:model encryptionAdminDeleteKeyRequest
type EncryptionAdminDeleteKeyRequest struct {

	// Id of the key to delete
	KeyID string `json:"KeyID,omitempty"`
}

// Validate validates this encryption admin delete key request
func (m *EncryptionAdminDeleteKeyRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this encryption admin delete key request based on context it is used
func (m *EncryptionAdminDeleteKeyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionAdminDeleteKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionAdminDeleteKeyRequest) UnmarshalBinary(b []byte) error {
	var res EncryptionAdminDeleteKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
