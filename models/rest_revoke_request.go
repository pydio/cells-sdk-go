// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestRevokeRequest Rest request for revocation. Token is not mandatory, if not set
// request will use current JWT token
// swagger:model restRevokeRequest
type RestRevokeRequest struct {

	// Pass a specific Token ID to be revoked. If empty, request will use current JWT
	TokenID string `json:"TokenId,omitempty"`
}

// Validate validates this rest revoke request
func (m *RestRevokeRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestRevokeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestRevokeRequest) UnmarshalBinary(b []byte) error {
	var res RestRevokeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
