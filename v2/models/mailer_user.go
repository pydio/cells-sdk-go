// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MailerUser mailer user
// swagger:model mailerUser
type MailerUser struct {

	// address
	Address string `json:"Address,omitempty"`

	// language
	Language string `json:"Language,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// Uuid
	UUID string `json:"Uuid,omitempty"`
}

// Validate validates this mailer user
func (m *MailerUser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MailerUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MailerUser) UnmarshalBinary(b []byte) error {
	var res MailerUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
