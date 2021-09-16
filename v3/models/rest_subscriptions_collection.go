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

// RestSubscriptionsCollection rest subscriptions collection
//
// swagger:model restSubscriptionsCollection
type RestSubscriptionsCollection struct {

	// subscriptions
	Subscriptions []*ActivitySubscription `json:"subscriptions"`
}

// Validate validates this rest subscriptions collection
func (m *RestSubscriptionsCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubscriptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSubscriptionsCollection) validateSubscriptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Subscriptions) { // not required
		return nil
	}

	for i := 0; i < len(m.Subscriptions); i++ {
		if swag.IsZero(m.Subscriptions[i]) { // not required
			continue
		}

		if m.Subscriptions[i] != nil {
			if err := m.Subscriptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rest subscriptions collection based on the context it is used
func (m *RestSubscriptionsCollection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubscriptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSubscriptionsCollection) contextValidateSubscriptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Subscriptions); i++ {

		if m.Subscriptions[i] != nil {
			if err := m.Subscriptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestSubscriptionsCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestSubscriptionsCollection) UnmarshalBinary(b []byte) error {
	var res RestSubscriptionsCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
