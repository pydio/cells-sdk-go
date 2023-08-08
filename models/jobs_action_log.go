// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// JobsActionLog jobs action log
//
// swagger:model jobsActionLog
type JobsActionLog struct {

	// action
	Action *JobsAction `json:"Action,omitempty"`

	// input message
	InputMessage *JobsActionMessage `json:"InputMessage,omitempty"`

	// output message
	OutputMessage *JobsActionMessage `json:"OutputMessage,omitempty"`
}

// Validate validates this jobs action log
func (m *JobsActionLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInputMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsActionLog) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	if m.Action != nil {
		if err := m.Action.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Action")
			}
			return err
		}
	}

	return nil
}

func (m *JobsActionLog) validateInputMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.InputMessage) { // not required
		return nil
	}

	if m.InputMessage != nil {
		if err := m.InputMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InputMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InputMessage")
			}
			return err
		}
	}

	return nil
}

func (m *JobsActionLog) validateOutputMessage(formats strfmt.Registry) error {
	if swag.IsZero(m.OutputMessage) { // not required
		return nil
	}

	if m.OutputMessage != nil {
		if err := m.OutputMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OutputMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OutputMessage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this jobs action log based on the context it is used
func (m *JobsActionLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInputMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutputMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsActionLog) contextValidateAction(ctx context.Context, formats strfmt.Registry) error {

	if m.Action != nil {

		if swag.IsZero(m.Action) { // not required
			return nil
		}

		if err := m.Action.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Action")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Action")
			}
			return err
		}
	}

	return nil
}

func (m *JobsActionLog) contextValidateInputMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.InputMessage != nil {

		if swag.IsZero(m.InputMessage) { // not required
			return nil
		}

		if err := m.InputMessage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InputMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("InputMessage")
			}
			return err
		}
	}

	return nil
}

func (m *JobsActionLog) contextValidateOutputMessage(ctx context.Context, formats strfmt.Registry) error {

	if m.OutputMessage != nil {

		if swag.IsZero(m.OutputMessage) { // not required
			return nil
		}

		if err := m.OutputMessage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("OutputMessage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("OutputMessage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsActionLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsActionLog) UnmarshalBinary(b []byte) error {
	var res JobsActionLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
