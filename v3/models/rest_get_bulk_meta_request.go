// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestGetBulkMetaRequest rest get bulk meta request
//
// swagger:model restGetBulkMetaRequest
type RestGetBulkMetaRequest struct {

	// Whether to query all services for the metadata they can contribute to enrich the node
	AllMetaProviders bool `json:"AllMetaProviders,omitempty"`

	// Limit number of results
	Limit int32 `json:"Limit,omitempty"`

	// List of node paths to query (use paths ending with /* to load the children)
	NodePaths []string `json:"NodePaths"`

	// Start listing at a given position
	Offset int32 `json:"Offset,omitempty"`

	// Load Versions of the given node
	Versions bool `json:"Versions,omitempty"`
}

// Validate validates this rest get bulk meta request
func (m *RestGetBulkMetaRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest get bulk meta request based on context it is used
func (m *RestGetBulkMetaRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestGetBulkMetaRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestGetBulkMetaRequest) UnmarshalBinary(b []byte) error {
	var res RestGetBulkMetaRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
