// Code generated by go-swagger; DO NOT EDIT.

package graph_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// UserStateReader is a Reader for the UserState structure.
type UserStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUserStateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUserStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUserStateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /graph/state/{Segment}] UserState", response, response.Code())
	}
}

// NewUserStateOK creates a UserStateOK with default headers values
func NewUserStateOK() *UserStateOK {
	return &UserStateOK{}
}

/*
UserStateOK describes a response with status code 200, with default header values.

A successful response.
*/
type UserStateOK struct {
	Payload *models.RestUserStateResponse
}

// IsSuccess returns true when this user state o k response has a 2xx status code
func (o *UserStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user state o k response has a 3xx status code
func (o *UserStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user state o k response has a 4xx status code
func (o *UserStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user state o k response has a 5xx status code
func (o *UserStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user state o k response a status code equal to that given
func (o *UserStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user state o k response
func (o *UserStateOK) Code() int {
	return 200
}

func (o *UserStateOK) Error() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateOK  %+v", 200, o.Payload)
}

func (o *UserStateOK) String() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateOK  %+v", 200, o.Payload)
}

func (o *UserStateOK) GetPayload() *models.RestUserStateResponse {
	return o.Payload
}

func (o *UserStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestUserStateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserStateUnauthorized creates a UserStateUnauthorized with default headers values
func NewUserStateUnauthorized() *UserStateUnauthorized {
	return &UserStateUnauthorized{}
}

/*
UserStateUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type UserStateUnauthorized struct {
}

// IsSuccess returns true when this user state unauthorized response has a 2xx status code
func (o *UserStateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user state unauthorized response has a 3xx status code
func (o *UserStateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user state unauthorized response has a 4xx status code
func (o *UserStateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this user state unauthorized response has a 5xx status code
func (o *UserStateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this user state unauthorized response a status code equal to that given
func (o *UserStateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the user state unauthorized response
func (o *UserStateUnauthorized) Code() int {
	return 401
}

func (o *UserStateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateUnauthorized ", 401)
}

func (o *UserStateUnauthorized) String() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateUnauthorized ", 401)
}

func (o *UserStateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUserStateForbidden creates a UserStateForbidden with default headers values
func NewUserStateForbidden() *UserStateForbidden {
	return &UserStateForbidden{}
}

/*
UserStateForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type UserStateForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this user state forbidden response has a 2xx status code
func (o *UserStateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user state forbidden response has a 3xx status code
func (o *UserStateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user state forbidden response has a 4xx status code
func (o *UserStateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this user state forbidden response has a 5xx status code
func (o *UserStateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this user state forbidden response a status code equal to that given
func (o *UserStateForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the user state forbidden response
func (o *UserStateForbidden) Code() int {
	return 403
}

func (o *UserStateForbidden) Error() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateForbidden  %+v", 403, o.Payload)
}

func (o *UserStateForbidden) String() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateForbidden  %+v", 403, o.Payload)
}

func (o *UserStateForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UserStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserStateNotFound creates a UserStateNotFound with default headers values
func NewUserStateNotFound() *UserStateNotFound {
	return &UserStateNotFound{}
}

/*
UserStateNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type UserStateNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this user state not found response has a 2xx status code
func (o *UserStateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user state not found response has a 3xx status code
func (o *UserStateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user state not found response has a 4xx status code
func (o *UserStateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this user state not found response has a 5xx status code
func (o *UserStateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this user state not found response a status code equal to that given
func (o *UserStateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the user state not found response
func (o *UserStateNotFound) Code() int {
	return 404
}

func (o *UserStateNotFound) Error() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateNotFound  %+v", 404, o.Payload)
}

func (o *UserStateNotFound) String() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateNotFound  %+v", 404, o.Payload)
}

func (o *UserStateNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UserStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserStateInternalServerError creates a UserStateInternalServerError with default headers values
func NewUserStateInternalServerError() *UserStateInternalServerError {
	return &UserStateInternalServerError{}
}

/*
UserStateInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type UserStateInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this user state internal server error response has a 2xx status code
func (o *UserStateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this user state internal server error response has a 3xx status code
func (o *UserStateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user state internal server error response has a 4xx status code
func (o *UserStateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this user state internal server error response has a 5xx status code
func (o *UserStateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this user state internal server error response a status code equal to that given
func (o *UserStateInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the user state internal server error response
func (o *UserStateInternalServerError) Code() int {
	return 500
}

func (o *UserStateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateInternalServerError  %+v", 500, o.Payload)
}

func (o *UserStateInternalServerError) String() string {
	return fmt.Sprintf("[GET /graph/state/{Segment}][%d] userStateInternalServerError  %+v", 500, o.Payload)
}

func (o *UserStateInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *UserStateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
