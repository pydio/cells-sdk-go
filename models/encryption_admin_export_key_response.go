// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EncryptionAdminExportKeyResponse encryption admin export key response
//
// swagger:model encryptionAdminExportKeyResponse
type EncryptionAdminExportKeyResponse struct {

	// key
	Key *EncryptionKey `json:"Key,omitempty"`
}

// Validate validates this encryption admin export key response
func (m *EncryptionAdminExportKeyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionAdminExportKeyResponse) validateKey(formats strfmt.Registry) error {
	if swag.IsZero(m.Key) { // not required
		return nil
	}

	if m.Key != nil {
		if err := m.Key.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this encryption admin export key response based on the context it is used
func (m *EncryptionAdminExportKeyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionAdminExportKeyResponse) contextValidateKey(ctx context.Context, formats strfmt.Registry) error {

	if m.Key != nil {

		if swag.IsZero(m.Key) { // not required
			return nil
		}

		if err := m.Key.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionAdminExportKeyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionAdminExportKeyResponse) UnmarshalBinary(b []byte) error {
	var res EncryptionAdminExportKeyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
