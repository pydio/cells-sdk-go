// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateApplyUpdateResponse update apply update response
//
// swagger:model updateApplyUpdateResponse
type UpdateApplyUpdateResponse struct {

	// message
	Message string `json:"Message,omitempty"`

	// success
	Success bool `json:"Success,omitempty"`
}

// Validate validates this update apply update response
func (m *UpdateApplyUpdateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update apply update response based on context it is used
func (m *UpdateApplyUpdateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateApplyUpdateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateApplyUpdateResponse) UnmarshalBinary(b []byte) error {
	var res UpdateApplyUpdateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}