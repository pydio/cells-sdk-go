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

// GetVersioningPolicyReader is a Reader for the GetVersioningPolicy structure.
type GetVersioningPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersioningPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersioningPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVersioningPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVersioningPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVersioningPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVersioningPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVersioningPolicyOK creates a GetVersioningPolicyOK with default headers values
func NewGetVersioningPolicyOK() *GetVersioningPolicyOK {
	return &GetVersioningPolicyOK{}
}

/*
GetVersioningPolicyOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetVersioningPolicyOK struct {
	Payload *models.TreeVersioningPolicy
}

// IsSuccess returns true when this get versioning policy o k response has a 2xx status code
func (o *GetVersioningPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get versioning policy o k response has a 3xx status code
func (o *GetVersioningPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versioning policy o k response has a 4xx status code
func (o *GetVersioningPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get versioning policy o k response has a 5xx status code
func (o *GetVersioningPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get versioning policy o k response a status code equal to that given
func (o *GetVersioningPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVersioningPolicyOK) Error() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyOK  %+v", 200, o.Payload)
}

func (o *GetVersioningPolicyOK) String() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyOK  %+v", 200, o.Payload)
}

func (o *GetVersioningPolicyOK) GetPayload() *models.TreeVersioningPolicy {
	return o.Payload
}

func (o *GetVersioningPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TreeVersioningPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersioningPolicyUnauthorized creates a GetVersioningPolicyUnauthorized with default headers values
func NewGetVersioningPolicyUnauthorized() *GetVersioningPolicyUnauthorized {
	return &GetVersioningPolicyUnauthorized{}
}

/*
GetVersioningPolicyUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type GetVersioningPolicyUnauthorized struct {
}

// IsSuccess returns true when this get versioning policy unauthorized response has a 2xx status code
func (o *GetVersioningPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versioning policy unauthorized response has a 3xx status code
func (o *GetVersioningPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versioning policy unauthorized response has a 4xx status code
func (o *GetVersioningPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versioning policy unauthorized response has a 5xx status code
func (o *GetVersioningPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get versioning policy unauthorized response a status code equal to that given
func (o *GetVersioningPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetVersioningPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyUnauthorized ", 401)
}

func (o *GetVersioningPolicyUnauthorized) String() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyUnauthorized ", 401)
}

func (o *GetVersioningPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVersioningPolicyForbidden creates a GetVersioningPolicyForbidden with default headers values
func NewGetVersioningPolicyForbidden() *GetVersioningPolicyForbidden {
	return &GetVersioningPolicyForbidden{}
}

/*
GetVersioningPolicyForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type GetVersioningPolicyForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get versioning policy forbidden response has a 2xx status code
func (o *GetVersioningPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versioning policy forbidden response has a 3xx status code
func (o *GetVersioningPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versioning policy forbidden response has a 4xx status code
func (o *GetVersioningPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versioning policy forbidden response has a 5xx status code
func (o *GetVersioningPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get versioning policy forbidden response a status code equal to that given
func (o *GetVersioningPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetVersioningPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetVersioningPolicyForbidden) String() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetVersioningPolicyForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetVersioningPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersioningPolicyNotFound creates a GetVersioningPolicyNotFound with default headers values
func NewGetVersioningPolicyNotFound() *GetVersioningPolicyNotFound {
	return &GetVersioningPolicyNotFound{}
}

/*
GetVersioningPolicyNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type GetVersioningPolicyNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get versioning policy not found response has a 2xx status code
func (o *GetVersioningPolicyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versioning policy not found response has a 3xx status code
func (o *GetVersioningPolicyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versioning policy not found response has a 4xx status code
func (o *GetVersioningPolicyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get versioning policy not found response has a 5xx status code
func (o *GetVersioningPolicyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get versioning policy not found response a status code equal to that given
func (o *GetVersioningPolicyNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetVersioningPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetVersioningPolicyNotFound) String() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetVersioningPolicyNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetVersioningPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersioningPolicyInternalServerError creates a GetVersioningPolicyInternalServerError with default headers values
func NewGetVersioningPolicyInternalServerError() *GetVersioningPolicyInternalServerError {
	return &GetVersioningPolicyInternalServerError{}
}

/*
GetVersioningPolicyInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type GetVersioningPolicyInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this get versioning policy internal server error response has a 2xx status code
func (o *GetVersioningPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get versioning policy internal server error response has a 3xx status code
func (o *GetVersioningPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get versioning policy internal server error response has a 4xx status code
func (o *GetVersioningPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get versioning policy internal server error response has a 5xx status code
func (o *GetVersioningPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get versioning policy internal server error response a status code equal to that given
func (o *GetVersioningPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetVersioningPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVersioningPolicyInternalServerError) String() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVersioningPolicyInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetVersioningPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
