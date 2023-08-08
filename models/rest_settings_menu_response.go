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

// RestSettingsMenuResponse rest settings menu response
//
// swagger:model restSettingsMenuResponse
type RestSettingsMenuResponse struct {

	// sections
	Sections []*RestSettingsSection `json:"Sections"`

	// metadata
	Metadata *RestSettingsEntryMeta `json:"__metadata__,omitempty"`
}

// Validate validates this rest settings menu response
func (m *RestSettingsMenuResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSettingsMenuResponse) validateSections(formats strfmt.Registry) error {
	if swag.IsZero(m.Sections) { // not required
		return nil
	}

	for i := 0; i < len(m.Sections); i++ {
		if swag.IsZero(m.Sections[i]) { // not required
			continue
		}

		if m.Sections[i] != nil {
			if err := m.Sections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Sections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Sections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestSettingsMenuResponse) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("__metadata__")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("__metadata__")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rest settings menu response based on the context it is used
func (m *RestSettingsMenuResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSettingsMenuResponse) contextValidateSections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sections); i++ {

		if m.Sections[i] != nil {

			if swag.IsZero(m.Sections[i]) { // not required
				return nil
			}

			if err := m.Sections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Sections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Sections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestSettingsMenuResponse) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {

		if swag.IsZero(m.Metadata) { // not required
			return nil
		}

		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("__metadata__")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("__metadata__")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestSettingsMenuResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestSettingsMenuResponse) UnmarshalBinary(b []byte) error {
	var res RestSettingsMenuResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
