// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/pydio/cells-sdk-go/v4/models"
)

// FrontSessionReader is a Reader for the FrontSession structure.
type FrontSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrontSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFrontSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFrontSessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFrontSessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFrontSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFrontSessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /frontend/session] FrontSession", response, response.Code())
	}
}

// NewFrontSessionOK creates a FrontSessionOK with default headers values
func NewFrontSessionOK() *FrontSessionOK {
	return &FrontSessionOK{}
}

/*
FrontSessionOK describes a response with status code 200, with default header values.

A successful response.
*/
type FrontSessionOK struct {
	Payload *models.RestFrontSessionResponse
}

// IsSuccess returns true when this front session o k response has a 2xx status code
func (o *FrontSessionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this front session o k response has a 3xx status code
func (o *FrontSessionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front session o k response has a 4xx status code
func (o *FrontSessionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this front session o k response has a 5xx status code
func (o *FrontSessionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this front session o k response a status code equal to that given
func (o *FrontSessionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the front session o k response
func (o *FrontSessionOK) Code() int {
	return 200
}

func (o *FrontSessionOK) Error() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionOK  %+v", 200, o.Payload)
}

func (o *FrontSessionOK) String() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionOK  %+v", 200, o.Payload)
}

func (o *FrontSessionOK) GetPayload() *models.RestFrontSessionResponse {
	return o.Payload
}

func (o *FrontSessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestFrontSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontSessionUnauthorized creates a FrontSessionUnauthorized with default headers values
func NewFrontSessionUnauthorized() *FrontSessionUnauthorized {
	return &FrontSessionUnauthorized{}
}

/*
FrontSessionUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type FrontSessionUnauthorized struct {
}

// IsSuccess returns true when this front session unauthorized response has a 2xx status code
func (o *FrontSessionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front session unauthorized response has a 3xx status code
func (o *FrontSessionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front session unauthorized response has a 4xx status code
func (o *FrontSessionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this front session unauthorized response has a 5xx status code
func (o *FrontSessionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this front session unauthorized response a status code equal to that given
func (o *FrontSessionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the front session unauthorized response
func (o *FrontSessionUnauthorized) Code() int {
	return 401
}

func (o *FrontSessionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionUnauthorized ", 401)
}

func (o *FrontSessionUnauthorized) String() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionUnauthorized ", 401)
}

func (o *FrontSessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFrontSessionForbidden creates a FrontSessionForbidden with default headers values
func NewFrontSessionForbidden() *FrontSessionForbidden {
	return &FrontSessionForbidden{}
}

/*
FrontSessionForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type FrontSessionForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front session forbidden response has a 2xx status code
func (o *FrontSessionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front session forbidden response has a 3xx status code
func (o *FrontSessionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front session forbidden response has a 4xx status code
func (o *FrontSessionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this front session forbidden response has a 5xx status code
func (o *FrontSessionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this front session forbidden response a status code equal to that given
func (o *FrontSessionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the front session forbidden response
func (o *FrontSessionForbidden) Code() int {
	return 403
}

func (o *FrontSessionForbidden) Error() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionForbidden  %+v", 403, o.Payload)
}

func (o *FrontSessionForbidden) String() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionForbidden  %+v", 403, o.Payload)
}

func (o *FrontSessionForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontSessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontSessionNotFound creates a FrontSessionNotFound with default headers values
func NewFrontSessionNotFound() *FrontSessionNotFound {
	return &FrontSessionNotFound{}
}

/*
FrontSessionNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type FrontSessionNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front session not found response has a 2xx status code
func (o *FrontSessionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front session not found response has a 3xx status code
func (o *FrontSessionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front session not found response has a 4xx status code
func (o *FrontSessionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this front session not found response has a 5xx status code
func (o *FrontSessionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this front session not found response a status code equal to that given
func (o *FrontSessionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the front session not found response
func (o *FrontSessionNotFound) Code() int {
	return 404
}

func (o *FrontSessionNotFound) Error() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionNotFound  %+v", 404, o.Payload)
}

func (o *FrontSessionNotFound) String() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionNotFound  %+v", 404, o.Payload)
}

func (o *FrontSessionNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontSessionInternalServerError creates a FrontSessionInternalServerError with default headers values
func NewFrontSessionInternalServerError() *FrontSessionInternalServerError {
	return &FrontSessionInternalServerError{}
}

/*
FrontSessionInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type FrontSessionInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front session internal server error response has a 2xx status code
func (o *FrontSessionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front session internal server error response has a 3xx status code
func (o *FrontSessionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front session internal server error response has a 4xx status code
func (o *FrontSessionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this front session internal server error response has a 5xx status code
func (o *FrontSessionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this front session internal server error response a status code equal to that given
func (o *FrontSessionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the front session internal server error response
func (o *FrontSessionInternalServerError) Code() int {
	return 500
}

func (o *FrontSessionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *FrontSessionInternalServerError) String() string {
	return fmt.Sprintf("[POST /frontend/session][%d] frontSessionInternalServerError  %+v", 500, o.Payload)
}

func (o *FrontSessionInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontSessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
