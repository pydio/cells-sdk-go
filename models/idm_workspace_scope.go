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

// IdmWorkspaceScope idm workspace scope
//
// swagger:model idmWorkspaceScope
type IdmWorkspaceScope string

func NewIdmWorkspaceScope(value IdmWorkspaceScope) *IdmWorkspaceScope {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IdmWorkspaceScope.
func (m IdmWorkspaceScope) Pointer() *IdmWorkspaceScope {
	return &m
}

const (

	// IdmWorkspaceScopeANY captures enum value "ANY"
	IdmWorkspaceScopeANY IdmWorkspaceScope = "ANY"

	// IdmWorkspaceScopeADMIN captures enum value "ADMIN"
	IdmWorkspaceScopeADMIN IdmWorkspaceScope = "ADMIN"

	// IdmWorkspaceScopeROOM captures enum value "ROOM"
	IdmWorkspaceScopeROOM IdmWorkspaceScope = "ROOM"

	// IdmWorkspaceScopeLINK captures enum value "LINK"
	IdmWorkspaceScopeLINK IdmWorkspaceScope = "LINK"
)

// for schema
var idmWorkspaceScopeEnum []interface{}

func init() {
	var res []IdmWorkspaceScope
	if err := json.Unmarshal([]byte(`["ANY","ADMIN","ROOM","LINK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		idmWorkspaceScopeEnum = append(idmWorkspaceScopeEnum, v)
	}
}

func (m IdmWorkspaceScope) validateIdmWorkspaceScopeEnum(path, location string, value IdmWorkspaceScope) error {
	if err := validate.EnumCase(path, location, value, idmWorkspaceScopeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this idm workspace scope
func (m IdmWorkspaceScope) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIdmWorkspaceScopeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this idm workspace scope based on context it is used
func (m IdmWorkspaceScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}