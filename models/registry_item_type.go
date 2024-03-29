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

// RegistryItemType registry item type
//
// swagger:model registryItemType
type RegistryItemType string

func NewRegistryItemType(value RegistryItemType) *RegistryItemType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated RegistryItemType.
func (m RegistryItemType) Pointer() *RegistryItemType {
	return &m
}

const (

	// RegistryItemTypeALL captures enum value "ALL"
	RegistryItemTypeALL RegistryItemType = "ALL"

	// RegistryItemTypeEDGE captures enum value "EDGE"
	RegistryItemTypeEDGE RegistryItemType = "EDGE"

	// RegistryItemTypeNODE captures enum value "NODE"
	RegistryItemTypeNODE RegistryItemType = "NODE"

	// RegistryItemTypeSERVICE captures enum value "SERVICE"
	RegistryItemTypeSERVICE RegistryItemType = "SERVICE"

	// RegistryItemTypeSERVER captures enum value "SERVER"
	RegistryItemTypeSERVER RegistryItemType = "SERVER"

	// RegistryItemTypeDAO captures enum value "DAO"
	RegistryItemTypeDAO RegistryItemType = "DAO"

	// RegistryItemTypeGENERIC captures enum value "GENERIC"
	RegistryItemTypeGENERIC RegistryItemType = "GENERIC"

	// RegistryItemTypeADDRESS captures enum value "ADDRESS"
	RegistryItemTypeADDRESS RegistryItemType = "ADDRESS"

	// RegistryItemTypeTAG captures enum value "TAG"
	RegistryItemTypeTAG RegistryItemType = "TAG"

	// RegistryItemTypePROCESS captures enum value "PROCESS"
	RegistryItemTypePROCESS RegistryItemType = "PROCESS"

	// RegistryItemTypeENDPOINT captures enum value "ENDPOINT"
	RegistryItemTypeENDPOINT RegistryItemType = "ENDPOINT"

	// RegistryItemTypeSTATS captures enum value "STATS"
	RegistryItemTypeSTATS RegistryItemType = "STATS"
)

// for schema
var registryItemTypeEnum []interface{}

func init() {
	var res []RegistryItemType
	if err := json.Unmarshal([]byte(`["ALL","EDGE","NODE","SERVICE","SERVER","DAO","GENERIC","ADDRESS","TAG","PROCESS","ENDPOINT","STATS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		registryItemTypeEnum = append(registryItemTypeEnum, v)
	}
}

func (m RegistryItemType) validateRegistryItemTypeEnum(path, location string, value RegistryItemType) error {
	if err := validate.EnumCase(path, location, value, registryItemTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this registry item type
func (m RegistryItemType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRegistryItemTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this registry item type based on context it is used
func (m RegistryItemType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
