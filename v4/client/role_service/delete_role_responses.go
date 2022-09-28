// Code generated by go-swagger; DO NOT EDIT.

package role_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoleOK creates a DeleteRoleOK with default headers values
func NewDeleteRoleOK() *DeleteRoleOK {
	return &DeleteRoleOK{}
}

/*
DeleteRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteRoleOK struct {
	Payload *models.IdmRole
}

// IsSuccess returns true when this delete role o k response has a 2xx status code
func (o *DeleteRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete role o k response has a 3xx status code
func (o *DeleteRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role o k response has a 4xx status code
func (o *DeleteRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete role o k response has a 5xx status code
func (o *DeleteRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role o k response a status code equal to that given
func (o *DeleteRoleOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteRoleOK) Error() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleOK  %+v", 200, o.Payload)
}

func (o *DeleteRoleOK) String() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleOK  %+v", 200, o.Payload)
}

func (o *DeleteRoleOK) GetPayload() *models.IdmRole {
	return o.Payload
}

func (o *DeleteRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdmRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleUnauthorized creates a DeleteRoleUnauthorized with default headers values
func NewDeleteRoleUnauthorized() *DeleteRoleUnauthorized {
	return &DeleteRoleUnauthorized{}
}

/*
DeleteRoleUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteRoleUnauthorized struct {
}

// IsSuccess returns true when this delete role unauthorized response has a 2xx status code
func (o *DeleteRoleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete role unauthorized response has a 3xx status code
func (o *DeleteRoleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role unauthorized response has a 4xx status code
func (o *DeleteRoleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete role unauthorized response has a 5xx status code
func (o *DeleteRoleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role unauthorized response a status code equal to that given
func (o *DeleteRoleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleUnauthorized ", 401)
}

func (o *DeleteRoleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleUnauthorized ", 401)
}

func (o *DeleteRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoleForbidden creates a DeleteRoleForbidden with default headers values
func NewDeleteRoleForbidden() *DeleteRoleForbidden {
	return &DeleteRoleForbidden{}
}

/*
DeleteRoleForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteRoleForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete role forbidden response has a 2xx status code
func (o *DeleteRoleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete role forbidden response has a 3xx status code
func (o *DeleteRoleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role forbidden response has a 4xx status code
func (o *DeleteRoleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete role forbidden response has a 5xx status code
func (o *DeleteRoleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role forbidden response a status code equal to that given
func (o *DeleteRoleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoleForbidden) String() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoleForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleNotFound creates a DeleteRoleNotFound with default headers values
func NewDeleteRoleNotFound() *DeleteRoleNotFound {
	return &DeleteRoleNotFound{}
}

/*
DeleteRoleNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteRoleNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete role not found response has a 2xx status code
func (o *DeleteRoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete role not found response has a 3xx status code
func (o *DeleteRoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role not found response has a 4xx status code
func (o *DeleteRoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete role not found response has a 5xx status code
func (o *DeleteRoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role not found response a status code equal to that given
func (o *DeleteRoleNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRoleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoleNotFound) String() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoleNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleInternalServerError creates a DeleteRoleInternalServerError with default headers values
func NewDeleteRoleInternalServerError() *DeleteRoleInternalServerError {
	return &DeleteRoleInternalServerError{}
}

/*
DeleteRoleInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteRoleInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete role internal server error response has a 2xx status code
func (o *DeleteRoleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete role internal server error response has a 3xx status code
func (o *DeleteRoleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role internal server error response has a 4xx status code
func (o *DeleteRoleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete role internal server error response has a 5xx status code
func (o *DeleteRoleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete role internal server error response a status code equal to that given
func (o *DeleteRoleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoleInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /role/{Uuid}][%d] deleteRoleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoleInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteRoleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}