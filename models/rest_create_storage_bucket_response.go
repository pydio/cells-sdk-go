// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestCreateStorageBucketResponse rest create storage bucket response
//
// swagger:model restCreateStorageBucketResponse
type RestCreateStorageBucketResponse struct {

	// bucket name
	BucketName string `json:"BucketName,omitempty"`

	// success
	Success bool `json:"Success,omitempty"`
}

// Validate validates this rest create storage bucket response
func (m *RestCreateStorageBucketResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest create storage bucket response based on context it is used
func (m *RestCreateStorageBucketResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestCreateStorageBucketResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestCreateStorageBucketResponse) UnmarshalBinary(b []byte) error {
	var res RestCreateStorageBucketResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
