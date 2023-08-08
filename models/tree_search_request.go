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

// TreeSearchRequest tree search request
//
// swagger:model treeSearchRequest
type TreeSearchRequest struct {

	// Load node details
	Details bool `json:"Details,omitempty"`

	// Start at given position
	From int32 `json:"From,omitempty"`

	// The query object
	Query *TreeQuery `json:"Query,omitempty"`

	// Limit the number of results
	Size int32 `json:"Size,omitempty"`

	// Generic Details Flags
	StatFlags []int64 `json:"StatFlags"`
}

// Validate validates this tree search request
func (m *TreeSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeSearchRequest) validateQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Query")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tree search request based on the context it is used
func (m *TreeSearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeSearchRequest) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.Query != nil {

		if swag.IsZero(m.Query) { // not required
			return nil
		}

		if err := m.Query.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreeSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeSearchRequest) UnmarshalBinary(b []byte) error {
	var res TreeSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
