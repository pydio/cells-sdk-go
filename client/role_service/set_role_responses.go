// Code generated by go-swagger; DO NOT EDIT.

package role_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// SetRoleReader is a Reader for the SetRole structure.
type SetRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSetRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /role/{Uuid}] SetRole", response, response.Code())
	}
}

// NewSetRoleOK creates a SetRoleOK with default headers values
func NewSetRoleOK() *SetRoleOK {
	return &SetRoleOK{}
}

/*
SetRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type SetRoleOK struct {
	Payload *models.IdmRole
}

// IsSuccess returns true when this set role o k response has a 2xx status code
func (o *SetRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set role o k response has a 3xx status code
func (o *SetRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role o k response has a 4xx status code
func (o *SetRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set role o k response has a 5xx status code
func (o *SetRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set role o k response a status code equal to that given
func (o *SetRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set role o k response
func (o *SetRoleOK) Code() int {
	return 200
}

func (o *SetRoleOK) Error() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleOK  %+v", 200, o.Payload)
}

func (o *SetRoleOK) String() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleOK  %+v", 200, o.Payload)
}

func (o *SetRoleOK) GetPayload() *models.IdmRole {
	return o.Payload
}

func (o *SetRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleUnauthorized creates a SetRoleUnauthorized with default headers values
func NewSetRoleUnauthorized() *SetRoleUnauthorized {
	return &SetRoleUnauthorized{}
}

/*
SetRoleUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type SetRoleUnauthorized struct {
}

// IsSuccess returns true when this set role unauthorized response has a 2xx status code
func (o *SetRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set role unauthorized response has a 3xx status code
func (o *SetRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role unauthorized response has a 4xx status code
func (o *SetRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this set role unauthorized response has a 5xx status code
func (o *SetRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this set role unauthorized response a status code equal to that given
func (o *SetRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the set role unauthorized response
func (o *SetRoleUnauthorized) Code() int {
	return 401
}

func (o *SetRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleUnauthorized ", 401)
}

func (o *SetRoleUnauthorized) String() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleUnauthorized ", 401)
}

func (o *SetRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetRoleForbidden creates a SetRoleForbidden with default headers values
func NewSetRoleForbidden() *SetRoleForbidden {
	return &SetRoleForbidden{}
}

/*
SetRoleForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type SetRoleForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this set role forbidden response has a 2xx status code
func (o *SetRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set role forbidden response has a 3xx status code
func (o *SetRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role forbidden response has a 4xx status code
func (o *SetRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set role forbidden response has a 5xx status code
func (o *SetRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set role forbidden response a status code equal to that given
func (o *SetRoleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the set role forbidden response
func (o *SetRoleForbidden) Code() int {
	return 403
}

func (o *SetRoleForbidden) Error() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleForbidden  %+v", 403, o.Payload)
}

func (o *SetRoleForbidden) String() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleForbidden  %+v", 403, o.Payload)
}

func (o *SetRoleForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SetRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleNotFound creates a SetRoleNotFound with default headers values
func NewSetRoleNotFound() *SetRoleNotFound {
	return &SetRoleNotFound{}
}

/*
SetRoleNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type SetRoleNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this set role not found response has a 2xx status code
func (o *SetRoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set role not found response has a 3xx status code
func (o *SetRoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role not found response has a 4xx status code
func (o *SetRoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set role not found response has a 5xx status code
func (o *SetRoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set role not found response a status code equal to that given
func (o *SetRoleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set role not found response
func (o *SetRoleNotFound) Code() int {
	return 404
}

func (o *SetRoleNotFound) Error() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleNotFound  %+v", 404, o.Payload)
}

func (o *SetRoleNotFound) String() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleNotFound  %+v", 404, o.Payload)
}

func (o *SetRoleNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SetRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleInternalServerError creates a SetRoleInternalServerError with default headers values
func NewSetRoleInternalServerError() *SetRoleInternalServerError {
	return &SetRoleInternalServerError{}
}

/*
SetRoleInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type SetRoleInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this set role internal server error response has a 2xx status code
func (o *SetRoleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set role internal server error response has a 3xx status code
func (o *SetRoleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role internal server error response has a 4xx status code
func (o *SetRoleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set role internal server error response has a 5xx status code
func (o *SetRoleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set role internal server error response a status code equal to that given
func (o *SetRoleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set role internal server error response
func (o *SetRoleInternalServerError) Code() int {
	return 500
}

func (o *SetRoleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *SetRoleInternalServerError) String() string {
	return fmt.Sprintf("[PUT /role/{Uuid}][%d] setRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *SetRoleInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SetRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SetRoleBody Role represents a generic set of permissions that can be applied to any users or groups.
swagger:model SetRoleBody
*/
type SetRoleBody struct {

	// List of profiles (standard, shared, admin) on which the role will be automatically applied
	AutoApplies []string `json:"AutoApplies"`

	// Is used in a stack of roles, this one will always be applied last.
	ForceOverride bool `json:"ForceOverride,omitempty"`

	// Whether this role is attached to a Group object
	GroupRole bool `json:"GroupRole,omitempty"`

	// Whether this role represents a user team or not
	IsTeam bool `json:"IsTeam,omitempty"`

	// Label of this role
	Label string `json:"Label,omitempty"`

	// Last modification date of the role
	LastUpdated int32 `json:"LastUpdated,omitempty"`

	// List of policies for securing this role access
	Policies []*models.ServiceResourcePolicy `json:"Policies"`

	// Whether the policies resolve into an editable state
	PoliciesContextEditable bool `json:"PoliciesContextEditable,omitempty"`

	// Whether this role is attached to a User object
	UserRole bool `json:"UserRole,omitempty"`
}

// Validate validates this set role body
func (o *SetRoleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetRoleBody) validatePolicies(formats strfmt.Registry) error {
	if swag.IsZero(o.Policies) { // not required
		return nil
	}

	for i := 0; i < len(o.Policies); i++ {
		if swag.IsZero(o.Policies[i]) { // not required
			continue
		}

		if o.Policies[i] != nil {
			if err := o.Policies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "Policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "Policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this set role body based on the context it is used
func (o *SetRoleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SetRoleBody) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Policies); i++ {

		if o.Policies[i] != nil {

			if swag.IsZero(o.Policies[i]) { // not required
				return nil
			}

			if err := o.Policies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "Policies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "Policies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SetRoleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetRoleBody) UnmarshalBinary(b []byte) error {
	var res SetRoleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
