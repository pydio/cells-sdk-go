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

// IdmSearchUserMetaRequest Request for searching UserMeta by NodeUuid or by Namespace
//
// swagger:model idmSearchUserMetaRequest
type IdmSearchUserMetaRequest struct {

	// Look for meta by their unique identifier
	MetaUuids []string `json:"MetaUuids"`

	// Filter meta by their namespace
	Namespace string `json:"Namespace,omitempty"`

	// Look for all meta for a list of nodes
	NodeUuids []string `json:"NodeUuids"`

	// Filter meta by policies query
	ResourceQuery *ServiceResourcePolicyQuery `json:"ResourceQuery,omitempty"`

	// Filter meta by owner (in the sense of the policies)
	ResourceSubjectOwner string `json:"ResourceSubjectOwner,omitempty"`
}

// Validate validates this idm search user meta request
func (m *IdmSearchUserMetaRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmSearchUserMetaRequest) validateResourceQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceQuery) { // not required
		return nil
	}

	if m.ResourceQuery != nil {
		if err := m.ResourceQuery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceQuery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceQuery")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this idm search user meta request based on the context it is used
func (m *IdmSearchUserMetaRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdmSearchUserMetaRequest) contextValidateResourceQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceQuery != nil {

		if swag.IsZero(m.ResourceQuery) { // not required
			return nil
		}

		if err := m.ResourceQuery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResourceQuery")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ResourceQuery")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdmSearchUserMetaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdmSearchUserMetaRequest) UnmarshalBinary(b []byte) error {
	var res IdmSearchUserMetaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
