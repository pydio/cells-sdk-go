// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestResetPasswordResponse rest reset password response
//
// swagger:model restResetPasswordResponse
type RestResetPasswordResponse struct {

	// message
	Message string `json:"Message,omitempty"`

	// success
	Success bool `json:"Success,omitempty"`
}

// Validate validates this rest reset password response
func (m *RestResetPasswordResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest reset password response based on context it is used
func (m *RestResetPasswordResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestResetPasswordResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestResetPasswordResponse) UnmarshalBinary(b []byte) error {
	var res RestResetPasswordResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}