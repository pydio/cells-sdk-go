// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestUserJobResponse rest user job response
//
// swagger:model restUserJobResponse
type RestUserJobResponse struct {

	// job Uuid
	JobUUID string `json:"JobUuid,omitempty"`
}

// Validate validates this rest user job response
func (m *RestUserJobResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this rest user job response based on context it is used
func (m *RestUserJobResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestUserJobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestUserJobResponse) UnmarshalBinary(b []byte) error {
	var res RestUserJobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
