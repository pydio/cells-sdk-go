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

// RestUserMetaCollection Collection of UserMeta
//
// swagger:model restUserMetaCollection
type RestUserMetaCollection struct {

	// metadatas
	Metadatas []*IdmUserMeta `json:"Metadatas"`
}

// Validate validates this rest user meta collection
func (m *RestUserMetaCollection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadatas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestUserMetaCollection) validateMetadatas(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadatas) { // not required
		return nil
	}

	for i := 0; i < len(m.Metadatas); i++ {
		if swag.IsZero(m.Metadatas[i]) { // not required
			continue
		}

		if m.Metadatas[i] != nil {
			if err := m.Metadatas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Metadatas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rest user meta collection based on the context it is used
func (m *RestUserMetaCollection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMetadatas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestUserMetaCollection) contextValidateMetadatas(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Metadatas); i++ {

		if m.Metadatas[i] != nil {
			if err := m.Metadatas[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Metadatas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestUserMetaCollection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestUserMetaCollection) UnmarshalBinary(b []byte) error {
	var res RestUserMetaCollection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
