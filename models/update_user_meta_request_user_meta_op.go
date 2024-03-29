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

// UpdateUserMetaRequestUserMetaOp update user meta request user meta op
//
// swagger:model UpdateUserMetaRequestUserMetaOp
type UpdateUserMetaRequestUserMetaOp string

func NewUpdateUserMetaRequestUserMetaOp(value UpdateUserMetaRequestUserMetaOp) *UpdateUserMetaRequestUserMetaOp {
	return &value
}

// Pointer returns a pointer to a freshly-allocated UpdateUserMetaRequestUserMetaOp.
func (m UpdateUserMetaRequestUserMetaOp) Pointer() *UpdateUserMetaRequestUserMetaOp {
	return &m
}

const (

	// UpdateUserMetaRequestUserMetaOpPUT captures enum value "PUT"
	UpdateUserMetaRequestUserMetaOpPUT UpdateUserMetaRequestUserMetaOp = "PUT"

	// UpdateUserMetaRequestUserMetaOpDELETE captures enum value "DELETE"
	UpdateUserMetaRequestUserMetaOpDELETE UpdateUserMetaRequestUserMetaOp = "DELETE"
)

// for schema
var updateUserMetaRequestUserMetaOpEnum []interface{}

func init() {
	var res []UpdateUserMetaRequestUserMetaOp
	if err := json.Unmarshal([]byte(`["PUT","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateUserMetaRequestUserMetaOpEnum = append(updateUserMetaRequestUserMetaOpEnum, v)
	}
}

func (m UpdateUserMetaRequestUserMetaOp) validateUpdateUserMetaRequestUserMetaOpEnum(path, location string, value UpdateUserMetaRequestUserMetaOp) error {
	if err := validate.EnumCase(path, location, value, updateUserMetaRequestUserMetaOpEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this update user meta request user meta op
func (m UpdateUserMetaRequestUserMetaOp) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUpdateUserMetaRequestUserMetaOpEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this update user meta request user meta op based on context it is used
func (m UpdateUserMetaRequestUserMetaOp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
