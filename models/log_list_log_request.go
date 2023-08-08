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

// LogListLogRequest ListLogRequest launches a parameterised query in the log repository and streams the results.
//
// swagger:model logListLogRequest
type LogListLogRequest struct {

	// format
	Format *ListLogRequestLogFormat `json:"Format,omitempty"`

	// Start at page
	Page int32 `json:"Page,omitempty"`

	// Bleve-type Query stsring
	Query string `json:"Query,omitempty"`

	// Number of results
	Size int32 `json:"Size,omitempty"`
}

// Validate validates this log list log request
func (m *LogListLogRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogListLogRequest) validateFormat(formats strfmt.Registry) error {
	if swag.IsZero(m.Format) { // not required
		return nil
	}

	if m.Format != nil {
		if err := m.Format.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Format")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this log list log request based on the context it is used
func (m *LogListLogRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFormat(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogListLogRequest) contextValidateFormat(ctx context.Context, formats strfmt.Registry) error {

	if m.Format != nil {

		if swag.IsZero(m.Format) { // not required
			return nil
		}

		if err := m.Format.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Format")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Format")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogListLogRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogListLogRequest) UnmarshalBinary(b []byte) error {
	var res LogListLogRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
