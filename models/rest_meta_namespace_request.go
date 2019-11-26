// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestMetaNamespaceRequest rest meta namespace request
// swagger:model restMetaNamespaceRequest
type RestMetaNamespaceRequest struct {

	// List of namespaces to load
	Namespace []string `json:"Namespace"`

	// Path to the requested node
	NodePath string `json:"NodePath,omitempty"`
}

// Validate validates this rest meta namespace request
func (m *RestMetaNamespaceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestMetaNamespaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestMetaNamespaceRequest) UnmarshalBinary(b []byte) error {
	var res RestMetaNamespaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
