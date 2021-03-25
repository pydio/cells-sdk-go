// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestListPeerFoldersRequest rest list peer folders request
// swagger:model restListPeerFoldersRequest
type RestListPeerFoldersRequest struct {

	// Path to the parent folder for listing
	Path string `json:"Path,omitempty"`

	// Restrict listing to a given peer
	PeerAddress string `json:"PeerAddress,omitempty"`
}

// Validate validates this rest list peer folders request
func (m *RestListPeerFoldersRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestListPeerFoldersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestListPeerFoldersRequest) UnmarshalBinary(b []byte) error {
	var res RestListPeerFoldersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
