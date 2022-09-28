// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// ListUserMetaNamespaceReader is a Reader for the ListUserMetaNamespace structure.
type ListUserMetaNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserMetaNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserMetaNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUserMetaNamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUserMetaNamespaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListUserMetaNamespaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListUserMetaNamespaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListUserMetaNamespaceOK creates a ListUserMetaNamespaceOK with default headers values
func NewListUserMetaNamespaceOK() *ListUserMetaNamespaceOK {
	return &ListUserMetaNamespaceOK{}
}

/*
ListUserMetaNamespaceOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListUserMetaNamespaceOK struct {
	Payload *models.RestUserMetaNamespaceCollection
}

// IsSuccess returns true when this list user meta namespace o k response has a 2xx status code
func (o *ListUserMetaNamespaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list user meta namespace o k response has a 3xx status code
func (o *ListUserMetaNamespaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta namespace o k response has a 4xx status code
func (o *ListUserMetaNamespaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list user meta namespace o k response has a 5xx status code
func (o *ListUserMetaNamespaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list user meta namespace o k response a status code equal to that given
func (o *ListUserMetaNamespaceOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListUserMetaNamespaceOK) Error() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceOK  %+v", 200, o.Payload)
}

func (o *ListUserMetaNamespaceOK) String() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceOK  %+v", 200, o.Payload)
}

func (o *ListUserMetaNamespaceOK) GetPayload() *models.RestUserMetaNamespaceCollection {
	return o.Payload
}

func (o *ListUserMetaNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestUserMetaNamespaceCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserMetaNamespaceUnauthorized creates a ListUserMetaNamespaceUnauthorized with default headers values
func NewListUserMetaNamespaceUnauthorized() *ListUserMetaNamespaceUnauthorized {
	return &ListUserMetaNamespaceUnauthorized{}
}

/*
ListUserMetaNamespaceUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListUserMetaNamespaceUnauthorized struct {
}

// IsSuccess returns true when this list user meta namespace unauthorized response has a 2xx status code
func (o *ListUserMetaNamespaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user meta namespace unauthorized response has a 3xx status code
func (o *ListUserMetaNamespaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta namespace unauthorized response has a 4xx status code
func (o *ListUserMetaNamespaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user meta namespace unauthorized response has a 5xx status code
func (o *ListUserMetaNamespaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list user meta namespace unauthorized response a status code equal to that given
func (o *ListUserMetaNamespaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListUserMetaNamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceUnauthorized ", 401)
}

func (o *ListUserMetaNamespaceUnauthorized) String() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceUnauthorized ", 401)
}

func (o *ListUserMetaNamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUserMetaNamespaceForbidden creates a ListUserMetaNamespaceForbidden with default headers values
func NewListUserMetaNamespaceForbidden() *ListUserMetaNamespaceForbidden {
	return &ListUserMetaNamespaceForbidden{}
}

/*
ListUserMetaNamespaceForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListUserMetaNamespaceForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list user meta namespace forbidden response has a 2xx status code
func (o *ListUserMetaNamespaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user meta namespace forbidden response has a 3xx status code
func (o *ListUserMetaNamespaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta namespace forbidden response has a 4xx status code
func (o *ListUserMetaNamespaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user meta namespace forbidden response has a 5xx status code
func (o *ListUserMetaNamespaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list user meta namespace forbidden response a status code equal to that given
func (o *ListUserMetaNamespaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListUserMetaNamespaceForbidden) Error() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceForbidden  %+v", 403, o.Payload)
}

func (o *ListUserMetaNamespaceForbidden) String() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceForbidden  %+v", 403, o.Payload)
}

func (o *ListUserMetaNamespaceForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListUserMetaNamespaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserMetaNamespaceNotFound creates a ListUserMetaNamespaceNotFound with default headers values
func NewListUserMetaNamespaceNotFound() *ListUserMetaNamespaceNotFound {
	return &ListUserMetaNamespaceNotFound{}
}

/*
ListUserMetaNamespaceNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListUserMetaNamespaceNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list user meta namespace not found response has a 2xx status code
func (o *ListUserMetaNamespaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user meta namespace not found response has a 3xx status code
func (o *ListUserMetaNamespaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta namespace not found response has a 4xx status code
func (o *ListUserMetaNamespaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user meta namespace not found response has a 5xx status code
func (o *ListUserMetaNamespaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list user meta namespace not found response a status code equal to that given
func (o *ListUserMetaNamespaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListUserMetaNamespaceNotFound) Error() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceNotFound  %+v", 404, o.Payload)
}

func (o *ListUserMetaNamespaceNotFound) String() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceNotFound  %+v", 404, o.Payload)
}

func (o *ListUserMetaNamespaceNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListUserMetaNamespaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserMetaNamespaceInternalServerError creates a ListUserMetaNamespaceInternalServerError with default headers values
func NewListUserMetaNamespaceInternalServerError() *ListUserMetaNamespaceInternalServerError {
	return &ListUserMetaNamespaceInternalServerError{}
}

/*
ListUserMetaNamespaceInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListUserMetaNamespaceInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list user meta namespace internal server error response has a 2xx status code
func (o *ListUserMetaNamespaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user meta namespace internal server error response has a 3xx status code
func (o *ListUserMetaNamespaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta namespace internal server error response has a 4xx status code
func (o *ListUserMetaNamespaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list user meta namespace internal server error response has a 5xx status code
func (o *ListUserMetaNamespaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list user meta namespace internal server error response a status code equal to that given
func (o *ListUserMetaNamespaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListUserMetaNamespaceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceInternalServerError  %+v", 500, o.Payload)
}

func (o *ListUserMetaNamespaceInternalServerError) String() string {
	return fmt.Sprintf("[GET /user-meta/namespace][%d] listUserMetaNamespaceInternalServerError  %+v", 500, o.Payload)
}

func (o *ListUserMetaNamespaceInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListUserMetaNamespaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}