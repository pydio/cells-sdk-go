// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// ListUserMetaTagsReader is a Reader for the ListUserMetaTags structure.
type ListUserMetaTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserMetaTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserMetaTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListUserMetaTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUserMetaTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListUserMetaTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListUserMetaTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user-meta/tags/{Namespace}] ListUserMetaTags", response, response.Code())
	}
}

// NewListUserMetaTagsOK creates a ListUserMetaTagsOK with default headers values
func NewListUserMetaTagsOK() *ListUserMetaTagsOK {
	return &ListUserMetaTagsOK{}
}

/*
ListUserMetaTagsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListUserMetaTagsOK struct {
	Payload *models.RestListUserMetaTagsResponse
}

// IsSuccess returns true when this list user meta tags o k response has a 2xx status code
func (o *ListUserMetaTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list user meta tags o k response has a 3xx status code
func (o *ListUserMetaTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta tags o k response has a 4xx status code
func (o *ListUserMetaTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list user meta tags o k response has a 5xx status code
func (o *ListUserMetaTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list user meta tags o k response a status code equal to that given
func (o *ListUserMetaTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list user meta tags o k response
func (o *ListUserMetaTagsOK) Code() int {
	return 200
}

func (o *ListUserMetaTagsOK) Error() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsOK  %+v", 200, o.Payload)
}

func (o *ListUserMetaTagsOK) String() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsOK  %+v", 200, o.Payload)
}

func (o *ListUserMetaTagsOK) GetPayload() *models.RestListUserMetaTagsResponse {
	return o.Payload
}

func (o *ListUserMetaTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestListUserMetaTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserMetaTagsUnauthorized creates a ListUserMetaTagsUnauthorized with default headers values
func NewListUserMetaTagsUnauthorized() *ListUserMetaTagsUnauthorized {
	return &ListUserMetaTagsUnauthorized{}
}

/*
ListUserMetaTagsUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListUserMetaTagsUnauthorized struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list user meta tags unauthorized response has a 2xx status code
func (o *ListUserMetaTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user meta tags unauthorized response has a 3xx status code
func (o *ListUserMetaTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta tags unauthorized response has a 4xx status code
func (o *ListUserMetaTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user meta tags unauthorized response has a 5xx status code
func (o *ListUserMetaTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list user meta tags unauthorized response a status code equal to that given
func (o *ListUserMetaTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list user meta tags unauthorized response
func (o *ListUserMetaTagsUnauthorized) Code() int {
	return 401
}

func (o *ListUserMetaTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListUserMetaTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListUserMetaTagsUnauthorized) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListUserMetaTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserMetaTagsForbidden creates a ListUserMetaTagsForbidden with default headers values
func NewListUserMetaTagsForbidden() *ListUserMetaTagsForbidden {
	return &ListUserMetaTagsForbidden{}
}

/*
ListUserMetaTagsForbidden describes a response with status code 403, with default header values.

User has no permission to access this particular resource
*/
type ListUserMetaTagsForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list user meta tags forbidden response has a 2xx status code
func (o *ListUserMetaTagsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user meta tags forbidden response has a 3xx status code
func (o *ListUserMetaTagsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta tags forbidden response has a 4xx status code
func (o *ListUserMetaTagsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user meta tags forbidden response has a 5xx status code
func (o *ListUserMetaTagsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list user meta tags forbidden response a status code equal to that given
func (o *ListUserMetaTagsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list user meta tags forbidden response
func (o *ListUserMetaTagsForbidden) Code() int {
	return 403
}

func (o *ListUserMetaTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsForbidden  %+v", 403, o.Payload)
}

func (o *ListUserMetaTagsForbidden) String() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsForbidden  %+v", 403, o.Payload)
}

func (o *ListUserMetaTagsForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListUserMetaTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserMetaTagsNotFound creates a ListUserMetaTagsNotFound with default headers values
func NewListUserMetaTagsNotFound() *ListUserMetaTagsNotFound {
	return &ListUserMetaTagsNotFound{}
}

/*
ListUserMetaTagsNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListUserMetaTagsNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list user meta tags not found response has a 2xx status code
func (o *ListUserMetaTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user meta tags not found response has a 3xx status code
func (o *ListUserMetaTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta tags not found response has a 4xx status code
func (o *ListUserMetaTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list user meta tags not found response has a 5xx status code
func (o *ListUserMetaTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list user meta tags not found response a status code equal to that given
func (o *ListUserMetaTagsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list user meta tags not found response
func (o *ListUserMetaTagsNotFound) Code() int {
	return 404
}

func (o *ListUserMetaTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsNotFound  %+v", 404, o.Payload)
}

func (o *ListUserMetaTagsNotFound) String() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsNotFound  %+v", 404, o.Payload)
}

func (o *ListUserMetaTagsNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListUserMetaTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserMetaTagsInternalServerError creates a ListUserMetaTagsInternalServerError with default headers values
func NewListUserMetaTagsInternalServerError() *ListUserMetaTagsInternalServerError {
	return &ListUserMetaTagsInternalServerError{}
}

/*
ListUserMetaTagsInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListUserMetaTagsInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list user meta tags internal server error response has a 2xx status code
func (o *ListUserMetaTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list user meta tags internal server error response has a 3xx status code
func (o *ListUserMetaTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list user meta tags internal server error response has a 4xx status code
func (o *ListUserMetaTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list user meta tags internal server error response has a 5xx status code
func (o *ListUserMetaTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list user meta tags internal server error response a status code equal to that given
func (o *ListUserMetaTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list user meta tags internal server error response
func (o *ListUserMetaTagsInternalServerError) Code() int {
	return 500
}

func (o *ListUserMetaTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListUserMetaTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /user-meta/tags/{Namespace}][%d] listUserMetaTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListUserMetaTagsInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListUserMetaTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
