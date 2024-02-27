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

// OpenAPIDiscoveryReader is a Reader for the OpenAPIDiscovery structure.
type OpenAPIDiscoveryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenAPIDiscoveryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenAPIDiscoveryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewOpenAPIDiscoveryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOpenAPIDiscoveryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOpenAPIDiscoveryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewOpenAPIDiscoveryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /config/discovery/openapi] OpenApiDiscovery", response, response.Code())
	}
}

// NewOpenAPIDiscoveryOK creates a OpenAPIDiscoveryOK with default headers values
func NewOpenAPIDiscoveryOK() *OpenAPIDiscoveryOK {
	return &OpenAPIDiscoveryOK{}
}

/*
OpenAPIDiscoveryOK describes a response with status code 200, with default header values.

A successful response.
*/
type OpenAPIDiscoveryOK struct {
	Payload *models.RestOpenAPIResponse
}

// IsSuccess returns true when this open Api discovery o k response has a 2xx status code
func (o *OpenAPIDiscoveryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this open Api discovery o k response has a 3xx status code
func (o *OpenAPIDiscoveryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this open Api discovery o k response has a 4xx status code
func (o *OpenAPIDiscoveryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this open Api discovery o k response has a 5xx status code
func (o *OpenAPIDiscoveryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this open Api discovery o k response a status code equal to that given
func (o *OpenAPIDiscoveryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the open Api discovery o k response
func (o *OpenAPIDiscoveryOK) Code() int {
	return 200
}

func (o *OpenAPIDiscoveryOK) Error() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryOK  %+v", 200, o.Payload)
}

func (o *OpenAPIDiscoveryOK) String() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryOK  %+v", 200, o.Payload)
}

func (o *OpenAPIDiscoveryOK) GetPayload() *models.RestOpenAPIResponse {
	return o.Payload
}

func (o *OpenAPIDiscoveryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestOpenAPIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenAPIDiscoveryUnauthorized creates a OpenAPIDiscoveryUnauthorized with default headers values
func NewOpenAPIDiscoveryUnauthorized() *OpenAPIDiscoveryUnauthorized {
	return &OpenAPIDiscoveryUnauthorized{}
}

/*
OpenAPIDiscoveryUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type OpenAPIDiscoveryUnauthorized struct {
}

// IsSuccess returns true when this open Api discovery unauthorized response has a 2xx status code
func (o *OpenAPIDiscoveryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this open Api discovery unauthorized response has a 3xx status code
func (o *OpenAPIDiscoveryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this open Api discovery unauthorized response has a 4xx status code
func (o *OpenAPIDiscoveryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this open Api discovery unauthorized response has a 5xx status code
func (o *OpenAPIDiscoveryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this open Api discovery unauthorized response a status code equal to that given
func (o *OpenAPIDiscoveryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the open Api discovery unauthorized response
func (o *OpenAPIDiscoveryUnauthorized) Code() int {
	return 401
}

func (o *OpenAPIDiscoveryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryUnauthorized ", 401)
}

func (o *OpenAPIDiscoveryUnauthorized) String() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryUnauthorized ", 401)
}

func (o *OpenAPIDiscoveryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOpenAPIDiscoveryForbidden creates a OpenAPIDiscoveryForbidden with default headers values
func NewOpenAPIDiscoveryForbidden() *OpenAPIDiscoveryForbidden {
	return &OpenAPIDiscoveryForbidden{}
}

/*
OpenAPIDiscoveryForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type OpenAPIDiscoveryForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this open Api discovery forbidden response has a 2xx status code
func (o *OpenAPIDiscoveryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this open Api discovery forbidden response has a 3xx status code
func (o *OpenAPIDiscoveryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this open Api discovery forbidden response has a 4xx status code
func (o *OpenAPIDiscoveryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this open Api discovery forbidden response has a 5xx status code
func (o *OpenAPIDiscoveryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this open Api discovery forbidden response a status code equal to that given
func (o *OpenAPIDiscoveryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the open Api discovery forbidden response
func (o *OpenAPIDiscoveryForbidden) Code() int {
	return 403
}

func (o *OpenAPIDiscoveryForbidden) Error() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryForbidden  %+v", 403, o.Payload)
}

func (o *OpenAPIDiscoveryForbidden) String() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryForbidden  %+v", 403, o.Payload)
}

func (o *OpenAPIDiscoveryForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *OpenAPIDiscoveryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenAPIDiscoveryNotFound creates a OpenAPIDiscoveryNotFound with default headers values
func NewOpenAPIDiscoveryNotFound() *OpenAPIDiscoveryNotFound {
	return &OpenAPIDiscoveryNotFound{}
}

/*
OpenAPIDiscoveryNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type OpenAPIDiscoveryNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this open Api discovery not found response has a 2xx status code
func (o *OpenAPIDiscoveryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this open Api discovery not found response has a 3xx status code
func (o *OpenAPIDiscoveryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this open Api discovery not found response has a 4xx status code
func (o *OpenAPIDiscoveryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this open Api discovery not found response has a 5xx status code
func (o *OpenAPIDiscoveryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this open Api discovery not found response a status code equal to that given
func (o *OpenAPIDiscoveryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the open Api discovery not found response
func (o *OpenAPIDiscoveryNotFound) Code() int {
	return 404
}

func (o *OpenAPIDiscoveryNotFound) Error() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryNotFound  %+v", 404, o.Payload)
}

func (o *OpenAPIDiscoveryNotFound) String() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryNotFound  %+v", 404, o.Payload)
}

func (o *OpenAPIDiscoveryNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *OpenAPIDiscoveryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenAPIDiscoveryInternalServerError creates a OpenAPIDiscoveryInternalServerError with default headers values
func NewOpenAPIDiscoveryInternalServerError() *OpenAPIDiscoveryInternalServerError {
	return &OpenAPIDiscoveryInternalServerError{}
}

/*
OpenAPIDiscoveryInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type OpenAPIDiscoveryInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this open Api discovery internal server error response has a 2xx status code
func (o *OpenAPIDiscoveryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this open Api discovery internal server error response has a 3xx status code
func (o *OpenAPIDiscoveryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this open Api discovery internal server error response has a 4xx status code
func (o *OpenAPIDiscoveryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this open Api discovery internal server error response has a 5xx status code
func (o *OpenAPIDiscoveryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this open Api discovery internal server error response a status code equal to that given
func (o *OpenAPIDiscoveryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the open Api discovery internal server error response
func (o *OpenAPIDiscoveryInternalServerError) Code() int {
	return 500
}

func (o *OpenAPIDiscoveryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryInternalServerError  %+v", 500, o.Payload)
}

func (o *OpenAPIDiscoveryInternalServerError) String() string {
	return fmt.Sprintf("[GET /config/discovery/openapi][%d] openApiDiscoveryInternalServerError  %+v", 500, o.Payload)
}

func (o *OpenAPIDiscoveryInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *OpenAPIDiscoveryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
