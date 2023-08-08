// Code generated by go-swagger; DO NOT EDIT.

package log_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// SyslogReader is a Reader for the Syslog structure.
type SyslogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyslogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyslogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSyslogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSyslogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSyslogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSyslogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /log/sys] Syslog", response, response.Code())
	}
}

// NewSyslogOK creates a SyslogOK with default headers values
func NewSyslogOK() *SyslogOK {
	return &SyslogOK{}
}

/*
SyslogOK describes a response with status code 200, with default header values.

A successful response.
*/
type SyslogOK struct {
	Payload *models.RestLogMessageCollection
}

// IsSuccess returns true when this syslog o k response has a 2xx status code
func (o *SyslogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this syslog o k response has a 3xx status code
func (o *SyslogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this syslog o k response has a 4xx status code
func (o *SyslogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this syslog o k response has a 5xx status code
func (o *SyslogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this syslog o k response a status code equal to that given
func (o *SyslogOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the syslog o k response
func (o *SyslogOK) Code() int {
	return 200
}

func (o *SyslogOK) Error() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogOK  %+v", 200, o.Payload)
}

func (o *SyslogOK) String() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogOK  %+v", 200, o.Payload)
}

func (o *SyslogOK) GetPayload() *models.RestLogMessageCollection {
	return o.Payload
}

func (o *SyslogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestLogMessageCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyslogUnauthorized creates a SyslogUnauthorized with default headers values
func NewSyslogUnauthorized() *SyslogUnauthorized {
	return &SyslogUnauthorized{}
}

/*
SyslogUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type SyslogUnauthorized struct {
}

// IsSuccess returns true when this syslog unauthorized response has a 2xx status code
func (o *SyslogUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this syslog unauthorized response has a 3xx status code
func (o *SyslogUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this syslog unauthorized response has a 4xx status code
func (o *SyslogUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this syslog unauthorized response has a 5xx status code
func (o *SyslogUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this syslog unauthorized response a status code equal to that given
func (o *SyslogUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the syslog unauthorized response
func (o *SyslogUnauthorized) Code() int {
	return 401
}

func (o *SyslogUnauthorized) Error() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogUnauthorized ", 401)
}

func (o *SyslogUnauthorized) String() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogUnauthorized ", 401)
}

func (o *SyslogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyslogForbidden creates a SyslogForbidden with default headers values
func NewSyslogForbidden() *SyslogForbidden {
	return &SyslogForbidden{}
}

/*
SyslogForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type SyslogForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this syslog forbidden response has a 2xx status code
func (o *SyslogForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this syslog forbidden response has a 3xx status code
func (o *SyslogForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this syslog forbidden response has a 4xx status code
func (o *SyslogForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this syslog forbidden response has a 5xx status code
func (o *SyslogForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this syslog forbidden response a status code equal to that given
func (o *SyslogForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the syslog forbidden response
func (o *SyslogForbidden) Code() int {
	return 403
}

func (o *SyslogForbidden) Error() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogForbidden  %+v", 403, o.Payload)
}

func (o *SyslogForbidden) String() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogForbidden  %+v", 403, o.Payload)
}

func (o *SyslogForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SyslogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyslogNotFound creates a SyslogNotFound with default headers values
func NewSyslogNotFound() *SyslogNotFound {
	return &SyslogNotFound{}
}

/*
SyslogNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type SyslogNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this syslog not found response has a 2xx status code
func (o *SyslogNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this syslog not found response has a 3xx status code
func (o *SyslogNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this syslog not found response has a 4xx status code
func (o *SyslogNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this syslog not found response has a 5xx status code
func (o *SyslogNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this syslog not found response a status code equal to that given
func (o *SyslogNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the syslog not found response
func (o *SyslogNotFound) Code() int {
	return 404
}

func (o *SyslogNotFound) Error() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogNotFound  %+v", 404, o.Payload)
}

func (o *SyslogNotFound) String() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogNotFound  %+v", 404, o.Payload)
}

func (o *SyslogNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SyslogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyslogInternalServerError creates a SyslogInternalServerError with default headers values
func NewSyslogInternalServerError() *SyslogInternalServerError {
	return &SyslogInternalServerError{}
}

/*
SyslogInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type SyslogInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this syslog internal server error response has a 2xx status code
func (o *SyslogInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this syslog internal server error response has a 3xx status code
func (o *SyslogInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this syslog internal server error response has a 4xx status code
func (o *SyslogInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this syslog internal server error response has a 5xx status code
func (o *SyslogInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this syslog internal server error response a status code equal to that given
func (o *SyslogInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the syslog internal server error response
func (o *SyslogInternalServerError) Code() int {
	return 500
}

func (o *SyslogInternalServerError) Error() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogInternalServerError  %+v", 500, o.Payload)
}

func (o *SyslogInternalServerError) String() string {
	return fmt.Sprintf("[POST /log/sys][%d] syslogInternalServerError  %+v", 500, o.Payload)
}

func (o *SyslogInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SyslogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}