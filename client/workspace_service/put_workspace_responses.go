// Code generated by go-swagger; DO NOT EDIT.

package workspace_service

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
	"github.com/go-openapi/validate"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// PutWorkspaceReader is a Reader for the PutWorkspace structure.
type PutWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPutWorkspaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutWorkspaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutWorkspaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutWorkspaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /workspace/{Slug}] PutWorkspace", response, response.Code())
	}
}

// NewPutWorkspaceOK creates a PutWorkspaceOK with default headers values
func NewPutWorkspaceOK() *PutWorkspaceOK {
	return &PutWorkspaceOK{}
}

/*
PutWorkspaceOK describes a response with status code 200, with default header values.

A successful response.
*/
type PutWorkspaceOK struct {
	Payload *models.IdmWorkspace
}

// IsSuccess returns true when this put workspace o k response has a 2xx status code
func (o *PutWorkspaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put workspace o k response has a 3xx status code
func (o *PutWorkspaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put workspace o k response has a 4xx status code
func (o *PutWorkspaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put workspace o k response has a 5xx status code
func (o *PutWorkspaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put workspace o k response a status code equal to that given
func (o *PutWorkspaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put workspace o k response
func (o *PutWorkspaceOK) Code() int {
	return 200
}

func (o *PutWorkspaceOK) Error() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceOK  %+v", 200, o.Payload)
}

func (o *PutWorkspaceOK) String() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceOK  %+v", 200, o.Payload)
}

func (o *PutWorkspaceOK) GetPayload() *models.IdmWorkspace {
	return o.Payload
}

func (o *PutWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmWorkspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWorkspaceUnauthorized creates a PutWorkspaceUnauthorized with default headers values
func NewPutWorkspaceUnauthorized() *PutWorkspaceUnauthorized {
	return &PutWorkspaceUnauthorized{}
}

/*
PutWorkspaceUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PutWorkspaceUnauthorized struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put workspace unauthorized response has a 2xx status code
func (o *PutWorkspaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put workspace unauthorized response has a 3xx status code
func (o *PutWorkspaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put workspace unauthorized response has a 4xx status code
func (o *PutWorkspaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put workspace unauthorized response has a 5xx status code
func (o *PutWorkspaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put workspace unauthorized response a status code equal to that given
func (o *PutWorkspaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the put workspace unauthorized response
func (o *PutWorkspaceUnauthorized) Code() int {
	return 401
}

func (o *PutWorkspaceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutWorkspaceUnauthorized) String() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutWorkspaceUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutWorkspaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWorkspaceForbidden creates a PutWorkspaceForbidden with default headers values
func NewPutWorkspaceForbidden() *PutWorkspaceForbidden {
	return &PutWorkspaceForbidden{}
}

/*
PutWorkspaceForbidden describes a response with status code 403, with default header values.

User has no permission to access this particular resource
*/
type PutWorkspaceForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put workspace forbidden response has a 2xx status code
func (o *PutWorkspaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put workspace forbidden response has a 3xx status code
func (o *PutWorkspaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put workspace forbidden response has a 4xx status code
func (o *PutWorkspaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put workspace forbidden response has a 5xx status code
func (o *PutWorkspaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put workspace forbidden response a status code equal to that given
func (o *PutWorkspaceForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put workspace forbidden response
func (o *PutWorkspaceForbidden) Code() int {
	return 403
}

func (o *PutWorkspaceForbidden) Error() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceForbidden  %+v", 403, o.Payload)
}

func (o *PutWorkspaceForbidden) String() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceForbidden  %+v", 403, o.Payload)
}

func (o *PutWorkspaceForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutWorkspaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWorkspaceNotFound creates a PutWorkspaceNotFound with default headers values
func NewPutWorkspaceNotFound() *PutWorkspaceNotFound {
	return &PutWorkspaceNotFound{}
}

/*
PutWorkspaceNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PutWorkspaceNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put workspace not found response has a 2xx status code
func (o *PutWorkspaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put workspace not found response has a 3xx status code
func (o *PutWorkspaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put workspace not found response has a 4xx status code
func (o *PutWorkspaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put workspace not found response has a 5xx status code
func (o *PutWorkspaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put workspace not found response a status code equal to that given
func (o *PutWorkspaceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the put workspace not found response
func (o *PutWorkspaceNotFound) Code() int {
	return 404
}

func (o *PutWorkspaceNotFound) Error() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceNotFound  %+v", 404, o.Payload)
}

func (o *PutWorkspaceNotFound) String() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceNotFound  %+v", 404, o.Payload)
}

func (o *PutWorkspaceNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutWorkspaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutWorkspaceInternalServerError creates a PutWorkspaceInternalServerError with default headers values
func NewPutWorkspaceInternalServerError() *PutWorkspaceInternalServerError {
	return &PutWorkspaceInternalServerError{}
}

/*
PutWorkspaceInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PutWorkspaceInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this put workspace internal server error response has a 2xx status code
func (o *PutWorkspaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put workspace internal server error response has a 3xx status code
func (o *PutWorkspaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put workspace internal server error response has a 4xx status code
func (o *PutWorkspaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put workspace internal server error response has a 5xx status code
func (o *PutWorkspaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put workspace internal server error response a status code equal to that given
func (o *PutWorkspaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put workspace internal server error response
func (o *PutWorkspaceInternalServerError) Code() int {
	return 500
}

func (o *PutWorkspaceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutWorkspaceInternalServerError) String() string {
	return fmt.Sprintf("[PUT /workspace/{Slug}][%d] putWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutWorkspaceInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PutWorkspaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PutWorkspaceBody A Workspace is composed of a set of nodes UUIDs and is used to provide accesses to the tree via ACLs.
swagger:model PutWorkspaceBody
*/
type PutWorkspaceBody struct {

	// JSON-encoded list of attributes
	Attributes string `json:"Attributes,omitempty"`

	// Description of the workspace (max length 1000)
	Description string `json:"Description,omitempty"`

	// Label of the workspace (max length 500)
	Label string `json:"Label,omitempty"`

	// Last modification time
	LastUpdated int32 `json:"LastUpdated,omitempty"`

	// Policies for securing access
	Policies []*models.ServiceResourcePolicy `json:"Policies"`

	// Context-resolved to quickly check if workspace is editable or not
	PoliciesContextEditable bool `json:"PoliciesContextEditable,omitempty"`

	// List of the Root Nodes in the tree that compose this workspace
	RootNodes map[string]models.TreeNode `json:"RootNodes,omitempty"`

	// Quick list of the RootNodes uuids
	RootUUIDs []string `json:"RootUUIDs"`

	// Scope can be ADMIN, ROOM (=CELL) or LINK
	Scope *models.IdmWorkspaceScope `json:"Scope,omitempty"`

	// Unique identifier of the workspace
	UUID string `json:"UUID,omitempty"`
}

// Validate validates this put workspace body
func (o *PutWorkspaceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRootNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutWorkspaceBody) validatePolicies(formats strfmt.Registry) error {
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

func (o *PutWorkspaceBody) validateRootNodes(formats strfmt.Registry) error {
	if swag.IsZero(o.RootNodes) { // not required
		return nil
	}

	for k := range o.RootNodes {

		if err := validate.Required("body"+"."+"RootNodes"+"."+k, "body", o.RootNodes[k]); err != nil {
			return err
		}
		if val, ok := o.RootNodes[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "RootNodes" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "RootNodes" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *PutWorkspaceBody) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(o.Scope) { // not required
		return nil
	}

	if o.Scope != nil {
		if err := o.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "Scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "Scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put workspace body based on the context it is used
func (o *PutWorkspaceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRootNodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutWorkspaceBody) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

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

func (o *PutWorkspaceBody) contextValidateRootNodes(ctx context.Context, formats strfmt.Registry) error {

	for k := range o.RootNodes {

		if val, ok := o.RootNodes[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (o *PutWorkspaceBody) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if o.Scope != nil {

		if swag.IsZero(o.Scope) { // not required
			return nil
		}

		if err := o.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "Scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "Scope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutWorkspaceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutWorkspaceBody) UnmarshalBinary(b []byte) error {
	var res PutWorkspaceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
