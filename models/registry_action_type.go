// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// RegistryActionType registry action type
//
// swagger:model registryActionType
type RegistryActionType string

func NewRegistryActionType(value RegistryActionType) *RegistryActionType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated RegistryActionType.
func (m RegistryActionType) Pointer() *RegistryActionType {
	return &m
}

const (

	// RegistryActionTypeANY captures enum value "ANY"
	RegistryActionTypeANY RegistryActionType = "ANY"

	// RegistryActionTypeFULLDIFF captures enum value "FULL_DIFF"
	RegistryActionTypeFULLDIFF RegistryActionType = "FULL_DIFF"

	// RegistryActionTypeFULLLIST captures enum value "FULL_LIST"
	RegistryActionTypeFULLLIST RegistryActionType = "FULL_LIST"

	// RegistryActionTypeCREATE captures enum value "CREATE"
	RegistryActionTypeCREATE RegistryActionType = "CREATE"

	// RegistryActionTypeUPDATE captures enum value "UPDATE"
	RegistryActionTypeUPDATE RegistryActionType = "UPDATE"

	// RegistryActionTypeDELETE captures enum value "DELETE"
	RegistryActionTypeDELETE RegistryActionType = "DELETE"
)

// for schema
var registryActionTypeEnum []interface{}

func init() {
	var res []RegistryActionType
	if err := json.Unmarshal([]byte(`["ANY","FULL_DIFF","FULL_LIST","CREATE","UPDATE","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		registryActionTypeEnum = append(registryActionTypeEnum, v)
	}
}

func (m RegistryActionType) validateRegistryActionTypeEnum(path, location string, value RegistryActionType) error {
	if err := validate.EnumCase(path, location, value, registryActionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this registry action type
func (m RegistryActionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRegistryActionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this registry action type based on context it is used
func (m RegistryActionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
