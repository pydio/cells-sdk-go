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

// UpdatePackage update package
//
// swagger:model updatePackage
type UpdatePackage struct {

	// GOARCH value used at build time
	BinaryArch string `json:"BinaryArch,omitempty"`

	// Checksum of the binary to verify its integrity
	BinaryChecksum string `json:"BinaryChecksum,omitempty"`

	// Hash type used for the signature
	BinaryHashType string `json:"BinaryHashType,omitempty"`

	// GOOS value used at build time
	BinaryOS string `json:"BinaryOS,omitempty"`

	// Signature of the binary
	BinarySignature string `json:"BinarySignature,omitempty"`

	// Size of the binary to download
	BinarySize string `json:"BinarySize,omitempty"`

	// Https URL where to download the binary
	BinaryURL string `json:"BinaryURL,omitempty"`

	// List or public URL of change logs
	ChangeLog string `json:"ChangeLog,omitempty"`

	// Long human-readable description (markdown)
	Description string `json:"Description,omitempty"`

	// Not used : if binary is a patch
	IsPatch bool `json:"IsPatch,omitempty"`

	// Short human-readable description
	Label string `json:"Label,omitempty"`

	// License of this package
	License string `json:"License,omitempty"`

	// Name of the application
	PackageName string `json:"PackageName,omitempty"`

	// Not used : if a patch, how to patch (bsdiff support)
	PatchAlgorithm string `json:"PatchAlgorithm,omitempty"`

	// Release date of the binary
	ReleaseDate int32 `json:"ReleaseDate,omitempty"`

	// Not used : at a point we may deliver services only updates
	ServiceName string `json:"ServiceName,omitempty"`

	// status
	Status *PackagePackageStatus `json:"Status,omitempty"`

	// Version of this new binary
	Version string `json:"Version,omitempty"`
}

// Validate validates this update package
func (m *UpdatePackage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePackage) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update package based on the context it is used
func (m *UpdatePackage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePackage) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

		if swag.IsZero(m.Status) { // not required
			return nil
		}

		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePackage) UnmarshalBinary(b []byte) error {
	var res UpdatePackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
