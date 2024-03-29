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

// GetConfigReader is a Reader for the GetConfig structure.
type GetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /config/{FullPath}] GetConfig", response, response.Code())
	}
}

// NewGetConfigOK creates a GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {
	return &GetConfigOK{}
}

/*
GetConfigOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetConfigOK struct {
	Payload *models.RestConfiguration
}

// IsSuccess returns true when this get config o k response has a 2xx status code
func (o *GetConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get config o k response has a 3xx status code
func (o *GetConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get config o k response has a 4xx status code
func (o *GetConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get config o k response has a 5xx status code
func (o *GetConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get config o k response a status code equal to that given
func (o *GetConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get config o k response
func (o *GetConfigOK) Code() int {
	return 200
}

func (o *GetConfigOK) Error() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigOK  %+v", 200, o.Payload)
}

func (o *GetConfigOK) String() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigOK  %+v", 200, o.Payload)
}

func (o *GetConfigOK) GetPayload() *models.RestConfiguration {
	return o.Payload
}

func (o *GetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigUnauthorized creates a GetConfigUnauthorized with default headers values
func NewGetConfigUnauthorized() *GetConfigUnauthorized {
	return &GetConfigUnauthorized{}
}

/*
GetConfigUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type GetConfigUnauthorized struct {
}

// IsSuccess returns true when this get config unauthorized response has a 2xx status code
func (o *GetConfigUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get config unauthorized response has a 3xx status code
func (o *GetConfigUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get config unauthorized response has a 4xx status code
func (o *GetConfigUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get config unauthorized response has a 5xx status code
func (o *GetConfigUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get config unauthorized response a status code equal to that given
func (o *GetConfigUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get config unauthorized response
func (o *GetConfigUnauthorized) Code() int {
	return 401
}

func (o *GetConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigUnauthorized ", 401)
}

func (o *GetConfigUnauthorized) String() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigUnauthorized ", 401)
}

func (o *GetConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigForbidden creates a GetConfigForbidden with default headers values
func NewGetConfigForbidden() *GetConfigForbidden {
	return &GetConfigForbidden{}
}

/*
GetConfigForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type GetConfigForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get config forbidden response has a 2xx status code
func (o *GetConfigForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get config forbidden response has a 3xx status code
func (o *GetConfigForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get config forbidden response has a 4xx status code
func (o *GetConfigForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get config forbidden response has a 5xx status code
func (o *GetConfigForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get config forbidden response a status code equal to that given
func (o *GetConfigForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get config forbidden response
func (o *GetConfigForbidden) Code() int {
	return 403
}

func (o *GetConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigForbidden) String() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigNotFound creates a GetConfigNotFound with default headers values
func NewGetConfigNotFound() *GetConfigNotFound {
	return &GetConfigNotFound{}
}

/*
GetConfigNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type GetConfigNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get config not found response has a 2xx status code
func (o *GetConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get config not found response has a 3xx status code
func (o *GetConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get config not found response has a 4xx status code
func (o *GetConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get config not found response has a 5xx status code
func (o *GetConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get config not found response a status code equal to that given
func (o *GetConfigNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get config not found response
func (o *GetConfigNotFound) Code() int {
	return 404
}

func (o *GetConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigNotFound) String() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigInternalServerError creates a GetConfigInternalServerError with default headers values
func NewGetConfigInternalServerError() *GetConfigInternalServerError {
	return &GetConfigInternalServerError{}
}

/*
GetConfigInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type GetConfigInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get config internal server error response has a 2xx status code
func (o *GetConfigInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get config internal server error response has a 3xx status code
func (o *GetConfigInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get config internal server error response has a 4xx status code
func (o *GetConfigInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get config internal server error response has a 5xx status code
func (o *GetConfigInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get config internal server error response a status code equal to that given
func (o *GetConfigInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get config internal server error response
func (o *GetConfigInternalServerError) Code() int {
	return 500
}

func (o *GetConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigInternalServerError) String() string {
	return fmt.Sprintf("[GET /config/{FullPath}][%d] getConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
