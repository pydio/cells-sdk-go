// Code generated by go-swagger; DO NOT EDIT.

package acl_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v3/models"
)

// SearchAclsReader is a Reader for the SearchAcls structure.
type SearchAclsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAclsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchAclsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchAclsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchAclsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchAclsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchAclsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchAclsOK creates a SearchAclsOK with default headers values
func NewSearchAclsOK() *SearchAclsOK {
	return &SearchAclsOK{}
}

/* SearchAclsOK describes a response with status code 200, with default header values.

SearchAclsOK search acls o k
*/
type SearchAclsOK struct {
	Payload *models.RestACLCollection
}

func (o *SearchAclsOK) Error() string {
	return fmt.Sprintf("[POST /acl][%d] searchAclsOK  %+v", 200, o.Payload)
}
func (o *SearchAclsOK) GetPayload() *models.RestACLCollection {
	return o.Payload
}

func (o *SearchAclsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestACLCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAclsUnauthorized creates a SearchAclsUnauthorized with default headers values
func NewSearchAclsUnauthorized() *SearchAclsUnauthorized {
	return &SearchAclsUnauthorized{}
}

/* SearchAclsUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type SearchAclsUnauthorized struct {
	Payload *models.RestError
}

func (o *SearchAclsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /acl][%d] searchAclsUnauthorized  %+v", 401, o.Payload)
}
func (o *SearchAclsUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchAclsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAclsForbidden creates a SearchAclsForbidden with default headers values
func NewSearchAclsForbidden() *SearchAclsForbidden {
	return &SearchAclsForbidden{}
}

/* SearchAclsForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type SearchAclsForbidden struct {
	Payload *models.RestError
}

func (o *SearchAclsForbidden) Error() string {
	return fmt.Sprintf("[POST /acl][%d] searchAclsForbidden  %+v", 403, o.Payload)
}
func (o *SearchAclsForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchAclsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAclsNotFound creates a SearchAclsNotFound with default headers values
func NewSearchAclsNotFound() *SearchAclsNotFound {
	return &SearchAclsNotFound{}
}

/* SearchAclsNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type SearchAclsNotFound struct {
	Payload *models.RestError
}

func (o *SearchAclsNotFound) Error() string {
	return fmt.Sprintf("[POST /acl][%d] searchAclsNotFound  %+v", 404, o.Payload)
}
func (o *SearchAclsNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchAclsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchAclsInternalServerError creates a SearchAclsInternalServerError with default headers values
func NewSearchAclsInternalServerError() *SearchAclsInternalServerError {
	return &SearchAclsInternalServerError{}
}

/* SearchAclsInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type SearchAclsInternalServerError struct {
	Payload *models.RestError
}

func (o *SearchAclsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /acl][%d] searchAclsInternalServerError  %+v", 500, o.Payload)
}
func (o *SearchAclsInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *SearchAclsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
