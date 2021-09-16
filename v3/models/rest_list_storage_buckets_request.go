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

// RestListStorageBucketsRequest rest list storage buckets request
//
// swagger:model restListStorageBucketsRequest
type RestListStorageBucketsRequest struct {

	// buckets regexp
	BucketsRegexp string `json:"BucketsRegexp,omitempty"`

	// data source
	DataSource *ObjectDataSource `json:"DataSource,omitempty"`
}

// Validate validates this rest list storage buckets request
func (m *RestListStorageBucketsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestListStorageBucketsRequest) validateDataSource(formats strfmt.Registry) error {
	if swag.IsZero(m.DataSource) { // not required
		return nil
	}

	if m.DataSource != nil {
		if err := m.DataSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DataSource")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this rest list storage buckets request based on the context it is used
func (m *RestListStorageBucketsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestListStorageBucketsRequest) contextValidateDataSource(ctx context.Context, formats strfmt.Registry) error {

	if m.DataSource != nil {
		if err := m.DataSource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DataSource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestListStorageBucketsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestListStorageBucketsRequest) UnmarshalBinary(b []byte) error {
	var res RestListStorageBucketsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
