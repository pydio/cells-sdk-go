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

// RestSearchACLRequest Rest request for ACL's
//
// swagger:model restSearchACLRequest
type RestSearchACLRequest struct {

	// Return counts only, no actual results
	CountOnly bool `json:"CountOnly,omitempty"`

	// Group results
	GroupBy int32 `json:"GroupBy,omitempty"`

	// Limit the number of results
	Limit string `json:"Limit,omitempty"`

	// Start listing at a given position
	Offset string `json:"Offset,omitempty"`

	// Single queries will be combined using this operation AND or OR logic
	Operation *ServiceOperationType `json:"Operation,omitempty"`

	// Atomic queries that will be combined using the OperationType (AND or OR)
	Queries []*IdmACLSingleQuery `json:"Queries"`
}

// Validate validates this rest search ACL request
func (m *RestSearchACLRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSearchACLRequest) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	if m.Operation != nil {
		if err := m.Operation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Operation")
			}
			return err
		}
	}

	return nil
}

func (m *RestSearchACLRequest) validateQueries(formats strfmt.Registry) error {
	if swag.IsZero(m.Queries) { // not required
		return nil
	}

	for i := 0; i < len(m.Queries); i++ {
		if swag.IsZero(m.Queries[i]) { // not required
			continue
		}

		if m.Queries[i] != nil {
			if err := m.Queries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Queries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Queries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rest search ACL request based on the context it is used
func (m *RestSearchACLRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOperation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestSearchACLRequest) contextValidateOperation(ctx context.Context, formats strfmt.Registry) error {

	if m.Operation != nil {

		if swag.IsZero(m.Operation) { // not required
			return nil
		}

		if err := m.Operation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Operation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Operation")
			}
			return err
		}
	}

	return nil
}

func (m *RestSearchACLRequest) contextValidateQueries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Queries); i++ {

		if m.Queries[i] != nil {

			if swag.IsZero(m.Queries[i]) { // not required
				return nil
			}

			if err := m.Queries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Queries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Queries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestSearchACLRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestSearchACLRequest) UnmarshalBinary(b []byte) error {
	var res RestSearchACLRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
