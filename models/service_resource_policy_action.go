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

// ServiceResourcePolicyAction service resource policy action
//
// swagger:model serviceResourcePolicyAction
type ServiceResourcePolicyAction string

func NewServiceResourcePolicyAction(value ServiceResourcePolicyAction) *ServiceResourcePolicyAction {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ServiceResourcePolicyAction.
func (m ServiceResourcePolicyAction) Pointer() *ServiceResourcePolicyAction {
	return &m
}

const (

	// ServiceResourcePolicyActionANY captures enum value "ANY"
	ServiceResourcePolicyActionANY ServiceResourcePolicyAction = "ANY"

	// ServiceResourcePolicyActionOWNER captures enum value "OWNER"
	ServiceResourcePolicyActionOWNER ServiceResourcePolicyAction = "OWNER"

	// ServiceResourcePolicyActionREAD captures enum value "READ"
	ServiceResourcePolicyActionREAD ServiceResourcePolicyAction = "READ"

	// ServiceResourcePolicyActionWRITE captures enum value "WRITE"
	ServiceResourcePolicyActionWRITE ServiceResourcePolicyAction = "WRITE"

	// ServiceResourcePolicyActionEDITRULES captures enum value "EDIT_RULES"
	ServiceResourcePolicyActionEDITRULES ServiceResourcePolicyAction = "EDIT_RULES"
)

// for schema
var serviceResourcePolicyActionEnum []interface{}

func init() {
	var res []ServiceResourcePolicyAction
	if err := json.Unmarshal([]byte(`["ANY","OWNER","READ","WRITE","EDIT_RULES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		serviceResourcePolicyActionEnum = append(serviceResourcePolicyActionEnum, v)
	}
}

func (m ServiceResourcePolicyAction) validateServiceResourcePolicyActionEnum(path, location string, value ServiceResourcePolicyAction) error {
	if err := validate.EnumCase(path, location, value, serviceResourcePolicyActionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this service resource policy action
func (m ServiceResourcePolicyAction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateServiceResourcePolicyActionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this service resource policy action based on context it is used
func (m ServiceResourcePolicyAction) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}