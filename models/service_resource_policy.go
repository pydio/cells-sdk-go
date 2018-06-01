// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceResourcePolicy service resource policy
// swagger:model serviceResourcePolicy
type ServiceResourcePolicy struct {

	// action
	Action ServiceResourcePolicyAction `json:"Action,omitempty"`

	// effect
	Effect ServiceResourcePolicyPolicyEffect `json:"Effect,omitempty"`

	// Json conditions
	JSONConditions string `json:"JsonConditions,omitempty"`

	// resource
	Resource string `json:"Resource,omitempty"`

	// subject
	Subject string `json:"Subject,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`
}

// Validate validates this service resource policy
func (m *ServiceResourcePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEffect(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceResourcePolicy) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if err := m.Action.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Action")
		}
		return err
	}

	return nil
}

func (m *ServiceResourcePolicy) validateEffect(formats strfmt.Registry) error {

	if swag.IsZero(m.Effect) { // not required
		return nil
	}

	if err := m.Effect.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Effect")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceResourcePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceResourcePolicy) UnmarshalBinary(b []byte) error {
	var res ServiceResourcePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}