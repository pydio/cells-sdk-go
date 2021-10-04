// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
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

/* GetVersioningPolicyOK describes a response with status code 200, with default header values.

GetVersioningPolicyOK get versioning policy o k
*/
type GetVersioningPolicyOK struct {
	Payload *models.TreeVersioningPolicy
}

func (o *GetVersioningPolicyOK) Error() string {
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

/* GetVersioningPolicyUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type GetVersioningPolicyUnauthorized struct {
	Payload *models.RestError
}

func (o *GetVersioningPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config/versioning/{Uuid}][%d] getVersioningPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *GetVersioningPolicyUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *GetVersioningPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersioningPolicyForbidden creates a GetVersioningPolicyForbidden with default headers values
func NewGetVersioningPolicyForbidden() *GetVersioningPolicyForbidden {
	return &GetVersioningPolicyForbidden{}
}

/* GetVersioningPolicyForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type GetVersioningPolicyForbidden struct {
	Payload *models.RestError
}

func (o *GetVersioningPolicyForbidden) Error() string {
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

/* GetVersioningPolicyNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type GetVersioningPolicyNotFound struct {
	Payload *models.RestError
}

func (o *GetVersioningPolicyNotFound) Error() string {
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

/* GetVersioningPolicyInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type GetVersioningPolicyInternalServerError struct {
	Payload *models.RestError
}

func (o *GetVersioningPolicyInternalServerError) Error() string {
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
