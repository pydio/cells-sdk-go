// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestDocumentAccessTokenRequest rest document access token request
//
// swagger:model restDocumentAccessTokenRequest
type RestDocumentAccessTokenRequest struct {

	// client ID
	ClientID string `json:"ClientID,omitempty"`

	// path
	Path string `json:"Path,omitempty"`
}

// Validate validates this rest document access token request
func (m *RestDocumentAccessTokenRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest document access token request based on context it is used
func (m *RestDocumentAccessTokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestDocumentAccessTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestDocumentAccessTokenRequest) UnmarshalBinary(b []byte) error {
	var res RestDocumentAccessTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
