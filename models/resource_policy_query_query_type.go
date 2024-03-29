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

// ResourcePolicyQueryQueryType resource policy query query type
//
// swagger:model ResourcePolicyQueryQueryType
type ResourcePolicyQueryQueryType string

func NewResourcePolicyQueryQueryType(value ResourcePolicyQueryQueryType) *ResourcePolicyQueryQueryType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ResourcePolicyQueryQueryType.
func (m ResourcePolicyQueryQueryType) Pointer() *ResourcePolicyQueryQueryType {
	return &m
}

const (

	// ResourcePolicyQueryQueryTypeCONTEXT captures enum value "CONTEXT"
	ResourcePolicyQueryQueryTypeCONTEXT ResourcePolicyQueryQueryType = "CONTEXT"

	// ResourcePolicyQueryQueryTypeANY captures enum value "ANY"
	ResourcePolicyQueryQueryTypeANY ResourcePolicyQueryQueryType = "ANY"

	// ResourcePolicyQueryQueryTypeNONE captures enum value "NONE"
	ResourcePolicyQueryQueryTypeNONE ResourcePolicyQueryQueryType = "NONE"

	// ResourcePolicyQueryQueryTypeUSER captures enum value "USER"
	ResourcePolicyQueryQueryTypeUSER ResourcePolicyQueryQueryType = "USER"
)

// for schema
var resourcePolicyQueryQueryTypeEnum []interface{}

func init() {
	var res []ResourcePolicyQueryQueryType
	if err := json.Unmarshal([]byte(`["CONTEXT","ANY","NONE","USER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourcePolicyQueryQueryTypeEnum = append(resourcePolicyQueryQueryTypeEnum, v)
	}
}

func (m ResourcePolicyQueryQueryType) validateResourcePolicyQueryQueryTypeEnum(path, location string, value ResourcePolicyQueryQueryType) error {
	if err := validate.EnumCase(path, location, value, resourcePolicyQueryQueryTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this resource policy query query type
func (m ResourcePolicyQueryQueryType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateResourcePolicyQueryQueryTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this resource policy query query type based on context it is used
func (m ResourcePolicyQueryQueryType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
