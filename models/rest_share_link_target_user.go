// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestShareLinkTargetUser rest share link target user
//
// swagger:model restShareLinkTargetUser
type RestShareLinkTargetUser struct {

	// display
	Display string `json:"Display,omitempty"`

	// download count
	DownloadCount int32 `json:"DownloadCount,omitempty"`
}

// Validate validates this rest share link target user
func (m *RestShareLinkTargetUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest share link target user based on context it is used
func (m *RestShareLinkTargetUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestShareLinkTargetUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestShareLinkTargetUser) UnmarshalBinary(b []byte) error {
	var res RestShareLinkTargetUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
