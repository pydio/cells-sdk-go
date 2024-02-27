// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v5/models"
)

// ListSharedResourcesReader is a Reader for the ListSharedResources structure.
type ListSharedResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSharedResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSharedResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSharedResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListSharedResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListSharedResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListSharedResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /share/resources] ListSharedResources", response, response.Code())
	}
}

// NewListSharedResourcesOK creates a ListSharedResourcesOK with default headers values
func NewListSharedResourcesOK() *ListSharedResourcesOK {
	return &ListSharedResourcesOK{}
}

/*
ListSharedResourcesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListSharedResourcesOK struct {
	Payload *models.RestListSharedResourcesResponse
}

// IsSuccess returns true when this list shared resources o k response has a 2xx status code
func (o *ListSharedResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list shared resources o k response has a 3xx status code
func (o *ListSharedResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shared resources o k response has a 4xx status code
func (o *ListSharedResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list shared resources o k response has a 5xx status code
func (o *ListSharedResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list shared resources o k response a status code equal to that given
func (o *ListSharedResourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list shared resources o k response
func (o *ListSharedResourcesOK) Code() int {
	return 200
}

func (o *ListSharedResourcesOK) Error() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesOK  %+v", 200, o.Payload)
}

func (o *ListSharedResourcesOK) String() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesOK  %+v", 200, o.Payload)
}

func (o *ListSharedResourcesOK) GetPayload() *models.RestListSharedResourcesResponse {
	return o.Payload
}

func (o *ListSharedResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestListSharedResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSharedResourcesUnauthorized creates a ListSharedResourcesUnauthorized with default headers values
func NewListSharedResourcesUnauthorized() *ListSharedResourcesUnauthorized {
	return &ListSharedResourcesUnauthorized{}
}

/*
ListSharedResourcesUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type ListSharedResourcesUnauthorized struct {
}

// IsSuccess returns true when this list shared resources unauthorized response has a 2xx status code
func (o *ListSharedResourcesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shared resources unauthorized response has a 3xx status code
func (o *ListSharedResourcesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shared resources unauthorized response has a 4xx status code
func (o *ListSharedResourcesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shared resources unauthorized response has a 5xx status code
func (o *ListSharedResourcesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list shared resources unauthorized response a status code equal to that given
func (o *ListSharedResourcesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list shared resources unauthorized response
func (o *ListSharedResourcesUnauthorized) Code() int {
	return 401
}

func (o *ListSharedResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesUnauthorized ", 401)
}

func (o *ListSharedResourcesUnauthorized) String() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesUnauthorized ", 401)
}

func (o *ListSharedResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListSharedResourcesForbidden creates a ListSharedResourcesForbidden with default headers values
func NewListSharedResourcesForbidden() *ListSharedResourcesForbidden {
	return &ListSharedResourcesForbidden{}
}

/*
ListSharedResourcesForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type ListSharedResourcesForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list shared resources forbidden response has a 2xx status code
func (o *ListSharedResourcesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shared resources forbidden response has a 3xx status code
func (o *ListSharedResourcesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shared resources forbidden response has a 4xx status code
func (o *ListSharedResourcesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shared resources forbidden response has a 5xx status code
func (o *ListSharedResourcesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list shared resources forbidden response a status code equal to that given
func (o *ListSharedResourcesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list shared resources forbidden response
func (o *ListSharedResourcesForbidden) Code() int {
	return 403
}

func (o *ListSharedResourcesForbidden) Error() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesForbidden  %+v", 403, o.Payload)
}

func (o *ListSharedResourcesForbidden) String() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesForbidden  %+v", 403, o.Payload)
}

func (o *ListSharedResourcesForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListSharedResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSharedResourcesNotFound creates a ListSharedResourcesNotFound with default headers values
func NewListSharedResourcesNotFound() *ListSharedResourcesNotFound {
	return &ListSharedResourcesNotFound{}
}

/*
ListSharedResourcesNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type ListSharedResourcesNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list shared resources not found response has a 2xx status code
func (o *ListSharedResourcesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shared resources not found response has a 3xx status code
func (o *ListSharedResourcesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shared resources not found response has a 4xx status code
func (o *ListSharedResourcesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list shared resources not found response has a 5xx status code
func (o *ListSharedResourcesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list shared resources not found response a status code equal to that given
func (o *ListSharedResourcesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list shared resources not found response
func (o *ListSharedResourcesNotFound) Code() int {
	return 404
}

func (o *ListSharedResourcesNotFound) Error() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesNotFound  %+v", 404, o.Payload)
}

func (o *ListSharedResourcesNotFound) String() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesNotFound  %+v", 404, o.Payload)
}

func (o *ListSharedResourcesNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListSharedResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSharedResourcesInternalServerError creates a ListSharedResourcesInternalServerError with default headers values
func NewListSharedResourcesInternalServerError() *ListSharedResourcesInternalServerError {
	return &ListSharedResourcesInternalServerError{}
}

/*
ListSharedResourcesInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type ListSharedResourcesInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this list shared resources internal server error response has a 2xx status code
func (o *ListSharedResourcesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list shared resources internal server error response has a 3xx status code
func (o *ListSharedResourcesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list shared resources internal server error response has a 4xx status code
func (o *ListSharedResourcesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list shared resources internal server error response has a 5xx status code
func (o *ListSharedResourcesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list shared resources internal server error response a status code equal to that given
func (o *ListSharedResourcesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list shared resources internal server error response
func (o *ListSharedResourcesInternalServerError) Code() int {
	return 500
}

func (o *ListSharedResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSharedResourcesInternalServerError) String() string {
	return fmt.Sprintf("[POST /share/resources][%d] listSharedResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListSharedResourcesInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *ListSharedResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
