// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestListUserMetaTagsResponse rest list user meta tags response
// swagger:model restListUserMetaTagsResponse
type RestListUserMetaTagsResponse struct {

	// List of existing tags values
	Tags []string `json:"Tags"`
}

// Validate validates this rest list user meta tags response
func (m *RestListUserMetaTagsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestListUserMetaTagsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestListUserMetaTagsResponse) UnmarshalBinary(b []byte) error {
	var res RestListUserMetaTagsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
