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

// DeleteUserMetaTagsReader is a Reader for the DeleteUserMetaTags structure.
type DeleteUserMetaTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserMetaTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserMetaTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserMetaTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserMetaTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserMetaTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserMetaTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /user-meta/tags/{Namespace}/{Tags}] DeleteUserMetaTags", response, response.Code())
	}
}

// NewDeleteUserMetaTagsOK creates a DeleteUserMetaTagsOK with default headers values
func NewDeleteUserMetaTagsOK() *DeleteUserMetaTagsOK {
	return &DeleteUserMetaTagsOK{}
}

/*
DeleteUserMetaTagsOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteUserMetaTagsOK struct {
	Payload *models.RestDeleteUserMetaTagsResponse
}

// IsSuccess returns true when this delete user meta tags o k response has a 2xx status code
func (o *DeleteUserMetaTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user meta tags o k response has a 3xx status code
func (o *DeleteUserMetaTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user meta tags o k response has a 4xx status code
func (o *DeleteUserMetaTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user meta tags o k response has a 5xx status code
func (o *DeleteUserMetaTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user meta tags o k response a status code equal to that given
func (o *DeleteUserMetaTagsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete user meta tags o k response
func (o *DeleteUserMetaTagsOK) Code() int {
	return 200
}

func (o *DeleteUserMetaTagsOK) Error() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsOK  %+v", 200, o.Payload)
}

func (o *DeleteUserMetaTagsOK) String() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsOK  %+v", 200, o.Payload)
}

func (o *DeleteUserMetaTagsOK) GetPayload() *models.RestDeleteUserMetaTagsResponse {
	return o.Payload
}

func (o *DeleteUserMetaTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestDeleteUserMetaTagsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserMetaTagsUnauthorized creates a DeleteUserMetaTagsUnauthorized with default headers values
func NewDeleteUserMetaTagsUnauthorized() *DeleteUserMetaTagsUnauthorized {
	return &DeleteUserMetaTagsUnauthorized{}
}

/*
DeleteUserMetaTagsUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type DeleteUserMetaTagsUnauthorized struct {
}

// IsSuccess returns true when this delete user meta tags unauthorized response has a 2xx status code
func (o *DeleteUserMetaTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user meta tags unauthorized response has a 3xx status code
func (o *DeleteUserMetaTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user meta tags unauthorized response has a 4xx status code
func (o *DeleteUserMetaTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user meta tags unauthorized response has a 5xx status code
func (o *DeleteUserMetaTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user meta tags unauthorized response a status code equal to that given
func (o *DeleteUserMetaTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete user meta tags unauthorized response
func (o *DeleteUserMetaTagsUnauthorized) Code() int {
	return 401
}

func (o *DeleteUserMetaTagsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsUnauthorized ", 401)
}

func (o *DeleteUserMetaTagsUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsUnauthorized ", 401)
}

func (o *DeleteUserMetaTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserMetaTagsForbidden creates a DeleteUserMetaTagsForbidden with default headers values
func NewDeleteUserMetaTagsForbidden() *DeleteUserMetaTagsForbidden {
	return &DeleteUserMetaTagsForbidden{}
}

/*
DeleteUserMetaTagsForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type DeleteUserMetaTagsForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete user meta tags forbidden response has a 2xx status code
func (o *DeleteUserMetaTagsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user meta tags forbidden response has a 3xx status code
func (o *DeleteUserMetaTagsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user meta tags forbidden response has a 4xx status code
func (o *DeleteUserMetaTagsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user meta tags forbidden response has a 5xx status code
func (o *DeleteUserMetaTagsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user meta tags forbidden response a status code equal to that given
func (o *DeleteUserMetaTagsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete user meta tags forbidden response
func (o *DeleteUserMetaTagsForbidden) Code() int {
	return 403
}

func (o *DeleteUserMetaTagsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserMetaTagsForbidden) String() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserMetaTagsForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteUserMetaTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserMetaTagsNotFound creates a DeleteUserMetaTagsNotFound with default headers values
func NewDeleteUserMetaTagsNotFound() *DeleteUserMetaTagsNotFound {
	return &DeleteUserMetaTagsNotFound{}
}

/*
DeleteUserMetaTagsNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type DeleteUserMetaTagsNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete user meta tags not found response has a 2xx status code
func (o *DeleteUserMetaTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user meta tags not found response has a 3xx status code
func (o *DeleteUserMetaTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user meta tags not found response has a 4xx status code
func (o *DeleteUserMetaTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user meta tags not found response has a 5xx status code
func (o *DeleteUserMetaTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user meta tags not found response a status code equal to that given
func (o *DeleteUserMetaTagsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete user meta tags not found response
func (o *DeleteUserMetaTagsNotFound) Code() int {
	return 404
}

func (o *DeleteUserMetaTagsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserMetaTagsNotFound) String() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserMetaTagsNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteUserMetaTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserMetaTagsInternalServerError creates a DeleteUserMetaTagsInternalServerError with default headers values
func NewDeleteUserMetaTagsInternalServerError() *DeleteUserMetaTagsInternalServerError {
	return &DeleteUserMetaTagsInternalServerError{}
}

/*
DeleteUserMetaTagsInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type DeleteUserMetaTagsInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this delete user meta tags internal server error response has a 2xx status code
func (o *DeleteUserMetaTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user meta tags internal server error response has a 3xx status code
func (o *DeleteUserMetaTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user meta tags internal server error response has a 4xx status code
func (o *DeleteUserMetaTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user meta tags internal server error response has a 5xx status code
func (o *DeleteUserMetaTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete user meta tags internal server error response a status code equal to that given
func (o *DeleteUserMetaTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete user meta tags internal server error response
func (o *DeleteUserMetaTagsInternalServerError) Code() int {
	return 500
}

func (o *DeleteUserMetaTagsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserMetaTagsInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /user-meta/tags/{Namespace}/{Tags}][%d] deleteUserMetaTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserMetaTagsInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *DeleteUserMetaTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
