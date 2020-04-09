// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AuthToken auth token
// swagger:model authToken
type AuthToken struct {

	// access token
	AccessToken string `json:"AccessToken,omitempty"`

	// expires at
	ExpiresAt string `json:"ExpiresAt,omitempty"`

	// ID token
	IDToken string `json:"IDToken,omitempty"`

	// refresh token
	RefreshToken string `json:"RefreshToken,omitempty"`
}

// Validate validates this auth token
func (m *AuthToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthToken) UnmarshalBinary(b []byte) error {
	var res AuthToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
