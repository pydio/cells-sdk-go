// Code generated by go-swagger; DO NOT EDIT.

package install_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// PerformInstallCheckReader is a Reader for the PerformInstallCheck structure.
type PerformInstallCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformInstallCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformInstallCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPerformInstallCheckUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformInstallCheckForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformInstallCheckNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformInstallCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /install/check] PerformInstallCheck", response, response.Code())
	}
}

// NewPerformInstallCheckOK creates a PerformInstallCheckOK with default headers values
func NewPerformInstallCheckOK() *PerformInstallCheckOK {
	return &PerformInstallCheckOK{}
}

/*
PerformInstallCheckOK describes a response with status code 200, with default header values.

A successful response.
*/
type PerformInstallCheckOK struct {
	Payload *models.InstallPerformCheckResponse
}

// IsSuccess returns true when this perform install check o k response has a 2xx status code
func (o *PerformInstallCheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this perform install check o k response has a 3xx status code
func (o *PerformInstallCheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform install check o k response has a 4xx status code
func (o *PerformInstallCheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform install check o k response has a 5xx status code
func (o *PerformInstallCheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this perform install check o k response a status code equal to that given
func (o *PerformInstallCheckOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the perform install check o k response
func (o *PerformInstallCheckOK) Code() int {
	return 200
}

func (o *PerformInstallCheckOK) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckOK  %+v", 200, o.Payload)
}

func (o *PerformInstallCheckOK) String() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckOK  %+v", 200, o.Payload)
}

func (o *PerformInstallCheckOK) GetPayload() *models.InstallPerformCheckResponse {
	return o.Payload
}

func (o *PerformInstallCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstallPerformCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformInstallCheckUnauthorized creates a PerformInstallCheckUnauthorized with default headers values
func NewPerformInstallCheckUnauthorized() *PerformInstallCheckUnauthorized {
	return &PerformInstallCheckUnauthorized{}
}

/*
PerformInstallCheckUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type PerformInstallCheckUnauthorized struct {
}

// IsSuccess returns true when this perform install check unauthorized response has a 2xx status code
func (o *PerformInstallCheckUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform install check unauthorized response has a 3xx status code
func (o *PerformInstallCheckUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform install check unauthorized response has a 4xx status code
func (o *PerformInstallCheckUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform install check unauthorized response has a 5xx status code
func (o *PerformInstallCheckUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this perform install check unauthorized response a status code equal to that given
func (o *PerformInstallCheckUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the perform install check unauthorized response
func (o *PerformInstallCheckUnauthorized) Code() int {
	return 401
}

func (o *PerformInstallCheckUnauthorized) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckUnauthorized ", 401)
}

func (o *PerformInstallCheckUnauthorized) String() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckUnauthorized ", 401)
}

func (o *PerformInstallCheckUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPerformInstallCheckForbidden creates a PerformInstallCheckForbidden with default headers values
func NewPerformInstallCheckForbidden() *PerformInstallCheckForbidden {
	return &PerformInstallCheckForbidden{}
}

/*
PerformInstallCheckForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type PerformInstallCheckForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this perform install check forbidden response has a 2xx status code
func (o *PerformInstallCheckForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform install check forbidden response has a 3xx status code
func (o *PerformInstallCheckForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform install check forbidden response has a 4xx status code
func (o *PerformInstallCheckForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform install check forbidden response has a 5xx status code
func (o *PerformInstallCheckForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this perform install check forbidden response a status code equal to that given
func (o *PerformInstallCheckForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the perform install check forbidden response
func (o *PerformInstallCheckForbidden) Code() int {
	return 403
}

func (o *PerformInstallCheckForbidden) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckForbidden  %+v", 403, o.Payload)
}

func (o *PerformInstallCheckForbidden) String() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckForbidden  %+v", 403, o.Payload)
}

func (o *PerformInstallCheckForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PerformInstallCheckForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformInstallCheckNotFound creates a PerformInstallCheckNotFound with default headers values
func NewPerformInstallCheckNotFound() *PerformInstallCheckNotFound {
	return &PerformInstallCheckNotFound{}
}

/*
PerformInstallCheckNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type PerformInstallCheckNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this perform install check not found response has a 2xx status code
func (o *PerformInstallCheckNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform install check not found response has a 3xx status code
func (o *PerformInstallCheckNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform install check not found response has a 4xx status code
func (o *PerformInstallCheckNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this perform install check not found response has a 5xx status code
func (o *PerformInstallCheckNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this perform install check not found response a status code equal to that given
func (o *PerformInstallCheckNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the perform install check not found response
func (o *PerformInstallCheckNotFound) Code() int {
	return 404
}

func (o *PerformInstallCheckNotFound) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckNotFound  %+v", 404, o.Payload)
}

func (o *PerformInstallCheckNotFound) String() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckNotFound  %+v", 404, o.Payload)
}

func (o *PerformInstallCheckNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PerformInstallCheckNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformInstallCheckInternalServerError creates a PerformInstallCheckInternalServerError with default headers values
func NewPerformInstallCheckInternalServerError() *PerformInstallCheckInternalServerError {
	return &PerformInstallCheckInternalServerError{}
}

/*
PerformInstallCheckInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type PerformInstallCheckInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this perform install check internal server error response has a 2xx status code
func (o *PerformInstallCheckInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this perform install check internal server error response has a 3xx status code
func (o *PerformInstallCheckInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this perform install check internal server error response has a 4xx status code
func (o *PerformInstallCheckInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this perform install check internal server error response has a 5xx status code
func (o *PerformInstallCheckInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this perform install check internal server error response a status code equal to that given
func (o *PerformInstallCheckInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the perform install check internal server error response
func (o *PerformInstallCheckInternalServerError) Code() int {
	return 500
}

func (o *PerformInstallCheckInternalServerError) Error() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformInstallCheckInternalServerError) String() string {
	return fmt.Sprintf("[POST /install/check][%d] performInstallCheckInternalServerError  %+v", 500, o.Payload)
}

func (o *PerformInstallCheckInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *PerformInstallCheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
