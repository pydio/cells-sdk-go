// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// ListProcessesReader is a Reader for the ListProcesses structure.
type ListProcessesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProcessesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProcessesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListProcessesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProcessesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProcessesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProcessesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /config/processes] ListProcesses", response, response.Code())
	}
}

// NewListProcessesOK creates a ListProcessesOK with default headers values
func NewListProcessesOK() *ListProcessesOK {
	return &ListProcessesOK{}
}

/*
ListProcessesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListProcessesOK struct {
	Payload *models.RestListProcessesResponse
}

// IsSuccess returns true when this list processes o k response has a 2xx status code
func (o *ListProcessesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list processes o k response has a 3xx status code
func (o *ListProcessesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list processes o k response has a 4xx status code
func (o *ListProcessesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list processes o k response has a 5xx status code
func (o *ListProcessesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list processes o k response a status code equal to that given
func (o *ListProcessesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list processes o k response
func (o *ListProcessesOK) Code() int {
	return 200
}

func (o *ListProcessesOK) Error() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesOK  %+v", 200, o.Payload)
}

func (o *ListProcessesOK) String() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesOK  %+v", 200, o.Payload)
}

func (o *ListProcessesOK) GetPayload() *models.RestListProcessesResponse {
	return o.Payload
}

func (o *ListProcessesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestListProcessesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProcessesUnauthorized creates a ListProcessesUnauthorized with default headers values
func NewListProcessesUnauthorized() *ListProcessesUnauthorized {
	return &ListProcessesUnauthorized{}
}

/*
ListProcessesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListProcessesUnauthorized struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list processes unauthorized response has a 2xx status code
func (o *ListProcessesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list processes unauthorized response has a 3xx status code
func (o *ListProcessesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list processes unauthorized response has a 4xx status code
func (o *ListProcessesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list processes unauthorized response has a 5xx status code
func (o *ListProcessesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list processes unauthorized response a status code equal to that given
func (o *ListProcessesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list processes unauthorized response
func (o *ListProcessesUnauthorized) Code() int {
	return 401
}

func (o *ListProcessesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProcessesUnauthorized) String() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProcessesUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListProcessesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProcessesForbidden creates a ListProcessesForbidden with default headers values
func NewListProcessesForbidden() *ListProcessesForbidden {
	return &ListProcessesForbidden{}
}

/*
ListProcessesForbidden describes a response with status code 403, with default header values.

User has no permission to access this particular resource
*/
type ListProcessesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list processes forbidden response has a 2xx status code
func (o *ListProcessesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list processes forbidden response has a 3xx status code
func (o *ListProcessesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list processes forbidden response has a 4xx status code
func (o *ListProcessesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list processes forbidden response has a 5xx status code
func (o *ListProcessesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list processes forbidden response a status code equal to that given
func (o *ListProcessesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list processes forbidden response
func (o *ListProcessesForbidden) Code() int {
	return 403
}

func (o *ListProcessesForbidden) Error() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesForbidden  %+v", 403, o.Payload)
}

func (o *ListProcessesForbidden) String() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesForbidden  %+v", 403, o.Payload)
}

func (o *ListProcessesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListProcessesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProcessesNotFound creates a ListProcessesNotFound with default headers values
func NewListProcessesNotFound() *ListProcessesNotFound {
	return &ListProcessesNotFound{}
}

/*
ListProcessesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListProcessesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list processes not found response has a 2xx status code
func (o *ListProcessesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list processes not found response has a 3xx status code
func (o *ListProcessesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list processes not found response has a 4xx status code
func (o *ListProcessesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list processes not found response has a 5xx status code
func (o *ListProcessesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list processes not found response a status code equal to that given
func (o *ListProcessesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list processes not found response
func (o *ListProcessesNotFound) Code() int {
	return 404
}

func (o *ListProcessesNotFound) Error() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesNotFound  %+v", 404, o.Payload)
}

func (o *ListProcessesNotFound) String() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesNotFound  %+v", 404, o.Payload)
}

func (o *ListProcessesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListProcessesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProcessesInternalServerError creates a ListProcessesInternalServerError with default headers values
func NewListProcessesInternalServerError() *ListProcessesInternalServerError {
	return &ListProcessesInternalServerError{}
}

/*
ListProcessesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListProcessesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list processes internal server error response has a 2xx status code
func (o *ListProcessesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list processes internal server error response has a 3xx status code
func (o *ListProcessesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list processes internal server error response has a 4xx status code
func (o *ListProcessesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list processes internal server error response has a 5xx status code
func (o *ListProcessesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list processes internal server error response a status code equal to that given
func (o *ListProcessesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list processes internal server error response
func (o *ListProcessesInternalServerError) Code() int {
	return 500
}

func (o *ListProcessesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProcessesInternalServerError) String() string {
	return fmt.Sprintf("[POST /config/processes][%d] listProcessesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProcessesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListProcessesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
