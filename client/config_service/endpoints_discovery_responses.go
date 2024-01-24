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

// EndpointsDiscoveryReader is a Reader for the EndpointsDiscovery structure.
type EndpointsDiscoveryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EndpointsDiscoveryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEndpointsDiscoveryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEndpointsDiscoveryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEndpointsDiscoveryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEndpointsDiscoveryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEndpointsDiscoveryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /config/discovery] EndpointsDiscovery", response, response.Code())
	}
}

// NewEndpointsDiscoveryOK creates a EndpointsDiscoveryOK with default headers values
func NewEndpointsDiscoveryOK() *EndpointsDiscoveryOK {
	return &EndpointsDiscoveryOK{}
}

/*
EndpointsDiscoveryOK describes a response with status code 200, with default header values.

A successful response.
*/
type EndpointsDiscoveryOK struct {
	Payload *models.RestDiscoveryResponse
}

// IsSuccess returns true when this endpoints discovery o k response has a 2xx status code
func (o *EndpointsDiscoveryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this endpoints discovery o k response has a 3xx status code
func (o *EndpointsDiscoveryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoints discovery o k response has a 4xx status code
func (o *EndpointsDiscoveryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoints discovery o k response has a 5xx status code
func (o *EndpointsDiscoveryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoints discovery o k response a status code equal to that given
func (o *EndpointsDiscoveryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the endpoints discovery o k response
func (o *EndpointsDiscoveryOK) Code() int {
	return 200
}

func (o *EndpointsDiscoveryOK) Error() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryOK  %+v", 200, o.Payload)
}

func (o *EndpointsDiscoveryOK) String() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryOK  %+v", 200, o.Payload)
}

func (o *EndpointsDiscoveryOK) GetPayload() *models.RestDiscoveryResponse {
	return o.Payload
}

func (o *EndpointsDiscoveryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDiscoveryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointsDiscoveryUnauthorized creates a EndpointsDiscoveryUnauthorized with default headers values
func NewEndpointsDiscoveryUnauthorized() *EndpointsDiscoveryUnauthorized {
	return &EndpointsDiscoveryUnauthorized{}
}

/*
EndpointsDiscoveryUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type EndpointsDiscoveryUnauthorized struct {
	Payload *models.RestError
}

// IsSuccess returns true when this endpoints discovery unauthorized response has a 2xx status code
func (o *EndpointsDiscoveryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoints discovery unauthorized response has a 3xx status code
func (o *EndpointsDiscoveryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoints discovery unauthorized response has a 4xx status code
func (o *EndpointsDiscoveryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoints discovery unauthorized response has a 5xx status code
func (o *EndpointsDiscoveryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoints discovery unauthorized response a status code equal to that given
func (o *EndpointsDiscoveryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the endpoints discovery unauthorized response
func (o *EndpointsDiscoveryUnauthorized) Code() int {
	return 401
}

func (o *EndpointsDiscoveryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryUnauthorized  %+v", 401, o.Payload)
}

func (o *EndpointsDiscoveryUnauthorized) String() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryUnauthorized  %+v", 401, o.Payload)
}

func (o *EndpointsDiscoveryUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *EndpointsDiscoveryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointsDiscoveryForbidden creates a EndpointsDiscoveryForbidden with default headers values
func NewEndpointsDiscoveryForbidden() *EndpointsDiscoveryForbidden {
	return &EndpointsDiscoveryForbidden{}
}

/*
EndpointsDiscoveryForbidden describes a response with status code 403, with default header values.

User has no permission to access this particular resource
*/
type EndpointsDiscoveryForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this endpoints discovery forbidden response has a 2xx status code
func (o *EndpointsDiscoveryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoints discovery forbidden response has a 3xx status code
func (o *EndpointsDiscoveryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoints discovery forbidden response has a 4xx status code
func (o *EndpointsDiscoveryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoints discovery forbidden response has a 5xx status code
func (o *EndpointsDiscoveryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoints discovery forbidden response a status code equal to that given
func (o *EndpointsDiscoveryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the endpoints discovery forbidden response
func (o *EndpointsDiscoveryForbidden) Code() int {
	return 403
}

func (o *EndpointsDiscoveryForbidden) Error() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryForbidden  %+v", 403, o.Payload)
}

func (o *EndpointsDiscoveryForbidden) String() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryForbidden  %+v", 403, o.Payload)
}

func (o *EndpointsDiscoveryForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *EndpointsDiscoveryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointsDiscoveryNotFound creates a EndpointsDiscoveryNotFound with default headers values
func NewEndpointsDiscoveryNotFound() *EndpointsDiscoveryNotFound {
	return &EndpointsDiscoveryNotFound{}
}

/*
EndpointsDiscoveryNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type EndpointsDiscoveryNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this endpoints discovery not found response has a 2xx status code
func (o *EndpointsDiscoveryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoints discovery not found response has a 3xx status code
func (o *EndpointsDiscoveryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoints discovery not found response has a 4xx status code
func (o *EndpointsDiscoveryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this endpoints discovery not found response has a 5xx status code
func (o *EndpointsDiscoveryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this endpoints discovery not found response a status code equal to that given
func (o *EndpointsDiscoveryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the endpoints discovery not found response
func (o *EndpointsDiscoveryNotFound) Code() int {
	return 404
}

func (o *EndpointsDiscoveryNotFound) Error() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryNotFound  %+v", 404, o.Payload)
}

func (o *EndpointsDiscoveryNotFound) String() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryNotFound  %+v", 404, o.Payload)
}

func (o *EndpointsDiscoveryNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *EndpointsDiscoveryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEndpointsDiscoveryInternalServerError creates a EndpointsDiscoveryInternalServerError with default headers values
func NewEndpointsDiscoveryInternalServerError() *EndpointsDiscoveryInternalServerError {
	return &EndpointsDiscoveryInternalServerError{}
}

/*
EndpointsDiscoveryInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type EndpointsDiscoveryInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this endpoints discovery internal server error response has a 2xx status code
func (o *EndpointsDiscoveryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this endpoints discovery internal server error response has a 3xx status code
func (o *EndpointsDiscoveryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this endpoints discovery internal server error response has a 4xx status code
func (o *EndpointsDiscoveryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this endpoints discovery internal server error response has a 5xx status code
func (o *EndpointsDiscoveryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this endpoints discovery internal server error response a status code equal to that given
func (o *EndpointsDiscoveryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the endpoints discovery internal server error response
func (o *EndpointsDiscoveryInternalServerError) Code() int {
	return 500
}

func (o *EndpointsDiscoveryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryInternalServerError  %+v", 500, o.Payload)
}

func (o *EndpointsDiscoveryInternalServerError) String() string {
	return fmt.Sprintf("[GET /config/discovery][%d] endpointsDiscoveryInternalServerError  %+v", 500, o.Payload)
}

func (o *EndpointsDiscoveryInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *EndpointsDiscoveryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
