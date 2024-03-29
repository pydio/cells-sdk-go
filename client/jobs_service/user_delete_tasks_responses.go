// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// UserDeleteTasksReader is a Reader for the UserDeleteTasks structure.
type UserDeleteTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserDeleteTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserDeleteTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUserDeleteTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserDeleteTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserDeleteTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserDeleteTasksInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /jobs/tasks/delete] UserDeleteTasks", response, response.Code())
	}
}

// NewUserDeleteTasksOK creates a UserDeleteTasksOK with default headers values
func NewUserDeleteTasksOK() *UserDeleteTasksOK {
	return &UserDeleteTasksOK{}
}

/*
UserDeleteTasksOK describes a response with status code 200, with default header values.

A successful response.
*/
type UserDeleteTasksOK struct {
	Payload *models.JobsDeleteTasksResponse
}

// IsSuccess returns true when this user delete tasks o k response has a 2xx status code
func (o *UserDeleteTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user delete tasks o k response has a 3xx status code
func (o *UserDeleteTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete tasks o k response has a 4xx status code
func (o *UserDeleteTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user delete tasks o k response has a 5xx status code
func (o *UserDeleteTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user delete tasks o k response a status code equal to that given
func (o *UserDeleteTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user delete tasks o k response
func (o *UserDeleteTasksOK) Code() int {
	return 200
}

func (o *UserDeleteTasksOK) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksOK  %+v", 200, o.Payload)
}

func (o *UserDeleteTasksOK) String() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksOK  %+v", 200, o.Payload)
}

func (o *UserDeleteTasksOK) GetPayload() *models.JobsDeleteTasksResponse {
	return o.Payload
}

func (o *UserDeleteTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobsDeleteTasksResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserDeleteTasksUnauthorized creates a UserDeleteTasksUnauthorized with default headers values
func NewUserDeleteTasksUnauthorized() *UserDeleteTasksUnauthorized {
	return &UserDeleteTasksUnauthorized{}
}

/*
UserDeleteTasksUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type UserDeleteTasksUnauthorized struct {
}

// IsSuccess returns true when this user delete tasks unauthorized response has a 2xx status code
func (o *UserDeleteTasksUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user delete tasks unauthorized response has a 3xx status code
func (o *UserDeleteTasksUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete tasks unauthorized response has a 4xx status code
func (o *UserDeleteTasksUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user delete tasks unauthorized response has a 5xx status code
func (o *UserDeleteTasksUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user delete tasks unauthorized response a status code equal to that given
func (o *UserDeleteTasksUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user delete tasks unauthorized response
func (o *UserDeleteTasksUnauthorized) Code() int {
	return 401
}

func (o *UserDeleteTasksUnauthorized) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksUnauthorized ", 401)
}

func (o *UserDeleteTasksUnauthorized) String() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksUnauthorized ", 401)
}

func (o *UserDeleteTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserDeleteTasksForbidden creates a UserDeleteTasksForbidden with default headers values
func NewUserDeleteTasksForbidden() *UserDeleteTasksForbidden {
	return &UserDeleteTasksForbidden{}
}

/*
UserDeleteTasksForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type UserDeleteTasksForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this user delete tasks forbidden response has a 2xx status code
func (o *UserDeleteTasksForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user delete tasks forbidden response has a 3xx status code
func (o *UserDeleteTasksForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete tasks forbidden response has a 4xx status code
func (o *UserDeleteTasksForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user delete tasks forbidden response has a 5xx status code
func (o *UserDeleteTasksForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user delete tasks forbidden response a status code equal to that given
func (o *UserDeleteTasksForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user delete tasks forbidden response
func (o *UserDeleteTasksForbidden) Code() int {
	return 403
}

func (o *UserDeleteTasksForbidden) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksForbidden  %+v", 403, o.Payload)
}

func (o *UserDeleteTasksForbidden) String() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksForbidden  %+v", 403, o.Payload)
}

func (o *UserDeleteTasksForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UserDeleteTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserDeleteTasksNotFound creates a UserDeleteTasksNotFound with default headers values
func NewUserDeleteTasksNotFound() *UserDeleteTasksNotFound {
	return &UserDeleteTasksNotFound{}
}

/*
UserDeleteTasksNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type UserDeleteTasksNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this user delete tasks not found response has a 2xx status code
func (o *UserDeleteTasksNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user delete tasks not found response has a 3xx status code
func (o *UserDeleteTasksNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete tasks not found response has a 4xx status code
func (o *UserDeleteTasksNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user delete tasks not found response has a 5xx status code
func (o *UserDeleteTasksNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user delete tasks not found response a status code equal to that given
func (o *UserDeleteTasksNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user delete tasks not found response
func (o *UserDeleteTasksNotFound) Code() int {
	return 404
}

func (o *UserDeleteTasksNotFound) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksNotFound  %+v", 404, o.Payload)
}

func (o *UserDeleteTasksNotFound) String() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksNotFound  %+v", 404, o.Payload)
}

func (o *UserDeleteTasksNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UserDeleteTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserDeleteTasksInternalServerError creates a UserDeleteTasksInternalServerError with default headers values
func NewUserDeleteTasksInternalServerError() *UserDeleteTasksInternalServerError {
	return &UserDeleteTasksInternalServerError{}
}

/*
UserDeleteTasksInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type UserDeleteTasksInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this user delete tasks internal server error response has a 2xx status code
func (o *UserDeleteTasksInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user delete tasks internal server error response has a 3xx status code
func (o *UserDeleteTasksInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user delete tasks internal server error response has a 4xx status code
func (o *UserDeleteTasksInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user delete tasks internal server error response has a 5xx status code
func (o *UserDeleteTasksInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user delete tasks internal server error response a status code equal to that given
func (o *UserDeleteTasksInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user delete tasks internal server error response
func (o *UserDeleteTasksInternalServerError) Code() int {
	return 500
}

func (o *UserDeleteTasksInternalServerError) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *UserDeleteTasksInternalServerError) String() string {
	return fmt.Sprintf("[POST /jobs/tasks/delete][%d] userDeleteTasksInternalServerError  %+v", 500, o.Payload)
}

func (o *UserDeleteTasksInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UserDeleteTasksInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
