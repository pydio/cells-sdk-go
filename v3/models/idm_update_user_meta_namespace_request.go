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

// IdmUpdateUserMetaNamespaceRequest Modify UserMetaNamespaces
//
// swagger:model idmUpdateUserMetaNamespaceRequest
type IdmUpdateUserMetaNamespaceRequest struct {

	// namespaces
	Namespaces []*IdmUserMetaNamespace `json:"Namespaces"`

	// operation
	Operation *UpdateUserMetaNamespaceRequestUserMetaNsOp `json:"Operation,omitempty"`
}

// Validate validates this idm update user meta namespace request
func (m *IdmUpdateUserMetaNamespaceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmUpdateUserMetaNamespaceRequest) validateNamespaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Namespaces); i++ {
		if swag.IsZero(m.Namespaces[i]) { // not required
			continue
		}

		if m.Namespaces[i] != nil {
			if err := m.Namespaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Namespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IdmUpdateUserMetaNamespaceRequest) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Operation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this idm update user meta namespace request based on the context it is used
func (m *IdmUpdateUserMetaNamespaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNamespaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmUpdateUserMetaNamespaceRequest) contextValidateNamespaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Namespaces); i++ {

		if m.Namespaces[i] != nil {
			if err := m.Namespaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Namespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IdmUpdateUserMetaNamespaceRequest) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {
		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Operation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdmUpdateUserMetaNamespaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmUpdateUserMetaNamespaceRequest) UnmarshalBinary(b []byte) error {
	var res IdmUpdateUserMetaNamespaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
