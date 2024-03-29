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

// RestDeleteNodesResponse rest delete nodes response
//
// swagger:model restDeleteNodesResponse
type RestDeleteNodesResponse struct {

	// delete jobs
	DeleteJobs []*RestBackgroundJobResult `json:"DeleteJobs"`
}

// Validate validates this rest delete nodes response
func (m *RestDeleteNodesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeleteJobs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestDeleteNodesResponse) validateDeleteJobs(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteJobs) { // not required
		return nil
	}

	for i := 0; i < len(m.DeleteJobs); i++ {
		if swag.IsZero(m.DeleteJobs[i]) { // not required
			continue
		}

		if m.DeleteJobs[i] != nil {
			if err := m.DeleteJobs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeleteJobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeleteJobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this rest delete nodes response based on the context it is used
func (m *RestDeleteNodesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeleteJobs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestDeleteNodesResponse) contextValidateDeleteJobs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeleteJobs); i++ {

		if m.DeleteJobs[i] != nil {

			if swag.IsZero(m.DeleteJobs[i]) { // not required
				return nil
			}

			if err := m.DeleteJobs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeleteJobs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeleteJobs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestDeleteNodesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestDeleteNodesResponse) UnmarshalBinary(b []byte) error {
	var res RestDeleteNodesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
