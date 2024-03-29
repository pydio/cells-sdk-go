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

// ObjectDataSource DataSource Object description
//
// swagger:model objectDataSource
type ObjectDataSource struct {

	// Corresponding objects service api key
	APIKey string `json:"ApiKey,omitempty"`

	// Corresponding objects service api secret
	APISecret string `json:"ApiSecret,omitempty"`

	// Data Source creation date
	CreationDate int32 `json:"CreationDate,omitempty"`

	// Whether this data source is disabled or running
	Disabled bool `json:"Disabled,omitempty"`

	// Encryption key used for encrypting data
	EncryptionKey string `json:"EncryptionKey,omitempty"`

	// Type of encryption applied before sending data to storage
	EncryptionMode *ObjectEncryptionMode `json:"EncryptionMode,omitempty"`

	// Store data in flat format (object-storage like)
	FlatStorage bool `json:"FlatStorage,omitempty"`

	// Data Source last synchronization date
	LastSynchronizationDate int32 `json:"LastSynchronizationDate,omitempty"`

	// Name of the data source (max length 34)
	Name string `json:"Name,omitempty"`

	// Corresponding objects service base folder inside the bucket
	ObjectsBaseFolder string `json:"ObjectsBaseFolder,omitempty"`

	// Corresponding objects service bucket
	ObjectsBucket string `json:"ObjectsBucket,omitempty"`

	// Corresponding objects service host
	ObjectsHost string `json:"ObjectsHost,omitempty"`

	// Corresponding objects service port
	ObjectsPort int32 `json:"ObjectsPort,omitempty"`

	// Corresponding objects service connection type
	ObjectsSecure bool `json:"ObjectsSecure,omitempty"`

	// Corresponding objects service name (underlying s3 service)
	ObjectsServiceName string `json:"ObjectsServiceName,omitempty"`

	// Peer address of the data source
	PeerAddress string `json:"PeerAddress,omitempty"`

	// Do not trigger resync at start
	SkipSyncOnRestart bool `json:"SkipSyncOnRestart,omitempty"`

	// List of key values describing storage configuration
	StorageConfiguration map[string]string `json:"StorageConfiguration,omitempty"`

	// Type of underlying storage (LOCAL, S3, AZURE, GCS)
	StorageType *ObjectStorageType `json:"StorageType,omitempty"`

	// Versioning policy describes how files are kept in the versioning queue
	VersioningPolicyName string `json:"VersioningPolicyName,omitempty"`

	// Not implemented, whether to watch for underlying changes on the FS
	Watch bool `json:"Watch,omitempty"`
}

// Validate validates this object data source
func (m *ObjectDataSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryptionMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectDataSource) validateEncryptionMode(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionMode) { // not required
		return nil
	}

	if m.EncryptionMode != nil {
		if err := m.EncryptionMode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EncryptionMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EncryptionMode")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectDataSource) validateStorageType(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageType) { // not required
		return nil
	}

	if m.StorageType != nil {
		if err := m.StorageType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("StorageType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this object data source based on the context it is used
func (m *ObjectDataSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEncryptionMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ObjectDataSource) contextValidateEncryptionMode(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionMode != nil {

		if swag.IsZero(m.EncryptionMode) { // not required
			return nil
		}

		if err := m.EncryptionMode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EncryptionMode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("EncryptionMode")
			}
			return err
		}
	}

	return nil
}

func (m *ObjectDataSource) contextValidateStorageType(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageType != nil {

		if swag.IsZero(m.StorageType) { // not required
			return nil
		}

		if err := m.StorageType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("StorageType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("StorageType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ObjectDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectDataSource) UnmarshalBinary(b []byte) error {
	var res ObjectDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
