// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EncryptionKeyInfo encryption key info
//
// swagger:model encryptionKeyInfo
type EncryptionKeyInfo struct {

	// exports
	Exports []*EncryptionExport `json:"Exports"`

	// imports
	Imports []*EncryptionImport `json:"Imports"`
}

// Validate validates this encryption key info
func (m *EncryptionKeyInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImports(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionKeyInfo) validateExports(formats strfmt.Registry) error {
	if swag.IsZero(m.Exports) { // not required
		return nil
	}

	for i := 0; i < len(m.Exports); i++ {
		if swag.IsZero(m.Exports[i]) { // not required
			continue
		}

		if m.Exports[i] != nil {
			if err := m.Exports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Exports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Exports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EncryptionKeyInfo) validateImports(formats strfmt.Registry) error {
	if swag.IsZero(m.Imports) { // not required
		return nil
	}

	for i := 0; i < len(m.Imports); i++ {
		if swag.IsZero(m.Imports[i]) { // not required
			continue
		}

		if m.Imports[i] != nil {
			if err := m.Imports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Imports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Imports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this encryption key info based on the context it is used
func (m *EncryptionKeyInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImports(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EncryptionKeyInfo) contextValidateExports(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Exports); i++ {

		if m.Exports[i] != nil {
			if err := m.Exports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Exports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Exports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EncryptionKeyInfo) contextValidateImports(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Imports); i++ {

		if m.Imports[i] != nil {
			if err := m.Imports[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Imports" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Imports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EncryptionKeyInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EncryptionKeyInfo) UnmarshalBinary(b []byte) error {
	var res EncryptionKeyInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}