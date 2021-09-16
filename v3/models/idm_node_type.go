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

// IdmNodeType idm node type
//
// swagger:model idmNodeType
type IdmNodeType string

func NewIdmNodeType(value IdmNodeType) *IdmNodeType {
	v := value
	return &v
}

const (

	// IdmNodeTypeUNKNOWN captures enum value "UNKNOWN"
	IdmNodeTypeUNKNOWN IdmNodeType = "UNKNOWN"

	// IdmNodeTypeUSER captures enum value "USER"
	IdmNodeTypeUSER IdmNodeType = "USER"

	// IdmNodeTypeGROUP captures enum value "GROUP"
	IdmNodeTypeGROUP IdmNodeType = "GROUP"
)

// for schema
var idmNodeTypeEnum []interface{}

func init() {
	var res []IdmNodeType
	if err := json.Unmarshal([]byte(`["UNKNOWN","USER","GROUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		idmNodeTypeEnum = append(idmNodeTypeEnum, v)
	}
}

func (m IdmNodeType) validateIdmNodeTypeEnum(path, location string, value IdmNodeType) error {
	if err := validate.EnumCase(path, location, value, idmNodeTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this idm node type
func (m IdmNodeType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIdmNodeTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this idm node type based on context it is used
func (m IdmNodeType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
