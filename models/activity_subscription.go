// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ActivitySubscription activity subscription
// swagger:model activitySubscription
type ActivitySubscription struct {

	// events
	Events []string `json:"Events"`

	// object Id
	ObjectID string `json:"ObjectId,omitempty"`

	// object type
	ObjectType ActivityOwnerType `json:"ObjectType,omitempty"`

	// user Id
	UserID string `json:"UserId,omitempty"`
}

// Validate validates this activity subscription
func (m *ActivitySubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivitySubscription) validateObjectType(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectType) { // not required
		return nil
	}

	if err := m.ObjectType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ObjectType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivitySubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivitySubscription) UnmarshalBinary(b []byte) error {
	var res ActivitySubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}