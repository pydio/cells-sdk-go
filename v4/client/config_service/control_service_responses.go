// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// ControlServiceReader is a Reader for the ControlService structure.
type ControlServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControlServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewControlServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewControlServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewControlServiceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewControlServiceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewControlServiceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewControlServiceOK creates a ControlServiceOK with default headers values
func NewControlServiceOK() *ControlServiceOK {
	return &ControlServiceOK{}
}

/*
ControlServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type ControlServiceOK struct {
	Payload *models.CtlService
}

// IsSuccess returns true when this control service o k response has a 2xx status code
func (o *ControlServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this control service o k response has a 3xx status code
func (o *ControlServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this control service o k response has a 4xx status code
func (o *ControlServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this control service o k response has a 5xx status code
func (o *ControlServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this control service o k response a status code equal to that given
func (o *ControlServiceOK) IsCode(code int) bool {
	return code == 200
}

func (o *ControlServiceOK) Error() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceOK  %+v", 200, o.Payload)
}

func (o *ControlServiceOK) String() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceOK  %+v", 200, o.Payload)
}

func (o *ControlServiceOK) GetPayload() *models.CtlService {
	return o.Payload
}

func (o *ControlServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CtlService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControlServiceUnauthorized creates a ControlServiceUnauthorized with default headers values
func NewControlServiceUnauthorized() *ControlServiceUnauthorized {
	return &ControlServiceUnauthorized{}
}

/*
ControlServiceUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ControlServiceUnauthorized struct {
}

// IsSuccess returns true when this control service unauthorized response has a 2xx status code
func (o *ControlServiceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this control service unauthorized response has a 3xx status code
func (o *ControlServiceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this control service unauthorized response has a 4xx status code
func (o *ControlServiceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this control service unauthorized response has a 5xx status code
func (o *ControlServiceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this control service unauthorized response a status code equal to that given
func (o *ControlServiceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ControlServiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceUnauthorized ", 401)
}

func (o *ControlServiceUnauthorized) String() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceUnauthorized ", 401)
}

func (o *ControlServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControlServiceForbidden creates a ControlServiceForbidden with default headers values
func NewControlServiceForbidden() *ControlServiceForbidden {
	return &ControlServiceForbidden{}
}

/*
ControlServiceForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ControlServiceForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this control service forbidden response has a 2xx status code
func (o *ControlServiceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this control service forbidden response has a 3xx status code
func (o *ControlServiceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this control service forbidden response has a 4xx status code
func (o *ControlServiceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this control service forbidden response has a 5xx status code
func (o *ControlServiceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this control service forbidden response a status code equal to that given
func (o *ControlServiceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ControlServiceForbidden) Error() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceForbidden  %+v", 403, o.Payload)
}

func (o *ControlServiceForbidden) String() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceForbidden  %+v", 403, o.Payload)
}

func (o *ControlServiceForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ControlServiceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControlServiceNotFound creates a ControlServiceNotFound with default headers values
func NewControlServiceNotFound() *ControlServiceNotFound {
	return &ControlServiceNotFound{}
}

/*
ControlServiceNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ControlServiceNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this control service not found response has a 2xx status code
func (o *ControlServiceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this control service not found response has a 3xx status code
func (o *ControlServiceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this control service not found response has a 4xx status code
func (o *ControlServiceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this control service not found response has a 5xx status code
func (o *ControlServiceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this control service not found response a status code equal to that given
func (o *ControlServiceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ControlServiceNotFound) Error() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceNotFound  %+v", 404, o.Payload)
}

func (o *ControlServiceNotFound) String() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceNotFound  %+v", 404, o.Payload)
}

func (o *ControlServiceNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ControlServiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewControlServiceInternalServerError creates a ControlServiceInternalServerError with default headers values
func NewControlServiceInternalServerError() *ControlServiceInternalServerError {
	return &ControlServiceInternalServerError{}
}

/*
ControlServiceInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ControlServiceInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this control service internal server error response has a 2xx status code
func (o *ControlServiceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this control service internal server error response has a 3xx status code
func (o *ControlServiceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this control service internal server error response has a 4xx status code
func (o *ControlServiceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this control service internal server error response has a 5xx status code
func (o *ControlServiceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this control service internal server error response a status code equal to that given
func (o *ControlServiceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ControlServiceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *ControlServiceInternalServerError) String() string {
	return fmt.Sprintf("[POST /config/ctl][%d] controlServiceInternalServerError  %+v", 500, o.Payload)
}

func (o *ControlServiceInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ControlServiceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
