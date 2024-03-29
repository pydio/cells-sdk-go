// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestFrontEnrollAuthResponse rest front enroll auth response
//
// swagger:model restFrontEnrollAuthResponse
type RestFrontEnrollAuthResponse struct {

	// Any parameters can be returned
	Info map[string]string `json:"Info,omitempty"`
}

// Validate validates this rest front enroll auth response
func (m *RestFrontEnrollAuthResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest front enroll auth response based on context it is used
func (m *RestFrontEnrollAuthResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestFrontEnrollAuthResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestFrontEnrollAuthResponse) UnmarshalBinary(b []byte) error {
	var res RestFrontEnrollAuthResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
