// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobsCtrlCommandResponse Response to the CtrlCommand
//
// swagger:model jobsCtrlCommandResponse
type JobsCtrlCommandResponse struct {

	// msg
	Msg string `json:"Msg,omitempty"`
}

// Validate validates this jobs ctrl command response
func (m *JobsCtrlCommandResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this jobs ctrl command response based on context it is used
func (m *JobsCtrlCommandResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *JobsCtrlCommandResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsCtrlCommandResponse) UnmarshalBinary(b []byte) error {
	var res JobsCtrlCommandResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
