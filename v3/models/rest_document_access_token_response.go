// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestDocumentAccessTokenResponse rest document access token response
//
// swagger:model restDocumentAccessTokenResponse
type RestDocumentAccessTokenResponse struct {

	// access token
	AccessToken string `json:"AccessToken,omitempty"`
}

// Validate validates this rest document access token response
func (m *RestDocumentAccessTokenResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest document access token response based on context it is used
func (m *RestDocumentAccessTokenResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestDocumentAccessTokenResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestDocumentAccessTokenResponse) UnmarshalBinary(b []byte) error {
	var res RestDocumentAccessTokenResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
