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

// FrontStateReader is a Reader for the FrontState structure.
type FrontStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FrontStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFrontStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewFrontStateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFrontStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFrontStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFrontStateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewFrontStateOK creates a FrontStateOK with default headers values
func NewFrontStateOK() *FrontStateOK {
	return &FrontStateOK{}
}

/*
FrontStateOK describes a response with status code 200, with default header values.

A successful response.
*/
type FrontStateOK struct {
	Payload *models.RestFrontStateResponse
}

// IsSuccess returns true when this front state o k response has a 2xx status code
func (o *FrontStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this front state o k response has a 3xx status code
func (o *FrontStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front state o k response has a 4xx status code
func (o *FrontStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this front state o k response has a 5xx status code
func (o *FrontStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this front state o k response a status code equal to that given
func (o *FrontStateOK) IsCode(code int) bool {
	return code == 200
}

func (o *FrontStateOK) Error() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateOK  %+v", 200, o.Payload)
}

func (o *FrontStateOK) String() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateOK  %+v", 200, o.Payload)
}

func (o *FrontStateOK) GetPayload() *models.RestFrontStateResponse {
	return o.Payload
}

func (o *FrontStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestFrontStateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontStateUnauthorized creates a FrontStateUnauthorized with default headers values
func NewFrontStateUnauthorized() *FrontStateUnauthorized {
	return &FrontStateUnauthorized{}
}

/*
FrontStateUnauthorized describes a response with status code 401, with default header values.

User is not authenticated
*/
type FrontStateUnauthorized struct {
}

// IsSuccess returns true when this front state unauthorized response has a 2xx status code
func (o *FrontStateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front state unauthorized response has a 3xx status code
func (o *FrontStateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front state unauthorized response has a 4xx status code
func (o *FrontStateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this front state unauthorized response has a 5xx status code
func (o *FrontStateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this front state unauthorized response a status code equal to that given
func (o *FrontStateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *FrontStateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateUnauthorized ", 401)
}

func (o *FrontStateUnauthorized) String() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateUnauthorized ", 401)
}

func (o *FrontStateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFrontStateForbidden creates a FrontStateForbidden with default headers values
func NewFrontStateForbidden() *FrontStateForbidden {
	return &FrontStateForbidden{}
}

/*
FrontStateForbidden describes a response with status code 403, with default header values.

User has no permission to access this resource
*/
type FrontStateForbidden struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front state forbidden response has a 2xx status code
func (o *FrontStateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front state forbidden response has a 3xx status code
func (o *FrontStateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front state forbidden response has a 4xx status code
func (o *FrontStateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this front state forbidden response has a 5xx status code
func (o *FrontStateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this front state forbidden response a status code equal to that given
func (o *FrontStateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *FrontStateForbidden) Error() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateForbidden  %+v", 403, o.Payload)
}

func (o *FrontStateForbidden) String() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateForbidden  %+v", 403, o.Payload)
}

func (o *FrontStateForbidden) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontStateNotFound creates a FrontStateNotFound with default headers values
func NewFrontStateNotFound() *FrontStateNotFound {
	return &FrontStateNotFound{}
}

/*
FrontStateNotFound describes a response with status code 404, with default header values.

Resource does not exist in the system
*/
type FrontStateNotFound struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front state not found response has a 2xx status code
func (o *FrontStateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front state not found response has a 3xx status code
func (o *FrontStateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front state not found response has a 4xx status code
func (o *FrontStateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this front state not found response has a 5xx status code
func (o *FrontStateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this front state not found response a status code equal to that given
func (o *FrontStateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *FrontStateNotFound) Error() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateNotFound  %+v", 404, o.Payload)
}

func (o *FrontStateNotFound) String() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateNotFound  %+v", 404, o.Payload)
}

func (o *FrontStateNotFound) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFrontStateInternalServerError creates a FrontStateInternalServerError with default headers values
func NewFrontStateInternalServerError() *FrontStateInternalServerError {
	return &FrontStateInternalServerError{}
}

/*
FrontStateInternalServerError describes a response with status code 500, with default header values.

An internal error occurred in the backend
*/
type FrontStateInternalServerError struct {
	Payload *models.RestError
}

// IsSuccess returns true when this front state internal server error response has a 2xx status code
func (o *FrontStateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this front state internal server error response has a 3xx status code
func (o *FrontStateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this front state internal server error response has a 4xx status code
func (o *FrontStateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this front state internal server error response has a 5xx status code
func (o *FrontStateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this front state internal server error response a status code equal to that given
func (o *FrontStateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *FrontStateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateInternalServerError  %+v", 500, o.Payload)
}

func (o *FrontStateInternalServerError) String() string {
	return fmt.Sprintf("[GET /frontend/state][%d] frontStateInternalServerError  %+v", 500, o.Payload)
}

func (o *FrontStateInternalServerError) GetPayload() *models.RestError {
	return o.Payload
}

func (o *FrontStateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
