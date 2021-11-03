// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// ListTasksLogsReader is a Reader for the ListTasksLogs structure.
type ListTasksLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTasksLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTasksLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListTasksLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTasksLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTasksLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListTasksLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTasksLogsOK creates a ListTasksLogsOK with default headers values
func NewListTasksLogsOK() *ListTasksLogsOK {
	return &ListTasksLogsOK{}
}

/* ListTasksLogsOK describes a response with status code 200, with default header values.

ListTasksLogsOK list tasks logs o k
*/
type ListTasksLogsOK struct {
	Payload *models.RestLogMessageCollection
}

func (o *ListTasksLogsOK) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/logs][%d] listTasksLogsOK  %+v", 200, o.Payload)
}
func (o *ListTasksLogsOK) GetPayload() *models.RestLogMessageCollection {
	return o.Payload
}

func (o *ListTasksLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestLogMessageCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksLogsUnauthorized creates a ListTasksLogsUnauthorized with default headers values
func NewListTasksLogsUnauthorized() *ListTasksLogsUnauthorized {
	return &ListTasksLogsUnauthorized{}
}

/* ListTasksLogsUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListTasksLogsUnauthorized struct {
}

func (o *ListTasksLogsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/logs][%d] listTasksLogsUnauthorized ", 401)
}

func (o *ListTasksLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListTasksLogsForbidden creates a ListTasksLogsForbidden with default headers values
func NewListTasksLogsForbidden() *ListTasksLogsForbidden {
	return &ListTasksLogsForbidden{}
}

/* ListTasksLogsForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListTasksLogsForbidden struct {
	Payload *models.RestError
}

func (o *ListTasksLogsForbidden) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/logs][%d] listTasksLogsForbidden  %+v", 403, o.Payload)
}
func (o *ListTasksLogsForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListTasksLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksLogsNotFound creates a ListTasksLogsNotFound with default headers values
func NewListTasksLogsNotFound() *ListTasksLogsNotFound {
	return &ListTasksLogsNotFound{}
}

/* ListTasksLogsNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListTasksLogsNotFound struct {
	Payload *models.RestError
}

func (o *ListTasksLogsNotFound) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/logs][%d] listTasksLogsNotFound  %+v", 404, o.Payload)
}
func (o *ListTasksLogsNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListTasksLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksLogsInternalServerError creates a ListTasksLogsInternalServerError with default headers values
func NewListTasksLogsInternalServerError() *ListTasksLogsInternalServerError {
	return &ListTasksLogsInternalServerError{}
}

/* ListTasksLogsInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListTasksLogsInternalServerError struct {
	Payload *models.RestError
}

func (o *ListTasksLogsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /jobs/tasks/logs][%d] listTasksLogsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListTasksLogsInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListTasksLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
